package main

import "fmt"

// ArraysAndSlices is the playground for arrays. And slices.
// Much of this is from the blog post on slices on the golang blog
// see http://blog.golang.org/go-slices-usage-and-internals
func ArraysAndSlices() {
	golang := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	a := golang[:2]
	b := golang[2:]
	c := golang[1:4]

	fmt.Printf("golang: %T %v\n", golang, golang)
	fmt.Printf("a: %T %v\n", a, a)
	fmt.Printf("b: %T %v\n", b, b)
	fmt.Printf("c: %T %v\n", c, c)

	a[0] = 'o'
	golang[2] = 'o'

	fmt.Printf("golang: %T %v\n", golang, golang)

	fmt.Println(string(golang)) //byte slice to string

	fmt.Printf("a: %T %v\n", a, a)
	fmt.Printf("b: %T %v\n", b, b)
	fmt.Printf("c: %T %v\n", c, c)

	d := a[:cap(a)]
	fmt.Printf("d: %T %v\n", d, d)

	// double slice length
	t := make([]byte, len(d), (cap(d)+1)*2)
	for i := range d {
		t[i] = d[i]
	}

	//or
	s := make([]byte, len(d), (cap(d)+1)*2)
	copy(d, s)

	fmt.Printf("d: cap(d) == %v, len(d) == %v\n", cap(d), len(d))
	fmt.Printf("t: cap(t) == %v, len(t) == %v\n", cap(t), len(t))

	fmt.Printf("t: %T %v\n", t, t)
	z := t[:cap(t)]
	fmt.Printf("z: %T %v\n", z, z)

}
