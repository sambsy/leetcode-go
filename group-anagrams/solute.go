package solute

// [字母异位词分组](https://leetcode.cn/problems/group-anagrams/)
func groupAnagrams(strs []string) [][]string {
	result := make(map[[26]int][]string, 0)
	for _, str := range strs {
		var ww [26]int
		for _, v := range str {
			ww[v-'a']++
		}
		result[ww] = append(result[ww], str)
	}
	rr := make([][]string, 0, len(result))
	for _, v := range result {
		rr = append(rr, v)
	}
	return rr
}

// 指数分解的特性
// 对于可能出现溢出的情况，可能出现问题呢
func groupAnagrams2(strs []string) [][]string {
	mapping := []int{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
		31, 37, 41, 43, 47, 53, 59, 61, 67,
		71, 73, 79, 83, 89, 97, 101,
	}

	result := make(map[int][]string, 0)
	for _, str := range strs {
		score := 1
		for _, v := range str {
			score *= mapping[v-'a']
		}
		result[score] = append(result[score], str)
	}

	rr := make([][]string, 0, len(result))
	for _, v := range result {
		rr = append(rr, v)
	}
	return rr
}
