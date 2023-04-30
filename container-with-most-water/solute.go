package solute

// [盛最多水的容器](https://leetcode.cn/problems/container-with-most-water/)
func maxArea(height []int) int {
	max := 0
	i, j := 0, len(height)-1
	for i < j {
		a := area(i, height[i], j, height[j])
		if a > max {
			max = a
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func area(a, b, c, d int) int {
	min := b
	if d < min {
		min = d
	}

	return min * (c - a)
}
