package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TUserModel = (*customTUserModel)(nil)

type (
	// TUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTUserModel.
	TUserModel interface {
		tUserModel
		withSession(session sqlx.Session) TUserModel
	}

	customTUserModel struct {
		*defaultTUserModel
	}
)

// NewTUserModel returns a model for the database table.
func NewTUserModel(conn sqlx.SqlConn) TUserModel {
	return &customTUserModel{
		defaultTUserModel: newTUserModel(conn),
	}
}

func (m *customTUserModel) withSession(session sqlx.Session) TUserModel {
	return NewTUserModel(sqlx.NewSqlConnFromSession(session))
}
