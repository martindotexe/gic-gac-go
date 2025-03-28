package game

import "fmt"

type Game uint32

// Game type is a 32-bit unsigned integer that represents the game state.
//
// 0000000000000 0 00 00 00 00 00 00 00 00 00
//                                   ^^ ^^ ^^ Row 1
//                          ^^ ^^ ^^ Row 2
//                 ^^ ^^ ^^ Row 3
//               ^ Bool player turn
// ^^^^^^^^^^^^^ Leftover bits

type Player uint8

const (
	Empty Player = iota
	Player1
	Player2
)

func NewGame() *Game {
	return new(Game)
}

func (g *Game) Player() Player {
	// Reads the 19th bit of the game state to determine the player
	return Player((*g >> 18) & 1)
}

func (g *Game) TogglePlayer() {
	// Toggles the 19th bit of the game state to switch players
	*g ^= 1 << 18
}

func (g *Game) Get(x, y int) (Player, error) {
	if 0 > x || x > 2 || 0 > y || y > 2 {
		return Empty, InvalidCoordinate
	}
	return Player((*g >> uint((x*3+y)*2)) & 3), nil
}

func (g *Game) Set(x, y int, p Player) error {
	if 0 > x || x > 2 || 0 > y || y > 2 {
		return InvalidCoordinate
	}
	*g = (*g &^ Game(3<<(uint(x*3+y)*2))) | Game(p)<<(uint(x*3+y)*2)
	return nil
}

func (p Player) String() string {
	switch p {
	case Player1:
		return "X"
	case Player2:
		return "O"
	default:
		return " "
	}
}

func (g Game) String() string {
	var s string

	for i := range 3 {
		x, _ := g.Get(i, 0)
		y, _ := g.Get(i, 1)
		z, _ := g.Get(i, 2)
		s += fmt.Sprintf(" %s | %s | %s\n", x, y, z)
		if i < 2 {
			s += fmt.Sprintf("---+---+---\n")
		}
	}
	return s
}
