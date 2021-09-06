package main

import "fmt"

func main() {
	fmt.Println(fib(6))
	fmt.Println(fib(10))
	fmt.Println(fib(50))
}

func fib(n int) int {
	table := make([]int, n+1)

	table[1] = 1

	for i := 0; i < n; i++ {
		if i+1 <= n {
			table[i+1] += table[i]
		}

		if i+2 <= n {
			table[i+2] += table[i]
		}
	}

	return table[n]
}
