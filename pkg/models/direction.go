package models

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

// Returns string corresponding to the given Direction
func (d Direction) String() string {
	if d < 0 || d > 3 {
		return "Invalid Direction"
	}
	return [...]string{"north", "south", "east", "west"}[d]
}

// Returns int corresponding to the given Direction
func (d Direction) Int() int {
	return [...]int{0, 1, 2, 3}[d]
}
