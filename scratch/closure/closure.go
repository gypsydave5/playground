package main

func closure() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
