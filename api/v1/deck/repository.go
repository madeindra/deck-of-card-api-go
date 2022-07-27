package deck

import (
	"context"
	"database/sql"
)

type (
	DeckRepo interface {
		Create(ctx context.Context, cards []string) (string, error)
		FindDeckByID(ctx context.Context, deck_id string) (Deck, error)
		FindCardsByDeckID(ctx context.Context, deck_id string) ([]Card, error)
		Update(ctx context.Context, deck_id string, count int64) error
	}

	deckRepoImpl struct {
		db        *sql.DB
		tableDeck string
		tableCard string
	}
)

func NewDeckRepo(db *sql.DB, tableDeck string, tableCard string) DeckRepo {
	return &deckRepoImpl{
		db:        db,
		tableDeck: tableDeck,
		tableCard: tableCard,
	}
}

func (repo *deckRepoImpl) Create(ctx context.Context, cards []string) (string, error) {
	// create deck

	// insert card for deck

	// return data
	return "", nil
}

func (repo *deckRepoImpl) FindDeckByID(ctx context.Context, deck_id string) (Deck, error) {
	// find deck

	// return response
	return Deck{}, nil
}

func (repo *deckRepoImpl) FindCardsByDeckID(ctx context.Context, deck_id string) ([]Card, error) {
	// find deck

	// return response
	return []Card{}, nil
}

func (repo *deckRepoImpl) Update(ctx context.Context, deck_id string, count int64) error {
	return nil
}
