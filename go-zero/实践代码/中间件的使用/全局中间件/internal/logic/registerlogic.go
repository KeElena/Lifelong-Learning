package logic

import (
	"context"
	"log"

	"middlewares/internal/svc"
	"middlewares/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	log.Println("处理注册业务")
	return
}
