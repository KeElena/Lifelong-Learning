package logic

import (
	"context"
	"user/internal/curd"

	"user/internal/svc"
	"user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line
	userInfo, err := curd.Query(l.svcCtx.MysqlDB, in.Id)
	if err != nil {
		return nil, err
	}
	return &user.UserResponse{Id: userInfo.Id, Name: userInfo.Name, Gender: userInfo.Gender}, nil
}
