package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/frasnym/go-boilerplate/app/entity"
	gorm "github.com/frasnym/go-boilerplate/app/infrastructure/gorm/database"
	"github.com/frasnym/go-boilerplate/app/infrastructure/gorm/repository"
	"github.com/frasnym/go-boilerplate/app/usecase/fanuser"
	"github.com/stretchr/testify/assert"
)

const (
	NAME string = "[TEST] Name 1"
)

var db, _ = gorm.NewConnectionDB("sqlite", "file::memory:?cache=shared", "localhost", "user", "password", 5432)

var (
	fanuserRepo       fanuser.FanuserRepository = repository.NewFanuserRepository(repository.NewBaseRepository(db))
	fanuserSrv        fanuser.FanuserUseCase    = fanuser.NewFanuserService(fanuserRepo)
	fanuserController FanuserController         = NewFanuserController(fanuserSrv)
)

func TestSignUpFanuser(t *testing.T) {
	// Create new HTTP POST request
	reqBody := []byte(`{"name": "` + NAME + `"}`)
	request, _ := http.NewRequest("POST", "/fanuser", bytes.NewBuffer(reqBody))

	// Assign HTTP Handler function (controller SignUpFanuser function)
	handler := http.HandlerFunc(fanuserController.SignUpFanuser)

	// Record HTTP Response (httptest)
	response := httptest.NewRecorder()

	// Dispatch the HTTP request
	handler.ServeHTTP(response, request)

	// Add assertions on the HTTP status code and the response
	status := response.Code
	if status != http.StatusCreated {
		t.Errorf("Handler returned invalid status code: got %v expected %v", status, http.StatusCreated)
	}

	// Decode the HTTP response
	var fanuser entity.Fanuser
	json.NewDecoder(io.Reader(response.Body)).Decode(&fanuser)

	// Assert HTTP response
	assert.NotNil(t, fanuser.Id)
	assert.Equal(t, NAME, fanuser.Name)
}
