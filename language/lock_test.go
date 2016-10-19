package test

import (
	"sync"
	"testing"
)

func Foo1() {
	i := 0
	i++
}

func Foo2() {
	lock := new(sync.Mutex)
	i := 0
	lock.Lock()
	i++
	lock.Unlock()
}

func BenchmarkFool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Foo1()
	}
}

func BenchmarkFoo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Foo2()
	}
}
