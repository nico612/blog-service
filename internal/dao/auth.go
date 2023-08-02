package dao

import "github.com/nico612/blog-service/internal/model"

// GetAuth 获取认证信息的方法
func (d Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{
		AppKey:    appKey,
		AppSecret: appSecret,
	}
	return auth.Get(d.engine)
}
