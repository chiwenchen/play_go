package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	result := twoSum(nums, 9)
	fmt.Println(result)

}

func twoSum(nums []int, target int) []int {
	// maiintain a hash map
	dict := make(map[int]int)
	ii := 0
	jj := 0

	for i, num := range nums {
		match := target - num

		if val, ok := dict[match]; ok {
			ii = i
			jj = val
		}
		dict[num] = i
	}

	return []int{ii, jj}
}
