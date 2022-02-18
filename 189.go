package main

import "fmt"

func main() {
	nums, k := []int{1, 2, 3, 4, 5, 6, 7}, 3
	rotatedNums := rotate(nums, k)
	fmt.Println(rotatedNums)
}

func rotate(nums []int, k int) []int {
	length := len(nums)
	rotatedNums := make([]int, length)

	if k >= length {
		k %= length
	}

	for i := 0; i < length; i++ {
		if i+k < length {
			rotatedNums[i+k] = nums[i]
		} else {
			rotatedNums[i+k-length] = nums[i]
		}
	}

	for i := 0; i < length; i++ {
		nums[i] = rotatedNums[i]
	}

	return nums
}
