package model

import "testing"

func TestPlayerProgression(t *testing.T) {
	player := NewPlayer("p1", "Nova")
	player.AddXP(100)
	if player.Level != 2 || player.XP != 0 {
		t.Fatalf("unexpected progression: %+v", player)
	}
	if err := player.SpendCredits(999); err == nil {
		t.Fatal("expected credit failure")
	}
}

func TestPlayerLives(t *testing.T) {
	player := NewPlayer("p1", "Nova")
	if !player.LoseLife() || player.Lives != 2 {
		t.Fatalf("life was not lost: %+v", player)
	}
	player.Lives = 0
	if player.LoseLife() {
		t.Fatal("life loss should stop at zero")
	}
}
