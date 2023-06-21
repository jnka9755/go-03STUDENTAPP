package course

import (
	"log"

	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(course *Course) error
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

func (r *repository) Create(course *Course) error {

	if err := r.db.Create(course).Error; err != nil {
		r.log.Println("Repository ->", err)
		return err
	}

	r.log.Println("Repository -> Create user with id: ", course.ID)

	return nil
}
