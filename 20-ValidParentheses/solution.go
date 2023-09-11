package main

import "fmt"

func main() {
    s := "(){}[]"    
    answer := isValid(s)
    fmt.Println(answer)
}

func isValid(s string) bool {
    parenStack := []string{}
    for _, c := range s {
        if c == '(' || c == '{' || c == '[' {
            parenStack = append(parenStack, string(c))
        } else {
            if (len(parenStack) == 0) {
                return false
            }
            check := parenStack[len(parenStack) - 1]
            parenStack = parenStack[0:len(parenStack) - 1]
            switch c {
                case ')':
                    if check != "(" {
                        return false
                    }
                case '}':
                    if check != "{" {
                        return false
                    }
                case ']':
                    if check != "[" {
                        return false
                    }
            }
        }
    }
    return len(parenStack) == 0 
}
