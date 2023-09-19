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
	goodsImageFieldNames          = builder.RawFieldNames(&GoodsImage{})
	goodsImageRows                = strings.Join(goodsImageFieldNames, ",")
	goodsImageRowsExpectAutoSet   = strings.Join(stringx.Remove(goodsImageFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	goodsImageRowsWithPlaceHolder = strings.Join(stringx.Remove(goodsImageFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	goodsImageModel interface {
		Insert(ctx context.Context, data *GoodsImage) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*GoodsImage, error)
		Update(ctx context.Context, data *GoodsImage) error
		Delete(ctx context.Context, id string) error
	}

	defaultGoodsImageModel struct {
		conn  sqlx.SqlConn
		table string
	}

	GoodsImage struct {
		Id         string `db:"id"`          // 主键id
		GoodsId    int64  `db:"goods_id"`    // 商品id
		ImageId    int64  `db:"image_id"`    // 图片id(关联图片记录表)
		CreateTime int64  `db:"create_time"` // 创建时间
	}
)

func newGoodsImageModel(conn sqlx.SqlConn) *defaultGoodsImageModel {
	return &defaultGoodsImageModel{
		conn:  conn,
		table: "`goods_image`",
	}
}

func (m *defaultGoodsImageModel) withSession(session sqlx.Session) *defaultGoodsImageModel {
	return &defaultGoodsImageModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`goods_image`",
	}
}

func (m *defaultGoodsImageModel) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultGoodsImageModel) FindOne(ctx context.Context, id string) (*GoodsImage, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", goodsImageRows, m.table)
	var resp GoodsImage
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

func (m *defaultGoodsImageModel) Insert(ctx context.Context, data *GoodsImage) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, goodsImageRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.GoodsId, data.ImageId)
	return ret, err
}

func (m *defaultGoodsImageModel) Update(ctx context.Context, data *GoodsImage) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, goodsImageRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.GoodsId, data.ImageId, data.Id)
	return err
}

func (m *defaultGoodsImageModel) tableName() string {
	return m.table
}
