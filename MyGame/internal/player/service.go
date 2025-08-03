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

func New(Id int) Player {
	var p Player = Player{Id, 0, *b.GetBusinesByName("None"), time.Now()}
	return p
}
