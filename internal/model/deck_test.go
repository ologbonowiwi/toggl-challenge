package model_test

import (
	"testing"

	"github.com/ologbonowiwi/toggl-challenge/internal/model"
)

func TestNewDeck(t *testing.T) {
	deck := model.NewDeck(false)
	if len(deck.Cards) != 52 {
		t.Errorf("NewDeck() length = %d, want 52", len(deck.Cards))
	}

	if deck.Suffled != false {
		t.Errorf("NewDeck().Suffled = %t, want false", deck.Suffled)
	}

	if deck.Remaining != 52 {
		t.Errorf("NewDeck().Remaining = %d, want 52", deck.Remaining)
	}
}

func TestNewDeckOrder(t *testing.T) {
	deck := model.NewDeck(false)
	for i, card := range deck.Cards {
		if card.Suit != model.Suits[i/13] {
			t.Errorf("NewDeck() = %s, want %s", card.Suit, model.Suits[i/13])
		}

		if card.Value != model.Values[i%13] {
			t.Errorf("NewDeck() = %s, want %s", card.Value, model.Values[i%13])
		}
	}
}

func TestNewDeckSuffledOrder(t *testing.T) {
	deck := model.NewDeck(true)
	if len(deck.Cards) != 52 {
		t.Errorf("NewDeck() = %d, want 52", len(deck.Cards))
	}

	if deck.Suffled != true {
		t.Errorf("NewDeck() = %t, want true", deck.Suffled)
	}

	deck2 := model.NewDeck(false)

	// compare the whole deck
	equals := true

	for i, card := range deck.Cards {
		// if any card is different, the decks are different
		if card != deck2.Cards[i] {
			equals = false
			break
		}
	}

	if equals {
		t.Errorf("NewDeck() = %v, want different", deck.Cards)
	}
}

func TestNewDeckValuePerSuit(t *testing.T) {
	deck := model.NewDeck(false)

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
	deck := model.NewDeck(false)

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
