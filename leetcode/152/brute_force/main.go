package main

import "fmt"

func main() {
	// nums := []int{2, 3, -2, 4}
	nums := []int{-2}
	fmt.Println("max product is", maxProduct(nums))
}

func maxProduct(nums []int) int {
	maxProduct := nums[0]

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			result := 1
			for _, num := range nums[i : j+1] {
				result *= num
			}

			if maxProduct < result {
				maxProduct = result
			}
		}
	}

	return maxProduct
}
