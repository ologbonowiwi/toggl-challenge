package model_test

import (
	"testing"

	"github.com/ologbonowiwi/toggl-challenge/internal/model"
)

func TestNewCard(t *testing.T) {
	tests := []struct {
		name  string
		value string
		suit  string
		want  model.Card
	}{
		{
			name:  "AS",
			value: "ACE",
			suit:  "SPADES",
			want: model.Card{
				Value: "ACE",
				Suit:  "SPADES",
				Code:  "AS",
			},
		},
		{
			name:  "2H",
			value: "2",
			suit:  "HEARTS",
			want: model.Card{
				Value: "2",
				Suit:  "HEARTS",
				Code:  "2H",
			},
		},
		{
			name:  "JD",
			value: "JACK",
			suit:  "DIAMONDS",
			want: model.Card{
				Value: "JACK",
				Suit:  "DIAMONDS",
				Code:  "JD",
			},
		},
		{
			name:  "QC",
			value: "QUEEN",
			suit:  "CLUBS",
			want: model.Card{
				Value: "QUEEN",
				Suit:  "CLUBS",
				Code:  "QC",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := model.NewCard(tt.value, tt.suit)
			if err != nil {
				t.Errorf("NewCard() error = %v, want nil", err)
			}
			if got != tt.want {
				t.Errorf("NewCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCardError(t *testing.T) {
	tests := []struct {
		name  string
		value string
		suit  string
		want  error
	}{
		{
			name:  "Invalid value",
			value: "70",
			suit:  "SPADES",
			want:  model.ErrInvalidValue,
		},
		{
			name:  "Invalid suit",
			value: "ACE",
			suit:  "SPADE",
			want:  model.ErrInvalidSuit,
		},
		{
			name:  "Invalid value and suit",
			value: "70",
			suit:  "SPADE",
			want:  model.ErrInvalidValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := model.NewCard(tt.value, tt.suit)
			if err != tt.want {
				t.Errorf("NewCard() error = %v, want %v", err, tt.want)
			}
		})
	}
}
