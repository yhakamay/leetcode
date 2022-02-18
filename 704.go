package main

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 4

	println(search(nums, target))
}

func search(nums []int, target int) (ans int) {
	var low, high int = 0, len(nums) - 1

	ans = -1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			ans = mid
			break
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return
}
