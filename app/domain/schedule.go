package domain

import (
	"time"
)

type Schedule struct {
	IdSchedule      int64     `json:"id"`
	IdGroup         int64     `json:"id_grupo"`
	DayTime         time.Time `json:"dia_hora"`
	DayWeek         string    `json:"dia_semana"`
	EventType       string    `json:"tipo_evento"`
	EventName       string    `json:"nome_evento"`
	MaximumQuantity int64     `json:"qtd_max_participantes"`
	StatusShedule   string    `json:"status_agenda"`
	distance        float64   `json:"distancia"`
}
