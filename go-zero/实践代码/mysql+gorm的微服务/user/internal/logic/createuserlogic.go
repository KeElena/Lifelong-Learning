package logic

import (
	"context"
	"user/internal/curd"

	"user/internal/svc"
	"user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.IdRequest) (*user.Info, error) {
	// todo: add your logic here and delete this line
	err := curd.Insert(l.svcCtx.MysqlDB, in.Id)
	if err != nil {
		return &user.Info{Info: err.Error()}, nil
	}
	return &user.Info{Info: "ok"}, nil
}
