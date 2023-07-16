package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"login/internal/svc"
	"login/internal/types"
	"time"

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
	payload := make(jwt.MapClaims)
	payload["uid"] = 1
	payload["name"] = "demo"
	payload["exp"] = time.Now().Unix() + 10
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = payload
	jwtToken, err := token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
	if err != nil {
		return nil, err
	}
	return &types.Response{Token: jwtToken}, nil
}
