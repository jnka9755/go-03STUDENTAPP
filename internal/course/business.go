package course

import "log"

type (
	Business interface {
		Create(course *Course) (*Course, error)
	}

	business struct {
		log        *log.Logger
		repository Repository
	}
)

func NewBusiness(log *log.Logger, repository Repository) Business {
	return &business{
		log:        log,
		repository: repository,
	}
}

func (b business) Create(course *Course) (*Course, error) {

	if err := b.repository.Create(course); err != nil {
		return nil, err
	}

	return course, nil
}
