package user

import (
	"github.com/arifbugaresa/go-starter/utils/common"
	"github.com/arifbugaresa/go-starter/utils/constant/dialect"
	. "github.com/arifbugaresa/go-starter/utils/constant/table"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	SignUpUser(ctx *gin.Context, model RegisterModel) (err error)
	RegisterUserTx(ctx *gin.Context, model RegisterModel, tx *sqlx.Tx) (err error)

	GetUserByEmailOrUsername(ctx *gin.Context, req LoginModel) (record LoginModel, err error)
	UpdateUserById(ctx *gin.Context, model UpdateProfileModel) (err error)

	GetRoleByRoleName(ctx *gin.Context, roleName string) (record GetRoleModel, err error)
}

type UserRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) SignUpUser(ctx *gin.Context, model RegisterModel) (err error) {
	return common.WrapperInsert(ctx, Users, model, r.db)
}

func (r *UserRepository) RegisterUserTx(ctx *gin.Context, model RegisterModel, tx *sqlx.Tx) (err error) {
	return common.WrapperInsertTx(ctx, Users, model, tx)
}

func (r *UserRepository) GetRoleByRoleName(ctx *gin.Context, roleName string) (record GetRoleModel, err error) {
	conn := goqu.New(dialect.Postgres, r.db)
	dataset := conn.
		Select(
			Roles.Col("role_id"),
			Roles.Col("role_name"),
		).
		From(Roles).
		Where(
			Roles.Col("role_name").Eq(roleName),
		)

	_, err = dataset.ScanStructContext(ctx, &record)
	if err != nil {
		return
	}

	return
}

func (r *UserRepository) GetUserByEmailOrUsername(ctx *gin.Context, req LoginModel) (record LoginModel, err error) {
	conn := goqu.New(dialect.Postgres, r.db)
	dataset := conn.
		Select(
			Roles.Col("role_name").As("role"),
			Roles.Col("role_id"),
			Users.Col("id"),
			Users.Col("username"),
			Users.Col("email"),
			Users.Col("full_name"),
			Users.Col("password"),
			Users.Col("photo"),
		).
		From(Users).
		Join(Roles,
			goqu.On(
				Roles.Col("role_id").Eq(Users.Col("role_id")),
			),
		).
		Where(
			goqu.Or(
				Users.Col("email").Eq(req.UserName),
				Users.Col("username").Eq(req.UserName),
			),
		)

	_, err = dataset.ScanStructContext(ctx, &record)
	if err != nil {
		return
	}

	return
}

func (r *UserRepository) UpdateUserById(ctx *gin.Context, model UpdateProfileModel) (err error) {
	conn := goqu.New(dialect.Postgres, r.db)
	_, err = conn.
		Update(Users).
		Set(model).
		Where(
			Users.Col("id").Eq(model.Id),
		).
		Executor().
		ExecContext(ctx)
	if err != nil {
		return err
	}

	return
}
