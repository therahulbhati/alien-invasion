package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWorld(t *testing.T) {
	world := NewWorld()
	assert.Equal(t, 0, len(world.Cities))
}

func TestWorld_String(t *testing.T) {
	world := NewWorld()
	fooCity := NewCity("Foo")
	barCity := NewCity("Bar")
	fooCity.Neighbour[North] = "Bar"
	barCity.Neighbour[South] = "Foo"
	world.Cities["Foo"] = fooCity
	world.Cities["Bar"] = barCity
	assert.Equal(t, "Bar south=Foo\nFoo north=Bar\n", world.String())
}
