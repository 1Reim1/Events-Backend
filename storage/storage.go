package storage

import "Events-Backend/data"

type Storage interface {
	GetEventList() (*[]data.Event, error)

	GetEventById(int) (*data.Event, error)

	GetEventListByCoords(float64, float64, float64) (*[]data.Event, error)
}
