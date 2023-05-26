package user

import "log"

type Business interface {
	Create(firstName, lastName, email, phone string) error
}

type business struct{}

func NewBusiness() Business {
	return &business{}
}

func (b business) Create(firstName, lastName, email, phone string) error {
	log.Println("Create user business")
	return nil
}
