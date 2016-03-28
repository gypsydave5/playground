package weightconv

import "testing"

func TestPToK(t *testing.T) {
	k := PToK(Pound(1))
	if k != 0.45359237 {
		t.Errorf("Expected 0.45359237, got %g", k)
	}
}

func TestKToP(t *testing.T) {
	p := KToP(Kilogram(0.45359237))
	if p != 1 {
		t.Errorf("Expected 1, got %g", p)
	}
}

func TestPoundString(t *testing.T) {
	p := Pound(1).String()
	if p != "1lb" {
		t.Errorf("Expected \"1lb\", got %v", p)
	}
}

func TestKilogramString(t *testing.T) {
	k := Kilogram(5).String()
	if k != "5kg" {
		t.Errorf("Expected 5kg, got %v", k)
	}
}
