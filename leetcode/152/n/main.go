package main

import "fmt"

func main() {
	nums := []int{2, 3, -2, 4}
	fmt.Println("max product is", maxProduct(nums))
}
func maxProduct(nums []int) int {
	// store the result that is the max we have found so far
	r := nums[0]

	// imax/imin stores the max/min product of
	// subarray that ends with the current number A[i]
	for i := 1; i < len(nums); i++ {
		imax := r
		imin := r
		// multiplied by a negative makes big number smaller, small number bigger
		// so we redefine the extremums by swapping them
		if nums[i] < 0 {
			imax, imin = imin, imax
		}

		// max/min product for the current number is either the current number itself
		// or the max/min by the previous number times the current one
		imax = Max(nums[i], imax*nums[i])
		imin = Min(nums[i], imin*nums[i])

		// the newly computed max value is a candidate for our global result
		r = Max(r, imax)
	}
	return r
}

func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
