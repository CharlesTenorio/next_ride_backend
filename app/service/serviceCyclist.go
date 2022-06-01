package service

import (
	"fmt"
	"time"

	"gitlab.com/charlestenorios/next_ride_backend/app/domain"
)

type CyclistService struct {
	Repository domain.CyclistsRepository
}

func NewCysistService(repository domain.CyclistsRepository) *CyclistService {
	return &CyclistService{Repository: repository}
}

func (c *CyclistService) findById(id string) (domain.Cyclist, error) {
	cyclist, error := c.Repository.Get(id)
	if error != nil {
		return domain.Cyclist{}, error
	}
	return cyclist, nil
}

func (c *CyclistService) findByName(name string) (domain.Cyclist, error) {
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

func (c *CyclistService) Create(id, idGroup, name, cpf string, birth time.Time, email, bloodType, healthPlan, contactEmergency, gotToKnow, img, participantType string) (domain.Cyclist, error) {
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

func (c *CyclistService) Update(id, idGroup, name, cpf string, birth time.Time,
	email, bloodType, healthPlan, contactEmergency, gotToKnow string,
	active bool, img, participantType string, pedaling,
	tours, travels int64) (domain.Cyclist, error) {

	cycli, err := c.Repository.Update(id, idGroup, name, cpf, birth, email, bloodType, healthPlan, contactEmergency, gotToKnow,
		active, img, participantType, pedaling, tours, travels)
	fmt.Println(name, id)

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
