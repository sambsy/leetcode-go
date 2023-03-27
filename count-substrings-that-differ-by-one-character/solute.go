package solute

// [统计只差一个字符的子串数目](https://leetcode.cn/problems/count-substrings-that-differ-by-one-character/)
func countSubstrings(s string, t string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			diff := 0
			for k := 0; i+k < len(s) && j+k < len(t); k++ {
				if s[i+k] != t[j+k] {
					diff++
				}
				if diff == 1 {
					sum += 1
				}
				if diff > 1 {
					break
				}
			}
		}
	}
	return sum
}
