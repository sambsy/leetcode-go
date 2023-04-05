package solute

var errVersion int

func isBadVersion(version int) bool {
	return version >= errVersion
}

// [第一个错误的版本](https://leetcode.cn/problems/first-bad-version/)
func firstBadVersion(n int) int {
	left, right := 1, n

	for left < right {
		mid := (left + right) / 2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
