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

	var s, t string

	fmt.Fscan(r, &s, &t)

	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(t))
	}

	for i := range dp {
		for j := range dp[i] {
			if 0 < i {
				dp[i][j] = dp[i-1][j]
			}
			if 0 < j {
				dp[i][j] = max(dp[i][j], dp[i][j-1])
			}
			if s[i] == t[j] {
				p := 0
				if 0 < i && 0 < j {
					p = dp[i-1][j-1]
				}
				p++
				dp[i][j] = max(dp[i][j], p)
			}
		}
	}

	i := len(s) - 1
	j := len(t) - 1
	ans := make([]byte, dp[i][j])

	for 0 <= i && 0 <= j && 0 < dp[i][j] {
		for 0 < i && dp[i][j] == dp[i-1][j] {
			i--
		}
		for 0 < j && dp[i][j] == dp[i][j-1] {
			j--
		}
		ans[dp[i][j]-1] = s[i]
		i--
		j--
	}

	fmt.Fprintln(w, string(ans))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
