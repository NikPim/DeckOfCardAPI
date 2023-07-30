package DeckOfCardAPI

import (
	"fmt"
	"math/rand"
	"time"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [4] Suit {Spade, Diamond, Club, Heart}

type Rank uint8

const (
	Ace Rank = iota + 1
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

var (
	min_rank = Ace
	max_rank = King
)

type Card struct {
	Suit
	Rank
}

func (s Suit) String() string{
	switch s{
	case 0:
		return "Spade"
	case 1:
		return "Diamond"
	case 2:
		return "Club"
	case 3:
		return "Heart"
	case 4:
		return "Joker\n"
	default:
		return "Error in suits"
	}
}

func (r Rank) String() string{
	switch r{
	case 1:
		return "Ace"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	case 10:
		return "Ten"
	case 11:
		return "Jack"
	case 12:
		return "Queen"
	case 13:
		return "King"
	default:
		return "Error in ranks"
	}
}

func (c Card) String() string {
	if c.Suit == Joker{
		return c.Suit.String()
	} else{
		return fmt.Sprintf("%s of %ss\n", c.Rank.String(), c.Suit.String())
	}
}

type Deck []Card

func NewDeck(options ...func(Deck) Deck) Deck {
	var d Deck
	for _, suit := range suits{
		for rank := min_rank; rank <= max_rank; rank++{
			d = append(d, Card{Suit: suit, Rank: rank})
		}
	}
	for _, option := range options{
		d = option(d)
	}
	return d
}

func (d Deck) Len() int {
	return len(d)
}

func (d Deck) Less(i, j int) bool {
	// Customize the sorting logic here
	if d[i].Suit != d[j].Suit {
		return d[i].Suit < d[j].Suit
	}
	// If ID values are the same, use Name as tiebreaker
	return d[i].Rank < d[j].Rank
}

func (d Deck) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d *Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	n := len(*d)
	for i := n-1; i > 0; i-- {
		j := r.Intn(i+1)
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	}
}

func (d *Deck) AddJoker(n int) {
		for i:=0; i<n; i++{
			*d = append(*d, Card{
				Suit: Joker,
				Rank: Rank(i),
			})
		}
	}

func (d *Deck) Filter(ranks ...Rank) {
	var filteredDeck Deck
	for _, card := range *d {
		found := false
		for _, rank := range ranks{
			if card.Rank == rank {
				found = true
				break
			}
		}
		if !found {
			filteredDeck = append(filteredDeck, card)
		}
	}
	*d = filteredDeck
}
