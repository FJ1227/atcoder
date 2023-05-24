package educatinaodpcontest

import (
	"fmt"
	"math"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
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
		for j := 1; j <= k; j++ {
			if i+j < n {
				chmin(&dp[i+j], dp[i]+abs(h[i]-h[i+j]))
			}
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
