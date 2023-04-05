package solute

// [同构字符串](https://leetcode.cn/problems/isomorphic-strings/)
func isIsomorphic(s string, t string) bool {
	mapping := make(map[byte]byte)
	revMapping := make(map[byte]byte)

	for i := 0; i < len(s) && i < len(t); i++ {
		value, ok := mapping[s[i]]
		if !ok {
			mapping[s[i]] = t[i]
		} else if value != t[i] {
			return false
		}

		value, ok = revMapping[t[i]]
		if !ok {
			revMapping[t[i]] = s[i]
		} else if value != s[i] {
			return false
		}
	}

	return true
}
