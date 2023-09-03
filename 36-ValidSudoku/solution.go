package main

import (
    "fmt"
)

/*
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
 

Example 1:


Input: board = 
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true
Example 2:

Input: board = 
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
*/

func main() {
    board := [][]string{
        {"8","3",".",".","7",".",".",".","."},
        {"6",".",".","1","9","5",".",".","."},
        {".","9","8",".",".",".",".","6","."},
        {"8",".",".",".","6",".",".",".","3"},
        {"4",".",".","8",".","3",".",".","1"},
        {"7",".",".",".","2",".",".",".","6"},
        {".","6",".",".",".",".","2","8","."},
        {".",".",".","4","1","9",".",".","5"},
        {".",".",".",".","8",".",".","7","9"},
    }
    answer := isValidSudoku(board)
    fmt.Println(answer)
}

func isValidSudoku(board [][]string) bool {
    checkMap := make(map[string] bool)
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            row := i
            col := j
            currentVal := string(board[row][col])

            if (currentVal == ".") {
                continue
            }

            _, rowCheck := checkMap[currentVal + "is in row: " + string(row)]
            _, colCheck := checkMap[currentVal + "is in col: " + string(col)]
            _, gridCheck := checkMap[currentVal + "is in grid: " + string(row/3) + "-" + string(col/3)]

            if rowCheck || colCheck || gridCheck {
                return false
            } else {
                checkMap[currentVal + "is in row: " + string(row)] = true
                checkMap[currentVal + "is in col: " + string(col)] = true
                checkMap[currentVal + "is in grid: " + string(row/3) + "-" + string(col/3)] = true
            }

        }
    }
    return true
}
