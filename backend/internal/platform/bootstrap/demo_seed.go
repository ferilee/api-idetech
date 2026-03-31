package bootstrap

import (
	"context"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func SeedDemoData(ctx context.Context, db *sql.DB) error {
	tenantID, err := ensureDemoTenant(ctx, db)
	if err != nil {
		return err
	}

	if err := ensureDemoUser(ctx, db, tenantID, "guru.demo", "guru.demo@idetech.local", "teacher", "demo123", "Guru Demo"); err != nil {
		return err
	}

	if err := ensureDemoUser(ctx, db, tenantID, "admin.demo", "admin.demo@idetech.local", "admin", "admin123", "Admin Demo"); err != nil {
		return err
	}

	return nil
}

func ensureDemoTenant(ctx context.Context, db *sql.DB) (string, error) {
	const query = `
INSERT INTO tenants (name, slug, domain, status, config)
VALUES (
  'IdeTech Demo School',
  'demo',
  'demo.idetech.local',
  'active',
  '{"theme":{"brandName":"IdeTech Demo","primary":"#1d4ed8","secondary":"#f59e0b","appearance":"junior"}}'::jsonb
)
ON CONFLICT (slug) DO UPDATE SET
  name = EXCLUDED.name,
  domain = EXCLUDED.domain,
  status = EXCLUDED.status,
  config = EXCLUDED.config,
  updated_at = NOW()
RETURNING id::text;
`

	var tenantID string
	if err := db.QueryRowContext(ctx, query).Scan(&tenantID); err != nil {
		return "", fmt.Errorf("seed demo tenant: %w", err)
	}

	return tenantID, nil
}

func ensureDemoUser(ctx context.Context, db *sql.DB, tenantID, username, email, role, password, displayName string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	const query = `
INSERT INTO users (tenant_id, username, email, password_hash, role, profile_data, is_active)
VALUES ($1::uuid, $2, $3, $4, $5, jsonb_build_object('display_name', $6::text), TRUE)
ON CONFLICT (tenant_id, username) DO UPDATE SET
  email = EXCLUDED.email,
  password_hash = EXCLUDED.password_hash,
  role = EXCLUDED.role,
  profile_data = EXCLUDED.profile_data,
  updated_at = NOW();
`

	if _, err := db.ExecContext(ctx, query, tenantID, username, email, string(hash), role, displayName); err != nil {
		return fmt.Errorf("seed demo user %s: %w", username, err)
	}

	return nil
}
