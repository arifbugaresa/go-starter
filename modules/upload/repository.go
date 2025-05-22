package upload

import (
	"github.com/jmoiron/sqlx"
)

type Repository interface{}

type UploadRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *UploadRepository {
	return &UploadRepository{db: db}
}
