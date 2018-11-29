package main

import (
	"math/rand"
	"sync"
	"testing"
)

type RegularIntMap struct {
	sync.RWMutex
	internal map[int]int
}

func NewRegularIntMap() *RegularIntMap {
	return &RegularIntMap{
		internal: make(map[int]int),
	}
}

func (rm *RegularIntMap) Load(key int) (value int, ok bool) {
	rm.RLock()
	result, ok := rm.internal[key]
	rm.RUnlock()
	return result, ok
}

func (rm *RegularIntMap) Delete(key int) {
	rm.RLock()
	delete(rm.internal, key)
	rm.RUnlock()
}

func (rm *RegularIntMap) Store(key, value int) {
	rm.Lock()
	rm.internal[key] = value
	rm.Unlock()
}

func nrand(n int) []int {
	i := make([]int, n)
	for ind := range i {
		i[ind] = rand.Int()
	}
	return i
}

func BenchmarkStoreRegular(b *testing.B) {
	nums := nrand(b.N)
	rm := NewRegularIntMap()
	b.ResetTimer()
	for _, v := range nums {
		rm.Store(v, v)
	}
}

func BenchmarkStoreSync(b *testing.B) {
	nums := nrand(b.N)
	var sm sync.Map
	b.ResetTimer()
	for _, v := range nums {
		sm.Store(v, v)
	}
}

func BenchmarkDeleteRegular(b *testing.B) {
	nums := nrand(b.N)
	rm := NewRegularIntMap()
	for _, v := range nums {
		rm.Store(v, v)
	}
	b.ResetTimer()

	for _, v := range nums {
		rm.Delete(v)
	}
}

func BenchmarkDeleteSync(b *testing.B) {
	nums := nrand(b.N)
	var sm sync.Map
	for _, v := range nums {
		sm.Store(v, v)
	}
	b.ResetTimer()
	for _, v := range nums {
		sm.Delete(v)
	}
}
