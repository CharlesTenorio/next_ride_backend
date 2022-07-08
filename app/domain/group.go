package domain

import (
	"errors"
	"strings"

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

func (group *Group) Prepere() error {
	if erro := group.validade(); erro != nil {
		return erro
	}
	group.formatSpace()
	return nil
}

func (g *Group) validade() error {
	if g.Name == "" {
		return errors.New("O Nome é o brigatório")
	}

	return nil

}

func (grupo *Group) formatSpace() {
	grupo.Name = strings.TrimSpace(grupo.Name)

}
