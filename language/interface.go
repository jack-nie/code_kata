package main

import (
	"fmt"
)

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
	//var i Integer = 1

	//var j1 LessAdder = &i

	// this will cause an error
	// var j2 LessAdder = i

	cache := New(12)
	fmt.Println(cache.MaxEntries)
	cache.Remove()

	s := make(map[string]int)
	if m, ok := s["jack"]; !ok {
		fmt.Println(m)
		fmt.Println(ok)
	}
}

type Cache struct {
	MaxEntries int
	OnEvicted  func(key Key, value interface{})
}

type Key interface{}

func New(maxEntries int) *Cache {
	return &Cache{
		MaxEntries: maxEntries,
		OnEvicted: func(key Key, value interface{}) {
			fmt.Println("Hello", key, value)
		},
	}
}

func (c *Cache) Remove() {
	if c.OnEvicted != nil {
		c.OnEvicted("jack", "nie")
	}
}
