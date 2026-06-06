# 977. Squares of a Sorted Array

> Topics: Array, Two Pointers

## Problem Statement

LeetCode: [977. Squares of a Sorted Array](https://leetcode.com/problems/squares-of-a-sorted-array/)

Given an integer array `nums` sorted in **non-decreasing order**, return an array containing the square of each number, also sorted in **non-decreasing order**.

### Example 1

```text
Input:
[-4,-1,0,3,10]

Output:
[0,1,9,16,100]
```

### Example 2

```text
Input:
[-7,-3,2,3,11]

Output:
[4,9,9,49,121]
```

## Visualization

Consider:

```text
[-4,-1,0,3,10]
```

After squaring:

```text
[16,1,0,9,100]
```

Notice that the array is no longer sorted.

The largest squared values come from numbers with the largest absolute values:

```text
|-4| = 4
|10| = 10
```

```text
-4  -1   0   3   10
 ^                 ^
left            right
```

Squares:

```text
16               100
```

Since:

```text
100 > 16
```

the largest remaining value belongs at the end of the result array.

## Key Observation

Although the original array is sorted:

```text
-7 < -3 < 2 < 3 < 11
```

their squares are not:

```text
49, 9, 4, 9, 121
```

The problem comes from negative values.

However, the largest square must always come from one of the ends:

```text
left  -> most negative number
right -> largest positive number
```

because those positions contain the largest absolute values.

Therefore:

```text
max square =
max(nums[left]², nums[right]²)
```

## Brute Force Approach

Square every number:

```go
for i := range nums {
    nums[i] *= nums[i]
}
```

Then sort the array:

```go
sort.Ints(nums)
```

### Complexity

```text
Time  : O(n log n)
Space : O(1) or O(n)
```

depending on the sorting implementation.

We can do better.

## Intuition

The input array is already sorted.

The largest absolute value must be at one of the two ends:

```text
left  = 0
right = n - 1
```

At each step:

1. Compare both squares.
2. Put the larger square into the last unused position.
3. Move the pointer that produced that square.
4. Fill the result array from right to left.

This avoids sorting entirely.

## Two-Pointer Strategy

Maintain:

```text
left  = 0
right = n - 1
```

And a position pointer:

```text
pos = n - 1
```

At every step:

```text
leftSq  = nums[left]²
rightSq = nums[right]²
```

If:

```text
leftSq > rightSq
```

place:

```text
leftSq
```

at:

```text
result[pos]
```

and move:

```text
left++
```

Otherwise place:

```text
rightSq
```

and move:

```text
right--
```

Then:

```text
pos--
```

Continue until:

```text
left > right
```

## Algorithm

### Step 1: Create Result Array

```go
n := len(nums)
result := make([]int, n)
```

### Step 2: Initialize Pointers

```go
left, right := 0, n-1
```

### Step 3: Fill Result From Right to Left

```go
pos := n - 1
```

### Step 4: Compare Squares

```go
leftSq := nums[left] * nums[left]
rightSq := nums[right] * nums[right]
```

### Step 5: Place Larger Square

```go
if leftSq > rightSq {
    result[pos] = leftSq
    left++
} else {
    result[pos] = rightSq
    right--
}
```

### Step 6: Move Result Position

```go
pos--
```

### Step 7: Return Result

```go
return result
```

## Dry Run

### Input

```text
[-4,-1,0,3,10]
```

Initial state:

```text
left  = 0
right = 4
pos   = 4
```

Squares:

```text
16 vs 100
```

Choose:

```text
100
```

Result:

```text
[_,_,_,_,100]
```

Move:

```text
right--
```

---

Now:

```text
left  = 0
right = 3
pos   = 3
```

Squares:

```text
16 vs 9
```

Choose:

```text
16
```

Result:

```text
[_,_,_,16,100]
```

Move:

```text
left++
```

---

Now:

```text
left  = 1
right = 3
pos   = 2
```

Squares:

```text
1 vs 9
```

Choose:

```text
9
```

Result:

```text
[_,_,9,16,100]
```

Move:

```text
right--
```

---

Now:

```text
left  = 1
right = 2
pos   = 1
```

Squares:

```text
1 vs 0
```

Choose:

```text
1
```

Result:

```text
[_,1,9,16,100]
```

Move:

```text
left++
```

---

Now:

```text
left = 2
right = 2
```

Squares:

```text
0
```

Result:

```text
[0,1,9,16,100]
```

Final answer:

```text
[0,1,9,16,100]
```

## Why This Works

The largest remaining square must always come from one of the ends.

Reason:

```text
Array is already sorted.
```

Therefore the largest magnitude value must be either:

```text
nums[left]
```

or

```text
nums[right]
```

No middle element can have a larger absolute value than both ends.

So at every step we can safely place the larger square into the largest remaining position in the result array.

Repeating this process builds the sorted answer without sorting.

## Correctness Argument

For every iteration:

1. The largest remaining square is located at either end.
2. The algorithm compares both candidates.
3. The larger square is placed into the largest unfilled position.
4. That value will never need to move again.
5. The corresponding pointer is advanced inward.

By repeatedly placing the largest remaining square, the result array is built in sorted order.

Therefore the algorithm always produces the correct answer.

## Complexity Analysis

### Time Complexity

```text
O(n)
```

Each pointer moves at most:

```text
n
```

times.

### Space Complexity

```text
O(n)
```

A result array of size:

```text
n
```

is created.

## Example Walkthrough

### Input

```text
[-5,-4,-3,-2,-1]
```

Squares:

```text
25,16,9,4,1
```

Result:

```text
[1,4,9,16,25]
```

The algorithm correctly builds the sorted order without calling a sorting function.

## Two Pointers vs Sorting

### Square + Sort

```text
Square every value
Sort the array
```

Complexity:

```text
Time  : O(n log n)
Space : Depends on sorting implementation
```

### Two Pointers

```text
Compare largest absolute values
Fill result from right to left
```

Complexity:

```text
Time  : O(n)
Space : O(n)
```

The two-pointer solution is optimal.

## Key Takeaways

* Squaring negative numbers breaks sorted order.
* The largest square always comes from one of the array ends.
* Compare absolute values using two pointers.
* Fill the result array from right to left.
* No sorting is required.
* Reduces complexity from:

```text
O(n log n)
```

to:

```text
O(n)
```

## Notes

* Classic Two Pointers problem.
* Excellent example of using the sorted property of the input.
* Core lesson:

```text
When the largest remaining value is always at an edge,
process the array from both ends inward.
```

For this problem:

```text
Compare squares at both ends and place the larger one first.
```
