/*
Given an integer array nums and an integer k, return the k most frequent elements.
You may return the answer in any order.



Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
Example 2:

Input: nums = [1], k = 1
Output: [1]
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	answer := topKFrequent(nums, k)
	fmt.Println(answer)
}

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	bucketMap := make([][]int, len(nums)+1)

	for _, num := range nums {
		count[num] += 1
	}

	for num, occurences := range count {
		bucketMap[occurences] = append(bucketMap[occurences], num)
	}
	toReturn := make([]int, 0, k)
Loop:
	for idx := len(bucketMap) - 1; idx >= 0 && len(toReturn) < k; idx-- {
		for _, num := range bucketMap[idx] {
			if len(toReturn) >= k {
				break Loop
			}
			toReturn = append(toReturn, num)
		}
	}
	return toReturn
}
