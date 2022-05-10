package storage

import "Events-Backend/data"

type Storage interface {
	GetEventList() *[]data.Event

	GetEventById(int) (*data.Event, error)

	GetEventListByCoords(float64, float64, float64, float64) *[]data.Event
}
