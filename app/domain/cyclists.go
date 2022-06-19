package domain

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
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
	UpdateCyclist(id, idGroup, name, Cpf string, birth time.Time, email, bloadType, healthPlan, contactEmergency, gotToKnow string,
		active bool, img, participantType string, pedaling, tours, travels int) (Cyclist, error)
	Delete(id string) (string, error)
	FindAll() ([]Cyclist, error)
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

func (c *Cyclist) ValidateName(r *http.Request) error {
	if c.Name == "" {
		return fmt.Errorf("Nome não informado", http.StatusBadRequest)
	}
	return nil
}

func (c *Cyclist) ValidateCpf(r *http.Request) error {
	if c.Cpf == "" {
		return fmt.Errorf("Cpf não informado", http.StatusBadRequest)
	}
	return nil
}

func (c *Cyclist) ValidateBirth(r *http.Request) error {
	if c.Birth.IsZero() {
		return fmt.Errorf("Nascimento não informado", http.StatusBadRequest)
	}
	return nil
}

func (c *Cyclist) ValidateEmail(r *http.Request) error {
	if c.Email == "" {
		return fmt.Errorf("Email não informado", http.StatusBadRequest)
	}
	return nil
}

func (c *Cyclist) ValidateBloodType(r *http.Request) error {
	if c.BloodType == "" {
		return fmt.Errorf("Tipo sanguineo não informado", http.StatusBadRequest)
	}
	return nil
}

func (c *Cyclist) ValidateHealthPlan(r *http.Request) error {
	if c.HealthPlan == "" {
		return fmt.Errorf("Plano de saúde não informado", http.StatusBadRequest)
	}
	return nil
}

func (c *Cyclist) ValidateContactEmergency(r *http.Request) error {
	if c.ContactEmergency == "" {
		return fmt.Errorf("Contato de emergência não informado", http.StatusBadRequest)
	}
	return nil
}

func (c *Cyclist) ValidarFicouSanbendo(r *http.Request) error {
	if c.GotToKnow == "" {
		return fmt.Errorf("Ficou sanbendo não informado", http.StatusBadRequest)
	}
	return nil
}

func (c *Cyclist) ValidateImg(r *http.Request) error {
	if c.Img == "" {
		return fmt.Errorf("Foto não informada", http.StatusBadRequest)
	}
	return nil
}

func (c *Cyclist) ValidarTipoParticipante(r *http.Request) error {
	if c.ParticipantType == "" {
		return fmt.Errorf("Tipo participante não informado", http.StatusBadRequest)
	}
	return nil
}
