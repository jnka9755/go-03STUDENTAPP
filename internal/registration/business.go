package registration

import (
	"log"

	"github.com/jnka9755/go-03STUDENTAPP/internal/domain"
)

type (
	Business interface {
		Create(register *CreateReq) (*domain.Registration, error)
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

func (b business) Create(request *CreateReq) (*domain.Registration, error) {

	register := domain.Registration{
		UserID:   request.UserID,
		CourseID: request.CourseID,
		Status:   "P",
	}

	if err := b.repository.Create(&register); err != nil {
		return nil, err
	}

	return &register, nil
}
