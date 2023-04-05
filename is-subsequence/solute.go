package solute

// [判断子序列](https://leetcode.cn/problems/is-subsequence/)
func isSubsequence(s string, t string) bool {
	idx := 0
	for i := 0; i < len(s); i++ {
		j := idx
		for ; j < len(t); j++ {
			if s[i] == t[j] {
				idx = j + 1
				break
			}
		}
		if j == len(t) {
			return false
		}
	}
	return true
}
