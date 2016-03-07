package dupes

// Dup keeps track of the repetitions of a particular line across a number of
// files
type Dup struct {
	Files []string
	Line  string
	Total int
}

// Dupes is a slice of Dup, with an interface to sort on
type Dupes []Dup

func (ds Dupes) Len() int {
	return len(ds)
}

func (ds Dupes) Less(i, j int) bool {
	return ds[i].Total > ds[j].Total
}

func (ds Dupes) Swap(i, j int) {
	ds[i], ds[j] = ds[j], ds[i]
}
