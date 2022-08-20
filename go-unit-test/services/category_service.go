package services

import (
	"errors"
	"go-unit-test/entity"
	"go-unit-test/repository"
)

type CategoryServie struct{
	Repository repository.CategoryRepository
}

func (service CategoryServie) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil{
		return nil, errors.New("Category not Found")
	}else{
		return category, nil
	}
}