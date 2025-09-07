package player

import (
	b "project/Lectures/MyGame/internal/buisnes"
	"time"
)

type Service interface {
	New(Id int) Player
}

type service struct {
	storage Storage
}

<<<<<<< HEAD
func New(Id int) *Player {
	var p Player = Player{Id, 0, *b.GetBusinesByName("None"), time.Now()}
	return &p
=======
func New(Id int) Player {
	var p Player = Player{Id, 0, *b.GetBusinesByName("None"), time.Now()}
	return p
>>>>>>> 31e8e98fb86b970935d45d13dc767d9c2a0c7f2c
}
