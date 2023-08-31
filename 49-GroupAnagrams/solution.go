/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

 

Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Example 2:

Input: strs = [""]
Output: [[""]]
Example 3:

Input: strs = ["a"]
Output: [["a"]]
*/

package main

import (
    "fmt"
)

func main() {
    strs := []string{"eat","tea","tan","ate","nat","bat"}
    answer := groupAnagrams(strs)
    fmt.Println(answer)
}

func groupAnagrams(strs []string) [][]string {
    groups := make(map[[26]int] []string)

    for _, str := range strs {
        var wordArray [26]int
        for _, char := range str {
            wordArray[char - 'a'] += 1
        }
        groups[wordArray] = append(groups[wordArray], str)
    }

    result := make([][]string, len(groups))
    index := 0
    for _, value := range groups {
        result[index] = value
        index++
    }
    return result
}
