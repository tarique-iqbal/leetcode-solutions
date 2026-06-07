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
			name:     "single character",
			input:    "a",
			expected: "a",
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
			name:     "nested multi character group",
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
			name:     "three level nesting",
			input:    "2[a3[b2[c]]]",
			expected: "abccbccbccabccbccbcc",
		},
		{
			name:     "adjacent encoded groups",
			input:    "2[a]3[b]4[c]",
			expected: "aabbbcccc",
		},
		{
			name:     "prefix and nested group",
			input:    "x3[a2[c]]",
			expected: "xaccaccacc",
		},
		{
			name:     "nested then suffix",
			input:    "3[a2[c]]x",
			expected: "accaccaccx",
		},
		{
			name:     "multiple plain segments",
			input:    "ab2[c]de3[f]gh",
			expected: "abccdefffgh",
		},
		{
			name:     "large repeat with nested content",
			input:    "2[ab3[c]]",
			expected: "abcccabccc",
		},
		{
			name:     "zero repeat mixed",
			input:    "abc0[xyz]def",
			expected: "abcdef",
		},
		{
			name:     "complex nested expression",
			input:    "3[z]2[2[y]pq4[2[jk]e1[f]]]ef",
			expected: "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef",
		},
		{
			name:     "zero repeat",
			input:    "0[a]",
			expected: "",
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			got := decodeString(tc.input)

			if got != tc.expected {
				t.Errorf(
					"decodeString(%q) = %q, want %q",
					tc.input,
					got,
					tc.expected,
				)
			}
		})
	}
}
