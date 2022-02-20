package main

func main() {
	nums := []int{5, 4, -1, 7, 8}
	maxSubArray := maxSubArray(nums)

	println(maxSubArray)
}

func maxSubArray(nums []int) (maxSubArray int) {
	tmpSum, maxSubArray := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		tmpSum = max(tmpSum+nums[i], nums[i])

		if tmpSum > maxSubArray {
			maxSubArray = tmpSum
		}
	}

	return
}

func max(num1 int, num2 int) (max int) {
	if num1 > num2 {
		max = num1
	} else {
		max = num2
	}

	return
}
