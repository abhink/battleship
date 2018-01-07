package battleship

import "strconv"

type Dimensions interface {
	TopLeft() Position
	BottomRight() Position
}

type Position interface {
	X() int
	Y() int
	Name() string
}

var positionRegex = regexp.MustCompile("^([A-Z])([0-9]+)$")

type position struct {
	X, Y int
}

func (p *position) X() int {
	return p.X
}

func (p *position) Y() int {
	return p.Y
}

func NewPosition(x, y int) *position {
	return &position{x, y}
}

func NewPositionByName(name string) *position {
	match := regexp.FindStringSubmatch(name)
	if match == nil || len(match) < 3 {
		return nil
	}
	x, _ := strconv.Atoi(match[2])
	y := match[1][0] - 'A'

	return &position{x, y}
}

type Ship interface {
	Destroyed() bool
	Health() int
	Hit(p Position) bool
	Dimensions
}

type Type uint8

var (
	P Type = iota
	Q Type
)

func (t Type) HitPoints() int {
	if t == P {
		return 1
	}
	return 2
}

type BattleShip struct {
	t      Type
	width  int
	height int
	health int
	body   [][]uint8
}

func NewBattleShip(t Type, y, x int) *BattleShip {
	body := make([][]bool, y)
	for i := 0; i < x; i++ {
		body[i] = make([]bool, x)
	}
	return &BattleShip{
		t:      t,
		width:  y,
		height: x,
		health: y * x * t.HitPoints(),
		body:   body,
	}
}

func (s *BattleShip) Hit(p Position) bool {
	if s.Body[p.Y()][p.X()] >= s.t.HitPoints() {
		return false
	}
	s.Body[p.Y()][p.X()]++
	s.Health--
	return true
}

func (s *BattleShip) Health() int {
	return s.health
}

func (s *BattleShip) Destroyed() bool {
	return s.health <= 0
}

func (s *BattleShip) TopLeft() Position {
	return &position{0, 0}
}

func (s *BattleShip) BottomRight() Position {
	return &position{s.width - 1, s.height - 1}
}
