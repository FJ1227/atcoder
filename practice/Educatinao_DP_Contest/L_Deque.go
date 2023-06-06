package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < n+1; i++ {
		dp[i][i] = 0
	}

	for width := 1; width <= n; width++ {
		for l := 0; l+width <= n; l++ {
			r := l + width
			if width%2 == n%2 {
				dp[l][r] = max(dp[l+1][r]+a[l], dp[l][r-1]+a[r-1])
			} else {
				dp[l][r] = min(dp[l+1][r]-a[l], dp[l][r-1]-a[r-1])
			}
		}
	}

	fmt.Fprintln(w, dp[0][n])

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
