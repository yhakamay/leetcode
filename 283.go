package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)

	for _, v := range nums {
		fmt.Printf("%v ", v)
	}
}

func moveZeroes(nums []int) {
	i := 0

	for j := range nums {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}
