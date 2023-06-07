package user

import (
	"fmt"
	"log"
)

type Business interface {
	Create(user *User) (*User, error)
	GetAll() ([]User, error)
	Get(id string) (*User, error)
	Delete(id string) error
	Update(id string, firstName, lastName, email, phone *string) error
}

type business struct {
	log        *log.Logger
	repository Repository
}

func NewBusiness(log *log.Logger, repository Repository) Business {
	return &business{
		log:        log,
		repository: repository,
	}
}

func (b business) Create(user *User) (*User, error) {

	b.log.Println("Create user Bussiness")
	if err := b.repository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (b business) GetAll() ([]User, error) {

	b.log.Println("GetAll user Bussiness")
	users, err := b.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (b business) Get(id string) (*User, error) {

	b.log.Println("Get user Bussiness")
	user, err := b.repository.Get(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (b business) Delete(id string) error {

	b.log.Println("Delete user Bussiness")
	return b.repository.Delete(id)
}

func (b business) Update(id string, firstName, lastName, email, phone *string) error {

	b.log.Println("Update user Bussiness")

	fmt.Println("ASDDASDASDa", firstName)
	fmt.Println("ASDDASDASDa", lastName)

	return b.repository.Update(id, firstName, lastName, email, phone)
}
