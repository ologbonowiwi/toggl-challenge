package api

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ologbonowiwi/toggl-challenge/internal/model"
	"github.com/ologbonowiwi/toggl-challenge/internal/services"
)

type APIHandler struct {
	deckService services.DeckService
}

func NewAPIHandler(deckService services.DeckService) *APIHandler {
	return &APIHandler{deckService: deckService}
}

func handleError(c *fiber.Ctx, err error) error {
	errStatusMap := map[error]int{
		services.ErrDeckNotFound: fiber.StatusNotFound,
		model.ErrNotEnoughCards:  fiber.StatusUnprocessableEntity,
		model.ErrInvalidAmount:   fiber.StatusBadRequest,
	}

	status, ok := errStatusMap[err]
	if !ok {
		status = fiber.StatusInternalServerError
	}

	return c.Status(status).JSON(fiber.Map{"error": err.Error()})
}

func (ah *APIHandler) SetupRoutes(app fiber.Router) {
	app.Post("/decks", ah.createDeck)
	app.Get("/decks/:id", ah.getDeck)
	app.Post("/decks/:id/draw", ah.drawCards)
}

func (ah *APIHandler) createDeck(c *fiber.Ctx) error {
	shuffle := c.QueryBool("shuffle")
	codes := c.Query("codes")

	var codesList []string

	if codes != "" {
		codesList = strings.Split(codes, ",")
	}

	deck := ah.deckService.CreateDeck(shuffle, codesList)

	return c.JSON(fiber.Map{
		"deck_id":   deck.ID,
		"shuffled":  deck.Shuffled,
		"remaining": deck.Remaining,
	})
}

func (ah *APIHandler) drawCards(c *fiber.Ctx) error {
	id := c.Params("id")
	amount := c.QueryInt("amount")

	cards, err := ah.deckService.DrawCards(id, amount)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(fiber.Map{"cards": cards})
}

func (ah *APIHandler) getDeck(c *fiber.Ctx) error {
	id := c.Params("id")

	deck, err := ah.deckService.GetDeck(id)

	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(deck)
}
