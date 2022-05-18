package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/jwtx"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
	"net/http"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取token
		token := c.Query("token")
		if token == "" || len(token) == 0 {
			rsp := service.Response{
				StatusCode: -1,
				StatusMsg:  "Invalid token",
			}
			c.JSON(http.StatusUnauthorized, rsp)
			c.Abort()
			return
		}

		// 解析token

		parsedToken, err := jwtx.ParseUserClaims(token)
		if err != nil {
			rsp := service.Response{
				StatusCode: -1,
				StatusMsg:  "Parse token failed",
			}
			c.JSON(http.StatusNonAuthoritativeInfo, rsp)
			c.Abort()
			return
		}

		c.Set("claims", parsedToken)
		c.Next()
	}
}
