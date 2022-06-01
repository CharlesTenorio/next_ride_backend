package domain

import (
	"time"
)

type Event struct {
	IdEvent         int64
	IdGroup         int64 `json:"id_grupo"`
	Descirpiton     string
	InitialDate     time.Time
	FinalDate       time.Time
	MaximumQuantity int64 `json:"qtd_max_participantes"`
	Totalcost       float64
	Local           string
	EventType       string
}

type EventDetail struct {
	IdDetail      int64     `json:"id_detalhe_evento"`
	IdEvente      int64     `json:"id_evento"`
	IdParther     int64     `json:"Parceiro"`
	AmounTDonated float64   `json:"valor doado"`
	Descirpiton   string    `json:"Descrição"`
	DateDonated   time.Time `json:"Data da doação"`
}
