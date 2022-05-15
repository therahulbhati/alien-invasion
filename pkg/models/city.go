package models

type City struct {
	Name        string
	Neighbour   map[Direction]string
	IsDestroyed bool
}

// Returns new City with the given name
func NewCity(name string) *City {
	return &City{
		Name:        name,
		Neighbour:   make(map[Direction]string),
		IsDestroyed: false,
	}
}
