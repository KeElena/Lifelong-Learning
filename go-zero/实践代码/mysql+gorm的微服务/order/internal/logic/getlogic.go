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
	userInfo, err := l.svcCtx.UserRpc.GetUser(context.Background(), &user.IdRequest{Id: req.Id})
	if err != nil {
		return &types.OrderReply{
			Id:     -1,
			Name:   "null",
			Gender: "null",
		}, nil
	}
	return &types.OrderReply{
		Id:     userInfo.Id,
		Name:   userInfo.Name,
		Gender: userInfo.Gender,
	}, nil
}
