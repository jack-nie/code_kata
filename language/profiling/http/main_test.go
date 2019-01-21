package main

import (
	"net/http/httptest"
	"testing"
)

func BenchmarkHiHandler(b *testing.B) {
	r := httptest.NewRequest("GET", "/", nil)
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		hiHandler(w, r)
	}
}
