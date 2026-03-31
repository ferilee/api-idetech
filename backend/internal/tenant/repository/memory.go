package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"

	"github.com/ferilee/api-idetech/backend/internal/tenant/domain"
)

type MemoryRepository struct {
	mu      sync.RWMutex
	tenants map[string]domain.Tenant
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		tenants: make(map[string]domain.Tenant),
	}
}

func (r *MemoryRepository) SeedDefaults() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tenants["demo"] = domain.Tenant{
		ID:     uuid.NewString(),
		Name:   "IdeTech Demo School",
		Slug:   "demo",
		Domain: "demo.idetech.local",
		Status: "active",
		Config: map[string]any{
			"theme": map[string]any{
				"brandName":  "IdeTech Demo",
				"primary":    "#1d4ed8",
				"secondary":  "#f59e0b",
				"appearance": "junior",
			},
		},
	}
}

func (r *MemoryRepository) FindBySlug(_ context.Context, slug string) (domain.Tenant, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tenant, ok := r.tenants[slug]
	if !ok {
		return domain.Tenant{}, fmt.Errorf("tenant %q not found", slug)
	}
	return tenant, nil
}
