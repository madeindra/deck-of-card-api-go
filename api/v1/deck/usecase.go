package deck

import (
	"context"
	"net/http"

	"github.com/madeindra/toggl-test/internal/response"
	"github.com/madeindra/toggl-test/internal/uuid"
)

type DeckUse interface {
	Create(ctx context.Context, isShuffled bool, cardsList []string) response.ResultData
	FindByID(ctx context.Context, deckUUID string) response.ResultData
	Draw(ctx context.Context, deckUUID string, count int) response.ResultData
}

type deckUser struct {
	repository DeckRepose
}

func NewDeckUsecase(repository DeckRepose) DeckUse {
	return &deckUser{
		repository: repository,
	}
}

func (uc *deckUser) Create(ctx context.Context, isShuffled bool, cards []string) response.ResultData {
	// check if list of cards is provided
	codes := []string{}
	if len(cards) != 0 {
		codes = append(codes, cards...)
	} else {
		codes = append(codes, standardDeck...)
	}

	// check if deck need to be shuffled
	if isShuffled {
		codes = shuffle(codes)
	}

	// create deck of card by decoding codes into values
	allCards, err := decode(codes)

	if err != nil {
		return response.ResultData{
			Code: http.StatusBadRequest,
			Data: response.FailedResult{
				Message: "Invalid card",
			},
		}
	}

	// craete uuid
	generator := uuid.UUIDGenerator{UUID: &uuid.GoogleUUID{}}

	// create deck
	deckUUID := generator.UUID.NewString()

	deckData := Deck{
		ID:        deckUUID,
		Shuffled:  isShuffled,
		Remaining: 0,
	}

	err = uc.repository.CreateDeck(ctx, deckData)
	if err != nil {
		return response.ResultData{
			Code: http.StatusInternalServerError,
			Data: response.FailedResult{
				Message: "Unable to create deck",
			},
		}
	}

	// create uuids for cards
	cardUUIDS := generator.NewStringSlice(len(allCards))

	// store deck of cards
	err = uc.repository.CreateCards(ctx, deckUUID, cardUUIDS, allCards)
	if err != nil {
		return response.ResultData{
			Code: http.StatusInternalServerError,
			Data: response.FailedResult{
				Message: "Unable to create cards",
			},
		}
	}

	// return response
	return response.ResultData{
		Code: http.StatusCreated,
		Data: Deck{
			ID:        deckUUID,
			Shuffled:  isShuffled,
			Remaining: len(allCards),
		},
	}
}

func (uc *deckUser) FindByID(ctx context.Context, deckUUID string) response.ResultData {
	// get decks by id
	deckData, err := uc.repository.FindDeckByID(ctx, deckUUID)

	// if deck does not exist, return error
	if err != nil {
		return response.ResultData{
			Code: http.StatusNotFound,
			Data: response.FailedResult{
				Message: "Unable to find deck",
			},
		}
	}

	if deckData.uuid == "" {
		return response.ResultData{
			Code: http.StatusBadRequest,
			Data: response.FailedResult{
				Message: "Invalid deck id",
			},
		}
	}

	// get cards by deck id
	cardData, err := uc.repository.FindCardsByDeckID(ctx, deckUUID)
	if err != nil {
		return response.ResultData{
			Code: http.StatusNotFound,
			Data: response.FailedResult{
				Message: "Unable to find cards",
			},
		}
	}

	// map data to struct
	deckWithCard := DeckWithCards{
		ID:        deckData.uuid,
		Shuffled:  deckData.shuffled,
		Remaining: len(cardData),
		Cards:     []Card{},
	}

	for _, card := range cardData {
		deckWithCard.Cards = append(deckWithCard.Cards, Card{
			Value: card.value,
			Suit:  card.suit,
			Code:  card.code,
		})
	}

	// return deck of cards
	return response.ResultData{
		Code: http.StatusOK,
		Data: deckWithCard,
	}
}

func (uc *deckUser) Draw(ctx context.Context, deckUUID string, count int) response.ResultData {
	// make sure deck is drawn at minimum of 1 card
	if count < 1 {
		return response.ResultData{
			Code: http.StatusBadRequest,
			Data: response.FailedResult{
				Message: "Invalid draw amount",
			},
		}
	}

	// get decks by id
	deckData, err := uc.repository.FindDeckByID(ctx, deckUUID)

	// if deck does not exist, return error
	if err != nil {
		return response.ResultData{
			Code: http.StatusNotFound,
			Data: response.FailedResult{
				Message: "Unable to find deck",
			},
		}
	}

	if deckData.uuid == "" {
		return response.ResultData{
			Code: http.StatusBadRequest,
			Data: response.FailedResult{
				Message: "Invalid deck id",
			},
		}
	}

	// get cards by deck id
	cardData, err := uc.repository.FindCardsWithLimit(ctx, deckUUID, count)
	if err != nil {
		return response.ResultData{
			Code: http.StatusNotFound,
			Data: response.FailedResult{
				Message: "Unable to find cards",
			},
		}
	}

	// map data to struct
	cardsRes := Cards{
		Cards: []Card{},
	}

	// store card uuid to be deleted from deck
	cardsUUID := []string{}

	for _, card := range cardData {
		cardsRes.Cards = append(cardsRes.Cards, Card{
			Value: card.value,
			Suit:  card.suit,
			Code:  card.code,
		})

		cardsUUID = append(cardsUUID, card.uuid)
	}

	// delete cards from deck
	err = uc.repository.DeleteCards(ctx, cardsUUID)
	if err != nil {
		return response.ResultData{
			Code: http.StatusInternalServerError,
			Data: response.FailedResult{
				Message: "Unable to draw cards from deck",
			},
		}
	}

	// return cards response
	return response.ResultData{
		Code: http.StatusOK,
		Data: cardsRes,
	}
}
