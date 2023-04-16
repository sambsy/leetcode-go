package solute

// [字符串的排列](https://leetcode.cn/problems/permutation-in-string/)
func checkInclusion(s1 string, s2 string) bool {
	charset1 := make(map[byte]int)

	for _, v := range []byte(s1) {
		charset1[v]++
	}

	charset2 := make(map[byte]int)
	for left, right := 0, 0; right < len(s2); right++ {
		charset2[s2[right]]++
		if charset2[s2[right]] < charset1[s2[right]] {
			continue
		}
		if charset2[s2[right]] == charset1[s2[right]] {
			if judge(charset1, charset2) {
				return true
			}
			continue
		}

		for charset2[s2[right]] > charset1[s2[right]] {
			charset2[s2[left]]--
			left++
		}

	}

	return false
}

func judge(set1, set2 map[byte]int) bool {
	for k, v := range set1 {
		if set2[k] != v {
			return false
		}
	}
	return true
}
