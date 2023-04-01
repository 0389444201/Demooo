package models

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserTableModel = (*customUserTableModel)(nil)

type (
	// UserTableModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserTableModel.
	UserTableModel interface {
		userTableModel
		FindByName(ctx context.Context, name string) (*UserTable, error)
		DeleteByName(ctx context.Context, name string) error
		UpdateByName(ctx context.Context, data *UserTable) error
		GetAll(ctx context.Context) ([]Table, error)
	}

	customUserTableModel struct {
		*defaultUserTableModel
	}
)

// NewUserTableModel returns a model for the database table.
func NewUserTableModel(conn sqlx.SqlConn) UserTableModel {
	return &customUserTableModel{
		defaultUserTableModel: newUserTableModel(conn),
	}
}

// /////
type Table struct {
	Id     int64  `db:"id"`
	Email  string `db:"email"`
	Name   string `db:"name"`
	Gender string `db:"gender"`
}

func (m *defaultUserTableModel) FindByName(ctx context.Context, name string) (*UserTable, error) {
	query := fmt.Sprintf("select %v from %v where `name` = ? limit 1", userTableRows, m.table)
	var resp UserTable
	err := m.conn.QueryRowCtx(ctx, &resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserTableModel) DeleteByName(ctx context.Context, name string) error {
	query := fmt.Sprintf("delete from %s where `name` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, name)
	return err
}

func (m *defaultUserTableModel) UpdateByName(ctx context.Context, data *UserTable) error {
	query := fmt.Sprintf("update %s set %s where `name` = ?", m.table, userTableRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Email, data.Name, data.Gender, data.Password, data.Id)
	return err
}

func (m *defaultUserTableModel) GetAll(ctx context.Context) ([]Table, error) {
	var resp []Table
	query := fmt.Sprintf("SELECT * FROM %v ;", m.table)
	m.conn.QueryRowsCtx(ctx, &resp, query)
	logx.Info(resp)
	return resp, nil
}
