package deck

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func shuffle(arr []string) []string {
	rand.Seed(time.Now().Unix())

	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	return arr
}

func decode(codes []string) ([]Card, error) {
	// initialize empty slice
	cardResult := []Card{}

	// split into 2 characters
	for _, c := range codes {
		parts := strings.Split(c, "")

		// alow at max 3 chars
		if len(parts) < 2 || len(parts) > 3 {
			return []Card{}, errors.New("Invalid Card Codes")
		}

		if len(parts) == 3 {
			// combine first and second parts into one string
			parts[0] += parts[1]

			// swap second parts with third parts
			parts[1] = parts[2]

			// trim last element
			parts = parts[:2]
		}

		// create empty card, value, and suite
		currentCard := Card{}
		value := ""
		suit := ""

		// first char is the value
		valueInt, err := strconv.Atoi(parts[0])

		// allow only 2-10 for integer value
		if valueInt < 11 && valueInt > 1 {
			value = parts[0]
		}

		// othern than that, check if it is ace or face card
		if err != nil {
			switch parts[0] {
			case "A":
				value = "ACE"
			case "J":
				value = "JACK"
			case "Q":
				value = "QUEEN"
			case "K":
				value = "KING"
			default:
				return []Card{}, errors.New("Invalid Card Codes")
			}
		}

		// second char is the suite
		switch parts[1] {
		case "H":
			suit = "HEARTS"
		case "D":
			suit = "DIAMONDS"
		case "C":
			suit = "CLUBS"
		case "S":
			suit = "SPADES"
		default:
			return []Card{}, errors.New("Invalid Card Codes")
		}

		currentCard.Code = c
		currentCard.Value = value
		currentCard.Suit = suit

		cardResult = append(cardResult, currentCard)
	}

	return cardResult, nil
}
