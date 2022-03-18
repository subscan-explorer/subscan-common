package jwt

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/subscan-explorer/subscan-common/core/util/xtime"
)

var cfg *Config

func init() {
	cfg = &Config{
		SigningKey: "somebody_else",
		Issuer:     "subscan",
		Subject:    "open platform",
		Audience:   []string{"develop"},
		Expires:    xtime.Duration(time.Hour * 24),
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
	strByte, _ := json.Marshal(u)
	tokenStr, err := GenToken(cfg, string(strByte))
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
	fmt.Println(res.Info)
}
