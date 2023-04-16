package solute

// [岛屿的最大面积](https://leetcode.cn/problems/max-area-of-island/)
func maxAreaOfIsland(grid [][]int) int {
	max := 0

	mark := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		mark[i] = make([]bool, len(grid[i]))
	}

	queue := make([][]int, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 && !mark[i][j] {
				queue = [][]int{{i, j}}
				mark[i][j] = true
				area := 0
				for len(queue) > 0 {
					area++
					x, y := queue[0][0], queue[0][1]
					queue = queue[1:]

					if x-1 >= 0 && grid[x-1][y] == 1 && !mark[x-1][y] {
						queue = append(queue, []int{x - 1, y})
						mark[x-1][y] = true
					}

					if x+1 < len(grid) && grid[x+1][y] == 1 && !mark[x+1][y] {
						queue = append(queue, []int{x + 1, y})
						mark[x+1][y] = true
					}

					if y-1 >= 0 && grid[x][y-1] == 1 && !mark[x][y-1] {
						queue = append(queue, []int{x, y - 1})
						mark[x][y-1] = true
					}

					if y+1 < len(grid[x]) && grid[x][y+1] == 1 && !mark[x][y+1] {
						queue = append(queue, []int{x, y + 1})
						mark[x][y+1] = true
					}
				}
				if area > max {
					max = area
				}
			}
		}
	}
	return max
}
