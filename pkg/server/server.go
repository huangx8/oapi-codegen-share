package server

import (
	"encoding/json"
	"net/http"
	"slices"

	"github.com/google/uuid"

	"github.com/huangx8/oapi-codegen-share/example/api"
)

type CardCenter struct {
	Store []api.Card
}

func NewServer() *CardCenter {
	return &CardCenter{
		Store: []api.Card{},
	}
}

func (c *CardCenter) ListCards(w http.ResponseWriter, r *http.Request, params api.ListCardsParams) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(c.Store)
}

func (c *CardCenter) AddCard(w http.ResponseWriter, r *http.Request) {
	var newCard api.NewCard
	if err := json.NewDecoder(r.Body).Decode(&newCard); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "invalid request"})
		return
	}

	card := api.Card{Id: uuid.NewString(), Owner: newCard.Owner}
	c.Store = append(c.Store, card)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(card)
}

func (c *CardCenter) FindCardByID(w http.ResponseWriter, r *http.Request, id string) {
	idx := slices.IndexFunc(c.Store, func(c api.Card) bool { return c.Id == id })
	if idx != -1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(c.Store[idx])
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": "card not found"})
}
