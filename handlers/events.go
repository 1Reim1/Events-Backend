package handlers

import (
	"Events-Backend/storage"
	"encoding/json"
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"strconv"
)

type EventsHandler struct {
	storage storage.Storage
}

func NewEventsHandler(storage storage.Storage) *EventsHandler {
	return &EventsHandler{storage: storage}
}

func (h *EventsHandler) GetList(w http.ResponseWriter, _ *http.Request) {
	events, err := h.storage.GetEventList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(events)
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

	event, err := h.storage.GetEventById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *EventsHandler) GetListByCoords(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	latitude, err := strconv.ParseFloat(vars["latitude"], 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	longitude, err := strconv.ParseFloat(vars["longitude"], 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	radius, err := strconv.ParseFloat(vars["radius"], 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	events, err := h.storage.GetEventListByCoords(latitude, longitude, math.Abs(radius))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
