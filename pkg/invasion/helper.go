package invasion

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"strings"

	"github.com/therahulbhati/alien-invasion/pkg/models"
)

func initializeWorld(file io.Reader) (*models.World, error) {
	scanner := bufio.NewScanner(file)
	cities := make(map[string]*models.City)

	for scanner.Scan() {
		cityInput := strings.Split(scanner.Text(), " ")
		originCityName := cityInput[0]
		if _, ok := cities[originCityName]; !ok {
			cities[originCityName] = models.NewCity(originCityName)
		}

		for _, v := range cityInput[1:] {
			directionCityMap := strings.Split(v, "=")
			if len(directionCityMap) != 2 {
				return nil, fmt.Errorf("invalid direction input %v", v)
			}
			direction, err := getDirection(directionCityMap[0])
			if err != nil {
				return nil, err
			}
			neighbourCityName := directionCityMap[1]
			if _, ok := cities[neighbourCityName]; !ok {
				cities[neighbourCityName] = models.NewCity(neighbourCityName)
			}

			cities[originCityName].Neighbour[direction] = neighbourCityName
			cities[neighbourCityName].Neighbour[getOppositeDirection(direction)] = originCityName
		}
	}

	return &models.World{
		Cities: cities,
	}, nil
}

func getAllCities(world *models.World) []string {
	cities := make([]string, 0, len(world.Cities))
	for city := range world.Cities {
		cities = append(cities, city)
	}
	sort.Strings(cities)
	return cities
}

func initializeAliens(world *models.World, numOfAliens int, randGen *rand.Rand) ([]*models.Alien, map[string]map[int]bool) {
	totalCities := len(world.Cities)
	cities := getAllCities(world)
	cityAliensMapping := make(map[string]map[int]bool)
	aliens := make([]*models.Alien, numOfAliens)

	for i := 0; i < numOfAliens; i++ {
		randomCity := cities[randGen.Intn(totalCities)]
		aliens[i] = models.NewAlien(i, randomCity)
		if _, ok := cityAliensMapping[randomCity]; !ok {
			cityAliensMapping[randomCity] = make(map[int]bool)
		}
		cityAliensMapping[randomCity][i] = true
		fmt.Printf("Alien %d placed in city %s\n", i, randomCity)
	}
	return aliens, cityAliensMapping
}

func getDirection(direction string) (models.Direction, error) {
	switch strings.ToLower(direction) {
	case "east":
		return models.East, nil
	case "west":
		return models.West, nil
	case "north":
		return models.North, nil
	case "south":
		return models.South, nil
	default:
		return -1, fmt.Errorf("invalid direction")
	}
}

func getOppositeDirection(direction models.Direction) models.Direction {
	switch direction {
	case models.East:
		return models.West
	case models.West:
		return models.East
	case models.South:
		return models.North
	case models.North:
		return models.South
	default:
		return -1
	}
}
