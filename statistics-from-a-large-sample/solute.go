package solute

// [大样本统计](https://leetcode.cn/problems/statistics-from-a-large-sample/)
func sampleStats(count []int) []float64 {
	sum := 0.0
	// 最小/最大/众数
	min, max, mode := -1, 0, 0
	for v, c := range count {
		sum += float64(c)
		if c > 0 {
			if min == -1 {
				min = v
			}
			max = v
		}
		if c > count[mode] {
			mode = v
		}
	}

	// 平均数
	mean := 0.0
	for v, c := range count {
		mean += float64(v*c) / sum
	}

	// 中位数
	median := 0.0
	for i, j := 0, len(count)-1; i <= j; {
		for i < len(count) && count[i] == 0 {
			i++
		}
		for j >= 0 && count[j] == 0 {
			j--
		}
		if i > j {
			break
		}
		if i == j {
			median = float64(i)
			break
		}
		if count[i] == count[j] {
			median = (float64(i+j) / 2)
			count[i] = 0
			count[j] = 0
		} else if count[i] < count[j] {
			count[j] -= count[i]
			count[i] = 0

		} else {
			count[i] -= count[j]
			count[j] = 0
		}
	}

	return []float64{
		float64(min),
		float64(max),
		float64(mean),
		float64(median),
		float64(mode),
	}
}
