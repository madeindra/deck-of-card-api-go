package deck

// standard deck
var standardDeck []string = []string{
	"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "10S", "JS", "QS", "KS",
	"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "10D", "JD", "QD", "KD",
	"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "10C", "JC", "QC", "KC",
	"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "10H", "JH", "QH", "KH",
}

// properties of deck
type Deck struct {
	ID        string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
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
	Remaining int    `json:"remaining"`
	Cards     []Card `json:"cards"`
}

// properties of cards
type Cards struct {
	Cards []Card `json:"cards"`
}

// column of Deck in Table
type DeckData struct {
	uuid     string
	shuffled bool
}

// column of Deck in Table
type CardData struct {
	uuid      string
	value     string
	suit      string
	code      string
	deck_uuid string
}

// for repository testing purpose
func NewDeckData(uuid string, shuffled bool) DeckData {
	return DeckData{
		uuid:     uuid,
		shuffled: shuffled,
	}
}

// for repository testing purpose
func NewCardData(uuid string, value string, suit string, code string, deck_uuid string) CardData {
	return CardData{
		uuid:      uuid,
		value:     value,
		suit:      suit,
		code:      code,
		deck_uuid: deck_uuid,
	}
}
