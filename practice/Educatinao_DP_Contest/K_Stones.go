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

	var n, k int
	fmt.Fscan(r, &n, &k)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	dp := make([]bool, k+1)

	for i := 0; i < k+1; i++ {
		for j := 0; j < n; j++ {
			if i-a[j] >= 0 && !dp[i-a[j]] {
				dp[i] = true
			}
		}
	}

	if dp[k] {
		fmt.Fprint(w, "First")
	} else {
		fmt.Fprint(w, "Second")
	}

}
