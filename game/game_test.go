package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGame(t *testing.T) {
	// Create a new game
	g := NewGame()

	// Ensure the game is empty
	assert.Equal(t, g, new(Game))

	// Ensure String() returns the correct string
	es := "   |   |   \n---+---+---\n   |   |   \n---+---+---\n   |   |   "
	assert.Equal(t, es, g.String())

	// Loop through all coordinates
	for x := range 3 {
		for y := range 3 {
			// Ensure the get method returns the correct value
			p, err := g.Get(x, y)
			require.NoError(t, err)
			assert.Equal(t, Empty, p)

			// Ensure the set method sets the correct value
			err = g.Set(x, y, Player1)
			require.NoError(t, err)

			// Ensure the get method returns the correct value
			p, err = g.Get(x, y)
			require.NoError(t, err)
			assert.Equal(t, Player1, p)

			// Ensure that overwriting gives an error
			err = g.Set(x, y, Player2)
			require.Error(t, err)
		}
	}

	// Ensure String() returns the correct string
	es = " X | X | X \n---+---+---\n X | X | X \n---+---+---\n X | X | X "
	assert.Equal(t, es, g.String())
	// Ensure the set method returns an error for invalid coordinates
	err := g.Set(3, 0, Player1)
	require.Error(t, err)

	err = g.Set(-1, 0, Player1)
	require.Error(t, err)

	// Ensure the get method returns an error for invalid coordinates
	_, err = g.Get(3, 0)
	require.Error(t, err)

	_, err = g.Get(-1, 0)
	require.Error(t, err)
}

func TestWinner(t *testing.T) {
	g := NewGame()

	// Ensure the game is not won
	assert.Equal(t, Empty, g.Winner())

	// Ensure the game is won by Player1
	e1 := g.Set(0, 0, Player1)
	e2 := g.Set(1, 0, Player1)
	e3 := g.Set(2, 0, Player1)

	assert.NoError(t, e1)
	assert.NoError(t, e2)
	assert.NoError(t, e3)
	assert.Equal(t, Player1, g.Winner())

	g = NewGame()

	// Ensure the game is won by Player2

	e1 = g.Set(0, 0, Player2)
	e2 = g.Set(1, 1, Player2)
	e3 = g.Set(2, 2, Player2)

	assert.NoError(t, e1)
	assert.NoError(t, e2)
	assert.NoError(t, e3)
	assert.Equal(t, Player2, g.Winner())

	// Ensure the game is won by Player1

	g = NewGame()

	e1 = g.Set(0, 0, Player1)
	e2 = g.Set(0, 1, Player1)
	e3 = g.Set(0, 2, Player1)

	assert.NoError(t, e1)
	assert.NoError(t, e2)
	assert.NoError(t, e3)
	assert.Equal(t, Player1, g.Winner())
}
