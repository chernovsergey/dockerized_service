package finder

import (
	"log"
	"reflect"
	"testing"
)

var benchResult int

func BenchmarkLookup(b *testing.B) {
	f := Finder{}

	err := f.BuildIndex([]byte("hannabananna"))
	if err != nil {
		log.Fatal("Failed to build index")
	}

	b.ResetTimer()

	var offsets []int
	for i := 0; i < b.N; i++ {
		offsets = f.Lookup([]byte("ana"), -1)
	}
	b.StopTimer()
	benchResult = len(offsets)
}

func TestLookup(t *testing.T) {
	f := Finder{}
	err := f.BuildIndex([]byte("banana"))
	if err != nil {
		t.Fatal("Failed to build index")
	}

	offsets := f.Lookup([]byte("ana"), -1)
	if !reflect.DeepEqual(offsets, []int{3, 1}) {
		t.Fatal("Wrong lookup answer.")
	}
}
