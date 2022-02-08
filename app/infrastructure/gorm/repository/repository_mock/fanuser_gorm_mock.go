package repository_mock

import (
	"github.com/frasnym/go-boilerplate/app/entity"
	"github.com/stretchr/testify/mock"
)

type FanuserRepositoryMock struct {
	Mock mock.Mock
}

func (repo *FanuserRepositoryMock) Create(fanuser *entity.Fanuser) (*entity.Fanuser, error) {
	args := repo.Mock.Called()
	result := args.Get(0)
	return result.(*entity.Fanuser), args.Error(1)
}
