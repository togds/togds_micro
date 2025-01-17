// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.3

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tUserFieldNames          = builder.RawFieldNames(&TUser{})
	tUserRows                = strings.Join(tUserFieldNames, ",")
	tUserRowsExpectAutoSet   = strings.Join(stringx.Remove(tUserFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tUserRowsWithPlaceHolder = strings.Join(stringx.Remove(tUserFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	tUserModel interface {
		Insert(ctx context.Context, data *TUser) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*TUser, error)
		FindOneUsername(ctx context.Context, Username string) (*TUser, error)
		Update(ctx context.Context, data *TUser) error
		Delete(ctx context.Context, userId int64) error
	}

	defaultTUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TUser struct {
		UserId     int64          `db:"user_id"`
		Username   string         `db:"username"`
		Password   string         `db:"password"`
		Email      sql.NullString `db:"email"`
		Phone      sql.NullString `db:"phone"`
		// Createtime time.Time      `db:"createtime"`
		// Updatetime time.Time      `db:"updatetime"`
		Createtime time.Time      `db:"createtime"`
		Updatetime time.Time      `db:"updatetime"`
		Status     int64          `db:"status"` // 1为激活0为未激活
	}
)

func newTUserModel(conn sqlx.SqlConn) *defaultTUserModel {
	return &defaultTUserModel{
		conn:  conn,
		table: "`t_user`",
	}
}

func (m *defaultTUserModel) Delete(ctx context.Context, userId int64) error {
	query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, userId)
	return err
}

func (m *defaultTUserModel) FindOne(ctx context.Context, userId int64) (*TUser, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", tUserRows, m.table)
	var resp TUser
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTUserModel) FindOneUsername(ctx context.Context, username string) (*TUser, error) {
	query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", tUserRows, m.table)
	var resp TUser
	err := m.conn.QueryRowCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTUserModel) Insert(ctx context.Context, data *TUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, tUserRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Username, data.Password, data.Email, data.Phone, data.Createtime, data.Updatetime, data.Status)
	return ret, err
}

func (m *defaultTUserModel) Update(ctx context.Context, data *TUser) error {
	query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, tUserRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Username, data.Password, data.Email, data.Phone, data.Createtime, data.Updatetime, data.Status, data.UserId)
	return err
}

func (m *defaultTUserModel) tableName() string {
	return m.table
}
