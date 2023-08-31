package main

import (
    "fmt"
    "math"
)

/*
    Given a string s, find the length of the longest 
    substring without repeating characters.
     

    Example 1:

    Input: s = "abcabcbb"
    Output: 3
    Explanation: The answer is "abc", with the length of 3.
    Example 2:

    Input: s = "bbbbb"
    Output: 1
    Explanation: The answer is "b", with the length of 1.
    Example 3:

    Input: s = "pwwkew"
    Output: 3
    Explanation: The answer is "wke", with the length of 3.
    Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func main() {
    s := "abcabcbb"
    answer := findLongestSubstring(s)
    fmt.Println(answer)
}

func findLongestSubstring(s string) int {
    charMap := make(map[byte]bool)
    start := 0
    max := 0

    for end := range s {
        char := s[end]

        if (!charMap[char]) {
            charMap[char] = true
            max = int(math.Max(float64(max), float64(end - start + 1)))
        } else {
            for charMap[char] {
                charMap[s[start]] = false
                start++
            }
            charMap[char] = true
        } 
    }
    return int(max)
}
