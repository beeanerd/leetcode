/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

 

Example 1:

Input: nums = [1,2,3,1]
Output: true
Example 2:

Input: nums = [1,2,3,4]
Output: false
Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
*/
package main

import (
    "fmt"
)

func main() {
    nums := []int{2,3,1}
    answer := containsDuplicate(nums)
    fmt.Println(answer)
}

func containsDuplicate(nums []int) bool {
    foundList := make(map[int] bool)

    for _, num := range nums {
        if (!foundList[num]) {
            foundList[num] = true
        } else {
            return true
        }
    }
    return false
}
