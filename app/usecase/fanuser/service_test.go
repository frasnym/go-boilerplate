package fanuser

import (
	"testing"

	"github.com/frasnym/go-boilerplate/app/entity"
	"github.com/frasnym/go-boilerplate/app/infrastructure/gorm/repository/repository_mock"
	"github.com/stretchr/testify/assert"
)

func TestValidateEmptyFanuser(t *testing.T) {
	testService := NewFanuserService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "fanuser is empty", err.Error())
}

func TestValidateEmptyFanuserName(t *testing.T) {
	testService := NewFanuserService(nil)

	fanuser := entity.Fanuser{}
	err := testService.Validate(&fanuser)

	assert.NotNil(t, err)
	assert.Equal(t, "name cannot be empty", err.Error())
}

func TestCreate(t *testing.T) {
	mockRepo := new(repository_mock.FanuserRepositoryMock)

	fanuser := entity.Fanuser{Name: "Name"}
	mockRepo.Mock.On("Create").Return(&fanuser, nil)

	testService := NewFanuserService(mockRepo)
	result, err := testService.Create(&fanuser)

	// Mock assertion: Behavioral
	mockRepo.Mock.AssertExpectations(t)

	// Data assertion
	assert.Equal(t, fanuser.Name, result.Name)
	assert.Nil(t, err)
}
