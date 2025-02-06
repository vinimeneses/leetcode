package main

import "fmt"

func main() {
	examples := []string{"()", "()[]{}", "(]", "([])"}

	for _, s := range examples {
		fmt.Printf("Input: %s, Output: %v\n", s, isValid(s))
	}
}

func isValid(s string) bool {

	var stack []rune

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
