package solute

// [总持续时间可被 60 整除的歌曲](https://leetcode.cn/problems/pairs-of-songs-with-total-durations-divisible-by-60/)
func numPairsDivisibleBy60(time []int) int {
	st := make(map[int]int)
	for _, v := range time {
		st[v%60]++
	}
	count := 0
	for k, v := range st {
		if k == 30 || k == 0 {
			if st[k] > 1 {
				count += st[k] * (st[k] - 1)
			}
			continue
		}
		vv, ok := st[60-k]
		if ok {
			count += vv * v
		}
	}
	count = count / 2
	return count
}
