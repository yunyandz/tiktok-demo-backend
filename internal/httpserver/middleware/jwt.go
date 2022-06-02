package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-backend/internal/jwtx"
	"github.com/yunyandz/tiktok-demo-backend/internal/service"
	"go.uber.org/zap"
)

func JWTAuth(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取token
		token := c.Query("token")
		if token == "" || len(token) == 0 {
			token = c.PostForm("token")
			if token == "" || len(token) == 0 {
				rsp := service.Response{
					StatusCode: -1,
					StatusMsg:  "Invalid token",
				}
				logger.Sugar().Errorf("Invalid token: %v", rsp)
				c.JSON(http.StatusUnauthorized, rsp)
				c.Abort()
				return
			}
		}

		// 解析token

		parsedToken, err := jwtx.ParseUserClaims(token)
		if err != nil {
			rsp := service.Response{
				StatusCode: -1,
				StatusMsg:  "Parse token failed",
			}
			logger.Sugar().Errorf("Parse token failed: %v", rsp)
			c.JSON(http.StatusNonAuthoritativeInfo, rsp)
			c.Abort()
			return
		}

		suid := c.Query("user_id")
		if suid != "" || len(suid) != 0 {
			uid, err := strconv.ParseUint(suid, 10, 64)
			if err != nil {
				rsp := service.Response{
					StatusCode: -1,
					StatusMsg:  "Invalid user_id",
				}
				logger.Sugar().Errorf("Invalid user_id: %v", rsp)
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
				logger.Sugar().Errorf("Invalid user_id: UserID != uid,%v", rsp)
				c.JSON(http.StatusNonAuthoritativeInfo, rsp)
				c.Abort()
				return
			}
		}

		c.Set("claims", parsedToken)
		c.Next()
	}
}
