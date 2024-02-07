package model

import (
	"errors"

	"github.com/ologbonowiwi/toggl-challenge/pkg"
)

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

// valid suits
var Suits = []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}

// valid values
var Values = []string{"ACE", "2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING"}

// errors
var (
	ErrInvalidSuit  = errors.New("invalid suit")
	ErrInvalidValue = errors.New("invalid value")
)

func NewCard(value, suit string) (Card, error) {
	if !pkg.Contains(Values, value) {
		return Card{}, ErrInvalidValue
	}

	if !pkg.Contains(Suits, suit) {
		return Card{}, ErrInvalidSuit
	}

	code := value[:1] + suit[:1]

	return Card{
		Value: value,
		Suit:  suit,
		Code:  code,
	}, nil
}
