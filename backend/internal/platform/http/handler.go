package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"

	authservice "github.com/ferilee/api-idetech/backend/internal/auth/service"
	"github.com/ferilee/api-idetech/backend/internal/platform/config"
	platformmiddleware "github.com/ferilee/api-idetech/backend/internal/platform/http/middleware"
	"github.com/ferilee/api-idetech/backend/internal/tenant/service"
)

type Handler struct {
	cfg           config.Config
	authService   *authservice.Service
	tenantService *service.Service
	router        chi.Router
}

func NewHandler(cfg config.Config, authService *authservice.Service, tenantService *service.Service) *Handler {
	h := &Handler{
		cfg:           cfg,
		authService:   authService,
		tenantService: tenantService,
		router:        chi.NewRouter(),
	}

	h.registerRoutes()
	return h
}

func (h *Handler) Router() http.Handler {
	return h.router
}

func (h *Handler) registerRoutes() {
	h.router.Use(chimiddleware.RequestID)
	h.router.Use(chimiddleware.RealIP)
	h.router.Use(chimiddleware.Recoverer)
	h.router.Use(chimiddleware.Timeout(30 * time.Second))
	h.router.Use(platformmiddleware.CORS(h.cfg.AllowedOrigins))

	h.router.Get("/healthz", h.handleHealth)

	h.router.Route("/api/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", h.handleLogin)
			r.With(platformmiddleware.RequireAuth(h.authService)).Get("/me", h.handleMe)
		})
		r.Get("/tenant/bootstrap", h.handleTenantBootstrap)
	})
}

func (h *Handler) handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"status": "ok",
		"env":    h.cfg.AppEnv,
	})
}

func (h *Handler) handleTenantBootstrap(w http.ResponseWriter, r *http.Request) {
	slug := platformmiddleware.ResolveTenantSlug(r)
	if slug == "" {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"error": "tenant slug could not be resolved from host or header",
		})
		return
	}

	tenant, err := h.tenantService.Bootstrap(r.Context(), slug)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]any{
			"error": err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"tenant": tenant,
	})
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var input authservice.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"error": "invalid login payload",
		})
		return
	}

	result, err := h.authService.Login(r.Context(), input)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, authservice.ErrInvalidCredentials) {
			status = http.StatusUnauthorized
		}
		writeJSON(w, status, map[string]any{
			"error": err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, result)
}

func (h *Handler) handleMe(w http.ResponseWriter, r *http.Request) {
	claims, ok := platformmiddleware.ClaimsFromContext(r.Context())
	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]any{
			"error": "missing auth context",
		})
		return
	}

	user, err := h.authService.Me(r.Context(), claims)
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, map[string]any{
			"error": "user not found",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"user": user,
	})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
