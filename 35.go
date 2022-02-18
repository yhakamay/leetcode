package main

func main() {
	nums := []int{1, 3, 5, 6}
	target := 7

	println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) (ans int) {
	var startIndex, endIndex int = 0, len(nums) - 1

	for startIndex <= endIndex {
		midIndex := (startIndex + endIndex) / 2

		if target < nums[0] {
			ans = 0
			break
		}

		if target > nums[len(nums)-1] {
			ans = len(nums)
			break
		}

		// Hit the target
		if nums[midIndex] == target {
			ans = midIndex
			break
		}

		// Insert the target between two nums
		if target > nums[midIndex] && target < nums[midIndex+1] {
			ans = midIndex + 1
			break
		}

		// Continue searching
		if target > nums[midIndex] {
			startIndex = midIndex + 1
		} else {
			endIndex = midIndex - 1
		}
	}

	return
}
