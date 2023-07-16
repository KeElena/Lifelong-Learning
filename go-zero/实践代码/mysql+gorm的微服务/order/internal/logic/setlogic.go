package logic

import (
	"context"
	"user/types/user"

	"order/internal/svc"
	"order/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetLogic {
	return &SetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetLogic) Set(req *types.OrderReq) (resp *types.HandleInfo, err error) {
	// todo: add your logic here and delete this line
	info, _ := l.svcCtx.UserRpc.CreateUser(context.Background(), &user.IdRequest{Id: req.Id})
	return &types.HandleInfo{Information: info.Info}, nil
}
