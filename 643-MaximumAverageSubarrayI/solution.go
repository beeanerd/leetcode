package main

import (
	"fmt"
	"math"
)

func main() {
	k := 4
	nums := []int{1, 12, -5, -6, 50, 3}

	answer := findMaxAverageSubarray(nums, k)
	fmt.Println(answer)
}

func findMaxAverageSubarray(nums []int, k int) float64 {
	var rollingSum int
	var start int
	max := math.Inf(-1)

	for end := 0; end < len(nums); end++ {
		rollingSum += nums[end]

		if (end - start + 1) == k {
			max = math.Max(max, float64(rollingSum)/float64(k))
			rollingSum -= nums[start]
			start += 1
		}
	}
	return max
}
