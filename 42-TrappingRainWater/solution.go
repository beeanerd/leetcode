package main

import (
    "fmt"
    "math"
)

/*
Given n non-negative integers representing an elevation map where the width of
each bar is 1, compute how much water it can trap after raining.


Example 1:

Input: height = [0,1,0,2,1,0,1,3,2,1,2,1] 
Output: 6 

Explanation: The above
elevation map (black section) is represented by array
[0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section)
are being trapped. Example 2:

Input: height = [4,2,0,3,2,5] Output: 9
*/

func main() {
    height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
    answer := trap(height)
    fmt.Println(answer)
}

func trap(height []int) int {
    maxLeft := make([]int, len(height))
    maxRight := make([]int, len(height))

    currLeft := 0
    currRight := 0
    for idx := range height {
        maxLeft[idx] = currLeft
        currLeft = int( math.Max(float64( currLeft ), float64( height[idx] )))
        maxRight[len(height) - idx - 1] = currRight
        currRight = int( math.Max(float64( currRight ), float64( height[len(height) - idx -1] )))
    } 

    rollingSum := 0

    for idx, num := range height {
        rollingSum += int(math.Max(0, float64(math.Min(float64(maxLeft[idx]), float64(maxRight[idx])) - float64(num))))
    }
    return rollingSum
}
