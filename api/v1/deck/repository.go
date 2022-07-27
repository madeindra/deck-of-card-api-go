package deck

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type (
	DeckRepo interface {
		CreateDeck(ctx context.Context, deck Deck) error
		CreateCards(ctx context.Context, deckUUID string, cards []Card) error
		FindDeckByID(ctx context.Context, deckUUID string) (DeckData, error)
		FindCardsByDeckID(ctx context.Context, deckUUID string) ([]CardData, error)
		Update(ctx context.Context, deckUUID string, count int) error
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

func (repo *deckRepoImpl) CreateDeck(ctx context.Context, deck Deck) error {
	// create deck
	query := fmt.Sprintf("INSERT INTO %s (uuid, shuffled) VALUES ($1, $2)", repo.tableDeck)
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// insert deck to db
	_, err = stmt.ExecContext(
		ctx,
		deck.ID,
		deck.Shuffled,
	)

	if err != nil {
		return err
	}

	// return data
	return nil
}

func (repo *deckRepoImpl) CreateCards(ctx context.Context, deckUUID string, cards []Card) error {
	// create card
	query := fmt.Sprintf("INSERT INTO %s (uuid, deck_uuid, value, suit, code) VALUES", repo.tableCard)

	// create multiple rows
	values := []interface{}{}
	for idx, row := range cards {
		query += fmt.Sprintf(" ($%d, $%d, $%d, $%d, $%d)", idx*5+1, idx*5+2, idx*5+3, idx*5+4, idx*5+5)

		// if not last value, append comma
		if idx < len(cards)-1 {
			query += ","
		}

		values = append(values, uuid.New().String(), deckUUID, row.Value, row.Suit, row.Code)
	}

	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// insert card for deck
	_, err = stmt.ExecContext(
		ctx,
		values...,
	)

	if err != nil {
		return err
	}

	// return data
	return nil
}

func (repo *deckRepoImpl) FindDeckByID(ctx context.Context, deckUUID string) (DeckData, error) {
	deckData := DeckData{}

	// get deck by id
	query := fmt.Sprintf("SELECT uuid, shuffled FROM %s where uuid=$1", repo.tableDeck)
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return deckData, err
	}
	defer stmt.Close()

	// map the data to struct
	row := stmt.QueryRowContext(ctx, deckUUID)

	err = row.Scan(
		&deckData.uuid,
		&deckData.shuffled,
	)

	if err != nil {
		return deckData, err
	}

	// return data
	return deckData, nil
}

func (repo *deckRepoImpl) FindCardsByDeckID(ctx context.Context, deckUUID string) ([]CardData, error) {
	cardData := []CardData{}

	// get deck by id
	query := fmt.Sprintf("SELECT uuid, deck_uuid, value, suit, code FROM %s WHERE deck_uuid = $1", repo.tableCard)
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return cardData, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, deckUUID)
	if err != nil {
		return cardData, err
	}

	// map data to struct
	for rows.Next() {
		card := CardData{}
		err = rows.Scan(
			&card.uuid,
			&card.deck_uuid,
			&card.value,
			&card.suit,
			&card.code,
		)

		if err != nil {
			return cardData, err
		}

		cardData = append(cardData, card)
	}

	// return data
	return cardData, nil
}

func (repo *deckRepoImpl) Update(ctx context.Context, deckUUID string, count int) error {
	return nil
}
