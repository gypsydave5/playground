package main

import "testing"

func TestZColorHexGreen(t *testing.T) {
	maxZ := 1.0
	minZ := -1.0
	z := 0.0
	c := hexColorByRange(maxZ, minZ, z)
	expectedC := "#00FF00"
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestZColorHexRed(t *testing.T) {
	maxZ := 1.0
	minZ := -1.0
	z := 1.0
	c := hexColorByRange(maxZ, minZ, z)
	expectedC := "#FF0000"
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestZColorHexRedAgain(t *testing.T) {
	maxZ := 0.9850673555377986
	minZ := -0.21285613860128652
	z := 0.9850673555377986
	c := hexColorByRange(maxZ, minZ, z)
	expectedC := "#FF0000"
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestZColorHexBlue(t *testing.T) {
	maxZ := 1.0
	minZ := -1.0
	z := -1.0
	c := hexColorByRange(maxZ, minZ, z)
	expectedC := "#0000FF"
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestZColorHexMid(t *testing.T) {
	maxZ := 1.0
	minZ := -1.0
	z := 0.5
	c := hexColorByRange(maxZ, minZ, z)
	expectedC := "#7F7F00"
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestZColorHexLowerMid(t *testing.T) {
	maxZ := 1.0
	minZ := -1.0
	z := -0.5
	c := hexColorByRange(maxZ, minZ, z)
	expectedC := "#007F7F"
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}
