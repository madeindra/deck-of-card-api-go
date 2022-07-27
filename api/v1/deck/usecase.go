package deck

import "context"

type DeckUsecase interface {
	Create(ctx context.Context, isShuffled bool, cardsList []string) interface{}
	FindByID(ctx context.Context, deck_id string) interface{}
	Draw(ctx context.Context, deck_id string) interface{}
}

type deckUsecaseImpl struct {
	repository DeckRepo
}

func NewDeckUsecase(repository DeckRepo) DeckUsecase {
	return &deckUsecaseImpl{
		repository: repository,
	}
}

func (uc *deckUsecaseImpl) Create(ctx context.Context, isShuffled bool, cardsList []string) interface{} {
	// read query param
	// if shuffled is present, read from it
	// if cards is present, read from it

	// create deck of card

	// if it is to be shuffled, shuffle it

	// store deck of cards

	// return response
	return &Deck{}
}

func (uc *deckUsecaseImpl) FindByID(ctx context.Context, deck_id string) interface{} {
	// read query param
	// if uuid is present, read from it

	// if uuid is invalid return error

	// return deck of cards
	return &DeckWithCards{}
}

func (uc *deckUsecaseImpl) Draw(ctx context.Context, deck_id string) interface{} {
	// read query param
	// if uuid is present, read from it

	// if uuid is invalid return error

	// get deck by id
	// if deck doesn't exist return error

	// if deck has no card left, return empty response

	// take some card(s)

	// update existing deck of cards

	// return cards response
	return Cards{}
}
