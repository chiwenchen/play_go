package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	pauseAndClear()
	// fmt.Println(uniquePaths(1, 1))
	// fmt.Println(uniquePaths(2, 2))
	fmt.Println(uniquePaths(3, 2))
	// fmt.Println(uniquePaths(7, 7))
	// fmt.Println(uniquePaths(10, 10))
}

type row []int

func uniquePaths(m int, n int) int {
	table := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		table[i] = make([]int, n+1)
	}

	table[1][1] = 1

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {

			if i < m {
				table[i+1][j] += table[i][j]
			}
			if j < n {
				table[i][j+1] += table[i][j]
			}

			printTable(table, i, j)
			pauseAndClear()
		}
	}

	return table[m][n]
}

func printTable(table [][]int, i, j int) {
	// this can print the full table
	fmt.Println("i:", i, "j:", j, "------------------")
	for _, row := range table {
		fmt.Println(row)
	}
}

func pauseAndClear() {
	time.Sleep(500 * time.Millisecond)
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
