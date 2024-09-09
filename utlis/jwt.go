package utlis

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	// 用于签名的密钥
	JwtKey = []byte(GetConfBody().Secret)
)

// JWT 结构体
type Claims struct {
	Username string
	Password string
	jwt.StandardClaims
}

// 生成 JWT 的函数 expirationMinutes 过期时间分钟
func GenerateJWT(username string, password string, expirationMinutes int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(expirationMinutes) * time.Minute)
	claims := &Claims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名
	return token.SignedString(JwtKey)
}

// 解析 JWT 的函数
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
