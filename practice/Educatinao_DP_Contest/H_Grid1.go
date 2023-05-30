package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1000000007

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var H, W int
	fmt.Fscan(r, &H, &W)

	grid := make([][]bool, H)

	for i := 0; i < H; i++ {
		var s string
		fmt.Fscan(r, &s)
		grid[i] = to_list(s)
	}

	dp := make([][]int, H)
	for i := 0; i < H; i++ {
		dp[i] = make([]int, W)
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			dp[0][0] = 1
			if i+1 < H && grid[i+1][j] {
				dp[i+1][j] = dp[i+1][j] + dp[i][j]
				dp[i+1][j] %= MOD
			}
			if j+1 < W && grid[i][j+1] {
				dp[i][j+1] = dp[i][j+1] + dp[i][j]
				dp[i][j+1] %= MOD
			}
		}
	}

	fmt.Fprint(w, dp[H-1][W-1])

}

func to_list(v string) []bool {
	ret := make([]bool, len(v))
	for i := 0; i < len(v); i++ {
		if v[i] == '.' {
			ret[i] = true
		}
	}
	return ret
}
