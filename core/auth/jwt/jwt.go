package jwt

import (
	"errors"
	"time"

	jwtV4 "github.com/golang-jwt/jwt/v4"
	"github.com/subscan-explorer/subscan-common/core/util/xtime"
)

type MyClaims struct {
	User interface{} `json:"user"`
	jwtV4.RegisteredClaims
}

type Config struct {
	ID         string
	SigningKey string
	Issuer     string
	Subject    string
	Audience   []string
	Expires    xtime.Duration
}

func GenToken(c *Config, user interface{}) (string, error) {
	claims := MyClaims{
		user,
		jwtV4.RegisteredClaims{
			ExpiresAt: jwtV4.NewNumericDate(time.Now().Add(time.Duration(c.Expires))),
			IssuedAt:  jwtV4.NewNumericDate(time.Now()),
			NotBefore: jwtV4.NewNumericDate(time.Now()),
			Issuer:    c.Issuer,
			Subject:   c.Subject,
			ID:        c.ID,
			Audience:  c.Audience,
		},
	}
	token := jwtV4.NewWithClaims(jwtV4.SigningMethodHS256, claims)
	return token.SignedString([]byte(c.SigningKey))
}

func ParseToken(c *Config, tokenStr string) (*MyClaims, error) {
	token, err := jwtV4.ParseWithClaims(tokenStr, &MyClaims{}, func(t *jwtV4.Token) (interface{}, error) {
		return []byte(c.SigningKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
