package battleship

import "errors"

type Board interface {
	Hit(p Position) bool
	AddShip(s Ship) error
	Dimesions
}

type BattleShipBoard struct {
	width  int
	height int
	ships  map[string]Ship
}

func NewBattleShipBoard(y, x int) *BattleShipBoard {
	return &BattleShipBoard{x, y}
}

func (b *BattleShipBoard) AddShip(s Ship, p Position) error {
	if b.CheckBounds(p, s.BottomRight()) {
		return errors.New("invalid ship placement")
	}
	ships[p.Name()] = s
}

func (b *BattleShipBoard) Hit(p Position) bool {
	ship, ok := b.ships[p.Name()]
	return ok && ship.Hit(p)
}

func (b *BattleShipBoard) checkBounds(p, q Position) bool {
	if p.X+q.X > b.width || p.Y+q.Y > b.height {
		return false
	}
	return true
}

func (b *BattleShipBoard) TopLeft() Position {
	return &position{0, 0}
}

func (b *BattleShipBoard) Height() int {
	return &position{b.width - 1, b.height - 1}
}
