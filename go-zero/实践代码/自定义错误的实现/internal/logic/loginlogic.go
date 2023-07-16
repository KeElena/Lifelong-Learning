package logic

import (
	"context"
	"errors/internal/errorx"

	"errors/internal/svc"
	"errors/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return nil, errorx.ParamsError
}
