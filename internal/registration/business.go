package registration

import (
	"errors"
	"log"

	"github.com/jnka9755/go-03STUDENTAPP/internal/course"
	"github.com/jnka9755/go-03STUDENTAPP/internal/domain"
	"github.com/jnka9755/go-03STUDENTAPP/internal/user"
)

type (
	Business interface {
		Create(register *CreateReq) (*domain.Registration, error)
	}

	business struct {
		log            *log.Logger
		userBusiness   user.Business
		courseBusiness course.Business
		repository     Repository
	}
)

func NewBusiness(log *log.Logger, userBusiness user.Business, courseBusiness course.Business, repository Repository) Business {
	return &business{
		log:            log,
		userBusiness:   userBusiness,
		courseBusiness: courseBusiness,
		repository:     repository,
	}
}

func (b business) Create(request *CreateReq) (*domain.Registration, error) {

	if _, err := b.userBusiness.Get(request.UserID); err != nil {
		return nil, errors.New("user_id doesn't exists")
	}

	if _, err := b.courseBusiness.Get(request.CourseID); err != nil {
		return nil, errors.New("course_id doesn't exists")
	}

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
