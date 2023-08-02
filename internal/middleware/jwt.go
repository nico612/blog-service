package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nico612/blog-service/pkg/app"
	"github.com/nico612/blog-service/pkg/errcode"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)

		// 查询token
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}

		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				//switch err.(*jwt.ClaimsValidator.Validate().Error()).Errors {
				//case jwt.ValidationErrorExpired:
				//	ecode = errcode.UnauthorizedTokenTimeout
				//default:
				//	ecode = errcode.UnauthorizedTokenError
				//}
				ecode = errcode.UnauthorizedTokenError
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		c.Next()

	}
}
