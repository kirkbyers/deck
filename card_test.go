package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Diamond})
	fmt.Println(Card{Rank: Nine, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Diamonds
	// Nine of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	expectedCardsLen := 13 * 4
	cards := New()
	// 13 ranks * 4 suits
	if len(cards) != expectedCardsLen {
		t.Errorf("Wrong number of cards in a new deck. Got %d, expected %d", len(cards), expectedCardsLen)
	}
}
