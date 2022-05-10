package handlers

import (
	"Events-Backend/storage"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type EventsHandler struct {
	storage *storage.Storage
}

func NewEventsHandler(storage *storage.Storage) *EventsHandler {
	return &EventsHandler{storage: storage}
}

func (h *EventsHandler) GetList(w http.ResponseWriter, _ *http.Request) {
	list := (*h.storage).GetEventList()

	err := json.NewEncoder(w).Encode(list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *EventsHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	event, err := (*h.storage).GetEventById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
