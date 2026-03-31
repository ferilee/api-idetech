package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/ferilee/api-idetech/backend/internal/tenant/domain"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) FindBySlug(ctx context.Context, slug string) (domain.Tenant, error) {
	const query = `
SELECT id::text, name, slug, COALESCE(domain, ''), status, config
FROM tenants
WHERE slug = $1
LIMIT 1;
`

	var tenant domain.Tenant
	var rawConfig []byte
	if err := r.db.QueryRowContext(ctx, query, slug).Scan(
		&tenant.ID,
		&tenant.Name,
		&tenant.Slug,
		&tenant.Domain,
		&tenant.Status,
		&rawConfig,
	); err != nil {
		if err == sql.ErrNoRows {
			return domain.Tenant{}, fmt.Errorf("tenant %q not found", slug)
		}
		return domain.Tenant{}, fmt.Errorf("find tenant by slug: %w", err)
	}

	if len(rawConfig) > 0 {
		if err := json.Unmarshal(rawConfig, &tenant.Config); err != nil {
			return domain.Tenant{}, fmt.Errorf("decode tenant config: %w", err)
		}
	}

	return tenant, nil
}
