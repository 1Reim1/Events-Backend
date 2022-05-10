package storage

import "Events-Backend/data"

type Storage interface {
	GetEventList() *[]data.Event
	GetEventById(int) (*data.Event, error)
}
