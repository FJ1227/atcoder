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
		fmt.Scan(&ws[i], &vs[i])
	}

	const INF = 1 << 60
	size := N*1000 + 1

	dp := make([][]int, N+1)

	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, size)
	}

	for i := 0; i < N+1; i++ {
		for j := 0; j < size; j++ {
			dp[i][j] = INF
		}
	}

	dp[0][0] = 0

	for i := 0; i < N; i++ {
		for j := 0; j < size; j++ {
			if j-vs[i] >= 0 {
				dp[i+1][j] = min(dp[i][j], dp[i][j-vs[i]]+ws[i])
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}

	var res int
	for i := 0; i < size; i++ {
		if dp[N][i] <= W {
			res = i
		}
	}

	fmt.Fprintln(w, res)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
