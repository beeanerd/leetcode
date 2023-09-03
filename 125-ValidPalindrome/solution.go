package main

import (
	"fmt"
    "unicode"
)

/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase 
letters and removing all non-alphanumeric characters, it reads the same forward and 
backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

 

Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
*/

func main() {
    s := "A man, a plan, a canal: Panama"
    answer := isPalindrome(s)
    fmt.Println(answer)
}

func isPalindrome(s string) bool {
	end := len(s) - 1
    start := 0
	if end == 0 {
		return true
	}
	for start < end {
        if !unicode.IsDigit(rune(s[start])) && !unicode.IsLetter(rune(s[start])) {
            start++
            continue
        }
        if !unicode.IsDigit(rune(s[end])) && !unicode.IsLetter(rune(s[end])) {
            end--
            continue
        }
		if unicode.ToLower(rune(s[start])) != unicode.ToLower(rune(s[end])) {
			return false
		}
        start++
		end--
	}
	return true
}
