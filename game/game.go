package game

import "fmt"

// Dimension defines the size of the tic-tac-toe board (Dimension x Dimension)
const Dimension = 3

type Game uint64

// Example of the game board representation for a 3x3 board:
// 000000000000000000000000000000000000000000000 0 00 00 00 00 00 00 00 00 00
//                                                                   ^^ ^^ ^^ Row 1
//                                                          ^^ ^^ ^^ Row 2
//                                                 ^^ ^^ ^^ Row 3
//                                               ^ Bool player turn

type Player uint8

const (
	Empty Player = iota
	Player1
	Player2
)

func NewGame() *Game {
	return new(Game)
}

func (g *Game) Winner() Player {
	// Check rows
	for i := 0; i < Dimension; i++ {
		firstPos, _ := g.Get(0, i)
		if firstPos == Empty {
			continue
		}

		win := true
		for j := 1; j < Dimension; j++ {
			pos, _ := g.Get(j, i)
			if pos != firstPos {
				win = false
				break
			}
		}

		if win {
			return firstPos
		}
	}

	// Check columns
	for i := 0; i < Dimension; i++ {
		firstPos, _ := g.Get(i, 0)
		if firstPos == Empty {
			continue
		}

		win := true
		for j := 1; j < Dimension; j++ {
			pos, _ := g.Get(i, j)
			if pos != firstPos {
				win = false
				break
			}
		}

		if win {
			return firstPos
		}
	}

	// Check main diagonal (top-left to bottom-right)
	firstPos, _ := g.Get(0, 0)
	if firstPos != Empty {
		win := true
		for i := 1; i < Dimension; i++ {
			pos, _ := g.Get(i, i)
			if pos != firstPos {
				win = false
				break
			}
		}

		if win {
			return firstPos
		}
	}

	// Check other diagonal (top-right to bottom-left)
	firstPos, _ = g.Get(0, Dimension-1)
	if firstPos != Empty {
		win := true
		for i := 1; i < Dimension; i++ {
			pos, _ := g.Get(i, Dimension-1-i)
			if pos != firstPos {
				win = false
				break
			}
		}

		if win {
			return firstPos
		}
	}

	return Empty
}

func (g *Game) Player() Player {
	// Reads the bit after the board representation to determine the player
	return Player((*g>>(uint(Dimension*Dimension*2)))&1) + 1
}

func (g *Game) TogglePlayer() {
	// Toggles the bit after the board representation to switch players
	*g ^= 1 << (uint(Dimension * Dimension * 2))
}

func (g *Game) Get(x, y int) (Player, error) {
	if 0 > x || x >= Dimension || 0 > y || y >= Dimension {
		return Empty, InvalidCoordinate
	}
	return Player((*g >> uint((x*Dimension+y)*2)) & 3), nil
}

func (g *Game) validMove(x, y int) bool {
	p, err := g.Get(x, y)
	return err == nil && p == Empty
}

func (g *Game) Set(x, y int, p Player) error {
	if 0 > x || x >= Dimension || 0 > y || y >= Dimension {
		return InvalidCoordinate
	} else if !g.validMove(x, y) {
		return InvalidMove
	}
	*g = (*g &^ Game(3<<(uint(x*Dimension+y)*2))) | Game(p)<<(uint(x*Dimension+y)*2)
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

	for i := 0; i < Dimension; i++ {
		for j := 0; j < Dimension; j++ {
			p, _ := g.Get(j, i)
			s += fmt.Sprintf(" %s ", p)
			if j < Dimension-1 {
				s += "|"
			}
		}
		if i < Dimension-1 {
			s += "\n"
		}
		if i < Dimension-1 {
			for j := 0; j < Dimension; j++ {
				s += "---"
				if j < Dimension-1 {
					s += "+"
				}
			}
			if i < Dimension-1 {
				s += "\n"
			}
		}
	}
	return s
}
