package DeckOfCardAPI

import (
	"testing"
)

func TestCardString(t *testing.T) {
	cards := []Card{
		{Rank: Ace,Suit: Spade},
		{Suit: Joker},
	}

	expected := []string{
		"Ace of Spades\n",
		"Joker\n",
	}
	for i, val := range cards{
		if val.String() != expected[i]{
			t.Errorf("Unexpected result. Got: %s, Expected: %s", val, expected)
		}
	}

}


func TestFilter(t *testing.T) {
	deck := NewDeck()
	deck.Filter(Two, Jack)
	for _, card := range(deck) {
		if card.Rank == Jack || card.Rank == Two {
			t.Error("Expected no Two or Jacks")
		}
	}
	if len(deck) != 44 {
		t.Error("Expected 44 cards, got", len(deck))
	}
}