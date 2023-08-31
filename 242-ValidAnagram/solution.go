/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

 

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
*/

package main

import (
    "fmt"
)

func main() {
    s := "anagram"
    t := "nagaram"
    answer := isAnagram(s, t)
    fmt.Println(answer)
}

func isAnagram(s string, t string) bool {
    checkString := make(map[byte] int)
    for _, char := range s {
        checkString[byte(char)] += 1
    }
    for _, char := range t {
        checkString[byte(char)] -= 1
        if checkString[byte(char)] < 0 {
            return false
        }
    }
    return true
}
