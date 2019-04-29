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
		{"23123123", "23,123,123"},
		{"+123123123", "+123,123,123"},
		{"+12311.23123456", "+12,311.23123456"},
	}

	for _, c := range cases {
		t.Run(c.input+"->"+c.want, func(t *testing.T) {
			got := comma(c.input)
			if got != c.want {
				t.Errorf("Got %#v, wanted %#v", got, c.want)
			}
		})
	}
}
