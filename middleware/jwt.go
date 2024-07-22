package middleware

import (
	"strings"
	"time"
	"xwya/model"
	"xwya/utils"
	"xwya/utils/msg"
	"xwya/utils/msgjson"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JwtKey = []byte(utils.JwtKey)

// MyClaims 结构表示 JWT 中的声明
type MyClaims struct {
	User model.User
	jwt.StandardClaims
}

// SetToken 生成 JWT Token 并返回
func SetToken(user model.User) (string, int) {
	// 设置 Token 过期时间为 3 天
	expireTime := time.Now().Add(time.Hour * 72).Unix()
	// 创建自定义声明结构
	claims := MyClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireTime, // 过期时间
			Issuer:    "xwya",
		},
	}
	// 使用 HS256 签名算法创建 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签名 Token 并返回
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", msg.Error
	}
	return tokenString, msg.Success
}

// CheckToken 验证 JWT Token 并返回声明
func CheckToken(tokenString string) *MyClaims {
	// 解析 Token 并验证签名
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	// 处理解析和验证错误
	if err != nil || err == jwt.ErrSignatureInvalid || !token.Valid {
		return nil // Token 格式错误
	}
	// 检查 Token 是否有效并返回声明
	claims, ok := token.Claims.(*MyClaims)
	if !ok || !token.Valid {
		return nil
	}

	return claims
}

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("xwya_admin_token")
		if err != nil {
			msgjson.ErrorMsg(c, msg.Error_TokenExist) // token 错误
			c.Abort()
			return
		}

		// 将 Token 头分成两部分，"Bearer" 和 Token 字符串
		checkToken := strings.SplitN(token, " ", 2)
		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			msgjson.ErrorMsg(c, msg.Error_TokenMalformed) // token 格式错误
			c.Abort()
			return
		}
		// 验证 Token 并获取声明
		key := CheckToken(checkToken[1])
		if key == nil {
			msgjson.ErrorMsg(c, msg.Error_TokenInvalid) // token 无效
			c.Abort()
			return
		}

		// 检查 Token 是否已过期
		if time.Now().Unix() > key.ExpiresAt {
			msgjson.ErrorMsg(c, msg.Error_TokenTimeout) // token 过期
			c.Abort()
			return
		}
		c.Set("user", key.User)
		c.Next()
	}
}
