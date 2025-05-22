package common

import (
	"context"
	"github.com/arifbugaresa/go-starter/utils/common/message"
	"github.com/arifbugaresa/go-starter/utils/constant/dialect"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"strings"
)

func BuildDatasetGetListWithParams(dataset *goqu.SelectDataset, param DefaultListRequest) *goqu.SelectDataset {
	if param.Search.Field != "" && param.Search.Value != "" {
		if strings.Contains(param.Search.Field, "id") || strings.Contains(param.Search.Field, "is") {
			dataset = dataset.Where(
				goqu.I(param.Search.Field).Eq(param.Search.Value),
			)
		} else {
			dataset = dataset.Where(
				goqu.I(param.Search.Field).ILike("%" + param.Search.Value + "%"),
			)
		}
	}

	if param.Sort.Field != "" && param.Sort.Order != "" {
		if param.Sort.Order == "asc" {
			dataset = dataset.Order(goqu.I(param.Sort.Field).Asc())
		} else {
			dataset = dataset.Order(goqu.I(param.Sort.Field).Desc())
		}
	}

	if param.Page != 0 && param.Limit != 0 {
		offset := (param.Page - 1) * param.Limit
		dataset = dataset.Limit(uint(param.Limit)).Offset(uint(offset))
	}

	return dataset
}

func WrapperInsert(ctx *gin.Context, tableName exp.IdentifierExpression, model interface{}, db *sqlx.DB) (err error) {
	conn := goqu.New(dialect.Postgres, db)
	dataset := conn.Insert(tableName).Rows(model)

	_, err = dataset.Executor().ExecContext(ctx)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return message.DuplicateData
		}
		return err
	}

	return nil
}

func WrapperInsertTx(ctx *gin.Context, tableName exp.IdentifierExpression, model interface{}, tx *sqlx.Tx) (err error) {
	conn := goqu.NewTx(dialect.Postgres, tx)
	dataset := conn.Insert(tableName).Rows(model)

	_, err = dataset.Executor().ExecContext(ctx)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return message.DuplicateData
		}
		return err
	}

	return nil
}

func WrapperInsertTxReturningId(ctx *gin.Context, tableName exp.IdentifierExpression, model interface{}, tx *sqlx.Tx) (id int64, err error) {
	conn := goqu.NewTx(dialect.Postgres, tx)
	dataset := conn.Insert(tableName).Rows(model).Returning(goqu.C("id")).Executor()

	if _, err = dataset.ScanVal(&id); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			err = message.DuplicateData
			return
		}

		err = message.InternalServerError
		return
	}

	return
}

func WrapperSelect[T any](ctx *gin.Context, tableName exp.IdentifierExpression, conditions []exp.Expression, db *sqlx.DB) (output T, err error) {
	conn := goqu.New(dialect.Postgres, db)
	_, err = conn.
		From(tableName).
		Where(conditions...).
		ScanStruct(&output)

	return
}

func WrapperUpdate(ctx *gin.Context, tableName exp.IdentifierExpression, model interface{}, conditions []exp.Expression, db *sqlx.DB) (err error) {
	conn := goqu.New(dialect.Postgres, db)
	dataset := conn.
		Update(tableName).
		Set(model).
		Where(conditions...)

	_, err = dataset.Executor().ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func WrapperUpdateTx(ctx *gin.Context, tableName exp.IdentifierExpression, model interface{}, conditions []exp.Expression, tx *sqlx.Tx) (err error) {
	conn := goqu.NewTx(dialect.Postgres, tx)
	dataset := conn.
		Update(tableName).
		Set(model).
		Where(conditions...)

	_, err = dataset.Executor().ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func WrapperDelete(ctx *gin.Context, tableName exp.IdentifierExpression, conditions []exp.Expression, db *sqlx.DB) (err error) {
	conn := goqu.New(dialect.Postgres, db)
	dataset := conn.
		Delete(tableName).
		Where(conditions...)

	_, err = dataset.Executor().ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func WrapperTx(ctx context.Context, db *sqlx.DB, fn func(tx *sqlx.Tx) error) error {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return message.InternalServerError
	}
	defer tx.Rollback()

	if err = fn(tx); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return message.InternalServerError
	}

	return nil
}
