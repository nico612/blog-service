package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nico612/blog-service/global"
	"github.com/nico612/blog-service/internal/service"
	"github.com/nico612/blog-service/pkg/app"
	"github.com/nico612/blog-service/pkg/convert"
	"github.com/nico612/blog-service/pkg/errcode"
	"github.com/nico612/blog-service/pkg/upload"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

// UploadFile handle
func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	// 通过c.Request.FormFile方法来获取上传的文件，“file”：是表单中文件上传字段的名称
	// 返回值：
	// file: 上传文件的 multipart.File 类型的对象，可以通过这个对象来读取上传的文件内容
	// fileHeader: 表示上传文件的文件头信息的"multipart.FileHeader"类型的对象，它包含有关上传文件的一些元数据， 例如：文件名、文件大小等信息
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	// 获取文件类型
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.uploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})

}
