package repository

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"golang.org/x/crypto/bcrypt"

	"github.com/ferilee/api-idetech/backend/internal/auth/domain"
)

type MemoryRepository struct {
	mu    sync.RWMutex
	users map[string]domain.User
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		users: make(map[string]domain.User),
	}
}

func (r *MemoryRepository) SeedDefaults() error {
	teacherHash, err := bcrypt.GenerateFromPassword([]byte("demo123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	adminHash, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[keyFor("demo", "guru.demo")] = domain.User{
		ID:           "user-demo-teacher",
		TenantSlug:   "demo",
		Username:     "guru.demo",
		Email:        "guru.demo@idetech.local",
		Role:         "teacher",
		PasswordHash: string(teacherHash),
		Profile: map[string]any{
			"display_name": "Guru Demo",
		},
	}

	r.users[keyFor("demo", "admin.demo")] = domain.User{
		ID:           "user-demo-admin",
		TenantSlug:   "demo",
		Username:     "admin.demo",
		Email:        "admin.demo@idetech.local",
		Role:         "admin",
		PasswordHash: string(adminHash),
		Profile: map[string]any{
			"display_name": "Admin Demo",
		},
	}

	return nil
}

func (r *MemoryRepository) FindByTenantAndIdentity(_ context.Context, tenantSlug, identity string) (domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[keyFor(tenantSlug, identity)]
	if ok {
		return user, nil
	}

	for _, user := range r.users {
		if user.TenantSlug == tenantSlug && strings.EqualFold(user.Email, identity) {
			return user, nil
		}
	}

	return domain.User{}, fmt.Errorf("user not found")
}

func (r *MemoryRepository) FindByID(_ context.Context, id string) (domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}

	return domain.User{}, fmt.Errorf("user not found")
}

func keyFor(tenantSlug, identity string) string {
	return strings.ToLower(strings.TrimSpace(tenantSlug)) + ":" + strings.ToLower(strings.TrimSpace(identity))
}
