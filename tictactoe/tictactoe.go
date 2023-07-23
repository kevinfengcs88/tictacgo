package tictactoe

import "fmt"

type TicTacToe struct {
    Board [3][3]rune
}

func (t TicTacToe) PrintBoard () {
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
            if (c == 0 || c == 1) {
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

func GameOver (t TicTacToe) (bool, bool) {
    b := t.Board

    for i := 0; i < len(b); i++ {
        unique := make(map[rune]bool)
        for _, value := range b[i] {
            unique[value] = true
        }
        if (len(unique) == 1 && unique[' '] != true) {
            return false, false
        }
    }

    for j := 0; j < len(b[0]); j++ {
        unique := make(map[rune]bool)
        for i := 0; i < len(b); i++ {
            unique[b[i][j]] = true
        }
        if (len(unique) == 1 && unique[' '] != true) {
            return false, false
        }
    }

    unique := make(map[rune]bool)
    unique[b[0][0]] = true
    unique[b[1][1]] = true
    unique[b[2][2]] = true
    if (len(unique) == 1 && unique[' '] != true) {
        return false, false
    }
    delete(unique, b[0][0])
    delete(unique, b[2][2])
    unique[b[1][1]] = true
    unique[b[2][0]] = true
    unique[b[0][2]] = true
    if (len(unique) == 1 && unique[' '] != true) {
        return false, false
    }

    delete(unique, b[1][1])
    delete(unique, b[0][2])
    delete(unique, b[2][0])

    for _, row := range b {
        for _, item := range row {
            unique[item] = true
        }
    }

    if unique[' '] == false {
        return false, true
    }

    return true, true
}

