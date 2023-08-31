package main

import (
    "fmt"
    "math"
)

/*Given an array of positive integers nums and a positive integer target, return the minimal length of a */
/*subarray*/
/*whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.*/

 

/*Example 1:*/

/*Input: target = 7, nums = [2,3,1,2,4,3]*/
/*Output: 2*/
/*Explanation: The subarray [4,3] has the minimal length under the problem constraint.*/
/*Example 2:*/

/*Input: target = 4, nums = [1,4,4]*/
/*Output: 1*/
/*Example 3:*/

/*Input: target = 11, nums = [1,1,1,1,1,1,1,1]*/
/*Output: 0*/


func main() {
    target := 7
    nums := []int{2,3,1,2,4,3}
    answer := findMinimumSizeSubarraySum(target, nums)
    fmt.Println(answer)
}

func findMinimumSizeSubarraySum(target int, nums []int) int {
    start := 0
    min := 0
    rollingSum := 0

    for end := 0; end < len(nums); end++ {
        rollingSum += nums[end]

        for rollingSum >= target {
            if min == 0 {
                min = end - start + 1
            }
            min = int(math.Min(float64(min), float64(end - start + 1)))
            rollingSum -= nums[start]
            start++
        }
    }

    if (min == 0) {
        return 0
    }

    return int(min)
}
