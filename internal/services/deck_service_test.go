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

	deck := service.NewDeck(false)
	if deck == nil {
		t.Errorf("NewDeck() = nil, want %v", deck)
	}
}

func TestDeckServiceDrawCards(t *testing.T) {
	repo := storage.NewLocalDeckRepository()
	service := services.NewDeckService(repo)

	deck := service.NewDeck(false)
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

	deck := service.NewDeck(false)
	got := service.GetDeck(deck.ID)
	if got == nil {
		t.Errorf("GetDeck() = nil, want %v", deck)
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

	deck := service.NewDeck(false)
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

	deck := service.NewDeck(false)
	cards, err := service.DrawCards(deck.ID, -1)
	if cards != nil {
		t.Errorf(drawCardsWantNil, cards)
	}

	if err != model.ErrInvalidAmount {
		t.Errorf("DrawCards() error = %v, want %v", err, model.ErrInvalidAmount)
	}
}
