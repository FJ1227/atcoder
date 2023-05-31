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

	var N int
	fmt.Fscan(r, &N)

	p := make([]float64, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &p[i])
	}

	dp := make([][]float64, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]float64, N+1)
	}

	dp[0][0] = 1.0

	for i := 0; i < N; i++ {
		for j := 0; j < i+1; j++ {
			dp[i+1][j+1] += dp[i][j] * p[i]
			dp[i+1][j] += dp[i][j] * (1.0 - p[i])
		}
	}

	ans := 0.0

	for i := N/2 + 1; i < N+1; i++ {
		ans += dp[N][i]
	}

	fmt.Fprintln(w, ans)

}
