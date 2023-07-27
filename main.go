package main

import (
	"fmt"
	"github.com/briandowns/spinner"
	_ "math/rand"
	"strconv"
	"tictacgo/tictactoe"
	"time"
)

func main() {
	var t tictactoe.TicTacToe
	t.Board = [3][3]rune{
		{' ', ' ', ' '},
		{' ', ' ', ' '},
		{' ', ' ', ' '},
	}

Loop:
	for {
		var input string
		fmt.Println("Enter the player count (1 or 2):")
		fmt.Scanln(&input)

		playerCount, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("ERROR: Please enter a valid integer")
			continue Loop
		}

		switch playerCount {
		case 1:
			playerVSAI(&t)
			break Loop
		case 2:
			playerVsPlayer(&t)
			break Loop
		default:
			fmt.Println("ERROR: Please enter either 1 or 2")
			continue Loop
		}
	}
}

func playerVsPlayer(t *tictactoe.TicTacToe) {
	var gameStatus bool = true
	var player int = 1
	var tie int

	m := make(map[int]rune)
	m[1] = 'x'
	m[2] = 'o'

	t.PrintBoard()

	for gameStatus {
		var move string
		var row int
		var col int

		fmt.Printf("Choose where you would like to place your mark (%c), player %d\n", m[player], player)
		fmt.Scanln(&move)

		moveNum, err := strconv.Atoi(move)

		if err != nil {
			fmt.Println("ERROR: Please enter a valid integer")
			continue
		}

		if moveNum < 0 || moveNum > 8 {
			fmt.Println("ERROR: Please enter an integer from 0 to 8")
			continue
		}

		row = moveNum / 3
		col = moveNum % 3

		if t.Board[row][col] != ' ' {
			fmt.Println("ERROR: Please select an empty square")
			continue
		}

		t.Board[row][col] = m[player]
		t.PrintBoard()

		gameStatus, tie = tictactoe.GameStatus(*t)

		if player == 1 {
			player = 2
		} else {
			player = 1
		}
	}
	fmt.Println("-----GAME OVER-----")

	switch tie {
	case 0:
		fmt.Println("It was a tie...")
	case 1:
		fmt.Println("Player 1 (X) won!")
	case 2:
		fmt.Println("Player 2 (O) won!")
	}
}

func playerVSAI(t *tictactoe.TicTacToe) {
	var gameStatus bool = true
	var player int = 1
	var tie int

	m := make(map[int]rune)
	m[1] = 'x'
	m[2] = 'o'

	t.PrintBoard()

	for gameStatus {
		switch player {
		case 1:
			var move string
			var row int
			var col int

			fmt.Printf("Choose where you would like to place your mark (%c), player %d\n", m[player], player)
			fmt.Scanln(&move)

			moveNum, err := strconv.Atoi(move)

			if err != nil {
				fmt.Println("ERROR: Please enter a valid integer")
				continue
			}

			if moveNum < 0 || moveNum > 8 {
				fmt.Println("ERROR: Please enter an integer from 0 to 8")
				continue
			}

			row = moveNum / 3
			col = moveNum % 3

			if t.Board[row][col] != ' ' {
				fmt.Println("ERROR: Please select an empty square")
				continue
			}

			t.Board[row][col] = m[player]
			t.PrintBoard()

			gameStatus, tie = tictactoe.GameStatus(*t)

			player = 2
		case 2:
			s := spinner.New(spinner.CharSets[4], 100*time.Millisecond)
			fmt.Println("Calculating move...")
			s.Start()
			time.Sleep(time.Second)
			s.Stop()
			// var row int
			// var col int
			//
			// random := rand.Intn(9)
			// row = random / 3
			// col = random % 3
			//
			// if (t.Board[row][col] != ' '){
			//     continue
			// } else {
			//     t.Board[row][col] = m[player]
			//     t.PrintBoard()
			// }

			_, bestRow, bestCol := tictactoe.MiniMax(*t, 0, true)
			fmt.Println("bestRow is", bestRow)
			fmt.Println("bestCol is", bestCol)
			t.Board[bestRow][bestCol] = 'o'

			t.PrintBoard()

			// END
			gameStatus, tie = tictactoe.GameStatus(*t)
			player = 1
		}
	}
	fmt.Println("-----GAME OVER-----")

	switch tie {
	case 0:
		fmt.Println("It was a tie...")
	case 1:
		fmt.Println("Player 1 (X) won!")
	case 2:
		fmt.Println("Player 2 (O) won!")
	}
}
