package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ylinyang/vue-demo-go/utils"
	"log"
	"net/http"
	"strings"
	"time"
)

var jwtKey = []byte("abc")

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// SetToken 生成token
func SetToken(u string) string {
	setClaims := MyClaims{
		u,
		jwt.StandardClaims{
			// token过期时间
			ExpiresAt: time.Now().Add(10 * time.Hour).Unix(),
			Issuer:    "demo",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, setClaims)
	signedString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Panicln("签发token失败", err)
	}
	return signedString
}

// 验证token
func checkToken(token string) (*MyClaims, int) {
	parseWithClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		log.Println(err)
		return nil, utils.ErrorToken
	}
	if claims, ok := parseWithClaims.Claims.(*MyClaims); ok && parseWithClaims.Valid {
		return claims, utils.SUCCESS
	} else {
		return nil, utils.ErrorToken
	}
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenHeader := context.Request.Header.Get("Authorization")
		// token为空
		if tokenHeader == "" {
			context.JSON(http.StatusOK, gin.H{
				"code":    utils.SUCCESS,
				"message": "token不存在",
			})
			context.Abort()
			return
		}
		n := strings.SplitN(tokenHeader, " ", 2)
		// [Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InJvb3QxIiwiZXhwIjoxNjYzNTQ3MTMzLCJpc3MiOiJkZW1vIn0.XwpQ5yXR9cVRmNCr6BPET8zopg89bt4J_kiR15UyOvI]
		if len(n) != 2 && n[0] != "Bearer" {
			context.JSON(http.StatusOK, gin.H{
				"code":    utils.ErrorToken,
				"message": "token格式异常",
			})
			context.Abort()
			return
		}
		token, i := checkToken(n[1])
		if i == utils.ErrorToken {
			context.JSON(http.StatusOK, gin.H{
				"code":    utils.ErrorToken,
				"message": "token不正确",
			})
			context.Abort()
			return
		}
		context.Set("username", token.Username)
		context.Next()
	}
}
