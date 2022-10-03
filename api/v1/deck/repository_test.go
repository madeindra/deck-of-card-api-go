package deck_test

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/madeindra/toggl-test/api/v1/deck"
	"github.com/madeindra/toggl-test/internal/constant"
	"github.com/madeindra/toggl-test/internal/mock"
	"github.com/stretchr/testify/assert"
)

var deckMock = &deck.Deck{
	ID:        "some-id",
	Shuffled:  true,
	Remaining: 52,
}

var cardsArgs = []driver.Value{"first-card", "second-card", "third-card"}
var cardsUUID = []string{"first-card", "second-card", "third-card"}

var allCardsArgs = []driver.Value{
	"first-card", "some-id", "JACK", "SPADE", "JS",
	"second-card", "some-id", "QUEEN", "HEART", "QH",
	"third-card", "some-id", "KING", "CLUB", "KC",
}
var allCards = []deck.Card{
	{Value: "JACK", Suit: "SPADE", Code: "JS"},
	{Value: "QUEEN", Suit: "HEART", Code: "QH"},
	{Value: "KING", Suit: "CLUB", Code: "KC"},
}

func TestCreateDeckSuccess(t *testing.T) {
	db, mock := mock.NewMock()
	repo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	defer db.Close()

	query := fmt.Sprintf(regexp.QuoteMeta("INSERT INTO %s (uuid, shuffled) VALUES ($1, $2)"), constant.TableDeck)
	mock.ExpectExec(query).WithArgs(deckMock.ID, deckMock.Shuffled).WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.TODO()
	err := repo.CreateDeck(ctx, *deckMock)
	assert.NoError(t, err)
}

func TestCreateDeckExecFailed(t *testing.T) {
	db, mock := mock.NewMock()
	repo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	defer db.Close()

	query := fmt.Sprintf(regexp.QuoteMeta("INSERT INTO %s (uuid, shuffled) VALUES ($1, $2)"), constant.TableDeck)
	mock.ExpectExec(query).WithArgs(deckMock.ID, deckMock.Shuffled).WillReturnError(errors.New("an error occured"))

	ctx := context.TODO()
	err := repo.CreateDeck(ctx, *deckMock)
	assert.Error(t, err)
}

func TestCreateDeckPrepareFailed(t *testing.T) {
	db, mock := mock.NewMock()
	repo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	defer db.Close()

	query := fmt.Sprintf(regexp.QuoteMeta("INSERT INTO %s (uuid, shuffled) VALUES ($1, $2)"), constant.TableDeck)
	mock.ExpectExec(query).WillReturnError(errors.New("an error occured"))

	ctx := context.TODO()
	err := repo.CreateDeck(ctx, *deckMock)
	assert.Error(t, err)
}

func TestCreateCardsSuccess(t *testing.T) {
	db, mock := mock.NewMock()
	repo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	defer db.Close()

	query := fmt.Sprintf(regexp.QuoteMeta("INSERT INTO %s (uuid, deck_uuid, value, suit, code) VALUES ($1, $2, $3, $4, $5), ($6, $7, $8, $9, $10), ($11, $12, $13, $14, $15)"), constant.TableCards)
	mock.ExpectExec(query).WithArgs(allCardsArgs...).WillReturnResult(sqlmock.NewResult(3, 3))

	ctx := context.TODO()
	err := repo.CreateCards(ctx, deckMock.ID, cardsUUID, allCards)
	assert.NoError(t, err)
}

func TestCreateCardsPrepareSuccess(t *testing.T) {
	db, mock := mock.NewMock()
	repo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	defer db.Close()

	query := fmt.Sprintf(regexp.QuoteMeta("INSERT INTO %s (uuid, deck_uuid, value, suit, code) VALUES ($1, $2, $3, $4, $5), ($6, $7, $8, $9, $10), ($11, $12, $13, $14, $15)"), constant.TableCards)
	mock.ExpectExec(query).WillReturnError(errors.New("an error occured"))

	ctx := context.TODO()
	err := repo.CreateCards(ctx, deckMock.ID, cardsUUID, allCards)
	assert.Error(t, err)
}

func TestCreateCardsExecSuccess(t *testing.T) {
	db, mock := mock.NewMock()
	repo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	defer db.Close()

	query := fmt.Sprintf(regexp.QuoteMeta("INSERT INTO %s (uuid, deck_uuid, value, suit, code) VALUES ($1, $2, $3, $4, $5), ($6, $7, $8, $9, $10), ($11, $12, $13, $14, $15)"), constant.TableCards)
	mock.ExpectExec(query).WithArgs(allCardsArgs...).WillReturnError(errors.New("an error occured"))

	ctx := context.TODO()
	err := repo.CreateCards(ctx, deckMock.ID, cardsUUID, allCards)
	assert.Error(t, err)
}

func TestFindDeckSuccess(t *testing.T) {
	db, mock := mock.NewMock()
	repo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	defer db.Close()

	query := fmt.Sprintf(regexp.QuoteMeta("SELECT uuid, shuffled FROM %s where uuid=$1"), constant.TableDeck)
	rows := sqlmock.NewRows([]string{"uuid", "shuffle"}).AddRow("some-id", true)

	mock.ExpectQuery(query).WithArgs(deckMock.ID).WillReturnRows(rows)

	ctx := context.TODO()
	_, err := repo.FindDeckByID(ctx, deckMock.ID)
	assert.NoError(t, err)
}

func TestFindDeckQueryFailed(t *testing.T) {
	db, mock := mock.NewMock()
	repo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	defer db.Close()

	query := fmt.Sprintf(regexp.QuoteMeta("SELECT uuid, shuffled FROM %s where uuid=$1"), constant.TableDeck)

	mock.ExpectQuery(query).WithArgs(deckMock.ID).WillReturnError(errors.New("an error occured"))

	ctx := context.TODO()
	_, err := repo.FindDeckByID(ctx, deckMock.ID)
	assert.Error(t, err)
}

func TestFindDeckPrepareFailed(t *testing.T) {
	db, mock := mock.NewMock()
	repo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	defer db.Close()

	query := fmt.Sprintf(regexp.QuoteMeta("SELECT uuid, shuffled FROM %s where uuid=$1"), constant.TableDeck)

	mock.ExpectQuery(query).WillReturnError(errors.New("an error occured"))

	ctx := context.TODO()
	_, err := repo.FindDeckByID(ctx, deckMock.ID)
	assert.Error(t, err)
}

func TestFindCardsByDeckIDSuccess(t *testing.T) {
	db, mock := mock.NewMock()
	repo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	defer db.Close()

	query := fmt.Sprintf(regexp.QuoteMeta("SELECT uuid, deck_uuid, value, suit, code FROM %s WHERE deck_uuid = $1"), constant.TableCards)
	rows := sqlmock.NewRows([]string{"uuid", "deck_uuid", "value", "suit", "code"}).AddRow("first-card", "some-id", "JACK", "SPADE", "JS")

	mock.ExpectQuery(query).WithArgs(deckMock.ID).WillReturnRows(rows)

	ctx := context.TODO()
	_, err := repo.FindCardsByDeckID(ctx, deckMock.ID)
	assert.NoError(t, err)
}

func TestFindCardsByDeckIDScanFailed(t *testing.T) {
	db, mock := mock.NewMock()
	repo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	defer db.Close()

	query := fmt.Sprintf(regexp.QuoteMeta("SELECT uuid, deck_uuid, value, suit, code FROM %s WHERE deck_uuid = $1"), constant.TableCards)
	rows := sqlmock.NewRows([]string{"column"}).AddRow(nil)

	mock.ExpectQuery(query).WithArgs(deckMock.ID).WillReturnRows(rows)

	ctx := context.TODO()
	_, err := repo.FindCardsByDeckID(ctx, deckMock.ID)
	assert.Error(t, err)
}

func TestFindCardsByDeckIDPrepareFailed(t *testing.T) {
	db, mock := mock.NewMock()
	repo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	defer db.Close()

	query := fmt.Sprintf(regexp.QuoteMeta("SELECT uuid, deck_uuid, value, suit, code FROM %s WHERE deck_uuid = $1"), constant.TableCards)

	mock.ExpectQuery(query).WillReturnError(errors.New("an error occured"))

	ctx := context.TODO()
	_, err := repo.FindCardsByDeckID(ctx, deckMock.ID)
	assert.Error(t, err)
}

func TestFindCardsByDeckIDExecFailed(t *testing.T) {
	db, mock := mock.NewMock()
	repo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	defer db.Close()

	query := fmt.Sprintf(regexp.QuoteMeta("SELECT uuid, deck_uuid, value, suit, code FROM %s WHERE deck_uuid = $1"), constant.TableCards)

	mock.ExpectQuery(query).WithArgs(deckMock.ID).WillReturnError(errors.New("an error occured"))

	ctx := context.TODO()
	_, err := repo.FindCardsByDeckID(ctx, deckMock.ID)
	assert.Error(t, err)
}

func TestFindCardsWithLimitSuccess(t *testing.T) {
	db, mock := mock.NewMock()
	repo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	defer db.Close()

	query := fmt.Sprintf(regexp.QuoteMeta("SELECT uuid, deck_uuid, value, suit, code FROM %s WHERE deck_uuid = $1 LIMIT %d"), constant.TableCards, 1)
	rows := sqlmock.NewRows([]string{"uuid", "deck_uuid", "value", "suit", "code"}).AddRow("first-card", "some-id", "JACK", "SPADE", "JS")

	mock.ExpectQuery(query).WithArgs(deckMock.ID).WillReturnRows(rows)

	ctx := context.TODO()
	_, err := repo.FindCardsWithLimit(ctx, deckMock.ID, 1)
	assert.NoError(t, err)
}

func TestFindCardsWithLimitScanFailed(t *testing.T) {
	db, mock := mock.NewMock()
	repo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	defer db.Close()

	query := fmt.Sprintf(regexp.QuoteMeta("SELECT uuid, deck_uuid, value, suit, code FROM %s WHERE deck_uuid = $1 LIMIT %d"), constant.TableCards, 1)
	rows := sqlmock.NewRows([]string{"column"}).AddRow(nil)

	mock.ExpectQuery(query).WithArgs(deckMock.ID).WillReturnRows(rows)

	ctx := context.TODO()
	_, err := repo.FindCardsWithLimit(ctx, deckMock.ID, 1)
	assert.Error(t, err)
}

func TestFindCardsWithLimitPrepareFailed(t *testing.T) {
	db, mock := mock.NewMock()
	repo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	defer db.Close()

	query := fmt.Sprintf(regexp.QuoteMeta("SELECT uuid, deck_uuid, value, suit, code FROM %s WHERE deck_uuid = $1 LIMIT %d"), constant.TableCards, 1)

	mock.ExpectQuery(query).WillReturnError(errors.New("an error occured"))

	ctx := context.TODO()
	_, err := repo.FindCardsWithLimit(ctx, deckMock.ID, 1)
	assert.Error(t, err)
}

func TestFindCardsWithLimitExecFailed(t *testing.T) {
	db, mock := mock.NewMock()
	repo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	defer db.Close()

	query := fmt.Sprintf(regexp.QuoteMeta("SELECT uuid, deck_uuid, value, suit, code FROM %s WHERE deck_uuid = $1 LIMIT %d"), constant.TableCards, 1)

	mock.ExpectQuery(query).WithArgs(deckMock.ID).WillReturnError(errors.New("an error occured"))

	ctx := context.TODO()
	_, err := repo.FindCardsWithLimit(ctx, deckMock.ID, 1)
	assert.Error(t, err)
}

func TestDeleteCardsSuccess(t *testing.T) {
	db, mock := mock.NewMock()
	repo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	defer db.Close()

	query := fmt.Sprintf(regexp.QuoteMeta("DELETE FROM %s WHERE uuid IN ($1,$2,$3)"), constant.TableCards)
	mock.ExpectExec(query).WithArgs(cardsArgs...).WillReturnResult(sqlmock.NewResult(3, 3))

	ctx := context.TODO()
	err := repo.DeleteCards(ctx, cardsUUID)
	assert.NoError(t, err)
}

func TestDeleteCardsExecFailed(t *testing.T) {
	db, mock := mock.NewMock()
	repo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	defer db.Close()

	query := fmt.Sprintf(regexp.QuoteMeta("DELETE FROM %s WHERE uuid IN ($1,$2,$3)"), constant.TableCards)
	mock.ExpectExec(query).WithArgs(cardsArgs...).WillReturnError(errors.New("an error occured"))

	ctx := context.TODO()
	err := repo.DeleteCards(ctx, cardsUUID)
	assert.Error(t, err)
}

func TestDeleteCardsPrepareFailed(t *testing.T) {
	db, mock := mock.NewMock()
	repo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	defer db.Close()

	query := fmt.Sprintf(regexp.QuoteMeta("DELETE FROM %s WHERE uuid IN ($1,$2,$3)"), constant.TableCards)
	mock.ExpectExec(query).WillReturnError(errors.New("an error occured"))

	ctx := context.TODO()
	err := repo.DeleteCards(ctx, cardsUUID)
	assert.Error(t, err)
}
