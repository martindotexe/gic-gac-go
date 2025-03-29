package game

import "testing"

func BenchmarkNewGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewGame(3)
	}
}

func BenchmarkGetDimension(b *testing.B) {
	g := NewGame(3)
	for i := 0; i < b.N; i++ {
		g.GetDimension()
	}
}

func BenchmarkWinner(b *testing.B) {
	g := NewGame(3)
	for i := 0; i < b.N; i++ {
		g.Winner()
	}
}

func BenchmarkSet(b *testing.B) {
	g := NewGame(3)
	for i := 0; i < b.N; i++ {
		_ = g.Set(0, 0, Player1)
	}
}

func BenchmarkGet(b *testing.B) {
	g := NewGame(3)
	for i := 0; i < b.N; i++ {
		_, _ = g.Get(0, 0)
	}
}

func BenchmarkTogglePlayer(b *testing.B) {
	g := NewGame(3)
	for i := 0; i < b.N; i++ {
		g.TogglePlayer()
	}
}

func BenchmarkString(b *testing.B) {
	g := NewGame(3)
	for i := 0; i < b.N; i++ {
		_ = g.String()
	}
}

func BenchmarkValidMove(b *testing.B) {
	g := NewGame(3)
	for i := 0; i < b.N; i++ {
		g.validMove(0, 0)
	}
}

func BenchmarkBoardBitPosition(b *testing.B) {
	g := NewGame(3)
	for i := 0; i < b.N; i++ {
		g.boardBitPosition(0, 0)
	}
}

func BenchmarkPlayerString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Player1.String()
	}
}

func BenchmarkPlayerString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Player2.String()
	}
}

func BenchmarkPlayerStringEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Empty.String()
	}
}
