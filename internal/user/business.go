package user

import "log"

type Business interface {
	Create(user *User) (*User, error)
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
