package main

func main() {
	numbers, target := []int{-1, 0}, -1
	res := twoSum(numbers, target)
	println(res[0], res[1])
}

func twoSum(numbers []int, target int) (res []int) {
	res = make([]int, 2)

	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				res[0], res[1] = i+1, j+1
				break
			}
		}
	}

	return
}
