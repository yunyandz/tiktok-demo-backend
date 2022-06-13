package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/jwtx"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
	"go.uber.org/zap"
)

func JWTAuth(logger *zap.Logger, strict bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		token := c.Query("token")
		if token == "" || len(token) == 0 {
			token = c.PostForm("token")
			if token == "" || len(token) == 0 {
				if strict {
					rsp := service.Response{
						StatusCode: -1,
						StatusMsg:  "Invalid token",
					}
					logger.Sugar().Debugf("Invalid token: %v", rsp)
					c.JSON(http.StatusUnauthorized, rsp)
					c.Abort()
				} else {
					c.Next()
				}
				return
			}
		}

		// 解析token

		parsedToken, err := jwtx.ParseUserClaims(token)
		if err != nil {
			if strict {
				rsp := service.Response{
					StatusCode: -1,
					StatusMsg:  "Parse token failed",
				}
				logger.Sugar().Debugf("Parse token failed: %v", rsp)
				c.JSON(http.StatusNonAuthoritativeInfo, rsp)
				c.Abort()
			} else {
				c.Next()
			}
			return
		}

		c.Set("claims", parsedToken)
		c.Next()
	}
}
