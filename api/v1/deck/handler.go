package deck

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/madeindra/toggl-test/internal/response"
)

type DeckHandler struct {
	Usecase DeckUsecase
}

func NewDeckHandler(router *mux.Router, usecase DeckUsecase) {
	handler := &DeckHandler{
		Usecase: usecase,
	}

	router.HandleFunc("/v1/decks", handler.Create).Methods(http.MethodPost)
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

	response.JSON(w, result)
}

func (handler *DeckHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uuid := mux.Vars(r)["uuid"]
	result := handler.Usecase.FindByID(ctx, uuid)

	response.JSON(w, result)
}

func (handler *DeckHandler) Draw(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uuid := mux.Vars(r)["uuid"]

	countParam := r.URL.Query().Get("count")
	count, _ := strconv.Atoi(countParam)

	result := handler.Usecase.Draw(ctx, uuid, count)

	response.JSON(w, result)
}
