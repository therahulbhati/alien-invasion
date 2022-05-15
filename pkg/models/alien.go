package models

type Alien struct {
	Id          int
	CurrentCity string
	IsAlive     bool
	IsTrapped   bool
	TotalMoves  int
}

func NewAlien(id int, currentCity string) *Alien {
	return &Alien{
		Id:          id,
		CurrentCity: currentCity,
		IsAlive:     true,
		IsTrapped:   false,
		TotalMoves:  0,
	}
}
