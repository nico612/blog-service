package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nico612/blog-service/global"
	"github.com/nico612/blog-service/internal/service"
	"github.com/nico612/blog-service/pkg/app"
	"github.com/nico612/blog-service/pkg/errcode"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	vaild, errs := app.BindAndValid(c, &param)
	if !vaild {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())

	// 检查app_key 和 app_secret
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
