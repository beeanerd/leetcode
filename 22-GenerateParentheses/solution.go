package main

import (
    "fmt"
    "strings"
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
    var res []string
    var stack []string

    var backtrack func(int, int)

    backtrack = func (openParen int, closeParen int) {
        if openParen == n && closeParen == n && openParen == closeParen {
            res = append(res, strings.Join(stack, ""))
            return
        }
        if openParen < n {
            stack = append(stack, "(")
            backtrack(openParen + 1, closeParen)
            stack = stack[:len(stack) - 1]
        }
        if closeParen < openParen {
            stack = append(stack, ")")
            backtrack(openParen, closeParen + 1)
            stack = stack[:len(stack) - 1]
        }
    }
    backtrack(0, 0)

    return res
}


/*
    Base Case: Length of index = 2n
    
                      ""
                      /  
                    "("  
                 /      \  
               "(("     "()"
               / \      /
            "(((""(()""()("

*/
