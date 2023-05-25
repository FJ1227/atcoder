package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var a = make([]int, n)
	var b = make([]int, n)
	var c = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i], &b[i], &c[i])
	}

	var dp = make([][3]int, n+1)

	for i := 1; i < n+1; i++ {
		if i == 1 {
			dp[1][0] = a[0]
			dp[1][1] = b[0]
			dp[1][2] = c[0]
		} else {
			dp[i][0] = max(dp[i-1][1], dp[i-1][2]) + a[i-1]
			dp[i][1] = max(dp[i-1][0], dp[i-1][2]) + b[i-1]
			dp[i][2] = max(dp[i-1][0], dp[i-1][1]) + c[i-1]
		}
	}

	fmt.Println(max(dp[n][0], max(dp[n][1], dp[n][2])))

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
