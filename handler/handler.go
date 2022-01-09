package handler

import "github.com/sygns13/golang-api/model"

// Storage es una interface de almacenamiento
type Storage interface {
	Create(person *model.Person) error
	Update(ID int, person *model.Person) error
	Delete(ID int) error
	GetByID(ID int) (model.Person, error)
	GetAll() (model.Persons, error)
}
