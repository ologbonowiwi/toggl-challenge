package storage

import (
	"errors"
	"sync"

	"github.com/ologbonowiwi/toggl-challenge/internal/model"
)

// LocalDeckRepository is an in-memory storage for decks.
// It is thread-safe, as it uses a mutex to control read and write operations.
type LocalDeckRepository struct {
	mu    sync.RWMutex
	decks map[string]*model.Deck
}

func NewLocalDeckRepository() *LocalDeckRepository {
	return &LocalDeckRepository{
		decks: make(map[string]*model.Deck),
	}
}

// GetDeck retrieves a deck by its ID.
func (r *LocalDeckRepository) GetDeck(id string) *model.Deck {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if deck, ok := r.decks[id]; ok {
		return deck
	}
	return nil
}

// SaveDeck stores or updates a deck.
func (r *LocalDeckRepository) SaveDeck(deck *model.Deck) error {
	if deck == nil || deck.ID == "" {
		return errors.New("invalid deck to save")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.decks[deck.ID] = deck
	return nil
}

// GetAllDecks retrieves all decks.
func (r *LocalDeckRepository) GetAllDecks() []*model.Deck {
	r.mu.RLock()
	defer r.mu.RUnlock()

	decks := make([]*model.Deck, 0, len(r.decks))
	for _, deck := range r.decks {
		decks = append(decks, deck)
	}
	return decks
}
