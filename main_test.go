package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestMergeInPlace(t *testing.T) {
	cases := []struct {
		a, b         []int
		wantA, wantB []int
	}{
		{[]int{1, 4, 8, 10}, []int{2, 3, 9}, []int{1, 2, 3, 4}, []int{8, 9, 10}},
		{[]int{1, 3, 5, 7}, []int{0, 2, 6, 8, 9}, []int{0, 1, 2, 3}, []int{5, 6, 7, 8, 9}},
		{[]int{10, 12}, []int{5, 18, 20}, []int{5, 10}, []int{12, 18, 20}},
		{[]int{1}, []int{}, []int{1}, []int{}},
		{[]int{}, []int{}, []int{}, []int{}},
		{[]int{2}, []int{1}, []int{1}, []int{2}},
	}
	for _, c := range cases {
		a := append([]int{}, c.a...)
		b := append([]int{}, c.b...)
		MergeInPlace(a, b)
		if !reflect.DeepEqual(a, c.wantA) || !reflect.DeepEqual(b, c.wantB) {
			t.Errorf("MergeInPlace(%v,%v) = %v,%v want %v,%v", c.a, c.b, a, b, c.wantA, c.wantB)
		}
	}
}

// Property check: result must equal a full sort of the concatenation.
func TestMergeInPlaceAgainstSort(t *testing.T) {
	a := []int{3, 27, 38, 43}
	b := []int{9, 10, 82}
	all := append(append([]int{}, a...), b...)
	sort.Ints(all)
	MergeInPlace(a, b)
	got := append(append([]int{}, a...), b...)
	if !reflect.DeepEqual(got, all) {
		t.Errorf("merged concat %v != sorted concat %v", got, all)
	}
}
