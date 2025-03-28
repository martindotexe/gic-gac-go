package main

import (
	"fmt"
	"martindotexe/tic-tac-toe/game"
)

func main() {
	g := game.NewGame()

	for g.Winner() == game.Empty {
		fmt.Println(g)
		fmt.Println("Enter your move:")
		var x, y int
		_, err := fmt.Scan(&x, &y)
		if err != nil {
			panic(err)
		}
		err = g.Set(x, y, g.Player())
		if err != nil {
			panic(err)
		}
		g.TogglePlayer()
	}
	fmt.Println(g.Winner(), "wins!")
}
