package services

import (
	"errors"

	"github.com/ologbonowiwi/toggl-challenge/internal/model"
)

var (
	ErrDeckNotFound = errors.New("deck not found")
)

type deckRepository interface {
	GetDeck(id string) *model.Deck
	SaveDeck(deck *model.Deck) error
}

type DeckService struct {
	deckRepository deckRepository
}

func NewDeckService(deckRepository deckRepository) *DeckService {
	return &DeckService{deckRepository: deckRepository}
}

func (ds *DeckService) CreateDeck(shuffled bool, codes []string) *model.Deck {
	deck := model.NewDeck()

	if shuffled {
		deck.Shuffle()
	}

	if len(codes) > 0 {
		codesMap := make(map[string]bool)
		for _, code := range codes {
			codesMap[code] = true
		}

		filtered := make([]model.Card, 0, len(deck.Cards))
		for _, card := range deck.Cards {
			if _, exists := codesMap[card.Code]; exists {
				filtered = append(filtered, card)
			}
		}

		deck.Cards = filtered
		deck.Remaining = len(filtered)
	}

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

func (ds *DeckService) GetDeck(id string) (*model.Deck, error) {
	deck := ds.deckRepository.GetDeck(id)

	if deck == nil {
		return nil, ErrDeckNotFound
	}

	return deck, nil
}
