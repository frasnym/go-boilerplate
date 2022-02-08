package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/frasnym/go-boilerplate/app/entity"
	"github.com/frasnym/go-boilerplate/app/error"
	"github.com/frasnym/go-boilerplate/app/usecase/fanuser"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type fanuserCtrl struct {
	fanuserSrv fanuser.FanuserUseCase
}

type FanuserController interface {
	SignUpFanuser(res http.ResponseWriter, req *http.Request)
}

func NewFanuserController(fanuserService fanuser.FanuserUseCase) FanuserController {
	return &fanuserCtrl{
		fanuserSrv: fanuserService,
	}
}

func (ctrl *fanuserCtrl) SignUpFanuser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var fanuser entity.Fanuser

	err := json.NewDecoder(req.Body).Decode(&fanuser)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(error.ServiceError{Message: "Error while decoding data"})
		return
	}

	err = ctrl.fanuserSrv.Validate(&fanuser)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(error.ServiceError{Message: err.Error()})
		return
	}

	uuid, err := gonanoid.New()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(error.ServiceError{Message: "Error while creating uid"})
		return
	}

	fanuser.Uid = uuid
	fanuser.CreatedAt = time.Now().UTC().Format("2006-01-02 15:04:05")
	fanuser.UpdatedAt = time.Now().UTC().Format("2006-01-02 15:04:05")

	createdFanuser, err := ctrl.fanuserSrv.Create(&fanuser)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(error.ServiceError{Message: "Error while creating fanuser"})
		return
	}

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(createdFanuser)
}
