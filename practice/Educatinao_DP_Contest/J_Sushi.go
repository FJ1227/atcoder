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

	sushis := [4]int{}
	for i := 0; i < n; i++ {
		var t int
		fmt.Fscan(r, &t)
		sushis[t]++
	}

	dp := [301][301][301]float64{}

	for k := 0; k <= n; k++ {
		for j := 0; j <= n; j++ {
			for i := 0; i <= n; i++ {
				if i+j+k == 0 {
					continue
				}
				exp := 1.0 * float64(n) / float64(i+j+k)
				if i > 0 {
					exp += dp[i-1][j][k] * float64(i) / float64(i+j+k)
				}
				if j > 0 {
					exp += dp[i+1][j-1][k] * float64(j) / float64(i+j+k)
				}
				if k > 0 {
					exp += dp[i][j+1][k-1] * float64(k) / float64(i+j+k)
				}
				dp[i][j][k] = exp
			}
		}
	}

	fmt.Fprint(w, dp[sushis[1]][sushis[2]][sushis[3]])

}
