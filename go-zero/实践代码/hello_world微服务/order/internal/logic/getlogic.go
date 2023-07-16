package logic

import (
	"context"
	"user/types/user"

	"order/internal/svc"
	"order/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.OrderReq) (resp *types.OrderReply, err error) {
	// todo: add your logic here and delete this line
	user,err:=l.svcCtx.UserRpc.GetUser(l.ctx,&user.IdRequest{Id: req.Id})
	if err!=nil{
		return nil, err
	}
	return &types.OrderReply{
		Id:     user.Id,
		Name:   user.Name,
		Gender: user.Gender,
	}, nil
}
