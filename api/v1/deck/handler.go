package deck

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

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
	ctx := r.Context()

	shuffleParam := r.URL.Query().Get("shuffled")
	cardsParam := r.URL.Query().Get("cards")

	isShuffled := false
	cardsList := []string{}

	if shuffleParam != "" {
		isShuffled, _ = strconv.ParseBool(shuffleParam)
	}

	if cardsParam != "" {
		cards := strings.Split(cardsParam, ",")
		cardsList = append(cardsList, cards...)
	}

	result := handler.Usecase.Create(ctx, isShuffled, cardsList)

	response, err := json.Marshal(result)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("{message: \"An Error Occured\"}"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(response)
}

func (handler *DeckHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uuid := mux.Vars(r)["uuid"]
	result := handler.Usecase.FindByID(ctx, uuid)

	response, err := json.Marshal(result)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("{message: \"An Error Occured\"}"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(response)
}

func (handler *DeckHandler) Draw(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uuid := mux.Vars(r)["uuid"]
	result := handler.Usecase.Draw(ctx, uuid)

	response, err := json.Marshal(result)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("{message: \"An Error Occured\"}"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(response)
}
