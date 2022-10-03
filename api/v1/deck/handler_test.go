package deck_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/madeindra/toggl-test/api/v1/deck"
	"github.com/madeindra/toggl-test/api/v1/deck/mocks"
	"github.com/madeindra/toggl-test/internal/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateHandlerDefault(t *testing.T) {
	deckUsecase := new(mocks.DeckUse)
	deckUsecase.On("Create", mock.Anything, mock.AnythingOfType("bool"), mock.AnythingOfType("[]string")).Return(response.ResultData{
		Code: 201,
		Data: deck.Deck{
			ID:        "deck-id",
			Shuffled:  true,
			Remaining: 52,
		},
	})

	deckHandler := deck.DeckHandler{
		Usecase: deckUsecase,
	}

	r := httptest.NewRequest(http.MethodPost, "/v1/decks", bytes.NewReader([]byte{}))
	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(deckHandler.Create)
	handler.ServeHTTP(recorder, r)

	response := deck.Deck{}
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, response.ID, "deck-id")
	assert.Equal(t, response.Shuffled, true)
	assert.Equal(t, response.Remaining, 52)
}

func TestCreateHandler(t *testing.T) {
	deckUsecase := new(mocks.DeckUse)
	deckUsecase.On("Create", mock.Anything, mock.AnythingOfType("bool"), mock.AnythingOfType("[]string")).Return(response.ResultData{
		Code: 201,
		Data: deck.Deck{
			ID:        "deck-id",
			Shuffled:  true,
			Remaining: 52,
		},
	})

	deckHandler := deck.DeckHandler{
		Usecase: deckUsecase,
	}

	r := httptest.NewRequest(http.MethodPost, "/v1/decks?shuffled=true&cards=AC,JS,QH,KD,2C,10S", bytes.NewReader([]byte{}))
	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(deckHandler.Create)
	handler.ServeHTTP(recorder, r)

	response := deck.Deck{}
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, response.ID, "deck-id")
	assert.Equal(t, response.Shuffled, true)
	assert.Equal(t, response.Remaining, 52)
}

func TestFindByIDHandler(t *testing.T) {
	deckUsecase := new(mocks.DeckUse)
	deckUsecase.On("FindByID", mock.Anything, mock.AnythingOfType("string")).Return(response.ResultData{
		Code: 200,
		Data: deck.DeckWithCards{
			ID:        "deck-id",
			Shuffled:  true,
			Remaining: 1,
			Cards: []deck.Card{
				{
					Value: "ACE",
					Suit:  "CLUB",
					Code:  "AC",
				},
			},
		},
	})

	deckHandler := deck.DeckHandler{
		Usecase: deckUsecase,
	}

	r := httptest.NewRequest(http.MethodGet, "/v1/decks/deck-id", bytes.NewReader([]byte{}))
	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(deckHandler.FindByID)
	handler.ServeHTTP(recorder, r)

	response := deck.DeckWithCards{}
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, response.ID, "deck-id")
	assert.Equal(t, response.Shuffled, true)
	assert.Equal(t, response.Remaining, 1)
	assert.Equal(t, response.Cards[0].Value, "ACE")
	assert.Equal(t, response.Cards[0].Suit, "CLUB")
	assert.Equal(t, response.Cards[0].Code, "AC")
}

func TestDrawHandler(t *testing.T) {
	deckUsecase := new(mocks.DeckUse)
	deckUsecase.On("Draw", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(response.ResultData{
		Code: 200,
		Data: deck.Cards{
			Cards: []deck.Card{
				{
					Value: "ACE",
					Suit:  "CLUB",
					Code:  "AC",
				},
			},
		},
	})

	deckHandler := deck.DeckHandler{
		Usecase: deckUsecase,
	}

	r := httptest.NewRequest(http.MethodGet, "/v1/decks/deck-id/draw?count=1", bytes.NewReader([]byte{}))
	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(deckHandler.Draw)
	handler.ServeHTTP(recorder, r)

	response := deck.Cards{}
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, response.Cards[0].Value, "ACE")
	assert.Equal(t, response.Cards[0].Suit, "CLUB")
	assert.Equal(t, response.Cards[0].Code, "AC")
}
