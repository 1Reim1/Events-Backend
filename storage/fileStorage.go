package storage

import (
	"Events-Backend/data"
	"encoding/json"
	"errors"
	"io/ioutil"
)

type FileStorage struct {
	events []data.Event
}

func NewFileStorage(file string) (*FileStorage, error) {
	var events []data.Event

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &events)
	if err != nil {
		return nil, err
	}

	return &FileStorage{events: events}, nil
}

func (s *FileStorage) GetEventList() *[]data.Event {
	return &s.events
}

func (s *FileStorage) GetEventById(id int) (*data.Event, error) {
	for _, event := range s.events {
		if event.Id == id {
			return &event, nil
		}
	}
	return nil, errors.New("event not found")
}
