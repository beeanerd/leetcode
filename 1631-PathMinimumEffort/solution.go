/*
    You are a hiker preparing for an upcoming hike. You are given heights, a 2D
    array of size rows x columns, where heights[row][col] represents the height
    of cell (row, col). You are situated in the top-left cell, (0, 0), and you
    hope to travel to the bottom-right cell, (rows-1, columns-1) (i.e.,
    0-indexed). You can move up, down, left, or right, and you wish to find a
    route that requires the minimum effort.

    A route's effort is the maximum absolute difference in heights between two
    consecutive cells of the route.

    Return the minimum effort required to travel from the top-left cell to the
    bottom-right cell.
*/

package main

import (
    "fmt"
)

func main() {

}

func minimumEffortPath(heights [][]int) int {
    var findMinimumPath func(int, int, [][]int) int

    findMinimumPath = func(row int, col int, heights[][]int) int {
        if row == len(heights) - 1 && col == len(heights[0]) - 1 {
            return heights[row][col]
        }
        // Up Case
        if row > 0 {

        }
        // Down Case
        // Right Case
        // Left Case
    }    
}
