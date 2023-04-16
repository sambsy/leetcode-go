package solute

// func lengthOfLongestSubstring(s string) int {
// 	length := 0
// 	count := make(map[byte]int)
// 	for left, right := 0, 0; right < len(s); right++ {
// 		count[s[right]]++
// 		if count[s[right]] == 1 {
// 			if cc := right - left + 1; cc > length {
// 				length = cc
// 			}
// 		} else {
// 			for count[s[right]] != 1 {
// 				count[s[left]]--
// 				left++
// 			}
// 		}
// 	}

// 	return length
// }

// [无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)
func lengthOfLongestSubstring(s string) int {
	max := 0
	for left, right := 0, 0; right < len(s); right++ {
		for i := left; i < right; i++ {
			if s[i] == s[right] {
				left = i + 1
				break
			}
		}
		if right-left+1 > max {
			max = right - left + 1
		}
	}

	return max
}
