package game

import (
	"strings"
)

// Max dimension is 5x5 (5*5*2 = 50 bits) => 64 bits - 50 bits = 14 bits for metadata
const MetaBits = 14

type Game uint64

// Example of the game board representation for a 5x5 board:
// +---------50 bits---------+---14 bits---+
// | 00 00 00 00 00 00 00 00 | 0 0 0 0 0 0 |
// +-------------------------+-------------+
// | Rows & cols             | Meta data   |
// +-------------------------+-------------+
//
// The first 50 bits are used to store the board cells (5x5 cells * 2 bits per cell = 50 bits)
// The last 14 bits are used for metadata:
// - Bit 0: Player turn (0 or 1)
// - Bits 1-2: Dimension (2-5)
// - Bits 3-14: Unused

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
	dim := g.GetDimension()
	// Estimate capacity to avoid reallocations
	// For a 3x3 board, we need roughly 80 bytes
	// For larger boards, scale accordingly
	capacity := dim * dim * 10

	var sb strings.Builder
	sb.Grow(capacity)

	for i := 0; i < dim; i++ {
		// Print cell values with dividers
		for j := 0; j < dim; j++ {
			p, _ := g.Get(j, i)
			sb.WriteString(" ")
			sb.WriteString(p.String())
			sb.WriteString(" ")

			if j < dim-1 {
				sb.WriteString("|")
			}
		}
		if i < dim-1 {
			sb.WriteString("\n")
		}

		// Print horizontal dividers
		if i < dim-1 {
			for j := 0; j < dim; j++ {
				sb.WriteString("---")

				if j < dim-1 {
					sb.WriteString("+")
				}
			}
			if i < dim-1 {
				sb.WriteString("\n")
			}
		}
	}

	return sb.String()
}
