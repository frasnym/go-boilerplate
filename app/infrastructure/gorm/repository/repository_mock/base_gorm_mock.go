package repository_mock

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type BaseRepositoryMock struct {
	Mock mock.Mock
}

func (b *BaseRepositoryMock) GetDB() *gorm.DB {
	return nil
}

func (b *BaseRepositoryMock) BeginTx() {

}

func (b *BaseRepositoryMock) CommitTx() {

}

func (b *BaseRepositoryMock) RollbackTx() {

}
