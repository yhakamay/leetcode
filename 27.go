package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2

	fmt.Println(removeElement(nums, val))
}

func removeElement(nums []int, val int) (ans int) {
	for _, num := range nums {
		if num != val {
			nums[ans] = num
			ans++
		}
	}

	return
}
