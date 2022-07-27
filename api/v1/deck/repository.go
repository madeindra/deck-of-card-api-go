package deck

import (
	"context"
	"database/sql"
)

type (
	DeckRepo interface {
		Create(ctx context.Context, cards string) (DeckWithCards, error)
		FindByID(ctx context.Context, ID int64) (Deck, error)
		Draw(ctx context.Context, count int64) ([]Card, error)
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

func (repo *deckRepoImpl) Create(ctx context.Context, cards string) (DeckWithCards, error) {
	return DeckWithCards{}, nil
}

func (repo *deckRepoImpl) FindByID(ctx context.Context, ID int64) (Deck, error) {
	return Deck{}, nil
}

func (repo *deckRepoImpl) Draw(ctx context.Context, count int64) ([]Card, error) {
	return []Card{}, nil
}
