package solute

// [交替合并字符串](https://leetcode.cn/problems/merge-strings-alternately/)
func mergeAlternately(word1 string, word2 string) string {
	result := make([]byte, 0, len(word1)+len(word2))

	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			result = append(result, word1[i])
		}
		if i < len(word2) {
			result = append(result, word2[i])
		}
	}

	return string(result)
}
