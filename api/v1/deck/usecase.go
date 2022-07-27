package deck

import "context"

type DeckUsecase interface {
	Create(ctx context.Context) interface{}
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

func (uc *deckUsecaseImpl) Create(ctx context.Context) interface{} {
	return &Deck{}
}

func (uc *deckUsecaseImpl) FindByID(ctx context.Context, deck_id string) interface{} {
	return &DeckWithCards{}
}

func (uc *deckUsecaseImpl) Draw(ctx context.Context, deck_id string) interface{} {
	return []Card{}
}
