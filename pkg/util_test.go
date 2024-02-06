package pkg_test

import (
	"testing"

	"github.com/ologbonowiwi/toggl-challenge/pkg"
)

func TestContainsString(t *testing.T) {
	tests := []struct {
		name string
		slice []string
		elem string
		want bool
	}{
		{
			name: "contains",
			slice: []string{"a", "b", "c"},
			elem: "a",
			want: true,
		},
		{
			name: "not contains",
			slice: []string{"a", "b", "c"},
			elem: "d",
			want: false,
		},
		{
			name: "empty slice",
			slice: []string{},
			elem: "a",
			want: false,
		},
		{
			name: "nil slice",
			slice: nil,
			elem: "a",
			want: false,
		},
		{
			name: "empty elem",
			slice: []string{"a", "b", "c"},
			elem: "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pkg.Contains(tt.slice, tt.elem)
			if got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsInt(t *testing.T) {
	tests := []struct {
		name string
		slice []int
		elem int
		want bool
	}{
		{
			name: "contains",
			slice: []int{1, 2, 3},
			elem: 1,
			want: true,
		},
		{
			name: "not contains",
			slice: []int{1, 2, 3},
			elem: 4,
			want: false,
		},
		{
			name: "empty slice",
			slice: []int{},
			elem: 1,
			want: false,
		},
		{
			name: "nil slice",
			slice: nil,
			elem: 1,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pkg.Contains(tt.slice, tt.elem)
			if got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}