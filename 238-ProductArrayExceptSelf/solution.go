package main

import (
    "fmt"
)

func main() {
    nums := []int{-1,1,0,-3,3}
    answer := productExceptSelf(nums)
    fmt.Println(answer)
}

func productExceptSelf(nums []int) []int {
    length := len(nums) 
    if (length == 1) {
        return nums
    }
    prefixArray := make([]int, length)
    postfixArray := make([]int, length)
    solutionArray := make([]int, length)

    for idx := 0; idx < length; idx++ {
        if (idx == 0) {
            prefixArray[idx] = nums[idx]
            postfixArray[length - idx - 1] = nums[length - idx - 1]
        } else {
            prefixArray[idx] = prefixArray[idx - 1] * nums[idx]
            postfixArray[length - idx - 1] = postfixArray[length - idx] * nums[length - idx - 1]
        }
    }
    solutionArray[0] = postfixArray[1] 
    solutionArray[length - 1] = prefixArray[length - 2]
    for idx := 1; idx < length - 1; idx++ {
        solutionArray[idx] = prefixArray[idx - 1] * postfixArray[idx + 1]
    }
    return solutionArray
}
