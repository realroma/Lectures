package player

type Storage interface {
	GetOne()
	GetAll()
	Create()
	Edit(p *Player)
	Delete(p *Player)
}
