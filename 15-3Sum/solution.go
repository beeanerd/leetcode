package main

import (
    "fmt"
    "slices"
)

/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that 
i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

 

Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation: 
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.
*/

func main() {
    nums := []int{-1,0,1,2,-1,-4}
    answer := threeSum(nums)
    fmt.Println(answer)
}

func threeSum(nums []int) [][]int {
    slices.Sort(nums)
    toReturn := make([][]int, 0)
    for idx, num := range nums {
        if idx > 0 && num == nums[idx - 1] {
            continue
        }
        start := idx + 1
        end := len(nums) - 1
        for start < end {
            threeSum := num + nums[start] + nums[end]
            if threeSum < 0 {
                start++
            } else if threeSum > 0 {
                end--
            } else {
                toReturn = append(toReturn, []int{num, nums[start], nums[end]})
                start++
                for nums[start] == nums[start - 1] && start < end {
                    start++

                }
            }
        }
    }
    return toReturn
}
