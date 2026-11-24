# Merge Two Sorted Arrays Without Extra Space

> Striver SDE Sheet 009 · Day 2 Arrays — the shrinking-gap (shell sort) trick over the virtual concatenation of both arrays. Go, stdlib only, tested.

## Problem

Given two sorted arrays, merge them so the first array holds the smallest values and the second the rest — both still sorted — using O(1) extra space. (With a scratch buffer this is the trivial two-pointer merge; the whole point is doing without one.)

## Approach

Treat the two arrays as one virtual array of length n+m, addressed through a helper that maps an index to the right slice. Then run shell-sort's gap pass: compare elements `gap` apart and swap when out of order, starting with gap = ⌈(n+m)/2⌉ and halving (ceiling division) until a final gap-1 pass. Because both halves were already sorted, these passes are enough to finish the merge — each pass moves every element close to its final home, and the gap-1 pass settles stragglers.

**Complexity:** O((n+m) · log(n+m)) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
