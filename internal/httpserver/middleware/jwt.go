package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/jwtx"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
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

		suid := c.Query("user_id")
		uid, err := strconv.ParseUint(suid, 10, 64)
		if err != nil {
			rsp := service.Response{
				StatusCode: -1,
				StatusMsg:  "Invalid user_id",
			}
			c.JSON(http.StatusNonAuthoritativeInfo, rsp)
			c.Abort()
			return
		}

		// 检查userid是否正确
		if parsedToken.UserID != uid {
			rsp := service.Response{
				StatusCode: -1,
				StatusMsg:  "Invalid user_id",
			}
			c.JSON(http.StatusNonAuthoritativeInfo, rsp)
			c.Abort()
			return
		}

		c.Set("claims", parsedToken)
		c.Next()
	}
}
