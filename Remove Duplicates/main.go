package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1
}

func main() {
	nums := []int{1, 2, 2, 3, 4, 4, 5}
	fmt.Println(removeDuplicates(nums))
}
