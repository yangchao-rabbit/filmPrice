package auth

import (
	"errors"
	"filmPrice/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaims struct {
	User string `json:"user"`
	jwt.RegisteredClaims
}

// GenToken 生成JWT Token
func GenToken(user string) (string, error) {
	claims := CustomClaims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 7 天有效
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Get().App.Secret))
}

func ParseToken(token string) (*CustomClaims, error) {
	tokenStr, err := jwt.ParseWithClaims(
		token,
		&CustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			// 可选校验签名算法
			if t.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(config.Get().App.Secret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := tokenStr.Claims.(*CustomClaims)
	if !ok || !tokenStr.Valid {
		return nil, errors.New("invalid or expired token")
	}

	return claims, nil
}
