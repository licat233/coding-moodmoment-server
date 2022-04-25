package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	standardFieldNames          = builder.RawFieldNames(&Standard{})
	standardRows                = strings.Join(standardFieldNames, ",")
	standardRowsExpectAutoSet   = strings.Join(stringx.Remove(standardFieldNames, "`id`", "`create_at`", "`update_at`"), ",")
	standardRowsWithPlaceHolder = strings.Join(stringx.Remove(standardFieldNames, "`id`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	StandardModel interface {
		Insert(ctx context.Context, data *Standard) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Standard, error)
		Update(ctx context.Context, data *Standard) error
		Delete(ctx context.Context, id int64) error
	}

	defaultStandardModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Standard struct {
		Id       int64     `db:"id"`        // 主键
		Channel  string    `db:"channel"`   // 渠道，如line，popin
		Account  string    `db:"account"`   // 账号,如鸡皮肤账号，祛斑账号
		Imp      int64     `db:"imp"`       // 展现
		Cost     float64   `db:"cost"`      // 消耗
		Click    int64     `db:"click"`     // 点击
		Ctr      float64   `db:"ctr"`       // 点击率
		Cv       int64     `db:"cv"`        // 转化
		Cvr      float64   `db:"cvr"`       // 转化率
		Cpa      float64   `db:"cpa"`       // 转化成本
		CreateAt time.Time `db:"create_at"` // 修改日期
		UpdateAt time.Time `db:"update_at"` // 更新日期
	}
)

func NewStandardModel(conn sqlx.SqlConn) StandardModel {
	return &defaultStandardModel{
		conn:  conn,
		table: "`standard`",
	}
}

func (m *defaultStandardModel) Insert(ctx context.Context, data *Standard) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, standardRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Channel, data.Account, data.Imp, data.Cost, data.Click, data.Ctr, data.Cv, data.Cvr, data.Cpa, data.CreateAt, data.UpdateAt)
	return ret, err
}

func (m *defaultStandardModel) FindOne(ctx context.Context, id int64) (*Standard, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", standardRows, m.table)
	var resp Standard
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultStandardModel) Update(ctx context.Context, data *Standard) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, standardRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Channel, data.Account, data.Imp, data.Cost, data.Click, data.Ctr, data.Cv, data.Cvr, data.Cpa, data.CreateAt, data.UpdateAt, data.Id)
	return err
}

func (m *defaultStandardModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}
