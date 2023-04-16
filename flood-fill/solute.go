package solute

// [图像渲染](https://leetcode.cn/problems/flood-fill/)
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}

	queue := [][]int{{sr, sc}}
	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]

		if image[x][y] == color {
			continue
		}

		list := [][]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
		for _, v := range list {
			if v[0] >= 0 && v[0] < len(image) && v[1] >= 0 && v[1] < len(image[v[0]]) && image[v[0]][v[1]] == image[x][y] {
				queue = append(queue, v)
			}
		}
		image[x][y] = color
	}

	return image
}
