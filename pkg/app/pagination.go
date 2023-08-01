package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nico612/blog-service/global"
	"github.com/nico612/blog-service/pkg/convert"
)

// 处理分页查询

// GetPage 获取请求参数中的page值
func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

// GetPageSize 获取请求参数中的page_size值
func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.DefaultPageSize
	}
	return pageSize
}

// GetPageOffset 计算数据库查询的Offset
func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
