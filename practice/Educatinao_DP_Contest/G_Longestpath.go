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

	var n, m int
	fmt.Fscan(r, &n, &m)

	paths := make([][]int, n)
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(r, &x, &y)
		paths[x-1] = append(paths[x-1], y-1)
	}

	memo := make([]int, n)

	for i := 0; i < n; i++ {
		dfs(i, paths, memo)
	}

	ans := 0
	for _, v := range memo {
		ans = max(ans, v)
	}

	fmt.Fprint(w, ans)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func dfs(i int, paths [][]int, memo []int) int {
	if memo[i] != 0 {
		return memo[i]
	}

	for _, to := range paths[i] {
		memo[i] = max(memo[i], dfs(to, paths, memo)+1)
	}

	return memo[i]
}
