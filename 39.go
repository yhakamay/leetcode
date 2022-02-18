package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{2, 7, 6, 3}
	target := 7
	res := combinationSum(candidates, target)

	fmt.Println(res)
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	res := make([][]int, 0)

	for i := 0; i < len(candidates)-1; i++ {
		tmpSum := 0
		tmpRes := make([]int, 0)
		for j := i + 1; j < len(candidates); j++ {
			tmpRes = append(tmpRes, candidates[j])

			if tmpSum == target {
				res = append(res, tmpRes)
				break
			}

			if tmpSum+candidates[j] > target {
				break
			}

			tmpSum += candidates[j]
		}
	}

	return res
}
