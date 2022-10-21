package practice

import (
	"testing"
)

func TestSortInts(t *testing.T) {
	SortInts()
}

func TestSortFloat64(t *testing.T) {
	SortFloat64s()
}
func TestSortStrings(t *testing.T) {
	SortStrings()
}

//[][]int
func TestSort2dSlice(t *testing.T) {
	Sort2dInts()
}

//[]map[string]float64
func TestSortMaps(t *testing.T) {
	SortMaps()
}

//[]struct{}
func TestSortStructs(t *testing.T) {
	SortStructs()
}
