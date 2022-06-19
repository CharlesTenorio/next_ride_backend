package service

import (
	"fmt"
	"gitlab.com/charlestenorios/next_ride_backend/app/domain"
	"gitlab.com/charlestenorios/next_ride_backend/app/repository"
	"time"
)

type CyclistService struct {
	Repository domain.CyclistsRepository
}

func NewCycistService(repository *repository.CysclistRepositoryPsql) *CyclistService {
	return &CyclistService{Repository: repository}
}

func (c *CyclistService) Get(id string) (domain.Cyclist, error) {
	cyclist, error := c.Repository.Get(id)
	if error != nil {
		return domain.Cyclist{}, error
	}
	return cyclist, nil
}

func (c *CyclistService) GetByName(name string) (domain.Cyclist, error) {
	name = "'" + name + "%'"
	cyclist, error := c.Repository.GetByName(name)
	if error != nil {
		return domain.Cyclist{}, error
	}
	return cyclist, nil
}

func (c *CyclistService) GetByCpf(cpf string) (domain.Cyclist, error) {

	cyclist, error := c.Repository.GetByCpf(cpf)
	if error != nil {
		return domain.Cyclist{}, error
	}
	return cyclist, nil
}

func (c *CyclistService) Create(idGroup, name, cpf string, birth time.Time, email, bloodType, healthPlan, contactEmergency, gotToKnow, img, participantType string) (domain.Cyclist, error) {
	cyclist := domain.NewCyclist()
	cyclist.IdGroup = idGroup
	cyclist.Name = name
	cyclist.Cpf = cpf
	cyclist.Birth = birth
	cyclist.Email = email
	cyclist.BloodType = bloodType
	cyclist.HealthPlan = healthPlan
	cyclist.ContactEmergency = contactEmergency
	cyclist.GotToKnow = gotToKnow
	cyclist.Img = img
	cyclist.ParticipantType = participantType

	cycl, err := c.Repository.Create(*cyclist)

	if err != nil {
		return domain.Cyclist{}, err
	}
	return cycl, nil
}

func (c *CyclistService) FindAll() ([]domain.Cyclist, error) {
	cyclist, error := c.Repository.FindAll()
	if error != nil {
		return []domain.Cyclist{}, error
	}
	return cyclist, nil
}

func (c *CyclistService) UpdateCyclist(id string, idGroup string, name string, cpf string, birth time.Time, email string, bloodType string, healthPlan string, contactEmergency string, gotToKnow string, active bool, img string, participantType string, pedaling int, tours int, travels int) (domain.Cyclist, error) {

	cycli, err := c.Repository.UpdateCyclist(id, idGroup, name, cpf, birth, email, bloodType, healthPlan, contactEmergency, gotToKnow,
		active, img, participantType, pedaling, tours, travels)

	if err != nil {
		return domain.Cyclist{}, err
	}
	return cycli, nil
}

func (c *CyclistService) Delete(id string) (string, error) {

	_, err := c.Repository.Delete(id)

	if err != nil {
		fmt.Print(err)
		return "DEU RIIM", err
	}

	return "Dado excluido com sucesso", nil
}
