package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	authdomain "github.com/ferilee/api-idetech/backend/internal/auth/domain"
	authrepo "github.com/ferilee/api-idetech/backend/internal/auth/repository"
	authservice "github.com/ferilee/api-idetech/backend/internal/auth/service"
	"github.com/ferilee/api-idetech/backend/internal/platform/bootstrap"
	"github.com/ferilee/api-idetech/backend/internal/platform/config"
	"github.com/ferilee/api-idetech/backend/internal/platform/database"
	apphttp "github.com/ferilee/api-idetech/backend/internal/platform/http"
	tenantdomain "github.com/ferilee/api-idetech/backend/internal/tenant/domain"
	tenantrepo "github.com/ferilee/api-idetech/backend/internal/tenant/repository"
	tenantservice "github.com/ferilee/api-idetech/backend/internal/tenant/service"
	userservice "github.com/ferilee/api-idetech/backend/internal/user/service"
)

type tenantRepository interface {
	FindBySlug(ctx context.Context, slug string) (tenantdomain.Tenant, error)
}

type authRepository interface {
	FindByTenantAndIdentity(ctx context.Context, tenantSlug, identity string) (authdomain.User, error)
	FindByID(ctx context.Context, id string) (authdomain.User, error)
	ListByTenant(ctx context.Context, tenantSlug string) ([]authdomain.User, error)
}

func main() {
	cfg := config.MustLoad()

	memoryTenantRepository := tenantrepo.NewMemoryRepository()
	memoryTenantRepository.SeedDefaults()
	var tenantRepository tenantRepository = memoryTenantRepository

	memoryAuthRepository := authrepo.NewMemoryRepository()
	if err := memoryAuthRepository.SeedDefaults(); err != nil {
		log.Fatalf("failed to seed auth repository: %v", err)
	}
	var authRepository authRepository = memoryAuthRepository

	if dsn := cfg.PostgresDSN(); dsn != "" {
		db, err := database.OpenPostgres(context.Background(), dsn)
		if err != nil {
			log.Fatalf("failed to connect postgres: %v", err)
		}
		defer db.Close()

		if err := bootstrap.SeedDemoData(context.Background(), db); err != nil {
			log.Fatalf("failed to seed postgres demo data: %v", err)
		}

		tenantRepository = tenantrepo.NewPostgresRepository(db)
		authRepository = authrepo.NewPostgresRepository(db)
		log.Printf("postgres repositories enabled")
	}

	tenantService := tenantservice.NewService(tenantRepository)
	authService := authservice.NewService(
		authRepository,
		tenantRepository,
		cfg.JWTIssuer,
		cfg.JWTAudience,
		cfg.JWTSecret,
	)
	userService := userservice.NewService(authRepository)
	handler := apphttp.NewHandler(cfg, authService, tenantService, userService)

	server := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           handler.Router(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("idetech api listening on :%s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed: %v", err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	}
}
