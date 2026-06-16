package service

import (
	"testing"
	"time"
)

func TestCalcAge(t *testing.T) {
	tests := []struct {
		name string
		dob  time.Time
		want int
	}{
		{"exact birthday today", time.Now().AddDate(-25, 0, 0), 25},
		{"birthday not yet this year", time.Now().AddDate(-25, 0, 1), 24},
		{"birthday passed this year", time.Now().AddDate(-25, 0, -1), 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcAge(tt.dob); got != tt.want {
				t.Errorf("calcAge() = %d, want %d", got, tt.want)
			}
		})
	}
}