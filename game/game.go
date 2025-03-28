package game

import "fmt"

// Dimension defines the size of the tic-tac-toe board (Dimension x Dimension)
const MetaBits = 3

type Game uint64

// Example of the game board representation for a 3x3 board:
// 0000000000000000000000000000000000000000000 00 00 00 00 00 00 00 00 00 00 00 0
//                                                                           ^^ ^ MetaBits
//                                                                              ^ Player turn (bit 0)
//                                                                           ^^__ Dimension (bits 1-2)
//                                                               ^^ ^^ ^^________ Row 1
//                                                      ^^ ^^ ^^_________________ Row 2
//                                             ^^ ^^ ^^__________________________ Row 3

type Player uint8

const (
	Empty Player = iota
	Player1
	Player2
)

func NewGame(dimension uint64) *Game {
	if 2 > dimension || dimension > 5 {
		return nil
	}
	dimension -= 2
	// Store dimension in bits 1-2 (bit 0 is for player turn)
	o := Game(dimension << 1)
	return &o
}

func (g *Game) GetDimension() int {
	// Extract the dimension bits (bits 1-2)
	return int((*g>>1)&3) + 2
}

func (g *Game) Winner() Player {
	dim := g.GetDimension()

	// Check rows
	for i := 0; i < dim; i++ {
		firstPos, _ := g.Get(0, i)
		if firstPos == Empty {
			continue
		}

		win := true
		for j := 1; j < dim; j++ {
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
	for i := 0; i < dim; i++ {
		firstPos, _ := g.Get(i, 0)
		if firstPos == Empty {
			continue
		}

		win := true
		for j := 1; j < dim; j++ {
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
		for i := 1; i < dim; i++ {
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
	firstPos, _ = g.Get(0, dim-1)
	if firstPos != Empty {
		win := true
		for i := 1; i < dim; i++ {
			pos, _ := g.Get(i, dim-1-i)
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
	// Reads bit 0 (the player turn bit)
	return Player((*g & 1)) + 1
}

func (g *Game) TogglePlayer() {
	// Toggles bit 0 (the player turn bit)
	*g ^= 1
}

func (g *Game) boardBitPosition(x, y int) uint {
	// Calculate bit position for board cells, starting after the MetaBits
	return uint(MetaBits + (x*g.GetDimension()+y)*2)
}

func (g *Game) Get(x, y int) (Player, error) {
	dim := g.GetDimension()
	if 0 > x || x >= dim || 0 > y || y >= dim {
		return Empty, InvalidCoordinate
	}
	pos := g.boardBitPosition(x, y)
	return Player((*g >> pos) & 3), nil
}

func (g *Game) validMove(x, y int) bool {
	p, err := g.Get(x, y)
	return err == nil && p == Empty
}

func (g *Game) Set(x, y int, p Player) error {
	dim := g.GetDimension()
	if 0 > x || x >= dim || 0 > y || y >= dim {
		return InvalidCoordinate
	} else if !g.validMove(x, y) {
		return InvalidMove
	}
	pos := g.boardBitPosition(x, y)
	*g = (*g &^ Game(3<<pos)) | Game(p)<<pos
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
	dim := g.GetDimension()

	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			p, _ := g.Get(j, i)
			s += fmt.Sprintf(" %s ", p)
			if j < dim-1 {
				s += "|"
			}
		}
		if i < dim-1 {
			s += "\n"
		}
		if i < dim-1 {
			for j := 0; j < dim; j++ {
				s += "---"
				if j < dim-1 {
					s += "+"
				}
			}
			if i < dim-1 {
				s += "\n"
			}
		}
	}
	return s
}
