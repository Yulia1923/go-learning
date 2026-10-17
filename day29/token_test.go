package main

import "testing"

func TestGenerateAndParseToken(t *testing.T) {
	tokenString, err := GenerateToken(1001)
	if err != nil {
		t.Fatal("生成 Token 失败: ", err)
	}

	claims, err := ParseToken(tokenString)
	if err != nil {
		t.Fatal("解析 Token 失败: ", err)
	}

	if claims.UserID != 1001 {
		t.Fatalf("用户 ID 不正确, 期望1001, 实际 %d", claims.UserID)
	}

	t.Log("Token 验证成功, 用户 ID: ", claims.UserID)
	t.Log("Token: ", tokenString)
}
