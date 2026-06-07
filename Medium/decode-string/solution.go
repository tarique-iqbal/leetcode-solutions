package main

import (
	"fmt"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}

	currentNum := 0
	currentStr := ""

	for _, ch := range s {
		switch {
		case unicode.IsDigit(ch):
			currentNum = currentNum*10 + int(ch-'0')

		case ch == '[':
			numStack = append(numStack, currentNum)
			strStack = append(strStack, currentStr)

			currentNum = 0
			currentStr = ""

		case ch == ']':
			n := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			prev := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			currentStr = prev + strings.Repeat(currentStr, n)

		default:
			currentStr += string(ch)
		}
	}

	return currentStr
}

func main() {
	testCases := []string{
		"3[a]2[bc]",
		"3[a2[c]]",
		"2[abc]3[cd]ef",
		"10[a]",
		"3[ab2[c]]",
	}

	for _, s := range testCases {
		fmt.Printf("Input : %s\n", s)
		fmt.Printf("Output: %s\n\n", decodeString(s))
	}
}
