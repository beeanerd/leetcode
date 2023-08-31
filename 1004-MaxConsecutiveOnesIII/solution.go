/*
Given a binary array nums and an integer k, return the maximum number 
of consecutive 1's in the array if you can flip at most k 0's.

 
Example 1:

Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Explanation: [1,1,1,0,0,1,1,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

Example 2:

Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
Output: 10
Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
*/

/*
    Store 1 count and 0 count, when 0 count exceeds k, start at the next index after the first 0
*/

package main

import (
    "fmt"
    "math"
)
func main() {
    oneArray := []int{1,1,1,0,0,0,1,1,1,1,0}
    k := 2
    answer := longestOnes(oneArray, k)
    fmt.Println(answer)
}

func longestOnes(nums []int, k int) int {
    zeroCount := 0
    firstSwappedIndex := 0
    start := 0

    max := math.Inf(-1)

    for end := 0; end < len(nums); end++ {
        if nums[end] == 0 {
            if zeroCount == 0 {
                firstSwappedIndex = end
            } 
            if zeroCount >= k {
                firstSwappedIndex++
                start = firstSwappedIndex
                for firstSwappedIndex < len(nums) && nums[firstSwappedIndex] == 1 {
                    firstSwappedIndex++
                }
                zeroCount--
            }
            zeroCount++
        }
        max = math.Max(float64(max), float64(end - start + 1))
    }
    return int(max)
}
