package common

import (
	"github.com/arifbugaresa/go-starter/middlewares"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type DefaultTable struct {
	Id        int64  `db:"id" goqu:"skipinsert,skipupdate"`
	CreatedBy int64  `db:"created_by" goqu:"skipupdate"`
	CreatedAt string `db:"created_at" goqu:"skipupdate"`
	UpdatedBy int64  `db:"updated_by"`
	UpdatedAt string `db:"updated_at"`
}

func (d DefaultTable) GetDefaultTable(ctx *gin.Context) DefaultTable {
	var (
		timeNow = time.Now().Format("2006-01-02 15:04:05")
	)

	auth, err := middlewares.GetSession(ctx)
	if err != nil {
		return DefaultTable{}
	}

	return DefaultTable{
		CreatedBy: auth.Id,
		UpdatedBy: auth.Id,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
}

type DefaultListRequest struct {
	Page   int64
	Limit  int64
	Search Search
	Sort   Sort
}

type Sort struct {
	Field string
	Order string
}

type Search struct {
	Field string
	Value string
}

func (d DefaultListRequest) GetParamRequest(ctx *gin.Context) DefaultListRequest {
	page, _ := strconv.ParseInt(ctx.Query("page"), 10, 64)
	limit, _ := strconv.ParseInt(ctx.Query("limit"), 10, 64)

	return DefaultListRequest{
		Page:  page,
		Limit: limit,
		Sort: Sort{
			Field: ctx.Query("sort_field"),
			Order: ctx.Query("sort_order"),
		},
		Search: Search{
			Field: ctx.Query("search_field"),
			Value: ctx.Query("search_value"),
		},
	}
}
