package repository

import (
	"go-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(Id string) *entity.Category {
	arguments := repository.Mock.Called(Id)
	if arguments.Get(0) == nil {
		return nil
	}else {
		category := arguments.Get(0).(entity.Category)
		return &category
	}
}