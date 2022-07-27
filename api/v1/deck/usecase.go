package deck

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/madeindra/toggl-test/internal/response"
)

type DeckUsecase interface {
	Create(ctx context.Context, isShuffled bool, cardsList []string) response.ResultData
	FindByID(ctx context.Context, deckUUID string) response.ResultData
	Draw(ctx context.Context, deckUUID string) response.ResultData
}

type deckUsecaseImpl struct {
	repository DeckRepo
}

func NewDeckUsecase(repository DeckRepo) DeckUsecase {
	return &deckUsecaseImpl{
		repository: repository,
	}
}

func (uc *deckUsecaseImpl) Create(ctx context.Context, isShuffled bool, cards []string) response.ResultData {
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

	// create deck
	deckUUID := uuid.New().String()

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

	// store deck of cards
	err = uc.repository.CreateCards(ctx, deckUUID, allCards)
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

func (uc *deckUsecaseImpl) FindByID(ctx context.Context, deckUUID string) response.ResultData {
	// get decks by id
	deckData, err := uc.repository.FindDeckByID(ctx, deckUUID)

	// if deck does not exist, return error
	if err != nil {
		return response.ResultData{
			Code: http.StatusBadRequest,
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

func (uc *deckUsecaseImpl) Draw(ctx context.Context, deckUUID string) response.ResultData {
	// get decks by id

	// if deck does not exist, return error

	// get cards by deck id

	// if deck has no card left, return empty response

	// take some card(s)

	// update existing deck of cards

	// return cards response
	return response.ResultData{
		Code: http.StatusOK,
		Data: Cards{},
	}
}
