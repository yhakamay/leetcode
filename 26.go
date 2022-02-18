package main

func main() {
	nums := []int{1, 1, 2}

	removeDuplicates(nums)
}

func removeDuplicates(nums []int) (ans int) {
	length := len(nums)

	if length <= 1 {
		return length
	}

	start := 0
	end := 1

	for end < length {
		if nums[start] == nums[end] {
			end++
		} else {
			start++
			nums[start] = nums[end]
		}
	}

	return start + 1
}
