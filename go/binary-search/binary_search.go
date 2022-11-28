package binarysearch

import (
	"sort"
)

func SearchInts(list []int, key int) int {
	sort.Ints(list)
	L := 0
	R := len(list) - 1
	var m int
	for L <= R {
		m = (L + R) / 2
		if list[m] < key {
			L = m + 1
		} else if list[m] > key {
			R = m - 1
		} else {
			return m
		}
	}
	return -1
}
