package solute

// [01 矩阵](https://leetcode.cn/problems/01-matrix/)
func updateMatrix(mat [][]int) [][]int {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				mat[i][j] = -1
			}
		}
	}

	queue := make([][]int, 0)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]

		if x-1 >= 0 && mat[x-1][y] == -1 {
			queue = append(queue, []int{x - 1, y})
			mat[x-1][y] = mat[x][y] + 1
		}
		if x+1 < len(mat) && mat[x+1][y] == -1 {
			queue = append(queue, []int{x + 1, y})
			mat[x+1][y] = mat[x][y] + 1
		}
		if y-1 >= 0 && mat[x][y-1] == -1 {
			queue = append(queue, []int{x, y - 1})
			mat[x][y-1] = mat[x][y] + 1
		}
		if y+1 < len(mat[x]) && mat[x][y+1] == -1 {
			queue = append(queue, []int{x, y + 1})
			mat[x][y+1] = mat[x][y] + 1
		}
	}

	return mat
}
