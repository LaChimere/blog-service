package service

import (
	"context"

	"github.com/lachimere/blog-service/global"
	"github.com/lachimere/blog-service/internal/dao"
	otgorm "github.com/lachimere/opentracing-gorm"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
}
