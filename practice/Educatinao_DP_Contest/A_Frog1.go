package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	var h = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&h[i])
	}

	var MAX = math.MaxInt64

	var dp = make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = MAX
	}

	dp[0] = 0

	for i := 0; i < n; i++ {
		if i > 0 {
			chmin(&dp[i], dp[i-1]+abs(h[i]-h[i-1]))
		}
		if i > 1 {
			chmin(&dp[i], dp[i-2]+abs(h[i]-h[i-2]))
		}
	}

	fmt.Println(dp[n-1])
}

func chmin(a *int, b int) {
	if *a > b {
		*a = b
	}
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
