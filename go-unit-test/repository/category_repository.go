package repository

import "go-unit-test/entity"

type CategoryRepository interface{
	FindById(Id string) *entity.Category
}