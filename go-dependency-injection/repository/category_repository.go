package repository

import (
	"context"
	"database/sql"
	"go-restful-api/model/entity"
)

type CategoryRepository interface{
	Create(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	Delete(ctx context.Context, tx *sql.Tx, category entity.Category) 
	GetById(ctx context.Context, tx *sql.Tx, categoryId int) (entity.Category, error)
	GetAll(ctx context.Context, tx *sql.Tx) []entity.Category
}