package model_test

import (
	"fmt"
	"testing"

	"github.com/ologbonowiwi/toggl-challenge/internal/model"
)

const (
	wantFullDeck     = "len(deck.Cards) = %d, want 52"
	wantShuffledDeck = "deck.Shuffled = %t, want true"
)

func TestNewDeck(t *testing.T) {
	deck := model.NewDeck()
	if len(deck.Cards) != 52 {
		t.Errorf("NewDeck() length = %d, want 52", len(deck.Cards))
	}

	if deck.Shuffled != false {
		t.Errorf("NewDeck().Shuffled = %t, want false", deck.Shuffled)
	}

	if deck.Remaining != 52 {
		t.Errorf("NewDeck().Remaining = %d, want 52", deck.Remaining)
	}
}

func TestNewDeckOrder(t *testing.T) {
	deck := model.NewDeck()
	for i, card := range deck.Cards {
		if card.Suit != model.Suits[i/13] {
			t.Errorf("NewDeck() = %s, want %s", card.Suit, model.Suits[i/13])
		}

		if card.Value != model.Values[i%13] {
			t.Errorf("NewDeck() = %s, want %s", card.Value, model.Values[i%13])
		}
	}
}

func TestNewDeckShuffledOrder(t *testing.T) {
	deck := model.NewDeck()
	if len(deck.Cards) != 52 {
		t.Errorf(wantFullDeck, len(deck.Cards))
	}

	if deck.Shuffled {
		t.Errorf("NewDeck().Shuffled = %t, want false", deck.Shuffled)
	}

	deck.Shuffle()

	if !deck.Shuffled {
		t.Errorf(wantShuffledDeck, deck.Shuffled)
	}

	// starts considering the decks are equal
	// it should be proven wrong on comparison
	equals := true
	d := model.NewDeck()

	for i, card := range deck.Cards {
		// if any card is different, the decks are different
		if card != d.Cards[i] {
			equals = false
			break
		}
	}

	if equals {
		t.Errorf("NewDeck() = %v, want different", deck.Cards)
	}
}

func TestNewDeckValuePerSuit(t *testing.T) {
	deck := model.NewDeck()

	for _, suit := range model.Suits {
		count := 0
		for _, card := range deck.Cards {
			if card.Suit == suit {
				count++
			}
		}

		if count != 13 {
			t.Errorf("NewDeck() = %d, want 13", count)
		}
	}
}

func TestNewDeckAmountOfSuit(t *testing.T) {
	deck := model.NewDeck()

	for _, value := range model.Values {
		count := 0
		for _, card := range deck.Cards {
			if card.Value == value {
				count++
			}
		}

		if count != 4 {
			t.Errorf("NewDeck() = %d, want 4", count)
		}
	}
}

func TestShuffle(t *testing.T) {
	deck := model.NewDeck()
	deck.Shuffle()

	if !deck.Shuffled {
		t.Errorf(wantShuffledDeck, deck.Shuffled)
	}

	if len(deck.Cards) != 52 {
		t.Errorf(wantFullDeck, len(deck.Cards))
	}

	if !deck.Shuffled {
		t.Errorf(wantShuffledDeck, deck.Shuffled)
	}

	// starts considering the decks are equal
	// it should be proven wrong on comparison
	equals := true
	d := model.NewDeck()

	for i, card := range deck.Cards {
		// if any card is different, the decks are different
		if card != d.Cards[i] {
			equals = false
			break
		}
	}

	if equals {
		t.Errorf("NewDeck() = %v, want different", deck.Cards)
	}
}

func TestDrawCards(t *testing.T) {
	tests := []struct {
		name     string
		amount   int
		expected int
	}{
		{"draw 1 card", 1, 51},
		{"draw 5 cards", 5, 47},
		{"draw 52 cards", 52, 0},
		{"draw 0 cards", 0, 52},
	}

	for _, test := range tests {
		for _, shuffle := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s, shuffle=%v", test.name, shuffle), func(t *testing.T) {
				deck := model.NewDeck()
				_, err := deck.DrawCards(test.amount)

				if err != nil {
					t.Errorf("DrawCards() error = %v, want nil", err)
				}

				if deck.Remaining != test.expected {
					t.Errorf("deck.Remaining = %d, want %d", deck.Remaining, test.expected)
				}

				if len(deck.Cards) != test.expected {
					t.Errorf("len(deck.Cards) = %d, want %d", len(deck.Cards), test.expected)
				}
			})
		}
	}
}

func TestDrawCardsError(t *testing.T) {
	tests := []struct {
		name   string
		amount int
		err    error
	}{
		{"draw 53 cards", 53, model.ErrNotEnoughCards},
		{"draw -1 cards", -1, model.ErrInvalidAmount},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			deck := model.NewDeck()
			_, err := deck.DrawCards(test.amount)

			if err != test.err {
				t.Errorf("DrawCards() error = %v, want %v", err, test.err)
			}

			if err == nil {
				t.Errorf("DrawCards() error = nil, want error")
			}

			if deck.Remaining != 52 {
				t.Errorf("deck.Remaining = %d, want 52", deck.Remaining)
			}

			if len(deck.Cards) != 52 {
				t.Errorf(wantFullDeck, len(deck.Cards))
			}
		})
	}
}

func TestDrawAllCardsTwice(t *testing.T) {
	deck := model.NewDeck()
	_, err := deck.DrawCards(52)

	if err != nil {
		t.Errorf("DrawCards() error = %v, want nil", err)
	}

	_, err = deck.DrawCards(52)

	if err != model.ErrNotEnoughCards {
		t.Errorf("DrawCards() error = %v, want %v", err, model.ErrNotEnoughCards)
	}

	if err == nil {
		t.Errorf("DrawCards() error = nil, want error")
	}

	if deck.Remaining != 0 {
		t.Errorf("deck.Remaining = %d, want 0", deck.Remaining)
	}

	if len(deck.Cards) != 0 {
		t.Errorf("len(deck.Cards) = %d, want 0", len(deck.Cards))
	}
}
