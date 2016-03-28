package lengthconv

import "testing"

func TestFToM(t *testing.T) {
	m := FToM(Foot(3048))
	if m != Meter(10000) {
		t.Errorf("Expected 10000, got %v", m)
	}
}

func TestMeterString(t *testing.T) {
	m := Meter(50).String()
	if m != "50m" {
		t.Errorf("Expected 50m, got %v", m)
	}
}
