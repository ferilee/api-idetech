package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/idtoken"

	authdomain "github.com/ferilee/api-idetech/backend/internal/auth/domain"
	tenantdomain "github.com/ferilee/api-idetech/backend/internal/tenant/domain"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type userRepository interface {
	FindByTenantAndIdentity(ctx context.Context, tenantSlug, identity string) (authdomain.User, error)
	FindByID(ctx context.Context, id string) (authdomain.User, error)
	Create(ctx context.Context, user authdomain.User) error
}

type tenantRepository interface {
	FindBySlug(ctx context.Context, slug string) (tenantdomain.Tenant, error)
}

type TokenClaims struct {
	UserID     string `json:"user_id"`
	TenantSlug string `json:"tenant_slug"`
	Username   string `json:"username"`
	Role       string `json:"role"`
	jwt.RegisteredClaims
}

type Service struct {
	users       userRepository
	tenants     tenantRepository
	jwtIssuer   string
	jwtAudience string
	jwtSecret      []byte
	googleClientID string
}

type LoginInput struct {
	TenantSlug string `json:"tenant_slug"`
	Identity   string `json:"identity"`
	Password   string `json:"password"`
}

type LoginResult struct {
	AccessToken string              `json:"access_token"`
	TokenType   string              `json:"token_type"`
	ExpiresIn   int64               `json:"expires_in"`
	User        authdomain.User     `json:"user"`
	Tenant      tenantdomain.Tenant `json:"tenant"`
}

type LoginGoogleInput struct {
	TenantSlug string `json:"tenant_slug"`
	IDToken    string `json:"id_token"`
}

func NewService(users userRepository, tenants tenantRepository, jwtIssuer, jwtAudience, jwtSecret, googleClientID string) *Service {
	return &Service{
		users:          users,
		tenants:        tenants,
		jwtIssuer:      jwtIssuer,
		jwtAudience:    jwtAudience,
		jwtSecret:      []byte(jwtSecret),
		googleClientID: googleClientID,
	}
}

func (s *Service) Login(ctx context.Context, input LoginInput) (LoginResult, error) {
	tenant, err := s.tenants.FindBySlug(ctx, input.TenantSlug)
	if err != nil {
		return LoginResult{}, ErrInvalidCredentials
	}

	user, err := s.users.FindByTenantAndIdentity(ctx, input.TenantSlug, input.Identity)
	if err != nil {
		return LoginResult{}, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return LoginResult{}, ErrInvalidCredentials
	}

	return s.issueToken(user, tenant)
}

func (s *Service) LoginGoogle(ctx context.Context, input LoginGoogleInput) (LoginResult, error) {
	tenant, err := s.tenants.FindBySlug(ctx, input.TenantSlug)
	if err != nil {
		return LoginResult{}, ErrInvalidCredentials
	}

	payload, err := idtoken.Validate(ctx, input.IDToken, s.googleClientID)
	if err != nil {
		return LoginResult{}, errors.New("invalid google token: " + err.Error())
	}

	email, ok := payload.Claims["email"].(string)
	if !ok {
		return LoginResult{}, errors.New("email not found in google token")
	}

	user, err := s.users.FindByTenantAndIdentity(ctx, input.TenantSlug, email)
	if err != nil {
		// Just-In-Time Provisioning
		role := "student"
		if email == "the.real.ferilee@gmail.com" {
			role = "admin"
		} else if strings.HasSuffix(email, "@guru.smk.belajar.id") ||
			strings.HasSuffix(email, "@guru.sma.belajar.id") ||
			strings.HasSuffix(email, "@guru.smp.belajar.id") ||
			strings.HasSuffix(email, "@guru.sd.belajar.id") {
			role = "teacher"
		}

		name, _ := payload.Claims["name"].(string)
		if name == "" {
			name = email
		}

		user = authdomain.User{
			ID:         email, // Will be replaced by UUID in postgres, but fine for memory
			TenantSlug: input.TenantSlug,
			Username:   email,
			Email:      email,
			Role:       role,
			Profile: map[string]any{
				"display_name": name,
				"picture":      payload.Claims["picture"],
			},
		}

		if err := s.users.Create(ctx, user); err != nil {
			return LoginResult{}, errors.New("failed to provision user: " + err.Error())
		}
		
		// Re-fetch to get correct ID (especially for postgres)
		user, _ = s.users.FindByTenantAndIdentity(ctx, input.TenantSlug, email)
	}

	return s.issueToken(user, tenant)
}

func (s *Service) issueToken(user authdomain.User, tenant tenantdomain.Tenant) (LoginResult, error) {
	expiresAt := time.Now().Add(24 * time.Hour) // Longer session for mobile-first feel
	claims := TokenClaims{
		UserID:     user.ID,
		TenantSlug: user.TenantSlug,
		Username:   user.Username,
		Role:       user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID,
			Issuer:    s.jwtIssuer,
			Audience:  []string{s.jwtAudience},
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return LoginResult{}, err
	}

	return LoginResult{
		AccessToken: signedToken,
		TokenType:   "Bearer",
		ExpiresIn:   int64(time.Until(expiresAt).Seconds()),
		User:        sanitizeUser(user),
		Tenant:      tenant,
	}, nil
}

func (s *Service) ParseToken(tokenString string) (*TokenClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (any, error) {
		return s.jwtSecret, nil
	}, jwt.WithAudience(s.jwtAudience), jwt.WithIssuer(s.jwtIssuer))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	claims, ok := parsedToken.Claims.(*TokenClaims)
	if !ok || !parsedToken.Valid {
		return nil, ErrInvalidCredentials
	}

	return claims, nil
}

func (s *Service) Me(ctx context.Context, claims *TokenClaims) (authdomain.User, error) {
	user, err := s.users.FindByID(ctx, claims.UserID)
	if err != nil {
		return authdomain.User{}, err
	}
	return sanitizeUser(user), nil
}

func sanitizeUser(user authdomain.User) authdomain.User {
	user.PasswordHash = ""
	return user
}
