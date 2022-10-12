package solute

var mm = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// [题目描述](https://leetcode-cn.com/problems/roman-to-integer/)
func romanToInt(s string) int {
	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		if i == len(s)-1 {
			result += mm[s[i]]
			continue
		}
		if mm[s[i]] < mm[s[i+1]] {
			result -= mm[s[i]]
		} else {
			result += mm[s[i]]
		}
	}

	return result
}
