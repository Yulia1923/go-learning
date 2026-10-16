package main

import (
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateAndParseToken(t *testing.T) {
	//设置仅供测试使用的密钥
	t.Setenv("JWT_SECRET", "testing-secret-key-for-jwt")

	//生成用户ID为123的Token
	tokenString, err := GenerateToken(123)
	if err != nil {
		t.Fatalf("生成 Token 失败 :%v", err)
	}

	//解析token
	claims, err := ParseToken(tokenString)
	if err != nil {
		t.Fatalf("解析 Token 失败 :%v", err)
	}

	//检查用户ID
	if claims.UserID != 123 {
		t.Errorf("期望 UserID 为123, 实际为 %d", claims.UserID)
	}

	//检查过期时间是否存在
	if claims.ExpiresAt == nil {
		t.Fatal("Token 没有设置过期时间")
	}

	//检查过期时间是否在未来
	if !claims.ExpiresAt.Time.After(time.Now()) {
		t.Error("Token 的过期时间应该在未来")
	}

	_ = os.Getenv("JWT_SECRET")

}

func TestParseTamperedToken(t *testing.T) {
	t.Setenv("JWT_SECRET", "test-secret-key-for-jwt")

	//先生成一个合法Token
	tokenString, err := GenerateToken(123)
	if err != nil {
		t.Fatalf("生成Token失败 : %v", err)
	}

	//篡改 Token 最后一个字符
	last := tokenString[len(tokenString)-1]
	replacement := byte('A')
	if last == replacement {
		replacement = 'B'
	}
	tamperedToken := tokenString[:len(tokenString)-1] +
		string(replacement)

	//解析被篡改的Token
	_, err = ParseToken(tamperedToken)
	if err == nil {
		t.Fatal("被篡改的 Token 不应该通过验证")
	}
}
func TestParseExpiredToken(t *testing.T) {
	secret := "test-secret-key-for-jwt"
	t.Setenv("JWT_SECRET", secret)

	now := time.Now()

	claims := Claims{
		UserID: 123,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now.Add(-2 * time.Minute)),
			ExpiresAt: jwt.NewNumericDate(now.Add(-1 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("签名失败: %v", err)
	}

	//尝试解析过期Token
	_, err = ParseToken(tokenString)
	if err == nil {
		t.Fatal("过期 Token 不应该通过验证")
	}
}
