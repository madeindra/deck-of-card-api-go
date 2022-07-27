package deck

// standard deck
var standardDeck [52]string = [52]string{
	"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "10H", "JH", "QH", "KH",
	"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "10D", "JD", "QD", "KD",
	"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "10C", "JC", "QC", "KC",
	"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "10S", "JS", "QS", "KS",
}

// card suites
var cardSuite map[string]string = map[string]string{"H": "HEARTS", "D": "DIAMONDS", "C": "CLUBS", "S": "SPADES"}

// non-number card value
var cardValue map[string]string = map[string]string{"A": "ACE", "J": "JACK", "Q": "QUEEN", "K": "KING"}

// properties of deck
type Deck struct {
	ID        string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int64  `json:"remaining"`
}

// properties of card
type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

// properties of deck with cards
type DeckWithCards struct {
	ID        string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int64  `json:"remaining"`
	Cards     []Card `json:"cards"`
}
