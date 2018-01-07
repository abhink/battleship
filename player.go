package battleship

type Player interface {
	Board() Board
	Ships() []Ship
	Fire(Position, Player) bool
}

type BattleShipPlayer struct {
	board Board
	ships []Ship
}

func (p *Player) Board() Board {
	return p.Board()
}

func (p *Player) Ships() Board {
	return p.ships
}

func (p *Player) Fire(ps Position, pl Player) bool {
	return pl.Board().Hit(p)
}
