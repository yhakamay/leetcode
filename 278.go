/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
package main

func main() {
	var n int = 5

	firstBadVersion(n)
}

func firstBadVersion(n int) (ans int) {
	var low, high int = 1, n
	ans = n

	for low < high {
		mid := (low + high) / 2
		isMiddleBad := isBadVersion(mid)

		if isMiddleBad {
			high, ans = mid, mid
		} else {
			low = mid + 1
		}
	}

	return
}
