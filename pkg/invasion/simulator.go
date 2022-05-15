package invasion

import (
	"fmt"
	"io"
	"math/rand"
	"sort"
	"strings"

	"github.com/therahulbhati/alien-invasion/pkg/models"
)

type Simulator struct {
	World             *models.World
	Aliens            []*models.Alien
	CityAliensMapping map[string]map[int]bool
	MaxIterations     int
	RandGen           *rand.Rand
}

// Returns a new Simulator object, build using given file descriptor, number of aliens, maximum iterations and seed value
func NewSimulator(file io.Reader, numOfAliens, maxIteration int, seedValue int64) (*Simulator, error) {
	randGen := rand.New(rand.NewSource(seedValue))
	world, err := initializeWorld(file)
	if err != nil {
		return nil, fmt.Errorf("failed to build simulator object %v", err)
	}
	aliens, cityAliensMapping := initializeAliens(world, numOfAliens, randGen)
	return &Simulator{
		World:             world,
		Aliens:            aliens,
		CityAliensMapping: cityAliensMapping,
		MaxIterations:     maxIteration,
		RandGen:           randGen,
	}, nil
}

// Runs Simulation unitl any termination condition is met
// Termination conditions:
// 1. All Aliens are dead
// 2. All Aliens are trapped
// 3. Max iterations are reached
// Print trapped Aliens if any
// Print world post-invasion
func (s *Simulator) Simulate() {
	// All the aliens in the same city fight and destroy the city
	fmt.Println("Iteration 0")
	s.fight()
	for i := 1; s.isAnyAlienAlive() && i <= s.MaxIterations; i++ {
		fmt.Printf("Iteration %d\n", i)
		s.move()
		s.fight()
		if s.areAllAliveAliensTrapped() {
			break
		}

	}
	s.printTrappedAliens()
	fmt.Println()
	fmt.Println(s.World.String())
}

// Move aliens in the neighbouring city if a path exists
func (s *Simulator) move() {
	for i := 0; i < len(s.Aliens); i++ {
		if s.Aliens[i].IsAlive && !s.Aliens[i].IsTrapped {
			if !s.canMove(*s.Aliens[i]) {
				s.Aliens[i].IsTrapped = true
				continue
			}
			currentCity := s.Aliens[i].CurrentCity
			nextCity := s.getRandomNeighbour(*s.World.Cities[currentCity])
			s.Aliens[i].CurrentCity = nextCity
			if _, ok := s.CityAliensMapping[nextCity]; !ok {
				s.CityAliensMapping[nextCity] = make(map[int]bool)
			}
			s.CityAliensMapping[nextCity][i] = true
			fmt.Printf("Alien %d moved to %s from %s\n", i, nextCity, currentCity)
		}
	}
}

// If more than two aliens are present in the same city they fight and kill each other destroying city
func (s *Simulator) fight() {
	destroyedCities := make([]string, 0) // track destroyed cities to remove linking
	allKilledAliens := make([]int, 0)    // track killed Aliens to mark them dead
	cityAlienMappingCopy := make(map[string]map[int]bool)
	cityNames := make([]string, 0)
	for cityName := range s.CityAliensMapping {
		cityNames = append(cityNames, cityName)
	}
	sort.Strings(cityNames)
	for _, city := range cityNames {
		aliens := s.CityAliensMapping[city]
		killedAliens := make([]int, 0)
		// Alien survived the fight
		if len(aliens) == 1 {
			cityAlienMappingCopy[city] = make(map[int]bool)
			for alien := range aliens {
				cityAlienMappingCopy[city][alien] = true
			}
			continue
		}
		for alien := range aliens {
			killedAliens = append(killedAliens, alien)
		}
		if len(killedAliens) > 1 {
			aliensString := ""
			sort.Ints(killedAliens)
			for _, v := range killedAliens {
				aliensString += fmt.Sprintf("%d, ", v)
			}
			destroyedCities = append(destroyedCities, city)
			fmt.Printf("%s has been destroyed by aliens %s!\n", city, strings.TrimRight(aliensString, ", "))
		}
		allKilledAliens = append(allKilledAliens, killedAliens...)
	}
	s.CityAliensMapping = cityAlienMappingCopy // update mapping after fight
	s.clean(allKilledAliens, destroyedCities)
}

// Mark killed aliens as dead and remove destroyed cities from world
func (s *Simulator) clean(killedAliens []int, destroyedCities []string) {
	for _, alien := range killedAliens {
		s.Aliens[alien].IsAlive = false
	}
	for _, city := range destroyedCities {
		delete(s.World.Cities, city)
	}
}

// return random neighbour city if any exists
func (s *Simulator) getRandomNeighbour(city models.City) string {
	neigbourCities := make([]string, 0)
	directions := make([]models.Direction, 0)
	for dir := range city.Neighbour {
		directions = append(directions, dir)
	}
	sort.SliceStable(directions, func(i, j int) bool {
		return directions[i].Int() > directions[j].Int()
	})
	for _, dir := range directions {
		neighbour := city.Neighbour[dir]
		if _, ok := s.World.Cities[neighbour]; ok {
			neigbourCities = append(neigbourCities, neighbour)
		}
	}
	return neigbourCities[s.RandGen.Intn(len(neigbourCities))]
}

// Returns false if all aliens are dead else true
func (s *Simulator) isAnyAlienAlive() bool {
	for i := 0; i < len(s.Aliens); i++ {
		if s.Aliens[i].IsAlive {
			return true
		}
	}
	return false
}

// Returns true if all aliens are trapped else false
func (s *Simulator) areAllAliveAliensTrapped() bool {
	for i := 0; i < len(s.Aliens); i++ {
		if s.Aliens[i].IsAlive && !s.Aliens[i].IsTrapped {
			return false
		}
	}
	return true
}

// Prints all trapped aliens
func (s *Simulator) printTrappedAliens() {
	for i := 0; i < len(s.Aliens); i++ {
		if s.Aliens[i].IsAlive && s.Aliens[i].IsTrapped {
			fmt.Printf("Alien %d trapped in city: %s\n", i, s.Aliens[i].CurrentCity)
		}
	}
}

func (s *Simulator) canMove(alien models.Alien) bool {
	currentCity := alien.CurrentCity
	for _, neighbour := range s.World.Cities[currentCity].Neighbour {
		if _, ok := s.World.Cities[neighbour]; ok {
			return true
		}
	}
	return false
}
