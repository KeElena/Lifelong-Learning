package logic

import (
	"context"
	"rpc-common/types/user"
	"user/internal/svc"

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

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserInfo, error) {
	// todo: add your logic here and delete this line
	u, err := l.svcCtx.Conn.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &user.UserInfo{Id: u.Id, Name: u.Name.String, Gender: u.Gender.String}, nil
}
