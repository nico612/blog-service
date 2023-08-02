package service

import (
	"errors"
	"github.com/nico612/blog-service/global"
	"github.com/nico612/blog-service/pkg/upload"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string //文件名
	AccessUrl string //访问路径
}

func (svc *Service) UploadFile(filetype upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)

	// 检查配置是否支持该文件类型
	if !upload.CheckContainExt(filetype, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	if upload.CheckMaxSize(filetype, file) {
		return nil, errors.New("exceeded maximum file limit")
	}

	uploadSavePath := upload.GetSavePath()

	// 检查文件路径是否存在
	if !upload.CheckSavePath(uploadSavePath) {
		// 创建文件路径rwx-rwx-rwx
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, err
		}
	}

	// 检查文件目录是否有权限
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions")
	}

	dst := uploadSavePath + "/" + fileName

	// 保存文件
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	// 资源访问路径
	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil

}
