package solute

func totalFruit(fruits []int) int {
	if len(fruits) == 0 {
		return 0
	}

	max := 0
	baskets := make(map[int]int)
	for left, right := 0, 0; right < len(fruits); right++ {
		for len(baskets) <= 2 && right < len(fruits) {
			baskets[fruits[right]] += 1
			right++
		}
		for len(baskets) > 2 {
			baskets[fruits[left]] -= 1
			if baskets[fruits[left]] == 0 {
				delete(baskets, fruits[left])
			}
			left++
		}
		if right >= len(fruits) {
			right = len(fruits) - 1
		}
		max = max2(max, right-left+1)
	}
	return max
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
