package main

import (
	"index/suffixarray"
	"testing"
)

var benchResult int

func BenchmarkLookup(b *testing.B) {
	arr := suffixarray.New([]byte("banana"))
	b.ResetTimer()
	var offsets []int
	for i := 0; i < b.N; i++ {
		offsets = arr.Lookup([]byte("ana"), -1)
	}
	b.StopTimer()
	benchResult = len(offsets)
}
