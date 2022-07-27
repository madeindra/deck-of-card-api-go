package deck

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type DeckHandler struct {
	Validate *validator.Validate
	Usecase  DeckUsecase
}

func NewDeckHandler(router *mux.Router, validate *validator.Validate, usecase DeckUsecase) {
	handler := &DeckHandler{
		Validate: validate,
		Usecase:  usecase,
	}

	router.HandleFunc("/v1/decks", handler.Create).Methods(http.MethodGet)
	router.HandleFunc("/v1/decks/{uuid}", handler.FindByID).Methods(http.MethodGet)
	router.HandleFunc("/v1/decks/{uuid}/draw", handler.Draw).Methods(http.MethodGet)
}

func (handler *DeckHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (handler *DeckHandler) FindByID(w http.ResponseWriter, r *http.Request) {

}

func (handler *DeckHandler) Draw(w http.ResponseWriter, r *http.Request) {

}
