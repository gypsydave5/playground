package main

import "testing"

func TestSquishSpace(t *testing.T) {
	b := []byte("Hello,	世界")
	b2 := squishSpace(b)

	want := "Hello, 世界"
	got := string(b2)
	if got != want {
		t.Errorf("Expected %v to equal %v", got, want)
	}

	b = []byte("Hello,	世  	  界")
	b2 = squishSpace(b)

	want = "Hello, 世 界"
	got = string(b2)
	if got != want {
		t.Errorf("Expected %v to equal %v", got, want)
	}
}

func TestReverse(t *testing.T) {
	cases := []struct {
		init string
		want string
	}{
		{"Hello,	世界", "界世	,olleH"},
		{"This is a test", "tset a si sihT"},
		{"123456", "654321"},
		{"1234567", "7654321"},
		{"界世	,olleH", "Hello,	世界"},
	}

	for _, c := range cases {
		t.Run(c.init, func(t *testing.T) {
			b := []byte(c.init)
			reverse2(b)
			got := string(b)
			if got != c.want {
				t.Errorf("Expected %v to equal %v", got, c.want)
			}
		})
	}
}
