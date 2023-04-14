package solute

// [反转字符串中的单词 III](https://leetcode.cn/problems/reverse-words-in-a-string-iii/)
func reverseWords(s string) string {
	rr := []byte(s)
	left, right := 0, 0
	for right < len(rr) {
		if rr[right] != ' ' {
			right++
			continue
		}

		if left != right {
			reverse(rr[left:right])
			left = right + 1
			right = right + 1
		}
	}
	reverse(rr[left:right])

	return string(rr)
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
