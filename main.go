package main

import (
	"fmt"
	"log/slog"
	"martindotexe/tic-tac-toe/game"
)

func main() {
	g := game.NewGame()

	err := g.Set(1, 1, game.Player1)

	if err != nil {
		slog.Error(err.Error())
	}

	fmt.Println(g)
}
