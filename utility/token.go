package utility

import (
	"SheeDrive/internal/consts"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 生成Token
func GenToken(username string) string {
	claim := jwt.RegisteredClaims{
		Subject: username,
		// 设置过期时间
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(consts.JwtTokenKey))

	if err != nil {
		panic("Token生成错误！")
	}

	return token
}
