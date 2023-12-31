package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nico612/blog-service/global"
	"github.com/nico612/blog-service/internal/middleware"
	"github.com/nico612/blog-service/internal/routers/api"
	v1 "github.com/nico612/blog-service/internal/routers/api/v1"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	// 路由中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.JWT())
	//r.Use(middleware.Translations())

	// 设置路由组
	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := api.NewUpload()

	// 文件上传
	r.POST("/upload/file", upload.UploadFile)
	// 提供静态资源的访问
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	// 获取认证
	r.POST("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)        //更新整条记录
		apiv1.PATCH("/tags:id/state", tag.Update) // 更新某个记录的某个字段
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("articles", article.List)
	}

	return r
}
