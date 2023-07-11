package registration

import (
	"log"

	"github.com/jnka9755/go-03STUDENTAPP/internal/domain"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(register *domain.Registration) error
	}

	repository struct {
		log *log.Logger
		db  *gorm.DB
	}
)

func NewRepository(l *log.Logger, db *gorm.DB) Repository {

	return &repository{
		db:  db,
		log: l,
	}
}

func (r *repository) Create(register *domain.Registration) error {

	if err := r.db.Create(register).Error; err != nil {
		r.log.Println("Repository ->", err)
		return err
	}

	r.log.Println("Repository -> Create register with id: ", register.ID)
	return nil
}
