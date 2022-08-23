package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-restful-api/helper"
	"go-restful-api/model/entity"
)

type CategoryRepositoryImpl struct{

}

func NewCategoryRepository() CategoryRepository  {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category  {
	query := "INSERT INTO categories(name) VALUES(?)"
	
	result, err := tx.ExecContext(ctx, query, category.Name)
	helper.Panic(err)

	id, err := result.LastInsertId()
	helper.Panic(err)

	category.Id = int(id)
	return category

}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category  {
	query := "UPDATE categories SET name = ? WHERE id = ?"
	
	_, err := tx.ExecContext(ctx, query, category.Name, category.Id)
	helper.Panic(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category entity.Category)  {
	query := "DELETE FROM categories WHERE id = ?"

	_, err := tx.ExecContext(ctx, query, category.Id)
	helper.Panic(err)
}

func (repository *CategoryRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, categoryId int) (entity.Category, error)  {
	query := "SELECT id, name FROM categories WHERE id = ?"

	rows, err := tx.QueryContext(ctx, query, categoryId)
	helper.Panic(err)
	defer rows.Close()

	category := entity.Category{}

	if rows.Next(){
		err := rows.Scan(&category.Id, &category.Name)
		helper.Panic(err)
		return category, nil
	}else {
		return category, errors.New("Category is not found")
	}
}

func (repository *CategoryRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []entity.Category {
	query := "SELECT id, name FROM categories"

	rows, err := tx.QueryContext(ctx, query)
	helper.Panic(err)
	defer rows.Close()


	var categories []entity.Category

	for rows.Next(){
		category := entity.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.Panic(err)
		categories = append(categories, category)
	}

	return categories
}