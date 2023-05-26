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

	var N, W int
	fmt.Fscan(r, &N, &W)

	ws := make([]int, N)
	vs := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(r, &ws[i], &vs[i])
	}

	dp := make([][]int, N+1)

	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, W+1)
	}

	for i := 0; i < N; i++ {
		for j := 0; j <= W; j++ {
			if j-ws[i] >= 0 {
				dp[i+1][j] = max(dp[i][j], dp[i][j-ws[i]]+vs[i])
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}
	fmt.Fprint(w, dp[N][W])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
