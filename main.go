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
    t.PrintBoard()
}

func game(t *tictactoe.TicTacToe) {
    var game bool = true
    var player int = 1
    var tie bool

    m := make(map[int]rune)
    m[1] = 'x'
    m[2] = 'o'

    for game {
        var move string
        var row int
        var col int
        // fmt.Println("Reference board (0-8 are viable moves):")
        // tictactoe.PrintReference()
        fmt.Printf("Choose where you would like to place your mark, player %d\n", player)
        fmt.Scanln(&move)

        moveNum, err := strconv.Atoi(move)

        if err != nil {
            fmt.Println("Please enter a valid integer")
            continue
        }

        fmt.Printf("%T\n", moveNum)
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

        game, tie = tictactoe.GameOver(*t)

        if player == 1 {
            player = 2
        } else {
            player = 1
        }
    }
    fmt.Println("-----GAME OVER-----")
    if tie {
        fmt.Println("It was a tie...")
    } else if player == 1 {
        fmt.Println("Player 2 (O) won!")
    } else if player == 2 {
        fmt.Println("Player 1 (X) won!")
    }
}

