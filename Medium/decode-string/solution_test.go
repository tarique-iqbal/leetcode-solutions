package main

import "testing"

func TestDecodeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "plain text only",
			input:    "abc",
			expected: "abc",
		},
		{
			name:     "single repeat",
			input:    "3[a]",
			expected: "aaa",
		},
		{
			name:     "multiple groups",
			input:    "3[a]2[bc]",
			expected: "aaabcbc",
		},
		{
			name:     "nested group",
			input:    "3[a2[c]]",
			expected: "accaccacc",
		},
		{
			name:     "mixed encoded and plain text",
			input:    "2[abc]3[cd]ef",
			expected: "abcabccdcdcdef",
		},
		{
			name:     "plain text around encoded group",
			input:    "abc3[cd]xyz",
			expected: "abccdcdcdxyz",
		},
		{
			name:     "nested multi-character group",
			input:    "3[a2[bd]]",
			expected: "abdbdabdbdabdbd",
		},
		{
			name:     "nested with explicit single repeat",
			input:    "abc3[de1[fg]]",
			expected: "abcdefgdefgdefg",
		},
		{
			name:     "multi digit repeat",
			input:    "10[a]",
			expected: "aaaaaaaaaa",
		},
		{
			name:     "multi digit nested repeat",
			input:    "2[10[a]]",
			expected: "aaaaaaaaaaaaaaaaaaaa",
		},
		{
			name:     "deep nesting",
			input:    "2[a2[b2[c]]]",
			expected: "abccbccabccbcc",
		},
		{
			name:     "adjacent encoded groups",
			input:    "2[a]3[b]4[c]",
			expected: "aabbbcccc",
		},
		{
			name:     "nested then suffix",
			input:    "3[a2[c]]x",
			expected: "accaccaccx",
		},
		{
			name:     "prefix and nested group",
			input:    "x3[a2[c]]",
			expected: "xaccaccacc",
		},
		{
			name:     "zero repeat",
			input:    "0[a]",
			expected: "",
		},
		{
			name:     "zero repeat mixed",
			input:    "abc0[xyz]def",
			expected: "abcdef",
		},
		{
			name:     "large repeat with nested content",
			input:    "2[ab3[c]]",
			expected: "abcccabccc",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := decodeString(tc.input)

			if got != tc.expected {
				t.Fatalf(
					"decodeString(%q) = %q, want %q",
					tc.input,
					got,
					tc.expected,
				)
			}
		})
	}
}
