package main

import (
	"fmt"
	"unicode"
)

func decodeString(s string) string {
	decoded, _ := dfs(s, 0)
	return decoded
}

func dfs(s string, i int) (string, int) {
	currentStr := ""
	num := 0

	for i < len(s) {
		ch := rune(s[i])

		switch {
		case unicode.IsDigit(ch):
			num = num*10 + int(ch-'0')
			i++

		case ch == '[':
			subStr, nextIndex := dfs(s, i+1)

			for j := 0; j < num; j++ {
				currentStr += subStr
			}

			num = 0
			i = nextIndex

		case ch == ']':
			return currentStr, i + 1

		default:
			currentStr += string(ch)
			i++
		}
	}

	return currentStr, i
}

func main() {
	testCases := []string{
		"3[a]",
		"3[a]2[bc]",
		"3[a2[c]]",
		"2[abc]3[cd]ef",
		"10[a]",
		"2[a2[b2[c]]]",
	}

	for _, tc := range testCases {
		fmt.Printf("Input : %s\n", tc)
		fmt.Printf("Output: %s\n\n", decodeString(tc))
	}
}
