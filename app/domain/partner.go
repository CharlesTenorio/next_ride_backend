package domain

import "time"

type Partner struct {
	Id          int64     `json:"id"`
	Name        string    `json:"nome"`
	Descirpiton string    `json:"descricao"`
	CreateAt    time.Time `json:"data"`
	TypeParther int64     `Json: "tipo"`
	img         string    `Json: "tipo"`
	Active      bool      `Json: "ativo"`
	Fone        string    `Json: "fone"`
	selo        string    `Json: "selo"`
}
