package fanuser

import (
	"errors"

	"github.com/frasnym/go-boilerplate/app/entity"
)

type fanuserService struct {
	fanuserRepo FanuserRepository
}

func NewFanuserService(repo FanuserRepository) FanuserUseCase {
	return &fanuserService{
		fanuserRepo: repo,
	}
}

func (*fanuserService) Validate(fanuser *entity.Fanuser) error {
	if fanuser == nil {
		return errors.New("fanuser is empty")
	}

	if fanuser.Name == "" {
		return errors.New("name cannot be empty")
	}

	return nil
}

func (srv *fanuserService) Create(fanuser *entity.Fanuser) (*entity.Fanuser, error) {
	createdFanuser, err := srv.fanuserRepo.Create(fanuser)
	if err != nil {
		return nil, err
	}

	return createdFanuser, nil
}
