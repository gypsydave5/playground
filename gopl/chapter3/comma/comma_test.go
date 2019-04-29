package main

import "testing"

// Exercise 3.10
func TestComma(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"10", "10"},
		{"9999", "9,999"},
		{"123123123", "123,123,123"},
	}

	for _, c := range cases {
		t.Run(c.input+"->"+c.want, func(t *testing.T) {
			got := comma(c.want)
			if got != c.want {
				t.Errorf("Got %#v, wanted %#v", got, c.want)
			}
		})
	}
}
