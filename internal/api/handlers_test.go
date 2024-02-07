package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/ologbonowiwi/toggl-challenge/cmd/app"
	"github.com/ologbonowiwi/toggl-challenge/internal/model"
	"github.com/stretchr/testify/assert"
)

const decksRoute = "/api/decks/"

func TestAPIHandlers(t *testing.T) {
	app := app.SetupApp()

	var id string

	t.Run("CreateDeck", func(t *testing.T) {
		for _, shuffled := range []bool{true, false} {
			query := "shuffle=" + strconv.FormatBool(shuffled)
			req := httptest.NewRequest(http.MethodPost, "/api/decks?"+query, nil)
			req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			var createDeckResponse struct {
				DeckID string `json:"deck_id"`
			}
			err = json.NewDecoder(resp.Body).Decode(&createDeckResponse)
			assert.NoError(t, err)
			id = createDeckResponse.DeckID
		}
	})

	t.Run("CreateDeckFiltered", func(t *testing.T) {
		for _, shuffled := range []bool{true, false} {
			query := "shuffle=" + strconv.FormatBool(shuffled) + "&codes=AS,2S,3S,4S,5S,6S,7S,8S,9S,10S,JS,QS,KS"
			req := httptest.NewRequest(http.MethodPost, "/api/decks?"+query, nil)
			req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)
		}
	})

	t.Run("DrawCards", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, decksRoute+id+"/draw?amount=2", nil)
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var drawResponse struct {
			Cards []struct {
				Value string `json:"value"`
				Suit  string `json:"suit"`
				Code  string `json:"code"`
			} `json:"cards"`
		}
		err = json.NewDecoder(resp.Body).Decode(&drawResponse)
		assert.NoError(t, err)
		assert.Len(t, drawResponse.Cards, 2) // Assert that 2 cards were drawn
	})

	t.Run("DrawCards ErrNotEnoughCards", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, decksRoute+id+"/draw?amount=53", nil)
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
	})

	t.Run("DrawCards ErrDeckNotFound", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/decks/invalid-id/draw?amount=1", nil)
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("GetDeck", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, decksRoute+id, nil)
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var getDeckResponse model.Deck
		err = json.NewDecoder(resp.Body).Decode(&getDeckResponse)
		assert.NoError(t, err)
		assert.Equal(t, id, getDeckResponse.ID)
	})

	t.Run("GetDeck ErrDeckNotFound", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/decks/invalid-id", nil)
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}
