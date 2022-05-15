package models

import (
	"reflect"
	"testing"
)

func TestNewAlien(t *testing.T) {
	type args struct {
		id          int
		currentCity string
	}
	tests := []struct {
		name string
		args args
		want *Alien
	}{
		{"basic1", args{10, "Foo"}, &Alien{
			Id:          10,
			CurrentCity: "Foo",
			IsAlive:     true,
			IsTrapped:   false,
			TotalMoves:  0,
		}},
		{"basic2", args{11, "Bar"}, &Alien{
			Id:          11,
			CurrentCity: "Bar",
			IsAlive:     true,
			IsTrapped:   false,
			TotalMoves:  0,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAlien(tt.args.id, tt.args.currentCity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAlien() = %v, want %v", got, tt.want)
			}
		})
	}
}
