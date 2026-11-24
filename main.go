package main

import "fmt"

// MergeInPlace merges two individually sorted slices in place, so that a ends
// up holding the smallest len(a) values and b the rest, both sorted. It uses
// the shrinking-gap (shell-sort style) technique: compare and swap elements
// `gap` apart across the virtual concatenation a++b, halving the gap each round.
func MergeInPlace(a, b []int) {
	n, m := len(a), len(b)
	total := n + m
	if total < 2 {
		return
	}
	at := func(i int) *int {
		if i < n {
			return &a[i]
		}
		return &b[i-n]
	}
	for gap := (total + 1) / 2; ; gap = (gap + 1) / 2 {
		for i := 0; i+gap < total; i++ {
			p, q := at(i), at(i+gap)
			if *p > *q {
				*p, *q = *q, *p
			}
		}
		if gap == 1 {
			break
		}
	}
}

func main() {
	a := []int{1, 4, 8, 10}
	b := []int{2, 3, 9}
	MergeInPlace(a, b)
	fmt.Println(a, b)
}
