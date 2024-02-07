package storage_test

import (
	"sync"
	"testing"

	"github.com/ologbonowiwi/toggl-challenge/internal/model"
	"github.com/ologbonowiwi/toggl-challenge/internal/storage"
)

func TestLocalDeckRepositorySaveAndGetDeck(t *testing.T) {
	repo := storage.NewLocalDeckRepository()
	deck := model.NewDeck(false)

	if err := repo.SaveDeck(&deck); err != nil {
		t.Errorf("SaveDeck() error = %v, want nil", err)
	}

	if got := repo.GetDeck(deck.ID); got == nil {
		t.Errorf("GetDeck() = nil, want %v", deck)
	}
}

func TestLocalDeckRepositorySaveDeckInvalid(t *testing.T) {
	repo := storage.NewLocalDeckRepository()
	deck := model.NewDeck(false)
	deck.ID = "" // invalid deck

	if err := repo.SaveDeck(&deck); err == nil {
		t.Errorf("SaveDeck() error = nil, want error")
	}

	if err := repo.SaveDeck(nil); err == nil {
		t.Errorf("SaveDeck() error = nil, want error")
	}
}

func TestLocalDeckRepositoryGetDeckNotFound(t *testing.T) {
	repo := storage.NewLocalDeckRepository()
	deck := repo.GetDeck("invalid")
	if deck != nil {
		t.Errorf("GetDeck() = %v, want nil", deck)
	}
}

func TestLocalDeckRepositorySaveAndGetDeckConcurrently(t *testing.T) {
	repo := storage.NewLocalDeckRepository()
	ids := make(chan string, 100)

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			deck := model.NewDeck(false)
			repo.SaveDeck(&deck)
			ids <- deck.ID
		}()
	}

	go func() {
		wg.Wait()
		close(ids)
	}()

	for id := range ids {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			deck := repo.GetDeck(id)
			if deck == nil {
				t.Errorf("GetDeck() = nil, want %v", deck)
			}
		}(id)
	}

	wg.Wait()

	decks := repo.GetAllDecks()

	if len(decks) != 100 {
		t.Errorf("GetAllDecks() length = %d, want 100", len(decks))
	}
}
