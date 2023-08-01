package dao

import "gorm.io/gorm"

// dao属于请求中间层，route.list -> dao.list -> model.list

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}
