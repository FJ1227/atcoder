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
	a := make([]int, n)
	b := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i], &b[i])
	}

	sum := 0

	for i := n - 1; i >= 0; i-- {
		a[i] += sum
		if a[i]%b[i] != 0 {
			sum += b[i] - a[i]%b[i]
		}
	}

	fmt.Fprintln(w, sum)
}
