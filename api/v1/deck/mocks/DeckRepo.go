// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	deck "github.com/madeindra/toggl-test/api/v1/deck"
	mock "github.com/stretchr/testify/mock"
)

// DeckRepo is an autogenerated mock type for the DeckRepo type
type DeckRepo struct {
	mock.Mock
}

// CreateCards provides a mock function with given fields: ctx, deckUUID, cardUUIDS, cards
func (_m *DeckRepo) CreateCards(ctx context.Context, deckUUID string, cardUUIDS []string, cards []deck.Card) error {
	ret := _m.Called(ctx, deckUUID, cardUUIDS, cards)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, []deck.Card) error); ok {
		r0 = rf(ctx, deckUUID, cardUUIDS, cards)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateDeck provides a mock function with given fields: ctx, _a1
func (_m *DeckRepo) CreateDeck(ctx context.Context, _a1 deck.Deck) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, deck.Deck) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCards provides a mock function with given fields: ctx, uuids
func (_m *DeckRepo) DeleteCards(ctx context.Context, uuids []string) error {
	ret := _m.Called(ctx, uuids)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) error); ok {
		r0 = rf(ctx, uuids)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindCardsByDeckID provides a mock function with given fields: ctx, deckUUID
func (_m *DeckRepo) FindCardsByDeckID(ctx context.Context, deckUUID string) ([]deck.CardData, error) {
	ret := _m.Called(ctx, deckUUID)

	var r0 []deck.CardData
	if rf, ok := ret.Get(0).(func(context.Context, string) []deck.CardData); ok {
		r0 = rf(ctx, deckUUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]deck.CardData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, deckUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindCardsWithLimit provides a mock function with given fields: ctx, deckUUID, count
func (_m *DeckRepo) FindCardsWithLimit(ctx context.Context, deckUUID string, count int) ([]deck.CardData, error) {
	ret := _m.Called(ctx, deckUUID, count)

	var r0 []deck.CardData
	if rf, ok := ret.Get(0).(func(context.Context, string, int) []deck.CardData); ok {
		r0 = rf(ctx, deckUUID, count)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]deck.CardData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int) error); ok {
		r1 = rf(ctx, deckUUID, count)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindDeckByID provides a mock function with given fields: ctx, deckUUID
func (_m *DeckRepo) FindDeckByID(ctx context.Context, deckUUID string) (deck.DeckData, error) {
	ret := _m.Called(ctx, deckUUID)

	var r0 deck.DeckData
	if rf, ok := ret.Get(0).(func(context.Context, string) deck.DeckData); ok {
		r0 = rf(ctx, deckUUID)
	} else {
		r0 = ret.Get(0).(deck.DeckData)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, deckUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
