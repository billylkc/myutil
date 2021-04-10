package myutil

import (
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"Hello"},
	}
	for _, tt := range tests {
		if got := Hello(); got != tt.want {
			t.Errorf("Hello() = %q, want %q", got, tt.want)
		}
	}
}
