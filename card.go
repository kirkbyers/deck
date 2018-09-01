//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
)

// Suit is the "family" of a card
type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

const (
	minSuit = Spade
	maxSuit = Heart
)

// Rank is the value of a card.
type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

// Card is a card in the deck. Has a Suit and Rank
type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

// New creates a new deck of Cards
func New() []Card {
	var cards []Card
	for suit := minSuit; suit <= maxSuit; suit++ {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Rank: rank, Suit: suit})
		}
	}
	return cards
}
