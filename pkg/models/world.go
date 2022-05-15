package models

import (
	"sort"
	"strings"
)

type World struct {
	Cities map[string]*City
}

// Returns new World
func NewWorld() *World {
	return &World{
		Cities: make(map[string]*City),
	}
}

// Returns World in the string format
func (w *World) String() string {
	cityNames := make([]string, 0)
	for cityName := range w.Cities {
		cityNames = append(cityNames, cityName)
	}
	sort.Strings(cityNames)

	var sb strings.Builder
	for _, cityName := range cityNames {
		sb.WriteString(cityName)
		city := w.Cities[cityName]
		for direction, neighbourCity := range city.Neighbour {
			if _, ok := w.Cities[neighbourCity]; ok {
				sb.WriteString(" ")
				sb.WriteString(direction.String())
				sb.WriteString("=")
				sb.WriteString(neighbourCity)
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
