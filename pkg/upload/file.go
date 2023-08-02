package upload

import (
	"github.com/nico612/blog-service/global"
	"github.com/nico612/blog-service/pkg/util"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

// TypeImage 从1开始递增
const (
	TypeImage FileType = iota + 1
	TypeExcel
	TypeTxt
)

// GetFileName 获取文件名
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

// GetFileExt 获取文件后缀
func GetFileExt(name string) string {
	return path.Ext(name)
}

// GetSavePath 获取文件保存路径
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

// CheckSavePath 检查文件目录是否存在, true： 存在，false 不存在
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsExist(err)
}

// CheckPermission 检查文件目录权限是否足够
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)      //获取文件信息
	return os.IsPermission(err) //判断给定的错误是否表示 "权限错误"（Permission Error）。
}

// CheckMaxSize 检查文件大小是否超出最大大小限制。 true: 超出限制
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := io.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}

	return false

}

// CheckContainExt 检查文件后缀是否包含在约定的后缀配置项中
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}

	return false
}

// CreateSavePath 创建保存路径
func CreateSavePath(dst string, perm os.FileMode) error {

	//若涉及的目录均已存在，则不会进行任何操作，直接返回 nil。
	return os.MkdirAll(dst, perm)
}

// SaveFile 保存文件，dst文件路径
func SaveFile(file *multipart.FileHeader, dst string) error {

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst) //创建文件并返回对应的文件对象和一个错误
	if err != nil {
		return err
	}
	defer out.Close()

	// 将src内容复制到out文件中
	_, err = io.Copy(out, src)
	return err
}
