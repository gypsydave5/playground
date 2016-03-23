package main

import "testing"

func TestClosure(t *testing.T) {
    counter := closure()
    a := counter()
    b := counter()
    if a != 1 || b != 2 {
        t.Failed()
    }
}
