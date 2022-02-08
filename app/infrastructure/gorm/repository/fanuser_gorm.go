package repository

import (
	"github.com/frasnym/go-boilerplate/app/entity"
	"github.com/frasnym/go-boilerplate/app/usecase/fanuser"
)

type fanuserRepo struct {
	base BaseRepository
}

func NewFanuserRepository(br BaseRepository) fanuser.FanuserRepository {
	return &fanuserRepo{br}
}

func (repo *fanuserRepo) Create(post *entity.Fanuser) (*entity.Fanuser, error) {
	err := repo.base.GetDB().
		Create(post).Error
	if err != nil {
		return nil, err
	}

	return post, nil
}
