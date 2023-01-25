package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SnapsModel = (*customSnapsModel)(nil)

type (
	// SnapsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSnapsModel.
	SnapsModel interface {
		snapsModel
	}

	customSnapsModel struct {
		*defaultSnapsModel
	}
)

// NewSnapsModel returns a model for the database table.
func NewSnapsModel(conn sqlx.SqlConn) SnapsModel {
	return &customSnapsModel{
		defaultSnapsModel: newSnapsModel(conn),
	}
}
