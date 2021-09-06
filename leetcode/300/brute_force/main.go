package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))

	result := sort.Search(len(nums), func(i int) bool { return nums[i] >= 23 })

	fmt.Println(nums[result])
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		maxval := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				maxval = Max(maxval, dp[j])
			}
		}

		dp[i] = maxval + 1
	}

	return MaxEle(dp)
}

func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func MaxEle(x []int) int {
	max := 0
	for _, e := range x {
		max = Max(max, e)
	}

	return max
}

// [10,9,2,5,3,7,101,18]
