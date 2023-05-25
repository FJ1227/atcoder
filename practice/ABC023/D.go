package main

import (
	"fmt"
	"math"
	"sort"
)

const MAX = math.MaxInt64

func main() {
	var n int
	fmt.Scan(&n)
	var h = make([]int, n)
	var s = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&h[i], &s[i])
	}

	var left = 0
	var right = MAX

	for right-left > 1 {
		var mid = (left + right) / 2
		ok := true
		var t = make([]int, n)
		for i := 0; i < n; i++ {
			if mid < h[i] {
				ok = false
			} else {
				t[i] = (mid - h[i]) / s[i]
			}
		}

		sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
		for i := 0; i < n; i++ {
			if t[i] < i {
				ok = false
			}
		}

		if ok {
			right = mid
		} else {
			left = mid
		}
	}
	fmt.Println(right)
}
