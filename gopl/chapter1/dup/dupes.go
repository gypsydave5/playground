package main

// Dup keeps track of the repetitions of a particular line across a number of
// files
type dup struct {
	files []string
	line  string
	total int
}

// Dupes is a slice of Dup, with an interface to sort on
type dupes []dup

func (ds dupes) Len() int {
	return len(ds)
}

func (ds dupes) Less(i, j int) bool {
	return ds[i].total > ds[j].total
}

func (ds dupes) Swap(i, j int) {
	ds[i], ds[j] = ds[j], ds[i]
}
