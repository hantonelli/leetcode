package tictactoe

type TicTacToe struct {
	board [][]bool    
}


func Constructor(n int) TicTacToe {
	board := make([][]bool, n)
	for i := 0;  i< n; i++ {
		board[i] = make([]bool, n)
	}
	return TicTacToe{board}    
}


func (t *TicTacToe) Move(row int, col int, player int) int {
	t.board[row][col] = true
}
