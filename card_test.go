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

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	expectedCard := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expectedCard {
		t.Errorf("Expected %s as first card. Received %s", expectedCard, cards[0])
	}
}
func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	expectedCard := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expectedCard {
		t.Errorf("Expected %s as first card. Received %s", expectedCard, cards[0])
	}
}

func TestShuffle(t *testing.T) {
	controlDeck := New(DefaultSort)
	testDeck := New(DefaultSort, Shuffle)
	for i, card := range testDeck {
		controlCard := controlDeck[i]
		if absRank(controlCard) != absRank(card) {
			return
		}
	}
	t.Error("Shuffled deck was in the same order as DefaultSort'd deck.")
}

func TestJokers(t *testing.T) {
	expectedJokerCount := 4
	cards := New(Jokers(expectedJokerCount))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != expectedJokerCount {
		t.Errorf("Expected %d Jokers. Found %d.", expectedJokerCount, count)
	}
}

func TestFilter(t *testing.T) {
	filterCards := []Rank{Ace, Jack, Queen, King}
	filterFn := func(card Card) bool {
		for _, rank := range filterCards {
			if card.Rank == rank {
				return true
			}
		}
		return false
	}
	cards := New(Filter(filterFn))
	for _, card := range cards {
		for _, rank := range filterCards {
			if card.Rank == rank {
				t.Errorf("Expected rank %s to be filtered from deck.", card.Rank.String())
			}
		}
	}
}

func TestDeck(t *testing.T) {
	expectedDeckCopies := 3
	controlDeck := New()
	actualDeck := New(Deck(expectedDeckCopies))
	if len(actualDeck) != expectedDeckCopies*len(controlDeck) {
		t.Errorf("Expected deck to be %d times larger than default. Got %d.", expectedDeckCopies, len(actualDeck)/len(controlDeck))
	}
}
