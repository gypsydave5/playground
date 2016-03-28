package lengthconv

import "testing"

func TestFToM(t *testing.T) {
	m := FToM(Foot(3048))
	if m != Meter(10000) {
		t.Errorf("Expected 10000, got %v", m)
	}
}

func TestMToF(t *testing.T) {
	f := MToF(Meter(1))
	if f != 0.3048 {
		t.Errorf("Expected 0.3048, got %v", f)
	}
}

func TestMeterString(t *testing.T) {
	m := Meter(50).String()
	if m != "50 m" {
		t.Errorf("Expected 50 m, got %v", m)
	}
}

func TestFootString(t *testing.T) {
	f := Foot(6).String()
	if f != "6 ft" {
		t.Errorf("Expected 6 ft, got %v", f)
	}
}
