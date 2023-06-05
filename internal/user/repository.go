package user

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error
}

type repo struct {
	log *log.Logger
	db  *gorm.DB
}

func NewRepository(log *log.Logger, db *gorm.DB) Repository {

	return &repo{
		log: log,
		db:  db,
	}
}

func (r *repo) Create(user *User) error {

	user.ID = uuid.New().String()

	if err := r.db.Create(user).Error; err != nil {
		r.log.Println("Repository ->", err)
		return err
	}

	r.log.Println("Repository -> Create user with id: ", user.ID)

	return nil
}
