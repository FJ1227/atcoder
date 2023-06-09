package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const MOD = 1000000007
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, k int
	fmt.Fscan(r, &n, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	dp := [110][100010]int{}

	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
			dp[i][j] %= MOD
			if j-a[i-1]-1 >= 0 {
				dp[i][j] -= dp[i-1][j-a[i-1]-1]
				dp[i][j] += MOD
				dp[i][j] %= MOD
			}
		}
	}

	fmt.Fprintln(w, dp[n][k])

}
