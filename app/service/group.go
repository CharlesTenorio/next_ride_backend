package service

import (
	"fmt"
	"gitlab.com/charlestenorios/next_ride_backend/app/domain"
)

type GroupService struct {
	Repository domain.GroupRepository
}

func NewGroupService(repository domain.GroupRepository) *GroupService {
	return &GroupService{Repository: repository}
}

func (g *GroupService) findById(id string) (domain.Group, error) {
	group, error := g.Repository.Get(id)
	if error != nil {
		return domain.Group{}, error
	}
	return group, nil
}

func (g *GroupService) Create(name string) (domain.Group, error) {
	group := domain.NewGroup()
	group.Name = name

	grp, err := g.Repository.Create(*group)

	if err != nil {
		fmt.Print(err)
		return domain.Group{}, err
	}
	return grp, nil
}

func (g *GroupService) FindAll() ([]domain.Group, error) {
	group, error := g.Repository.FindAll()
	if error != nil {
		return []domain.Group{}, error
	}
	return group, nil
}

func (g *GroupService) Update(id, name string) (domain.Group, error) {

	grp, err := g.Repository.Update(name, id)
	fmt.Println(name, id)

	if err != nil {
		fmt.Print(err)
		return domain.Group{}, err
	}

	return grp, nil
}

func (g *GroupService) Delete(id string) (string, error) {

	_, err := g.Repository.Delete(id)

	if err != nil {
		fmt.Print(err)
		return "DEU RIIM", err
	}

	return "Dado excluido com sucesso", nil
}
