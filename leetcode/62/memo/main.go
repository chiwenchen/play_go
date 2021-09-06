package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(uniquePaths(1, 1))
	fmt.Println(uniquePaths(2, 2))
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(7, 7))
	fmt.Println(uniquePaths(10, 10))
}

func uniquePaths(m int, n int) int {
	memo := make(map[string]int)

	return countAndMemoPath(m, n, memo)
}

func countAndMemoPath(m int, n int, memo map[string]int) int {
	com := strconv.Itoa(m) + "," + strconv.Itoa(n)

	if val, ok := memo[com]; ok {
		return val
	}

	if m == 0 || n == 0 {
		return 0
	}

	if m == 1 || n == 1 {
		return 1
	}

	memo[com] = countAndMemoPath(m-1, n, memo) + countAndMemoPath(m, n-1, memo)

	return memo[com]
}
