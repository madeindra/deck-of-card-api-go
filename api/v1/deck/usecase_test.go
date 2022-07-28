package deck_test

import (
	"context"
	"errors"
	"testing"

	"github.com/madeindra/toggl-test/api/v1/deck"
	"github.com/madeindra/toggl-test/api/v1/deck/mocks"
	uuid "github.com/madeindra/toggl-test/internal/uuid/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockDeckData = deck.NewDeckData("someid", true)
var mockCardsData = []deck.CardData{deck.NewCardData("otherid", "ACE", "SPACE", "AS", "someid")}

func TestCreateDefaultSuccess(t *testing.T) {
	deckRepo := new(mocks.DeckRepo)
	uuid := new(uuid.UUID)

	uuid.On("NewString").Return("random-deck-id", nil)
	uuid.On("NewStringSlice", mock.AnythingOfType("int")).Return([]string{"first-card-id", "until", "fifty-second-card-id"})

	deckRepo.On("CreateDeck", mock.Anything, mock.AnythingOfType("deck.Deck")).Return(nil)
	deckRepo.On("CreateCards", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("[]string"), mock.AnythingOfType("[]deck.Card")).Return(nil)

	deckUsecase := deck.NewDeckUsecase(deckRepo, uuid)

	ctx := context.TODO()

	result := deckUsecase.Create(ctx, true, []string{})

	assert.Equal(t, 201, result.Code)

	deckRepo.AssertExpectations(t)
}

func TestCreateSuccess(t *testing.T) {
	deckRepo := new(mocks.DeckRepo)
	uuid := new(uuid.UUID)

	uuid.On("NewString").Return("random-deck-id", nil)
	uuid.On("NewStringSlice", mock.AnythingOfType("int")).Return([]string{"first-card-id", "second-card-id", "third-card-id", "fourth-card-id", "fifth-card-id", "sixth-card-id"})

	deckRepo.On("CreateDeck", mock.Anything, mock.AnythingOfType("deck.Deck")).Return(nil)
	deckRepo.On("CreateCards", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("[]string"), mock.AnythingOfType("[]deck.Card")).Return(nil)

	deckUsecase := deck.NewDeckUsecase(deckRepo, uuid)

	ctx := context.TODO()

	result := deckUsecase.Create(ctx, true, []string{"AC", "JS", "QH", "KD", "2C", "10S"})

	assert.Equal(t, 201, result.Code)

	deckRepo.AssertExpectations(t)
}

func TestCreateInvalidCode1Failed(t *testing.T) {
	deckRepo := new(mocks.DeckRepo)
	uuid := new(uuid.UUID)

	uuid.On("NewString").Return("random-deck-id", nil)
	uuid.On("NewStringSlice", mock.AnythingOfType("int")).Return([]string{"first-card-id"})

	deckUsecase := deck.NewDeckUsecase(deckRepo, uuid)

	ctx := context.TODO()

	result := deckUsecase.Create(ctx, true, []string{"X"})

	assert.Equal(t, 400, result.Code)
}

func TestCreateInvalidCode2Failed(t *testing.T) {
	deckRepo := new(mocks.DeckRepo)
	uuid := new(uuid.UUID)

	uuid.On("NewString").Return("random-deck-id", nil)
	uuid.On("NewStringSlice", mock.AnythingOfType("int")).Return([]string{"first-card-id"})

	deckUsecase := deck.NewDeckUsecase(deckRepo, uuid)

	ctx := context.TODO()

	result := deckUsecase.Create(ctx, true, []string{"XC"})

	assert.Equal(t, 400, result.Code)
}

func TestCreateInvalidCode3Failed(t *testing.T) {
	deckRepo := new(mocks.DeckRepo)
	uuid := new(uuid.UUID)

	uuid.On("NewString").Return("random-deck-id", nil)
	uuid.On("NewStringSlice", mock.AnythingOfType("int")).Return([]string{"first-card-id"})

	deckUsecase := deck.NewDeckUsecase(deckRepo, uuid)

	ctx := context.TODO()

	result := deckUsecase.Create(ctx, true, []string{"AX"})

	assert.Equal(t, 400, result.Code)
}

func TestCreateDeckFailed(t *testing.T) {
	deckRepo := new(mocks.DeckRepo)
	uuid := new(uuid.UUID)

	uuid.On("NewString").Return("random-deck-id", nil)
	uuid.On("NewStringSlice", mock.AnythingOfType("int")).Return([]string{"first-card-id", "second-card-id", "third-card-id", "fourth-card-id", "fifth-card-id", "sixth-card-id"})

	deckRepo.On("CreateDeck", mock.Anything, mock.AnythingOfType("deck.Deck")).Return(errors.New("an error occured"))

	deckUsecase := deck.NewDeckUsecase(deckRepo, uuid)

	ctx := context.TODO()

	result := deckUsecase.Create(ctx, true, []string{"AC", "JS", "QH", "KD", "2C", "10S"})

	assert.Equal(t, 500, result.Code)

	deckRepo.AssertExpectations(t)
}

func TestCreateCardFailed(t *testing.T) {
	deckRepo := new(mocks.DeckRepo)
	uuid := new(uuid.UUID)

	uuid.On("NewString").Return("random-deck-id", nil)
	uuid.On("NewStringSlice", mock.AnythingOfType("int")).Return([]string{"first-card-id", "second-card-id", "third-card-id", "fourth-card-id", "fifth-card-id", "sixth-card-id"})

	deckRepo.On("CreateDeck", mock.Anything, mock.AnythingOfType("deck.Deck")).Return(nil)
	deckRepo.On("CreateCards", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("[]string"), mock.AnythingOfType("[]deck.Card")).Return(errors.New("an error occured"))

	deckUsecase := deck.NewDeckUsecase(deckRepo, uuid)

	ctx := context.TODO()

	result := deckUsecase.Create(ctx, true, []string{"AC", "JS", "QH", "KD", "2C", "10S"})

	assert.Equal(t, 500, result.Code)

	deckRepo.AssertExpectations(t)
}

func TestFindByIDSuccess(t *testing.T) {
	deckRepo := new(mocks.DeckRepo)
	uuid := new(uuid.UUID)

	deckRepo.On("FindDeckByID", mock.Anything, mock.AnythingOfType("string")).Return(mockDeckData, nil)
	deckRepo.On("FindCardsByDeckID", mock.Anything, mock.AnythingOfType("string")).Return(mockCardsData, nil)

	deckUsecase := deck.NewDeckUsecase(deckRepo, uuid)

	ctx := context.TODO()

	result := deckUsecase.FindByID(ctx, "someid")

	assert.Equal(t, 200, result.Code)

	deckRepo.AssertExpectations(t)
}

func TestFindByIDFindDeckFailed(t *testing.T) {
	deckRepo := new(mocks.DeckRepo)
	uuid := new(uuid.UUID)

	deckRepo.On("FindDeckByID", mock.Anything, mock.AnythingOfType("string")).Return(deck.NewDeckData("", false), errors.New("an error occured"))

	deckUsecase := deck.NewDeckUsecase(deckRepo, uuid)

	ctx := context.TODO()

	result := deckUsecase.FindByID(ctx, "someid")

	assert.Equal(t, 404, result.Code)

	deckRepo.AssertExpectations(t)
}

func TestFindByIDInvalidDeckFailed(t *testing.T) {
	deckRepo := new(mocks.DeckRepo)
	uuid := new(uuid.UUID)

	deckRepo.On("FindDeckByID", mock.Anything, mock.AnythingOfType("string")).Return(deck.NewDeckData("", false), nil)

	deckUsecase := deck.NewDeckUsecase(deckRepo, uuid)

	ctx := context.TODO()

	result := deckUsecase.FindByID(ctx, "someid")

	assert.Equal(t, 400, result.Code)

	deckRepo.AssertExpectations(t)
}

func TestFindByIDFindCardsFailed(t *testing.T) {
	deckRepo := new(mocks.DeckRepo)
	uuid := new(uuid.UUID)

	deckRepo.On("FindDeckByID", mock.Anything, mock.AnythingOfType("string")).Return(mockDeckData, nil)
	deckRepo.On("FindCardsByDeckID", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("an error occured"))

	deckUsecase := deck.NewDeckUsecase(deckRepo, uuid)

	ctx := context.TODO()

	result := deckUsecase.FindByID(ctx, "someid")

	assert.Equal(t, 404, result.Code)

	deckRepo.AssertExpectations(t)
}

func TestDrawSuccess(t *testing.T) {
	deckRepo := new(mocks.DeckRepo)
	uuid := new(uuid.UUID)

	deckRepo.On("FindDeckByID", mock.Anything, mock.AnythingOfType("string")).Return(mockDeckData, nil)
	deckRepo.On("FindCardsWithLimit", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(mockCardsData, nil)
	deckRepo.On("DeleteCards", mock.Anything, mock.AnythingOfType("[]string")).Return(nil)

	deckUsecase := deck.NewDeckUsecase(deckRepo, uuid)

	ctx := context.TODO()

	result := deckUsecase.Draw(ctx, "someid", 1)

	assert.Equal(t, 200, result.Code)

	deckRepo.AssertExpectations(t)
}

func TestDrawInvalidCountFailed(t *testing.T) {
	deckRepo := new(mocks.DeckRepo)
	uuid := new(uuid.UUID)

	deckUsecase := deck.NewDeckUsecase(deckRepo, uuid)

	ctx := context.TODO()

	result := deckUsecase.Draw(ctx, "someid", 0)

	assert.Equal(t, 400, result.Code)

	deckRepo.AssertExpectations(t)
}

func TestDrawFindDeckFailed(t *testing.T) {
	deckRepo := new(mocks.DeckRepo)
	uuid := new(uuid.UUID)

	deckRepo.On("FindDeckByID", mock.Anything, mock.AnythingOfType("string")).Return(deck.NewDeckData("", false), errors.New("an error occured"))

	deckUsecase := deck.NewDeckUsecase(deckRepo, uuid)

	ctx := context.TODO()

	result := deckUsecase.Draw(ctx, "someid", 1)

	assert.Equal(t, 404, result.Code)

	deckRepo.AssertExpectations(t)
}

func TestDrawInvalidDeckFailed(t *testing.T) {
	deckRepo := new(mocks.DeckRepo)
	uuid := new(uuid.UUID)

	deckRepo.On("FindDeckByID", mock.Anything, mock.AnythingOfType("string")).Return(deck.NewDeckData("", false), nil)

	deckUsecase := deck.NewDeckUsecase(deckRepo, uuid)

	ctx := context.TODO()

	result := deckUsecase.Draw(ctx, "someid", 1)

	assert.Equal(t, 400, result.Code)

	deckRepo.AssertExpectations(t)
}

func TestDrawFindCardsFailed(t *testing.T) {
	deckRepo := new(mocks.DeckRepo)
	uuid := new(uuid.UUID)

	deckRepo.On("FindDeckByID", mock.Anything, mock.AnythingOfType("string")).Return(mockDeckData, nil)
	deckRepo.On("FindCardsWithLimit", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(nil, errors.New("an error occured"))

	deckUsecase := deck.NewDeckUsecase(deckRepo, uuid)

	ctx := context.TODO()

	result := deckUsecase.Draw(ctx, "someid", 1)

	assert.Equal(t, 404, result.Code)

	deckRepo.AssertExpectations(t)
}

func TestDrawDeleteFailed(t *testing.T) {
	deckRepo := new(mocks.DeckRepo)
	uuid := new(uuid.UUID)

	deckRepo.On("FindDeckByID", mock.Anything, mock.AnythingOfType("string")).Return(mockDeckData, nil)
	deckRepo.On("FindCardsWithLimit", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(mockCardsData, nil)
	deckRepo.On("DeleteCards", mock.Anything, mock.AnythingOfType("[]string")).Return(errors.New("an error occured"))

	deckUsecase := deck.NewDeckUsecase(deckRepo, uuid)

	ctx := context.TODO()

	result := deckUsecase.Draw(ctx, "someid", 1)

	assert.Equal(t, 500, result.Code)

	deckRepo.AssertExpectations(t)
}
