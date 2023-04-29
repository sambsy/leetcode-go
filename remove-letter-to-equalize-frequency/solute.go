package solute

// [删除字符使频率相同](https://leetcode.cn/problems/remove-letter-to-equalize-frequency/)
func equalFrequency(word string) bool {
	freq := make(map[byte]int, 26)
	for _, v := range word {
		freq[byte(v)]++
	}

	tt := word[0]
	for t, v := range freq {
		if v != freq[tt] {
			return judge(freq, t, v-1) || judge(freq, tt, freq[tt]-1)
		}
		tt = t
	}

	return freq[tt] == 1 || len(freq) == 1
}

func judge(freq map[byte]int, dd byte, value int) bool {
	for d, v := range freq {
		if d == dd {
			continue
		}
		if value == 0 {
			value = v
		}
		if v != value {
			return false
		}
	}
	return true
}
