package solute

// [题目信息](https://leetcode.cn/problems/maximum-swap)
func maximumSwap(num int) int {
	fields := split(num)
	for i := 0; i < len(fields); i++ {
		max := fields[i]
		maxIndex := i
		for j := i + 1; j < len(fields); j++ {
			if fields[j] >= max {
				max = fields[j]
				maxIndex = j
			}
		}
		if max > fields[i] {
			fields[i], fields[maxIndex] = fields[maxIndex], fields[i]
			break
		}
	}
	return join(fields)
}

func split(num int) []uint8 {
	var result []uint8
	for num > 0 {
		result = append([]uint8{uint8(num % 10)}, result...)
		num = num / 10
	}

	return result
}

func join(fields []uint8) int {
	r := 0
	for _, v := range fields {
		r = r*10 + int(v)
	}
	return r
}
