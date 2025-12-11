package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager interface {
	GenerateTokens(userId string, role string) (accessToken string, refreshToken string, err error)
	ParseToken(tokenString string) (userClaims *UserClaims, err error)
}

type Manager struct {
	signingKey      string
	accessTokenTtl  time.Duration
	refreshTokenTtl time.Duration
}

type UserClaims struct {
	jwt.RegisteredClaims
	UserId    string `json:"user_id"`
	UserRole  string `json:"user_role"`
	TokenType string `json:"token_type"`
}

func NewManager(signingKey string, accessTokenTtl, refreshTokenTtl time.Duration) (*Manager, error) {
	if signingKey == "" {
		return nil, fmt.Errorf("empty signing key")
	}

	return &Manager{
		signingKey:      signingKey,
		accessTokenTtl:  accessTokenTtl,
		refreshTokenTtl: refreshTokenTtl,
	}, nil
}

func (m *Manager) createJWT(userId, userRole, tokenType string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&UserClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
			UserId:    userId,
			UserRole:  userRole,
			TokenType: tokenType,
		},
	)

	return token.SignedString([]byte(m.signingKey))
}

func (m *Manager) GenerateTokens(userId, userRole string) (string, string, error) {
	refresh, err := m.createJWT(userId, userRole, "refresh", m.refreshTokenTtl)
	if err != nil {
		return "", "", err
	}

	access, err := m.createJWT(userId, userRole, "access", m.accessTokenTtl)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

func (m *Manager) ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
