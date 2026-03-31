package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	authdomain "github.com/ferilee/api-idetech/backend/internal/auth/domain"
	tenantdomain "github.com/ferilee/api-idetech/backend/internal/tenant/domain"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type userRepository interface {
	FindByTenantAndIdentity(ctx context.Context, tenantSlug, identity string) (authdomain.User, error)
	FindByID(ctx context.Context, id string) (authdomain.User, error)
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
	jwtSecret   []byte
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

func NewService(users userRepository, tenants tenantRepository, jwtIssuer, jwtAudience, jwtSecret string) *Service {
	return &Service{
		users:       users,
		tenants:     tenants,
		jwtIssuer:   jwtIssuer,
		jwtAudience: jwtAudience,
		jwtSecret:   []byte(jwtSecret),
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

	expiresAt := time.Now().Add(15 * time.Minute)
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
