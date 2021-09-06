package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{10, 9, 2, 5, 4, 7, 3, 101, 18}
	fmt.Println(lengthOfLIS(nums1))
}

// TC is n * log n
// n is for looping over nums
// log n is for using BS to find first element that gte current number.

func lengthOfLIS(nums []int) int {
	dp := make([]int, 0)

	for i := 0; i < len(nums); i++ { // here is O(n)
		currentNum := nums[i]

		// condition 1: first element
		// condition 2: nums[i] gt than last element of dp
		if len(dp) == 0 || currentNum > dp[len(dp)-1] {
			dp = append(dp, currentNum)
		} else {
			// here using BS, so it is O(logn)
			iOfReplace := sort.Search(len(dp), func(i int) bool { return dp[i] >= currentNum })
			dp[iOfReplace] = currentNum
		}
		fmt.Println(dp)
	}
	return len(dp)
}
