package interfaces

import "gitlab.com/charlestenorios/next_ride_backend/app/domain"

type PedalGroupRepository interface {
	Get(id int64) (domain.Group, error)
	Create(pedalGroup domain.Group) (domain.Group, error)
	Update(id int16, group domain.Group) (domain.Group, error)
	Delete(id int64) (string, error)
	FindAll() ([]domain.Group, error)
}

type CyclistsRepository interface {
	Get(id int64) (domain.Cyclist, error)
	Create(cyclist domain.Cyclist) (domain.Cyclist, error)
	Update(id int16, cyclist domain.Cyclist) (domain.Cyclist, error)
	Delete(id int64) (string, error)
	FindAll() ([]domain.Cyclist, error)
}
