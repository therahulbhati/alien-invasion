package models

import "testing"

func TestDirection_String(t *testing.T) {
	tests := []struct {
		name string
		d    Direction
		want string
	}{
		{"test_east", East, "east"},
		{"test_west", West, "west"},
		{"test_south", South, "south"},
		{"test_north", North, "north"},
		{"test_invalid", -1, "Invalid Direction"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("Direction.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
