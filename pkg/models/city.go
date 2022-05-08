package models

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

type City struct {
	Name        string
	Neighbour   map[Direction]*City
	IsDestroyed bool
}
