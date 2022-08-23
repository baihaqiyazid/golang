package services

import (
	"context"
	"go-restful-api/model/web"
)

type CategoryService interface{
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	GetById(ctx context.Context, categoryId int) web.CategoryResponse
	GetAll(ctx context.Context) []web.CategoryResponse
}