package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGame3x3(t *testing.T) {
	// Create a new game
	g := NewGame(3)

	// Ensure the game is empty
	assert.Equal(t, Game(2), *g) // 3: 010 = 2

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

func TestGame4x4(t *testing.T) {
	// Create a new game
	g := NewGame(4)

	// Ensure the game is empty
	assert.Equal(t, Game(4), *g) // 4: 100 = 4

	// Ensure String() returns the correct string
	es := "   |   |   |   \n---+---+---+---\n   |   |   |   \n---+---+---+---\n   |   |   |   \n---+---+---+---\n   |   |   |   "
	assert.Equal(t, es, g.String())

	// Loop through all coordinates
	for x := range 4 {
		for y := range 4 {
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
	es = " X | X | X | X \n---+---+---+---\n X | X | X | X \n---+---+---+---\n X | X | X | X \n---+---+---+---\n X | X | X | X "
	assert.Equal(t, es, g.String())

	// Ensure the set method returns an error for invalid coordinates
	err := g.Set(4, 0, Player1)
	require.Error(t, err)

	err = g.Set(-1, 0, Player1)
	require.Error(t, err)

	// Ensure the get method returns an error for invalid coordinates
	_, err = g.Get(4, 0)
	require.Error(t, err)

	_, err = g.Get(-1, 0)
	require.Error(t, err)
}

func TestGame5x5(t *testing.T) {
	// Create a new game
	g := NewGame(5)

	// Ensure the game is empty
	assert.Equal(t, Game(6), *g) // 6: 110 = 6

	// Ensure String() returns the correct string
	es := "   |   |   |   |   \n---+---+---+---+---\n   |   |   |   |   \n---+---+---+---+---\n   |   |   |   |   \n---+---+---+---+---\n   |   |   |   |   \n---+---+---+---+---\n   |   |   |   |   "
	assert.Equal(t, es, g.String())

	// Loop through all coordinates
	for x := range 5 {
		for y := range 5 {
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
	es = " X | X | X | X | X \n---+---+---+---+---\n X | X | X | X | X \n---+---+---+---+---\n X | X | X | X | X \n---+---+---+---+---\n X | X | X | X | X \n---+---+---+---+---\n X | X | X | X | X "
	assert.Equal(t, es, g.String())

	// Ensure the set method returns an error for invalid coordinates
	err := g.Set(5, 0, Player1)
	require.Error(t, err)

	err = g.Set(-1, 0, Player1)
	require.Error(t, err)

	// Ensure the get method returns an error for invalid coordinates
	_, err = g.Get(5, 0)
	require.Error(t, err)

	_, err = g.Get(-1, 0)
	require.Error(t, err)
}

func TestWinner(t *testing.T) {
	g := NewGame(3)

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

	g = NewGame(3)

	// Ensure the game is won by Player2

	e1 = g.Set(0, 0, Player2)
	e2 = g.Set(1, 1, Player2)
	e3 = g.Set(2, 2, Player2)

	assert.NoError(t, e1)
	assert.NoError(t, e2)
	assert.NoError(t, e3)
	assert.Equal(t, Player2, g.Winner())

	// Ensure the game is won by Player1

	g = NewGame(3)

	e1 = g.Set(0, 0, Player1)
	e2 = g.Set(0, 1, Player1)
	e3 = g.Set(0, 2, Player1)

	assert.NoError(t, e1)
	assert.NoError(t, e2)
	assert.NoError(t, e3)
	assert.Equal(t, Player1, g.Winner())
}

func TestTogglePlayer(t *testing.T) {
	g := NewGame(3)

	// Ensure the game is empty
	assert.Equal(t, Player1, g.Player())

	// Ensure the game is Player2
	g.TogglePlayer()
	assert.Equal(t, Player2, g.Player())

	// Ensure the game is Player1
	g.TogglePlayer()
	assert.Equal(t, Player1, g.Player())
}

func TestNewGame(t *testing.T) {
	// Ensure the game is nil
	g := NewGame(1)
	assert.Nil(t, g)

	// Ensure the game is not nil
	g = NewGame(2)
	assert.NotNil(t, g)

	// Ensure the game is nil
	g = NewGame(6)
	assert.Nil(t, g)
}

func TestGetDimension(t *testing.T) {
	// Ensure the dimension is 2
	g := NewGame(2)
	assert.Equal(t, 2, g.GetDimension())

	// Ensure the dimension is 3
	g = NewGame(3)
	assert.Equal(t, 3, g.GetDimension())

	// Ensure the dimension is 4
	g = NewGame(4)
	assert.Equal(t, 4, g.GetDimension())

	// Ensure the dimension is 5
	g = NewGame(5)
	assert.Equal(t, 5, g.GetDimension())
}
