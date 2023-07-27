package tictactoe

import "fmt"

type TicTacToe struct {
	Board [3][3]rune
}

func (t TicTacToe) PrintBoard() {
	b := t.Board
	fmt.Println("  0 1 2  ")
	fmt.Println(" ------- ")
	for r := 0; r < len(b); r++ {
		if r != 1 {
			fmt.Print(" |")
		} else {
			fmt.Print("3|")
		}
		for c := 0; c < len(b[c]); c++ {
			fmt.Printf("%c", b[r][c])
			if c == 0 || c == 1 {
				fmt.Printf("|")
			}
		}
		if r != 1 {
			fmt.Print("| ")
		} else {
			fmt.Print("|5")
		}
		fmt.Println()
		if r != 2 {
			fmt.Println(" ------- ")
		}
	}
	fmt.Println(" ------- ")
	fmt.Println("  6 7 8  ")
}

// first is the game status (true = game can continue, false = game is over)
// second is the tie status (0 = tie, 1 = player 1 won, 2 = player 2 won)
func GameStatus(t TicTacToe) (bool, int) {
	b := t.Board
	var tie int

	for i := 0; i < len(b); i++ {
		unique := make(map[rune]bool)
		for _, value := range b[i] {
			unique[value] = true
		}
		if len(unique) == 1 && unique[' '] != true {
			_, ok := unique['x']
			if ok {
				tie = 1
			} else {
				tie = 2
			}
			return false, tie
		}
	}

	for j := 0; j < len(b[0]); j++ {
		unique := make(map[rune]bool)
		for i := 0; i < len(b); i++ {
			unique[b[i][j]] = true
		}
		if len(unique) == 1 && unique[' '] != true {
			_, ok := unique['x']
			if ok {
				tie = 1
			} else {
				tie = 2
			}
			return false, tie
		}
	}

	unique := make(map[rune]bool)
	unique[b[0][0]] = true
	unique[b[1][1]] = true
	unique[b[2][2]] = true
	if len(unique) == 1 && unique[' '] != true {
		_, ok := unique['x']
		if ok {
			tie = 1
		} else {
			tie = 2
		}
		return false, tie
	}
	delete(unique, b[0][0])
	delete(unique, b[2][2])
	unique[b[1][1]] = true
	unique[b[2][0]] = true
	unique[b[0][2]] = true
	if len(unique) == 1 && unique[' '] != true {
		_, ok := unique['x']
		if ok {
			tie = 1
		} else {
			tie = 2
		}
		return false, tie
	}

	delete(unique, b[1][1])
	delete(unique, b[0][2])
	delete(unique, b[2][0])

	// this can be made more efficient
	// just check if an x or an o exists
	// if so, return true, tie
	// afterwards (outside of scope), return false, tie

	// if we are able to loop through every square without encountering a space rune (' ')
	// then that means the game is over and it is a tie (false, tie)
	for _, row := range b {
		for _, item := range row {
			unique[item] = true
		}
	}

	if unique[' '] == false {
		return false, tie
	}

	return true, tie
}

// max being true means this call of the function is on the computer
// false means that the current call is on the player (and thus, we minimize the score)
func MiniMax(t TicTacToe, depth int, max bool) (int, int, int) {
	status, tie := GameStatus(t)

	// return 10 - depth if computer wins
	if status == false && tie == 2 {
		return 10 - depth, 0, 0
		// return -10 + depth if computer inevitably loses
	} else if status == false && tie == 1 {
		return -10 + depth, 0, 0
		// return 0 if it's a tie
	} else if status == false && tie == 0 {
		return 0, 0, 0
	}

	if max {
		bestScore := -69
		var row int
		var col int
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if t.Board[i][j] == ' ' {
					t.Board[i][j] = 'o'
					score, _, _ := MiniMax(t, depth+1, false)
					t.Board[i][j] = ' '
					if score > bestScore {
						bestScore = score
						row = i
						col = j
					}
				}
			}
		}
		return bestScore, row, col
	} else {
		bestScore := 69
		var row int
		var col int
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if t.Board[i][j] == ' ' {
					t.Board[i][j] = 'x'
					score, _, _ := MiniMax(t, depth+1, true)
					t.Board[i][j] = ' '
					if score < bestScore {
						bestScore = score
						row = i
						col = j
					}
				}
			}
		}
		return bestScore, row, col
	}
}
