package finder

import (
	"index/suffixarray"
)

type Finder struct {
	arr *suffixarray.Index
}

func (f *Finder) isEmpty() bool {
	return f.arr == nil
}

func (f *Finder) BuildIndex(text []byte) error {
	f.arr = suffixarray.New(text)
	return nil
}

func (f *Finder) Lookup(pat []byte, n int) []int {
	return f.arr.Lookup(pat, n)
}
