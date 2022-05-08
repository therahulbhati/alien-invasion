package models

type Alien struct {
	Id          int
	CurrentCity *City
	IsAlive     bool
	TotalMoves  int
}
