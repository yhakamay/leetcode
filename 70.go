package main

func main() {}

func climbStairs(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	}

	oneStepBefore, twoStepsBefore, allWays := 2, 1, 0

	for i := 2; i < n; i++ {
		allWays = oneStepBefore + twoStepsBefore
		twoStepsBefore = oneStepBefore
		oneStepBefore = allWays
	}

	return allWays
}
