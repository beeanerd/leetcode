package main

import (
    "fmt"
)

/*
Given n pairs of parentheses, write a function to generate all combinations of
well-formed parentheses.


Example 1:

Input: n = 3 Output: ["((()))","(()())","(())()","()(())","()()()"]

Example 2:

Input: n = 1 Output: ["()"]

*/

func main() {
    n := 3
    answer := generateParenthesis(n)
    fmt.Println(answer)
}

func generateParenthesis(n int) []string {
    
}

func generateParenthesisHelper(remaining int, toAppend *[]string) {
    if remaining == 0
}
