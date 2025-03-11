package leetcode

/*
 * @lc app=leetcode id=289 lang=golang
 *
 * [289] Game of Life
 */

// @lc code=start
func gameOfLife(board [][]int) {
	y, x := len(board), len(board[0])
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			alive := getLiveNeighbors(j, i, board)
			if board[i][j] == 1 {
				// -1 標記為要死掉
				if alive < 2 || alive > 3 {
					board[i][j] = -1
				}
			} else {
				// -2 標記為要復活
				if alive == 3 {
					board[i][j] = -2
				}
			}
		}
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if board[i][j] == -1 {
				board[i][j] = 0
			} else if board[i][j] == -2 {
				board[i][j] = 1
			}
		}
	}
}

func getLiveNeighbors(x, y int, board [][]int) int {
	neighbors := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	alive := 0
	// 遍歷鄰居
	for _, direction := range neighbors {
		nextX, nextY := x+direction[0], y+direction[1]

		// 確保鄰居在邊界內
		if nextY >= 0 && nextY < len(board) && nextX >= 0 && nextX < len(board[0]) {
			if board[nextY][nextX] == 1 || board[nextY][nextX] == -1 {
				alive++
			}
		}
	}
	return alive
}

// func print(board [][]int){
//     for i := 0; i < len(board); i++ {
// 		for j := 0; j < len(board[i]); j++ {
// 			fmt.Print(board[i][j])
// 			fmt.Print(" ")
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println()
// 	fmt.Println()
// }

// @lc code=end
