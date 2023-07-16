package logic

import (
	"context"
	"rpc-common/types/score"

	_ "github.com/dtm-labs/dtmdriver"
	"github.com/zeromicro/go-zero/core/logx"
	"rpc-common/types/user"
	"userapi/internal/svc"
	"userapi/internal/types"
)

var dtmServer = "etcd://127.0.0.1:2379/dtmservice"

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

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	userInfo, err := l.svcCtx.UserRpc.CreateUser(l.ctx, &user.UserInfo{Name: req.Name, Gender: req.Gender})
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.UserScore.SaveUserScore(l.ctx, &score.UserScoreRequest{Uid: userInfo.Id, Score: 10})
	if err != nil {
		return nil, err
	}
	return &types.RegisterResp{Status: "OK", Data: types.Data{Id: userInfo.Id, Name: userInfo.Name, Gender: userInfo.Gender}}, nil
}
