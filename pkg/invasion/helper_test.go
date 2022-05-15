package invasion

import (
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/therahulbhati/alien-invasion/pkg/models"
)

func Test_initializeWorld(t *testing.T) {
	expectedWorld := getFiveCityWorld()
	filePath := "../../test/testData/sample.txt"
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Error while opening file: %s, Error: %v", filePath, err)
	}
	defer file.Close()

	actualWorld, err := initializeWorld(file)
	if err != nil {
		t.Fatal("Error in initializing world", err)
	}
	assert.Equal(t, expectedWorld, actualWorld)

}

func Test_getAllCities(t *testing.T) {
	world := getFiveCityWorld()
	cities := getAllCities(world)
	assert.Equal(t, []string{"Bar", "Baz", "Bee", "Foo", "Qu-ux"}, cities)
}

func Test_initializeAliens(t *testing.T) {
	assert := assert.New(t)

	world := getFiveCityWorld()
	randGen := rand.New(rand.NewSource(42))
	aliens, cityAliensMapping := initializeAliens(world, 4, randGen)
	assert.Equal("Bar", aliens[0].CurrentCity)
	assert.Equal("Bee", aliens[1].CurrentCity)
	assert.Equal("Foo", aliens[2].CurrentCity)
	assert.Equal("Bar", aliens[3].CurrentCity)
	assert.True(cityAliensMapping["Bar"][0])
	assert.True(cityAliensMapping["Bar"][3])
	assert.True(cityAliensMapping["Bee"][1])
	assert.True(cityAliensMapping["Foo"][2])

}

func Test_getDirection(t *testing.T) {
	assert := assert.New(t)
	d, err := getDirection("east")
	handleError(assert, err)
	assert.Equal(models.East, d)

	d, err = getDirection("west")
	handleError(assert, err)
	assert.Equal(models.West, d)

	d, err = getDirection("north")
	handleError(assert, err)
	assert.Equal(models.North, d)

	d, err = getDirection("south")
	handleError(assert, err)
	assert.Equal(models.South, d)
}

func Test_getOppositeDirection(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(models.East, getOppositeDirection(models.West))
	assert.Equal(models.West, getOppositeDirection(models.East))
	assert.Equal(models.North, getOppositeDirection(models.South))
	assert.Equal(models.South, getOppositeDirection(models.North))
	assert.Equal(models.Direction(-1), getOppositeDirection(11))
}

func handleError(assert *assert.Assertions, err error) {
	if err != nil {
		assert.Fail("Got error", err)
	}
}
