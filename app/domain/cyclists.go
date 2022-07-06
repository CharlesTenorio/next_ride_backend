package domain

import (
	"errors"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Cyclist struct {
	Id               string    `json:"id"`
	IdGroup          string    `json:"id_grupo_pedal"`
	Name             string    `json:"nome"`
	Cpf              string    `json:"cpf"`
	Birth            time.Time `json:"nascimento"`
	Email            string    `json:"email"`
	BloodType        string    `json:"tipo_sanguineo"`
	HealthPlan       string    `json:"plano_saude"`
	ContactEmergency string    `json:"contato_emergencia"`
	GotToKnow        string    `json:"ficou_sanbendo"`
	CreateAt         time.Time `json:"data"`
	Active           bool      `json:"ativo"`
	Img              string    `json:"foto"`
	ParticipantType  string    `json:"tipo_participante"`
	Pedaling         int       `json:"qtd_pedaladas"`
	Tours            int       `json:"qtd_pedaladas_real"`
	Travels          int       `json:"qtd_viagens"`
}

type CyclistsRepository interface {
	Get(id string) (Cyclist, error)
	GetByName(name string) (Cyclist, error)
	GetByCpf(cpf string) (Cyclist, error)
	Create(Cyclist Cyclist) (Cyclist, error)
	UpdateCyclist(Cyclist Cyclist) (Cyclist, error)
	Delete(id string) (string, error)
	FindAll() ([]Cyclist, error)
	UpdatePedaling(id string, newPedaling int) (string, error)
	UpdateTours(id string, newTours int) (string, error)
	UpdateTravels(id string, newTravel int) (string, error)
}

func NewCyclist() *Cyclist {
	cyclist := Cyclist{
		Id:       uuid.NewV4().String(),
		CreateAt: time.Now(),
		Active:   true,
		Pedaling: 0,
		Tours:    0,
		Travels:  0,
	}

	return &cyclist
}

func (cyclist *Cyclist) Prepere() error {
	if erro := cyclist.Validade(); erro != nil {
		return erro
	}
	cyclist.formatSpace()
	return nil
}

func (c *Cyclist) Validade() error {
	if c.Name == "" {
		return errors.New("O Nome o brigatório")
	}

	if c.Cpf == "" {
		return errors.New("O CPF o brigatório")
	}

	if c.Birth.IsZero() {
		return errors.New(" Nascimento é o brigatório")
	}

	if c.Email == "" {
		return errors.New("Email não informado")
	}

	if c.BloodType == "" {
		return errors.New("Tipo sanguineo não informado")
	}

	if c.HealthPlan == "" {
		return errors.New("Plano de saúde não informado")
	}

	if c.ContactEmergency == "" {
		return errors.New("Contato de emergência não informado")
	}

	if c.GotToKnow == "" {
		return errors.New("Ficou sanbendo não informado")
	}

	if c.Img == "" {
		return errors.New("Foto não informada")
	}

	if c.ParticipantType == "" {
		return errors.New("Tipo participante não informado")
	}
	return nil

}

func (cyclist *Cyclist) formatSpace() {
	cyclist.Name = strings.TrimSpace(cyclist.Name)
	cyclist.Email = strings.TrimSpace(cyclist.Email)
}
