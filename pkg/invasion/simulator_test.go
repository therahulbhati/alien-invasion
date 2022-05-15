package invasion

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/therahulbhati/alien-invasion/pkg/models"
)

func TestNewSimulator(t *testing.T) {
	assert := assert.New(t)
	filePath := "../../test/testData/sample.txt"
	file, err := os.Open(filePath)

	if err != nil {
		t.Fatalf("Error while opening file: %s, Error: %v", filePath, err)

	}
	defer file.Close()

	s, err := NewSimulator(file, 5, 100, 42)
	if err != nil {
		t.Fatal("Error while initiating new Simulator", err)
	}
	assert.True(s.CityAliensMapping["Bar"][0])
	assert.True(s.CityAliensMapping["Bee"][1])
	assert.True(s.CityAliensMapping["Foo"][2])
	assert.True(s.CityAliensMapping["Bar"][3])
	assert.True(s.CityAliensMapping["Foo"][4])
}

func TestSimulator_Simulate(t *testing.T) {
	assert := assert.New(t)
	world := getFiveCityWorld()
	aliens := make([]*models.Alien, 3)
	aliens[0] = &models.Alien{
		Id:          0,
		IsAlive:     true,
		IsTrapped:   false,
		CurrentCity: "Foo",
	}

	aliens[1] = &models.Alien{
		Id:          1,
		IsAlive:     true,
		IsTrapped:   false,
		CurrentCity: "Qu-ux",
	}
	aliens[2] = &models.Alien{
		Id:          2,
		IsAlive:     true,
		IsTrapped:   false,
		CurrentCity: "Bee",
	}
	cityAliensMapping := make(map[string]map[int]bool)
	cityAliensMapping["Foo"] = make(map[int]bool)
	cityAliensMapping["Foo"][0] = true

	cityAliensMapping["Qu-ux"] = make(map[int]bool)
	cityAliensMapping["Qu-ux"][1] = true
	s := &Simulator{
		World:             world,
		Aliens:            aliens,
		CityAliensMapping: cityAliensMapping,
		MaxIterations:     1,
		RandGen:           rand.New(rand.NewSource(42)),
	}

	out := captureOutput(func() { s.Simulate() })
	expected := `Iteration 0
Iternation 1
Alien 0 moved to Bar from Foo
Alien 1 moved to Foo from Qu-ux
Alien 2 moved to Bar from Bee
Bar has been destroyed by aliens 0, 2!
Foo has been destroyed by aliens 0, 1!

Baz
Bee
Qu-ux

`
	fmt.Println(out)
	assert.Equal(expected, out)

}

func TestSimulator_move(t *testing.T) {
	assert := assert.New(t)
	world := getFiveCityWorld()
	aliens := make([]*models.Alien, 2)
	aliens[0] = &models.Alien{
		Id:          0,
		IsAlive:     true,
		IsTrapped:   false,
		CurrentCity: "Foo",
	}

	aliens[1] = &models.Alien{
		Id:          1,
		IsAlive:     true,
		IsTrapped:   false,
		CurrentCity: "Qu-ux",
	}
	cityAliensMapping := make(map[string]map[int]bool)
	cityAliensMapping["Foo"] = make(map[int]bool)
	cityAliensMapping["Foo"][0] = true

	cityAliensMapping["Qu-ux"] = make(map[int]bool)
	cityAliensMapping["Qu-ux"][1] = true
	s := &Simulator{
		World:             world,
		Aliens:            aliens,
		CityAliensMapping: cityAliensMapping,
		MaxIterations:     1,
		RandGen:           rand.New(rand.NewSource(42)),
	}

	out := captureOutput(func() { s.move() })
	assert.Equal("Alien 0 moved to Bar from Foo\nAlien 1 moved to Foo from Qu-ux\n", out)
}

func TestSimulator_fight(t *testing.T) {
	assert := assert.New(t)
	world := getFiveCityWorld()
	aliens := make([]*models.Alien, 2)
	aliens[0] = &models.Alien{
		Id:          0,
		IsAlive:     true,
		IsTrapped:   false,
		CurrentCity: "Foo",
	}

	aliens[1] = &models.Alien{
		Id:          1,
		IsAlive:     true,
		IsTrapped:   false,
		CurrentCity: "Foo",
	}
	cityAliensMapping := make(map[string]map[int]bool)
	cityAliensMapping["Foo"] = make(map[int]bool)
	// place alien 0 in city Foo
	cityAliensMapping["Foo"][0] = true
	cityAliensMapping["Foo"][1] = true
	s := &Simulator{
		World:             world,
		Aliens:            aliens,
		CityAliensMapping: cityAliensMapping,
		MaxIterations:     1,
		RandGen:           rand.New(rand.NewSource(42)),
	}
	out := captureOutput(func() { s.fight() })
	assert.Equal("Foo has been destroyed by aliens 0, 1!\n", out)
}

func TestSimulator_clean(t *testing.T) {
	assert := assert.New(t)
	world := getFiveCityWorld()
	aliens := make([]*models.Alien, 2)
	aliens[0] = &models.Alien{
		Id:          0,
		IsAlive:     true,
		IsTrapped:   false,
		CurrentCity: "Foo",
	}

	aliens[1] = &models.Alien{
		Id:          1,
		IsAlive:     true,
		IsTrapped:   false,
		CurrentCity: "Bar",
	}
	cityAliensMapping := make(map[string]map[int]bool)
	cityAliensMapping["Foo"] = make(map[int]bool)
	// place alien 0 in city Foo
	cityAliensMapping["Foo"][0] = true
	s := &Simulator{
		World:             world,
		Aliens:            aliens,
		CityAliensMapping: cityAliensMapping,
		MaxIterations:     1,
		RandGen:           rand.New(rand.NewSource(42)),
	}
	killedAliens := []int{0, 1}
	destroyedCiites := []string{"Rah"}
	s.clean(killedAliens, destroyedCiites)
	assert.False(s.Aliens[0].IsAlive)
	assert.False(s.Aliens[1].IsAlive)

	_, isPresent := s.World.Cities["Rah"]
	assert.False(isPresent)
}

func TestSimulator_getRandomNeighbour(t *testing.T) {
	assert := assert.New(t)

	world := getFiveCityWorld()
	aliens := make([]*models.Alien, 1)
	aliens[0] = &models.Alien{
		Id:          0,
		IsAlive:     true,
		IsTrapped:   false,
		CurrentCity: "Foo",
	}
	cityAliensMapping := make(map[string]map[int]bool)
	cityAliensMapping["Foo"] = make(map[int]bool)
	// place alien 0 in city Foo
	cityAliensMapping["Foo"][0] = true
	s := &Simulator{
		World:             world,
		Aliens:            aliens,
		CityAliensMapping: cityAliensMapping,
		MaxIterations:     1,
		RandGen:           rand.New(rand.NewSource(42)),
	}
	neighbourCity := s.getRandomNeighbour(*s.World.Cities["Foo"])
	assert.Equal("Bar", neighbourCity)

	delete(s.World.Cities, "Bar")
	neighbourCity = s.getRandomNeighbour(*s.World.Cities["Foo"])
	assert.Equal("Qu-ux", neighbourCity)
}

func TestSimulator_isAnyAlienAlive(t *testing.T) {
	assert := assert.New(t)

	world := getTwoCityWorld()
	aliens := make([]*models.Alien, 1)
	aliens[0] = &models.Alien{
		Id:          0,
		IsAlive:     true,
		IsTrapped:   false,
		CurrentCity: "Foo",
	}
	cityAliensMapping := make(map[string]map[int]bool)
	cityAliensMapping["Foo"] = make(map[int]bool)
	// place alien 0 in city Foo
	cityAliensMapping["Foo"][0] = true
	s := &Simulator{
		World:             world,
		Aliens:            aliens,
		CityAliensMapping: cityAliensMapping,
		MaxIterations:     1,
		RandGen:           rand.New(rand.NewSource(42)),
	}

	assert.True(s.isAnyAlienAlive())
	// Kill alien
	s.Aliens[0].IsAlive = false
	assert.False(s.isAnyAlienAlive())
}

func TestSimulator_areAllAliveAliensTrapped(t *testing.T) {
	assert := assert.New(t)

	world := getTwoCityWorld()
	aliens := make([]*models.Alien, 1)
	aliens[0] = &models.Alien{
		Id:          0,
		IsAlive:     true,
		IsTrapped:   false,
		CurrentCity: "Foo",
	}
	cityAliensMapping := make(map[string]map[int]bool)
	cityAliensMapping["Foo"] = make(map[int]bool)
	// place alien 0 in city Foo
	cityAliensMapping["Foo"][0] = true
	s := &Simulator{
		World:             world,
		Aliens:            aliens,
		CityAliensMapping: cityAliensMapping,
		MaxIterations:     1,
		RandGen:           rand.New(rand.NewSource(42)),
	}
	assert.False(s.areAllAliveAliensTrapped())

	// Destroy city bar and trap alien
	delete(s.World.Cities, "Bar")
	s.Aliens[0].IsTrapped = true

	assert.True(s.areAllAliveAliensTrapped())
}

func TestSimulator_printTrappedAliens(t *testing.T) {
	assert := assert.New(t)
	world := getTwoCityWorld()
	aliens := make([]*models.Alien, 1)
	aliens[0] = &models.Alien{
		Id:          0,
		IsAlive:     true,
		IsTrapped:   false,
		CurrentCity: "Foo",
	}
	cityAliensMapping := make(map[string]map[int]bool)
	cityAliensMapping["Foo"] = make(map[int]bool)
	// place alien 0 in city Foo
	cityAliensMapping["Foo"][0] = true
	s := &Simulator{
		World:             world,
		Aliens:            aliens,
		CityAliensMapping: cityAliensMapping,
		MaxIterations:     1,
		RandGen:           rand.New(rand.NewSource(42)),
	}

	out := captureOutput(func() { s.printTrappedAliens() })
	assert.Equal("", out)

	// Destroy city bar and trap alien
	delete(s.World.Cities, "Bar")
	s.Aliens[0].IsTrapped = true

	out = captureOutput(func() { s.printTrappedAliens() })
	assert.Equal("Alien 0 trapped in city: Foo\n", out)
}

func TestSimulator_canMove(t *testing.T) {
	assert := assert.New(t)

	world := getTwoCityWorld()
	aliens := make([]*models.Alien, 1)
	aliens[0] = &models.Alien{
		Id:          0,
		IsAlive:     true,
		IsTrapped:   false,
		CurrentCity: "Foo",
	}
	cityAliensMapping := make(map[string]map[int]bool)
	cityAliensMapping["Foo"] = make(map[int]bool)
	// place alien 0 in city Foo
	cityAliensMapping["Foo"][0] = true
	simulator := &Simulator{
		World:             world,
		Aliens:            aliens,
		CityAliensMapping: cityAliensMapping,
		MaxIterations:     1,
		RandGen:           rand.New(rand.NewSource(42)),
	}

	assert.Equal(true, simulator.canMove(*simulator.Aliens[0]))
	// Destroy city bar and trap alien
	delete(simulator.World.Cities, "Bar")
	assert.Equal(false, simulator.canMove(*simulator.Aliens[0]))

}

func getTwoCityWorld() *models.World {

	fooCity := &models.City{
		Name:        "Foo",
		Neighbour:   make(map[models.Direction]string),
		IsDestroyed: false,
	}

	barCity := &models.City{
		Name:        "Bar",
		Neighbour:   make(map[models.Direction]string),
		IsDestroyed: false,
	}

	fooCity.Neighbour[models.North] = "Bar"
	barCity.Neighbour[models.South] = "Foo"

	world := &models.World{
		Cities: make(map[string]*models.City),
	}
	world.Cities["Foo"] = fooCity
	world.Cities["Bar"] = barCity
	return world
}

func captureOutput(f func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	return string(out)
}

func getFiveCityWorld() *models.World {
	// 	Foo north=Bar west=Baz south=Qu-ux
	// Bar south=Foo west=Bee

	fooCity := &models.City{
		Name:        "Foo",
		Neighbour:   make(map[models.Direction]string),
		IsDestroyed: false,
	}

	barCity := &models.City{
		Name:        "Bar",
		Neighbour:   make(map[models.Direction]string),
		IsDestroyed: false,
	}

	bazCity := &models.City{
		Name:        "Baz",
		Neighbour:   make(map[models.Direction]string),
		IsDestroyed: false,
	}
	beeCity := &models.City{
		Name:        "Bee",
		Neighbour:   make(map[models.Direction]string),
		IsDestroyed: false,
	}

	quuxCity := &models.City{
		Name:        "Qu-ux",
		Neighbour:   make(map[models.Direction]string),
		IsDestroyed: false,
	}
	fooCity.Neighbour[models.North] = "Bar"
	barCity.Neighbour[models.South] = "Foo"
	fooCity.Neighbour[models.West] = "Baz"
	bazCity.Neighbour[models.East] = "Foo"
	fooCity.Neighbour[models.South] = "Qu-ux"
	quuxCity.Neighbour[models.North] = "Foo"

	barCity.Neighbour[models.West] = "Bee"
	beeCity.Neighbour[models.East] = "Bar"

	world := &models.World{
		Cities: make(map[string]*models.City),
	}
	world.Cities["Foo"] = fooCity
	world.Cities["Bar"] = barCity
	world.Cities["Bee"] = beeCity
	world.Cities["Baz"] = bazCity
	world.Cities["Qu-ux"] = quuxCity

	return world
}
