package fanuser

import "github.com/frasnym/go-boilerplate/app/entity"

type FanuserUseCase interface {
	Validate(post *entity.Fanuser) error
	Create(post *entity.Fanuser) (*entity.Fanuser, error)
}

type FanuserRepository interface {
	Create(post *entity.Fanuser) (*entity.Fanuser, error)
}
