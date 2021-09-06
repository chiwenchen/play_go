package main

import (
	"fmt"
	"sort"
)

func main() {
	// Search 可以找出符合 f 的條件的第一筆資料的index
	// 例如以下的資料餵給 Seach，會得到 3，因為 nums[3] = 99 是第一個大於 22的數字
	nums := []int{1, 2, 3, 99, 200}
	result := sort.Search(len(nums), func(i int) bool { return nums[i] >= 23 })

	fmt.Println("index is", result)
	fmt.Println("numer is", nums[result])

	table := makeSliceofSlice(3, 5)

	for _, row := range table {
		fmt.Println(row)
	}
}

// 創造一個 slice of slice of int
func makeSliceofSlice(row, col int) [][]int {
	table := make([][]int, row)

	for i := range table {
		table[i] = make([]int, col)
	}

	return table
}
