package main

func containsDuplicate(nums []int) bool {
	checked := make(map[int]bool)

	for _, v := range nums {
		if _, ok := checked[v]; ok {
			return true
		}
		checked[v] = true
	}

	return false

}
