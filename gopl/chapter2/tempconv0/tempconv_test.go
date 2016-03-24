package tempconv

import "testing"

func TestFToC(t *testing.T) {
	c := FToC(32)
	if c != Celsius(0) {
		t.Errorf("expected 0, got %v", c)
	}
}

func TestCToF(t *testing.T) {
	f := CToF(0)
	if f != Farenheit(32) {
		t.Errorf("expected 0, got %v", f)
	}
}

func TestCString(t *testing.T) {
	c := Celsius(32).String()
	if c != "32°C" {
		t.Errorf("expected 32°C, got %v", c)
	}
}

func TestFString(t *testing.T) {
	f := Farenheit(32).String()
	if f != "32°F" {
		t.Errorf("expected 32°F, got %v", f)
	}
}
