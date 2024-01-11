package solute

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	flags := make(map[byte]int)

	left, right := 0, 0

	for right < len(s) {
		for i := right; i < len(s); i++ {
			if _, exist := flags[s[i]]; !exist {
				flags[s[i]] = i
				continue
			}
			right = i
			break
		}

		if len(flags) > maxLength {
			maxLength = len(flags)
		}

		for ; left < right; left++ {
			if s[left] != s[right] {
				delete(flags, s[left])
				continue
			}
			break
		}
		if left != right {
			delete(flags, s[left])
			left++
		}
		flags[s[right]] = right
		right++
	}
	return maxLength
}

/*
pwwkew
pw
w
*/
