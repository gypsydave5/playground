package tempconv

import "testing"

func TestFToC(t *testing.T) {
	c := FToC(Fahrenheit(32))
	if c != Celsius(0) {
		t.Errorf("expected 0 Celsius, got %g %T", c, c)
	}
}

func TestCToF(t *testing.T) {
	f := CToF(Celsius(0))
	if f != Fahrenheit(32) {
		t.Errorf("expected 32 Fahrenheit, got %g %T", f, f)
	}
}

func TestCToK(t *testing.T) {
	k := CToK(Celsius(0))
	if k != Kelvin(273.15) {
		t.Errorf("Expected 273.15 Kelvin, got %g %T", k, k)
	}
}

func TestKToC(t *testing.T) {
	c := KToC(Kelvin(0))
	if c != Celsius(-273.15) {
		t.Errorf("Expected -273.15 Celsius, got %g %t", c, c)
	}
}

func TestFToK(t *testing.T) {
	k := FToK(Fahrenheit(212))
	if k != Kelvin(373.15) {
		t.Errorf("Expected 373.15 Kelvin, got %g %T", k, k)
	}
}

func TestKToF(t *testing.T) {
	f := KToF(Kelvin(273.15))
	if f != Fahrenheit(32) {
		t.Errorf("expected 32 Fahrenheit, got %g %T", f, f)
	}
}
