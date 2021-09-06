package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 0, 0, 0}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	answers := [][]int{}

	sort.Ints(nums)

	hash := make(map[int]int)

	for index, num := range nums {
		hash[num] = index
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			target := -(nums[i] + nums[j])

			if i != 0 && nums[i] == nums[i-1] || i != j-1 && nums[j] == nums[j-1] {
				// 避免重複的情況
				continue
			}

			if index, ok := hash[target]; ok {
				if index <= i || index <= j {
					continue
				}
				answers = append(answers, []int{nums[i], nums[j], target})
			}
		}
	}

	return answers

}
