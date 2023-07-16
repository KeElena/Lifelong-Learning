package logic

import (
	"business/internal/svc"
	"business/internal/types"
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIngoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIngoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIngoLogic {
	return &GetIngoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIngoLogic) GetIngo() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	id, _ := l.ctx.Value("uid").(json.Number).Int64()
	return &types.Response{Uid: id, Name: l.ctx.Value("name").(string)}, nil
}
