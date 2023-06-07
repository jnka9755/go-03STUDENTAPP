package user

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error
	GetAll() ([]User, error)
	Get(id string) (*User, error)
	Delete(id string) error
	Update(id string, firstName, lastName, email, phone *string) error
}

type repository struct {
	log *log.Logger
	db  *gorm.DB
}

func NewRepository(log *log.Logger, db *gorm.DB) Repository {

	return &repository{
		log: log,
		db:  db,
	}
}

func (r *repository) Create(user *User) error {

	user.ID = uuid.New().String()

	if err := r.db.Create(user).Error; err != nil {
		r.log.Println("Repository ->", err)
		return err
	}

	r.log.Println("Repository -> Create user with id: ", user.ID)

	return nil
}

func (r *repository) GetAll() ([]User, error) {

	r.log.Println("GetAll user Repository")
	var users []User

	if err := r.db.Model(&users).Order("created_at desc").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repository) Get(id string) (*User, error) {

	r.log.Println("Get user Repository")
	user := User{ID: id}

	if err := r.db.First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) Delete(id string) error {

	r.log.Println("Delete user Respository")
	user := User{ID: id}

	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}

	r.log.Println("Repository -> Delete user with id: ", user.ID)

	return nil
}

func (r *repository) Update(id string, firstName, lastName, email, phone *string) error {

	r.log.Println("Udate user Respository")

	fmt.Println("ASDDASDASDa", firstName)
	fmt.Println("ASDDASDASDa", lastName)

	values := make(map[string]interface{})

	if firstName != nil {
		values["first_name"] = *firstName
	}

	if lastName != nil {
		values["last_name"] = *lastName
	}

	if email != nil {
		values["email"] = *email
	}

	if phone != nil {
		values["phone"] = *phone
	}

	if err := r.db.Model(&User{}).Where("id = ?", id).Updates(values).Error; err != nil {
		return err
	}

	return nil
}
