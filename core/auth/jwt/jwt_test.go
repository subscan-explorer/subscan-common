package jwt

import (
	"fmt"
	"testing"
	"time"
)

var cfg *Config

func init() {
	cfg = &Config{
		SigningKey: "somebody_else",
		Issuer:     "subscan",
		Subject:    "open platform",
		Audience:   []string{"develop"},
		Expires:    time.Now().Add(24 * time.Hour),
	}
}

type User struct {
	ID   int64
	Name string
}

// go test -v -test.run TestGenToken
func TestGenToken(t *testing.T) {
	u := &User{
		ID:   1,
		Name: "xiequan",
	}
	tokenStr, err := GenToken(cfg, u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tokenStr)
	res, err := ParseToken(cfg, tokenStr)
	if err != nil {
		fmt.Println("===ParseToken====", err)
		return
	}
	fmt.Println(res.User)
}
