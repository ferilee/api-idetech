package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	authrepo "github.com/ferilee/api-idetech/backend/internal/auth/repository"
	authservice "github.com/ferilee/api-idetech/backend/internal/auth/service"
	"github.com/ferilee/api-idetech/backend/internal/platform/config"
	apphttp "github.com/ferilee/api-idetech/backend/internal/platform/http"
	tenantrepo "github.com/ferilee/api-idetech/backend/internal/tenant/repository"
	tenantservice "github.com/ferilee/api-idetech/backend/internal/tenant/service"
)

func main() {
	cfg := config.MustLoad()

	tenantRepository := tenantrepo.NewMemoryRepository()
	tenantRepository.SeedDefaults()

	authRepository := authrepo.NewMemoryRepository()
	if err := authRepository.SeedDefaults(); err != nil {
		log.Fatalf("failed to seed auth repository: %v", err)
	}

	tenantService := tenantservice.NewService(tenantRepository)
	authService := authservice.NewService(
		authRepository,
		tenantRepository,
		cfg.JWTIssuer,
		cfg.JWTAudience,
		cfg.JWTSecret,
	)
	handler := apphttp.NewHandler(cfg, authService, tenantService)

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
