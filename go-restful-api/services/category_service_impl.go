package services

import (
	"context"
	"database/sql"
	"go-restful-api/helper"
	"go-restful-api/model/entity"
	"go-restful-api/model/web"
	"go-restful-api/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	repository repository.CategoryRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {

	err :=  service.Validate.Struct(request) //Request Validation 
	helper.Panic(err)
	
	tx, err := service.DB.Begin()
	helper.Panic(err)
	defer helper.CommitOrRollback(tx)

	category := entity.Category{
		Name: request.Name,
	}

	category = service.repository.Create(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	
	err :=  service.Validate.Struct(request) //Request Validation 
	helper.Panic(err)
	
	tx, err := service.DB.Begin()
	helper.Panic(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.repository.GetById(ctx, tx, request.Id) //Check Category exist
	helper.Panic(err)

	category.Name = request.Name

	category = service.repository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.Panic(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.repository.GetById(ctx, tx, categoryId) //Check Category exist
	helper.Panic(err)

	service.repository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) GetById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.Panic(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.repository.GetById(ctx, tx, categoryId) //Check Category exist
	helper.Panic(err)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) GetAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.Panic(err)
	defer helper.CommitOrRollback(tx)

	categories := service.repository.GetAll(ctx, tx) //Check Category exist

	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, helper.ToCategoryResponse(category))
	}
	return categoryResponses
}
