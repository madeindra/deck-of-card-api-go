package deck

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type DeckUsecase interface {
	Create(ctx context.Context, isShuffled bool, cardsList []string) interface{}
	FindByID(ctx context.Context, deckUUID string) interface{}
	Draw(ctx context.Context, deckUUID string) interface{}
}

type deckUsecaseImpl struct {
	repository DeckRepo
}

func NewDeckUsecase(repository DeckRepo) DeckUsecase {
	return &deckUsecaseImpl{
		repository: repository,
	}
}

func (uc *deckUsecaseImpl) Create(ctx context.Context, isShuffled bool, cards []string) interface{} {
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
		return &Deck{}
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
		return errors.New("Error: unable to create deck")
	}

	// store deck of cards
	err = uc.repository.CreateCards(ctx, deckUUID, allCards)
	if err != nil {
		return errors.New("Error: unable to create card")
	}

	// return response
	return Deck{
		ID:        deckUUID,
		Shuffled:  isShuffled,
		Remaining: len(allCards),
	}
}

func (uc *deckUsecaseImpl) FindByID(ctx context.Context, deckUUID string) interface{} {
	// get decks by id

	// if deck does not exist, return error

	// get cards by deck id

	// return deck of cards
	return &DeckWithCards{}
}

func (uc *deckUsecaseImpl) Draw(ctx context.Context, deckUUID string) interface{} {
	// get decks by id

	// if deck does not exist, return error

	// get cards by deck id

	// if deck has no card left, return empty response

	// take some card(s)

	// update existing deck of cards

	// return cards response
	return Cards{}
}
