package main

type Integer int

func (i Integer) Less(j Integer) bool {
	return i < j
}

func (i *Integer) Add(j Integer) {
	*i += j
}

type LessAdder interface {
	Less(i Integer) bool
	Add(i Integer)
}

func main() {
	var i Integer = 1

	var j1 LessAdder = &i

	// this will cause an error
	var j2 LessAdder = i
}
