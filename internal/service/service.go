package service

import (
	"context"
	"github.com/nico612/blog-service/global"
	"github.com/nico612/blog-service/internal/dao"
)

// 处理业务逻辑

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	return Service{ctx: ctx, dao: dao.New(global.DBEngine)}
}
