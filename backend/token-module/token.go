package token_module

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	signingKey = []byte("lu@P$dM&Hw*43JC^bf0KegyIOk#YDsZtxmcLFU1Gz!n25aiWv6jqQE7%ATrSBp8h9oRXNV") // 用于签名和验证令牌的密钥
)

func GetToken(username string) string {
	claims := jwt.MapClaims{
		"username": username,                               //传输进来的 用户名
		"exp":      time.Now().Add(time.Hour * 100).Unix(), // 设置过期时间为100小时
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//签名
	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		fmt.Println("令牌签名失败:", err)
		return ""
	}

	return signedToken
}

// 解析并验证令牌
func parseToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法是否有效
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("无效的签名方法: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
