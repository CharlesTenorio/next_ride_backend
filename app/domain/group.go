package domain

import (
	uuid "github.com/satori/go.uuid"
)

type Group struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"nome" validate:"required"`
}

type GroupRepository interface {
	Get(id string) (Group, error)
	Create(Group Group) (Group, error)
	Update(id string, name string) (Group, error)
	Delete(id string) (string, error)
	FindAll() ([]Group, error)
}

func NewGroup() *Group {
	group := Group{
		Id: uuid.NewV4().String(),
	}

	return &group
}
