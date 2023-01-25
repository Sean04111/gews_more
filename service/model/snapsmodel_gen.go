// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	snapsFieldNames          = builder.RawFieldNames(&Snaps{})
	snapsRows                = strings.Join(snapsFieldNames, ",")
	snapsRowsExpectAutoSet   = strings.Join(stringx.Remove(snapsFieldNames, "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), ",")
	snapsRowsWithPlaceHolder = strings.Join(stringx.Remove(snapsFieldNames, "`sid`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), "=?,") + "=?"
)

type (
	snapsModel interface {
		Insert(ctx context.Context, data *Snaps) (sql.Result, error)
		FindOne(ctx context.Context, sid int64) (*Snaps, error)
		Update(ctx context.Context, data *Snaps) error
		Delete(ctx context.Context, sid int64) error
	}

	defaultSnapsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Snaps struct {
		Sid     int64          `db:"sid"`
		Speaker string         `db:"speaker"`
		Message string         `db:"message"`
		Date    string         `db:"date"`
		At      sql.NullString `db:"at"`
	}
)

func newSnapsModel(conn sqlx.SqlConn) *defaultSnapsModel {
	return &defaultSnapsModel{
		conn:  conn,
		table: "`snaps`",
	}
}

func (m *defaultSnapsModel) Delete(ctx context.Context, sid int64) error {
	query := fmt.Sprintf("delete from %s where `sid` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, sid)
	return err
}

func (m *defaultSnapsModel) FindOne(ctx context.Context, sid int64) (*Snaps, error) {
	query := fmt.Sprintf("select %s from %s where `sid` = ? limit 1", snapsRows, m.table)
	var resp Snaps
	err := m.conn.QueryRowCtx(ctx, &resp, query, sid)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSnapsModel) Insert(ctx context.Context, data *Snaps) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, snapsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Sid, data.Speaker, data.Message, data.Date, data.At)
	return ret, err
}

func (m *defaultSnapsModel) Update(ctx context.Context, data *Snaps) error {
	query := fmt.Sprintf("update %s set %s where `sid` = ?", m.table, snapsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Speaker, data.Message, data.Date, data.At, data.Sid)
	return err
}

func (m *defaultSnapsModel) tableName() string {
	return m.table
}