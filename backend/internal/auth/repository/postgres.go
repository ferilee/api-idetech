package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/ferilee/api-idetech/backend/internal/auth/domain"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) FindByTenantAndIdentity(ctx context.Context, tenantSlug, identity string) (domain.User, error) {
	const query = `
SELECT
  u.id::text,
  t.slug,
  u.username,
  COALESCE(u.email, ''),
  u.role,
  u.password_hash,
  u.profile_data
FROM users u
JOIN tenants t ON t.id = u.tenant_id
WHERE t.slug = $1
  AND (LOWER(u.username) = LOWER($2) OR LOWER(COALESCE(u.email, '')) = LOWER($2))
  AND u.is_active = TRUE
LIMIT 1;
`

	return scanUser(r.db.QueryRowContext(ctx, query, tenantSlug, identity))
}

func (r *PostgresRepository) FindByID(ctx context.Context, id string) (domain.User, error) {
	const query = `
SELECT
  u.id::text,
  t.slug,
  u.username,
  COALESCE(u.email, ''),
  u.role,
  u.password_hash,
  u.profile_data
FROM users u
JOIN tenants t ON t.id = u.tenant_id
WHERE u.id = $1::uuid
LIMIT 1;
`

	return scanUser(r.db.QueryRowContext(ctx, query, id))
}

type rowScanner interface {
	Scan(dest ...any) error
}

func scanUser(row rowScanner) (domain.User, error) {
	var user domain.User
	var rawProfile []byte
	if err := row.Scan(
		&user.ID,
		&user.TenantSlug,
		&user.Username,
		&user.Email,
		&user.Role,
		&user.PasswordHash,
		&rawProfile,
	); err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, fmt.Errorf("user not found")
		}
		return domain.User{}, fmt.Errorf("scan user: %w", err)
	}

	if len(rawProfile) > 0 {
		if err := json.Unmarshal(rawProfile, &user.Profile); err != nil {
			return domain.User{}, fmt.Errorf("decode user profile: %w", err)
		}
	}

	return user, nil
}
