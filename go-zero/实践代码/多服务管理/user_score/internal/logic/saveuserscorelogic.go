package logic

import (
	"context"
	"database/sql"
	"rpc-common/types/score"
	"user_score/internal/model"

	"user_score/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveUserScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveUserScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveUserScoreLogic {
	return &SaveUserScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveUserScoreLogic) SaveUserScore(in *score.UserScoreRequest) (*score.UserScoreResponse, error) {
	// todo: add your logic here and delete this line
	userScore := &model.UserScore{
		UserId: in.Uid,
		Score:  sql.NullInt64{Int64: in.Score, Valid: true},
	}
	_, err := l.svcCtx.Conn.Insert(context.Background(), userScore)
	if err != nil {
		return nil, err
	}
	return &score.UserScoreResponse{
		Uid:   userScore.UserId,
		Score: userScore.Score.Int64,
	}, nil
}
