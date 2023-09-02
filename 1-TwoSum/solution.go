package main

import (
    "fmt"
)

func main() {
    nums := []int{2,7,3,5,8}
    k := 9
    answer := findTwoSum(nums, k)
    fmt.Println(answer)
}

func findTwoSum(nums []int, target int) []int {
    visited := make(map[int] int)
    for count, val := range nums {
        if idx, found := visited[target - val]; found {
            return []int{count, idx}
        }
        visited[val] = count
    }
    return nil
}
