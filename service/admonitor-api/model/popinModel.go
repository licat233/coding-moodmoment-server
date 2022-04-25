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
	popinFieldNames          = builder.RawFieldNames(&Popin{})
	popinRows                = strings.Join(popinFieldNames, ",")
	popinRowsExpectAutoSet   = strings.Join(stringx.Remove(popinFieldNames, "`id`", "`create_at`", "`update_at`"), ",")
	popinRowsWithPlaceHolder = strings.Join(stringx.Remove(popinFieldNames, "`id`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	PopinModel interface {
		Insert(ctx context.Context, data *Popin) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Popin, error)
		Update(ctx context.Context, data *Popin) error
		Delete(ctx context.Context, id int64) error
		FindsByDate(ctx context.Context, startDate, endDate time.Time, account string) ([]*Popin, error)
	}

	defaultPopinModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Popin struct {
		Id         int64     `db:"id"`          // 主键
		Account    string    `db:"account"`     // 广告账号
		Cost       float64   `db:"cost"`        // 消耗额
		Click      int64     `db:"click"`       // 点击数
		Imp        int64     `db:"imp"`         // 展现量
		Ctr        float64   `db:"ctr"`         // 点击率
		Cv         int64     `db:"cv"`          // 转化数
		Cvr        float64   `db:"cvr"`         // 转化率
		Cpa        float64   `db:"cpa"`         // 转化成本
		Status     int64     `db:"status"`      // 状态
		Tp         string    `db:"type"`        // 数据类型：newest最新数据，node节点数据，day全天数据，logic逻辑数据
		Currency   string    `db:"currency"`    // 币种
		RecordDate time.Time `db:"record_date"` // 记录日期
		CreateAt   time.Time `db:"create_at"`   // 创建时间
		UpdateAt   time.Time `db:"update_at"`   // 更新时间
	}
)

func NewPopinModel(conn sqlx.SqlConn) PopinModel {
	return &defaultPopinModel{
		conn:  conn,
		table: "`popin`",
	}
}

func (m *defaultPopinModel) Insert(ctx context.Context, data *Popin) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, popinRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Account, data.Cost, data.Click, data.Imp, data.Ctr, data.Cv, data.Cvr, data.Cpa, data.Status, data.Tp, data.Currency, data.RecordDate, data.CreateAt, data.UpdateAt)
	return ret, err
}

func (m *defaultPopinModel) FindOne(ctx context.Context, id int64) (*Popin, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", popinRows, m.table)
	var resp Popin
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

func (m *defaultPopinModel) Update(ctx context.Context, data *Popin) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, popinRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Account, data.Cost, data.Click, data.Imp, data.Ctr, data.Cv, data.Cvr, data.Cpa, data.Status, data.Tp, data.Currency, data.RecordDate, data.CreateAt, data.UpdateAt, data.Id)
	return err
}

func (m *defaultPopinModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPopinModel) FindsByDate(ctx context.Context, startDate, endDate time.Time, account string) ([]*Popin, error) {
	query := fmt.Sprintf("select %s from %s where account = ? AND  DATE(record_date) between ? AND ? AND type = 'logic'", popinRows, m.table)
	var resp []*Popin
	err := m.conn.QueryRowsCtx(ctx, &resp, query, account, startDate, endDate)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
