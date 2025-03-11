package leetcode

/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	for _, v := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8} {
		verticle := validDirection(v, v, board, 1)   // 驗證縱向
		horizontal := validDirection(v, v, board, 0) // 驗證橫向
		if !verticle || !horizontal {
			return false
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !valid3x3(j*3, i*3, board) {
				return false
			}
		}
	}
	return true
}

func valid3x3(x, y int, board [][]byte) bool {
	hashset := make(map[byte]struct{}, 9)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[y+i][x+j] == '.' {
				continue
			}
			if _, ok := hashset[board[y+i][x+j]]; ok {
				return false
			}
			hashset[board[y+i][x+j]] = struct{}{}
		}
	}
	return true
}

func validDirection(x, y int, board [][]byte, direction int) bool {
	hashset := make(map[byte]struct{}, 9)
	if direction == 0 { // 橫向
		for i := 0; i < 9; i++ {
			if board[y][i] == '.' {
				continue
			}
			if _, ok := hashset[board[y][i]]; ok {
				return false
			}
			hashset[board[y][i]] = struct{}{}
		}
	} else {
		for i := 0; i < 9; i++ {
			if board[i][x] == '.' {
				continue
			}
			if _, ok := hashset[board[i][x]]; ok {
				return false
			}
			hashset[board[i][x]] = struct{}{}
		}
	}
	return true
}

// @lc code=end
