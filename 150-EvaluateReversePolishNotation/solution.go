package main

import (
    "fmt"
    "strconv"
)

func main() {
    token := []string{"2","1","+","3","*"}
    answer := evalRPN(token)
    fmt.Println(answer)
}

func evalRPN(tokens []string) int {
    digitStack := []int{}
    for _, token := range tokens {
        if val, err := strconv.Atoi(token); err == nil {
            digitStack = append(digitStack, val)
        } else {
            secondToken := digitStack[len(digitStack) - 1]
            firstToken := digitStack[len(digitStack) - 2]
            digitStack = digitStack[:len(digitStack) - 2]
            switch token {
                case "+":
                    digitStack = append(digitStack, int(firstToken) + int(secondToken))
                case "-":
                    digitStack = append(digitStack, int(firstToken) - int(secondToken))
                case "*":
                    digitStack = append(digitStack, int(firstToken) * int(secondToken))
                case "/":
                    digitStack = append(digitStack, int(firstToken) / int(secondToken))
            }
        }
    }
    return digitStack[len(digitStack) - 1]
}
