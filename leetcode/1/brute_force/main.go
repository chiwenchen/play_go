package main

func twoSum(nums []int, target int) []int {
	var ii, jj int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j <= i {
				continue
			}

			if nums[i]+nums[j] == target {
				ii = i
				jj = j
				break
			}

		}
	}

	return []int{ii, jj}
}
