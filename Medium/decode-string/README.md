# 394. Decode String

> Topics: String, Stack, Recursion

## Problem Statement

LeetCode: [394. Decode String](https://leetcode.com/problems/decode-string/)

Given an encoded string, return its decoded version.

The encoding rule is:

```text
k[encoded_string]
```

where:

```text
encoded_string inside the brackets is repeated exactly k times
```

You may assume:

```text
- k is always a positive integer
- Input is always valid
- No extra spaces exist
```

### Example 1

```text
Input:

"3[a]2[bc]"

Output:

"aaabcbc"
```

### Example 2

```text
Input:

"3[a2[c]]"

Output:

"accaccacc"
```

### Example 3

```text
Input:

"2[abc]3[cd]ef"

Output:

"abcabccdcdcdef"
```


## Visualization

Consider:

```text
3[a2[c]]
```

The innermost expression is:

```text
2[c]
```

which becomes:

```text
cc
```

Now:

```text
3[acc]
```

becomes:

```text
accaccacc
```

Result:

```text
"accaccacc"
```


## Key Observation

Whenever we encounter:

```text
[
```

we are entering a new nested decoding context.

We need to remember:

```text
1. The repetition count before '['
2. The string built before '['
```

because after decoding the nested portion we must return to the previous state.

A stack naturally stores this information.


## Intuition

Consider:

```text
3[a2[c]]
```

Processing:

```text
3
```

means:

```text
currentNum = 3
```

When we see:

```text
[
```

we save:

```text
count = 3
current string = ""
```

and start decoding the nested part.

Later:

```text
2[c]
```

creates another nested level.

Each new bracket introduces a new decoding context.

Stacks allow us to return to the correct previous state when a closing bracket appears.


## Stack Strategy

Maintain:

```text
numStack
strStack
```

and two working variables:

```text
currentNum
currentStr
```

### Digit

Build the repetition count.

```text
currentNum = currentNum * 10 + digit
```

This supports multi-digit values such as:

```text
12[a]
100[ab]
```


### Opening Bracket '['

Store the current state.

```text
push(currentNum)
push(currentStr)
```

Then start a fresh context.

```text
currentNum = 0
currentStr = ""
```


### Closing Bracket ']'

The current string has been completely decoded.

Retrieve:

```text
repeat count
previous string
```

Then combine:

```text
previous + repeat(current, count)
```


### Character

Append directly.

```text
currentStr += character
```


## Algorithm

### Step 1: Initialize Stacks

```go
numStack := []int{}
strStack := []string{}

currentNum := 0
currentStr := ""
```


### Step 2: Process Each Character

```go
for _, ch := range s
```


### Step 3: Handle Digits

```go
currentNum = currentNum*10 + int(ch-'0')
```

Build multi-digit numbers.


### Step 4: Handle '['

Save the current decoding state.

```go
numStack = append(numStack, currentNum)
strStack = append(strStack, currentStr)
```

Reset:

```go
currentNum = 0
currentStr = ""
```


### Step 5: Handle ']'

Retrieve previous state.

```go
n := pop(numStack)
prev := pop(strStack)
```

Expand:

```go
currentStr = prev + strings.Repeat(currentStr, n)
```


### Step 6: Handle Letters

```go
currentStr += string(ch)
```


### Step 7: Return Result

```go
return currentStr
```


## Dry Run

### Input

```text
3[a2[c]]
```

Initial state:

```text
currentNum = 0
currentStr = ""

numStack = []
strStack = []
```

---

Read:

```text
3
```

```text
currentNum = 3
```

---

Read:

```text
[
```

Push state:

```text
numStack = [3]
strStack = [""]
```

Reset:

```text
currentNum = 0
currentStr = ""
```

---

Read:

```text
a
```

```text
currentStr = "a"
```

---

Read:

```text
2
```

```text
currentNum = 2
```

---

Read:

```text
[
```

Push:

```text
numStack = [3,2]
strStack = ["","a"]
```

Reset.

---

Read:

```text
c
```

```text
currentStr = "c"
```

---

Read:

```text
]
```

Pop:

```text
n = 2
prev = "a"
```

Expand:

```text
currentStr = "a" + "cc"
```

Result:

```text
acc
```

---

Read:

```text
]
```

Pop:

```text
n = 3
prev = ""
```

Expand:

```text
currentStr = "" + "accaccacc"
```

Result:

```text
accaccacc
```

Traversal complete.


## Why This Works

Every opening bracket:

```text
[
```

creates a new decoding context.

The stacks remember:

```text
- how many times the nested string must repeat
- what string existed before the nested section
```

When:

```text
]
```

is reached, the nested substring is already fully decoded.

We simply:

```text
repeat it
append it to the previous string
continue
```

This exactly mirrors how nested expressions are evaluated.


## Correctness Argument

For every bracket pair:

1. The repetition count is stored before entering the nested section.
2. The current partial string is stored before entering the nested section.
3. Characters inside the brackets are fully decoded first.
4. Upon reaching `]`, the decoded substring is repeated the required number of times.
5. The result is appended to the correct previous string context.

Thus every encoded segment:

```text
k[encoded_string]
```

is transformed into:

```text
encoded_string repeated k times
```

and nested structures are decoded correctly.

Therefore the algorithm always returns the correct decoded string.


## Complexity Analysis

Let:

```text
n = length of input string
```

### Time Complexity

```text
O(n + output_size)
```

Each character is processed once.

String expansion contributes work proportional to the size of the final decoded output.


### Space Complexity

```text
O(n)
```

Stacks may store nested contexts for every opening bracket.

In the worst case:

```text
"1[1[1[1[a]]]]"
```

stack depth can be proportional to the input length.


## Example Walkthrough

### Input

```text
2[abc]3[cd]ef
```

Decode:

```text
2[abc] -> abcabc
```

Decode:

```text
3[cd] -> cdcdcd
```

Combine:

```text
abcabc + cdcdcd + ef
```

Result:

```text
abcabccdcdcdef
```


## Recursive vs Iterative

### Recursive Solution

```text
Uses call stack
Natural for nested structures
```

Complexity:

```text
Time  : O(n + output_size)
Space : O(depth)
```


### Iterative Stack Solution

```text
Uses explicit stacks
Avoids recursion
Handles deep nesting safely
```

Complexity:

```text
Time  : O(n + output_size)
Space : O(n)
```

Both approaches follow the same idea:

```text
Save context when entering brackets,
restore context when leaving brackets.
```


## Key Takeaways

* Nested encodings create nested decoding contexts.
* Two stacks store:

  * repetition counts
  * previous strings
* Push state when encountering `'['`.
* Pop state when encountering `']'`.
* Multi-digit numbers are handled naturally.
* Nested expressions are decoded from the inside out.
* The algorithm processes the string in a single pass.


## Notes

* Classic Stack problem.
* Excellent example of context restoration.
* Demonstrates how stacks can replace recursion.

Core lesson:

```text
Whenever nested structures require remembering previous state,
a stack is often the simplest solution.
```

For Decode String:

```text
Save state on '[',
restore and expand on ']'.
```
