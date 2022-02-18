package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-4, -1, 0, 3, 10}

	fmt.Println(sortedSquares(nums))
}

func sortedSquares(nums []int) (squaredNums []int) {
	for _, num := range nums {
		squaredNums = append(squaredNums, num*num)
	}

	sort.Ints(squaredNums)

	return
}
