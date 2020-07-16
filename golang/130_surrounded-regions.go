var lenb int
var lenb2 int

func solve(board [][]byte) {
	lenb = len(board)
	if lenb == 0 {
		return
	}

	lenb2 = len(board[0])
	if lenb2 == 0 {
		return
	}

	for i := 0; i < lenb2; i++ {
		change(board, 0, i)
		change(board, lenb-1, i)
	}
	for i := 0; i < lenb; i++ {
		change(board, i, 0)
		change(board, i, lenb2-1)
	}

	for i := 0; i < lenb; i++ {
		for j := 0; j < lenb2; j++ {
			switch board[i][j] {
			case 'P':
				board[i][j] = 'O'
			case 'O':
				board[i][j] = 'X'
			}
		}
	}

}

func change(board [][]byte, x int, y int) {

	if x >= 0 && y >= 0 && x < lenb && y < lenb2 && board[x][y] == 'O' {
		board[x][y] = 'P'
		change(board, x-1, y)
		change(board, x+1, y)
		change(board, x, y-1)
		change(board, x, y+1)
	}

}