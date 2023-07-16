package logic

import (
	"context"
	"database/sql"
	"user/model"

	"rpc-common/types/user"
	"user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.UserInfo) (*user.UserInfo, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.Conn.Insert(l.ctx, &model.Users{Name: sql.NullString{String: in.Name, Valid: true}, Gender: sql.NullString{String: in.Gender, Valid: true}})
	if err != nil {
		return nil, err
	}
	id, _ := res.LastInsertId()
	return &user.UserInfo{Id: id, Name: in.Name, Gender: in.Gender}, nil
}
