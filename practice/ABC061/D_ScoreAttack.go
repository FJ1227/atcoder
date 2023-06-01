package main

import (
	"bufio"
	"fmt"
	"os"
)

type Edge struct {
	to   int
	cost int
}

const INF = 1 << 60

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m int
	fmt.Fscan(r, &n, &m)

	edges := make([][]Edge, n)

	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(r, &a, &b, &c)
		edges[a-1] = append(edges[a-1], Edge{b - 1, -c})
	}

	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}
	dist[0] = 0
	negative_cycle := false

	for iter := 0; iter <= n*2; iter++ {
		for i := 0; i < n; i++ {
			if dist[i] == INF {
				continue
			}
			for _, e := range edges[i] {
				if dist[e.to] > dist[i]+e.cost {
					dist[e.to] = dist[i] + e.cost
					if e.to == n-1 && iter == n*2 {
						negative_cycle = true
					}
				}
			}
		}
	}
	if negative_cycle {
		fmt.Fprint(w, "inf")
	} else {
		fmt.Fprint(w, -dist[n-1])
	}

}
