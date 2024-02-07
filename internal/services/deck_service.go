package services

import (
	"errors"

	"github.com/ologbonowiwi/toggl-challenge/internal/model"
)

var (
	ErrDeckNotFound = errors.New("deck not found")
)

type DeckRepository interface {
	GetDeck(id string) *model.Deck
	SaveDeck(deck *model.Deck) error
}

type DeckService struct {
	deckRepository DeckRepository
}

func NewDeckService(deckRepository DeckRepository) *DeckService {
	return &DeckService{deckRepository: deckRepository}
}

func (ds *DeckService) NewDeck(shuffled bool) *model.Deck {
	deck := model.NewDeck(shuffled)
	ds.deckRepository.SaveDeck(&deck)
	return &deck
}

func (ds *DeckService) DrawCards(id string, amount int) ([]model.Card, error) {
	deck := ds.deckRepository.GetDeck(id)
	if deck == nil {
		return nil, ErrDeckNotFound
	}
	cards, err := deck.DrawCards(amount)
	if err != nil {
		return nil, err
	}
	ds.deckRepository.SaveDeck(deck)
	return cards, nil
}

func (ds *DeckService) GetDeck(id string) *model.Deck {
	return ds.deckRepository.GetDeck(id)
}
