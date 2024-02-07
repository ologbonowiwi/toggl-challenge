package services_test

import (
	"testing"

	"github.com/ologbonowiwi/toggl-challenge/internal/model"
	"github.com/ologbonowiwi/toggl-challenge/internal/services"
	"github.com/ologbonowiwi/toggl-challenge/internal/storage"
)

const drawCardsWantNil = "DrawCards() = %v, want nil"

func TestDeckServiceNewDeck(t *testing.T) {
	repo := storage.NewLocalDeckRepository()
	service := services.NewDeckService(repo)

	deck := service.NewDeck(false, []string{})
	if deck == nil {
		t.Errorf("NewDeck() = nil, want %v", deck)
	}
}

func TestDeckServiceNewDeckFiltered(t *testing.T) {
	repo := storage.NewLocalDeckRepository()
	service := services.NewDeckService(repo)

	codes := make([]string, 0, 52)

	for _, suit := range model.Suits {
		for _, value := range model.Values {
			card, _ := model.NewCard(value, suit)
			codes = append(codes, card.Code)
		}
	}

	tests := []struct {
		name  string
		codes []string
		want  int
	}{
		{"empty", []string{}, 52},
		{"one", []string{"AS"}, 1},
		{"two", []string{"AS", "2S"}, 2},
		{"duplicated", []string{"AS", "AS"}, 1},
		{"all", codes, 52},
		{"all twice", append(codes, codes...), 52},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deck := service.NewDeck(false, tt.codes)
			if len(deck.Cards) != tt.want {
				t.Errorf("len(NewDeck()) = %d, want %d", len(deck.Cards), tt.want)
			}
		})
	}
}

func TestDeckServiceNewShuffledDeck(t *testing.T) {
	repo := storage.NewLocalDeckRepository()
	service := services.NewDeckService(repo)

	deck := service.NewDeck(true, []string{})

	if deck.Shuffled != true {
		t.Errorf("ShuffleDeck() = %t, want true", deck.Shuffled)
	}
}

func TestDeckServiceDrawCards(t *testing.T) {
	repo := storage.NewLocalDeckRepository()
	service := services.NewDeckService(repo)

	deck := service.NewDeck(false, []string{})
	cards, err := service.DrawCards(deck.ID, 5)
	if err != nil {
		t.Errorf("DrawCards() error = %v, want nil", err)
	}

	if len(cards) != 5 {
		t.Errorf("len(DrawCards()) = %d, want 5", len(cards))
	}
}

func TestDeckServiceGetDeck(t *testing.T) {
	repo := storage.NewLocalDeckRepository()
	service := services.NewDeckService(repo)

	deck := service.NewDeck(false, []string{})
	got, err := service.GetDeck(deck.ID)
	if err != nil {
		t.Errorf("GetDeck() error = %v, want nil", err)
	}
	if got == nil {
		t.Errorf("GetDeck() = nil, want %v", deck)
	}
}

func TestDeckServiceGetDeckNotFound(t *testing.T) {
	repo := storage.NewLocalDeckRepository()
	service := services.NewDeckService(repo)

	deck, err := service.GetDeck("invalid")

	if err == nil {
		t.Errorf("GetDeck() error = nil, want error")
	}

	if deck != nil {
		t.Errorf("GetDeck() = %v, want nil", deck)
	}
}

func TestDeckServiceDrawCardsNotFound(t *testing.T) {
	repo := storage.NewLocalDeckRepository()
	service := services.NewDeckService(repo)

	cards, err := service.DrawCards("invalid", 5)
	if cards != nil {
		t.Errorf(drawCardsWantNil, cards)
	}

	if err == nil {
		t.Errorf("DrawCards() error = nil, want error")
	}
}

func TestDeckServiceDrawCardsError(t *testing.T) {
	repo := storage.NewLocalDeckRepository()
	service := services.NewDeckService(repo)

	deck := service.NewDeck(false, []string{})
	cards, err := service.DrawCards(deck.ID, 53)
	if cards != nil {
		t.Errorf(drawCardsWantNil, cards)
	}

	if err != model.ErrNotEnoughCards {
		t.Errorf("DrawCards() error = %v, want %v", err, model.ErrNotEnoughCards)
	}
}

func TestDeckServiceDrawCardsInvalidAmount(t *testing.T) {
	repo := storage.NewLocalDeckRepository()
	service := services.NewDeckService(repo)

	deck := service.NewDeck(false, []string{})
	cards, err := service.DrawCards(deck.ID, -1)
	if cards != nil {
		t.Errorf(drawCardsWantNil, cards)
	}

	if err != model.ErrInvalidAmount {
		t.Errorf("DrawCards() error = %v, want %v", err, model.ErrInvalidAmount)
	}
}
