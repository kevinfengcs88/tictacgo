package main

import (
	"fmt"
	"strconv"
	"tictacgo/tictactoe"
)

func main() {
    var t tictactoe.TicTacToe
    t.Board = [3][3]rune{
        {' ', ' ', ' '},
        {' ', ' ', ' '},
        {' ', ' ', ' '},
    }
    game(&t)
}

func game(t *tictactoe.TicTacToe) {
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
            fmt.Println("Please enter a valid integer")
            continue
        }

        if (moveNum < 0 || moveNum > 8) {
            fmt.Println("Please enter an integer from 0 to 8")
            continue
        }

        row = moveNum / 3
        col = moveNum % 3

        if (t.Board[row][col] != ' '){
            fmt.Println("Please select an empty square")
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
