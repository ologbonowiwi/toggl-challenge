package model

import (
	"errors"
	"math/rand"

	"github.com/google/uuid"
)

type Deck struct {
	ID        string `json:"id"`
	Cards     []Card `json:"cards"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

// errors
var (
	ErrNotEnoughCards = errors.New("not enough cards")
	ErrInvalidAmount  = errors.New("invalid amount")
)

func NewDeck() Deck {
	cards := make([]Card, 0, 52)

	for _, suit := range Suits {
		for _, value := range Values {
			card, _ := NewCard(value, suit)
			cards = append(cards, card)
		}
	}

	deck := Deck{
		ID:        uuid.NewString(),
		Cards:     cards,
		Shuffled:  false,
		Remaining: len(cards),
	}

	return deck
}

func (d *Deck) Shuffle() {
	for i := range d.Cards {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
	d.Shuffled = true
}

func (d *Deck) DrawCards(amount int) ([]Card, error) {
	if amount > d.Remaining {
		return nil, ErrNotEnoughCards
	}

	if amount < 0 {
		return nil, ErrInvalidAmount
	}

	cards := d.Cards[:amount]
	d.Cards = d.Cards[amount:]
	d.Remaining -= amount

	return cards, nil
}
