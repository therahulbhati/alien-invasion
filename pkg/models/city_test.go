package models

import (
	"reflect"
	"testing"
)

func TestNewCity(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *City
	}{
		{"basic1", args{"Foo"}, &City{Name: "Foo", IsDestroyed: false, Neighbour: make(map[Direction]string)}},
		{"basic2", args{"Bar"}, &City{Name: "Bar", IsDestroyed: false, Neighbour: make(map[Direction]string)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCity(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCity() = %v, want %v", got, tt.want)
			}
		})
	}
}
