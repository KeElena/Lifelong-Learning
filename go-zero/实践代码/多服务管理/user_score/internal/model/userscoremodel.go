package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserScoreModel = (*customUserScoreModel)(nil)

type (
	// UserScoreModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserScoreModel.
	UserScoreModel interface {
		userScoreModel
	}

	customUserScoreModel struct {
		*defaultUserScoreModel
	}
)

// NewUserScoreModel returns a model for the database table.
func NewUserScoreModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserScoreModel {
	return &customUserScoreModel{
		defaultUserScoreModel: newUserScoreModel(conn, c, opts...),
	}
}
