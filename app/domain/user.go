package domain

import (
	"errors"
	"time"

	"github.com/badoux/checkmail"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/charlestenorios/next_ride_backend/app/safety"
)

type User struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Admin    bool      `json:"admin"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	CreateAt time.Time `json:"data"`
	Active   bool      `json:"ativo"`
}

func NewUser() *User {
	user := User{
		Id:       uuid.NewV4().String(),
		CreateAt: time.Now(),
		Active:   true,
	}

	return &user
}

type UserRepository interface {
	Get(id string) (User, error)
	Create(User User) (User, error)
	Delete(id string) (string, error)
	FindAll() ([]User, error)
}

func (usr *User) ValidationUsr(stage string) error {
	if usr.Name == "" {
		return errors.New("Nome obrigatório")
	}

	if usr.Email == "" {
		return errors.New("Email obrigatório")
	}

	if usr.Password == "" {
		return errors.New("senha é obrigarotria")
	}

	if error := checkmail.ValidateFormat(usr.Email); error != nil {
		return errors.New("Formato de email e invalido")
	}

	if stage == "create" {
		pHash, erro := safety.Hash(usr.Password)
		if erro != nil {
			return erro
		}
		usr.Password = string(pHash)
	}

	return nil

}
