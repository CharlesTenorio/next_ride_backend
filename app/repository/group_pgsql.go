package repository

import (
	"database/sql"
	"fmt"

	"gitlab.com/charlestenorios/next_ride_backend/app/domain"
)

type GroupRepositoryPsql struct {
	db *sql.DB
}

var valIDation domain.Group

func NewGroupRepositoryPsql(db *sql.DB) *GroupRepositoryPsql {
	return &GroupRepositoryPsql{db: db}
}

func (g *GroupRepositoryPsql) Get(id string) (domain.Group, error) {
	var grupo domain.Group
	stmt, err := g.db.Prepare("select id_group, name_group from pedal_group where id_group =$1")
	if err != nil {
		fmt.Print(err)
		return domain.Group{}, err
	}
	err = stmt.QueryRow(id).Scan(&grupo.Id, &grupo.Name)
	if err != nil {
		fmt.Println(string(err.Error()))
		return domain.Group{}, err
	}
	return grupo, nil

}

func (g *GroupRepositoryPsql) Create(group domain.Group) (domain.Group, error) {
	err := valIDation.Prepere()
	if err != nil {
		return domain.Group{}, err
	}

	stmt, err := g.db.Prepare("INSERT INTO pedal_group (id_group, name_group) VALUES($1, $2)")
	if err != nil {
		fmt.Print(err.Error())

		return domain.Group{}, err

	}
	_, err = stmt.Exec(group.Id, group.Name)
	if err != nil {
		return domain.Group{}, err
	}
	err = stmt.Close()
	if err != nil {
		return domain.Group{}, err
	}
	return group, nil
}

func (g *GroupRepositoryPsql) FindAll() ([]domain.Group, error) {
	var group domain.Group
	var groups []domain.Group
	stmt, err := g.db.Query("select id_group, name_group from pedal_group")
	if err != nil {
		fmt.Print(err)
		return []domain.Group{}, err
	}

	for stmt.Next() {
		err = stmt.Scan(&group.Id, &group.Name)
		groups = append(groups, group)
	}

	return groups, nil
}

func (g *GroupRepositoryPsql) Update(name, id string) (domain.Group, error) {
	err := valIDation.Prepere()
	if err != nil {
		return domain.Group{}, err
	}
	var group domain.Group

	stmt, err := g.db.Prepare("update pedal_group set name_group=$1 where id_group=$2")
	if err != nil {
		fmt.Print(err)

		return domain.Group{}, err

	}
	_, err = stmt.Exec(name, id)
	if err != nil {
		return domain.Group{}, err
	}
	err = stmt.Close()
	if err != nil {
		return domain.Group{}, err
	}
	return group, nil
}

func (g *GroupRepositoryPsql) Delete(id string) (string, error) {

	stmt, err := g.db.Prepare("DELETE FROM pedal_group where id_group=$1")
	if err != nil {
		fmt.Print(err)

		return "erro ao prerar para exclusão", err

	}
	_, err = stmt.Exec(id)
	if err != nil {
		return "Error ao excluir", err
	}
	err = stmt.Close()
	if err != nil {
		return "erro ao fechar conexao", err
	}
	return "Registro excluido com sucesso", nil
}
