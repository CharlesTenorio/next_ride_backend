package domain

import "time"

type Checkin struct {
	Id              int64 `json:"id"`
	IdShedule       int64 `json:"id_agenda"`
	MaximumQuantity int64 `json:"qtd_max_participantes"`
	CyclistQuantity int64 `json:"qtd_participantes"`
}

type DetalheCheckin struct {
	Id          int64     `json:"id"`
	IdCheckin   int64     `json:"id_checkin"`
	IdCyclist   int64     `json:"id_participante"`
	CheckinDate time.Time `json:"deta_hora"`
}
