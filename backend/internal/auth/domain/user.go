package domain

type User struct {
	ID           string         `json:"id"`
	TenantSlug   string         `json:"tenant_slug"`
	Username     string         `json:"username"`
	Email        string         `json:"email"`
	Role         string         `json:"role"`
	PasswordHash string         `json:"-"`
	Profile      map[string]any `json:"profile"`
}
