# 34. Find First and Last Position of Element in Sorted Array

> Topics: Array, Binary Search

## Problem Statement

LeetCode: [34. Find First and Last Position of Element in Sorted Array](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/)

## Problem Overview

You are given a sorted array of integers `nums` and a target value `target`.

Return the starting and ending position of the target value.

If the target is not found, return:

```text
[-1, -1]
```

You must write an algorithm with:

```text
O(log n)
```

runtime complexity.

## Intuition

Since the array is sorted, Binary Search is the optimal approach.

The key observation:

- The first occurrence of the target can be found using a modified binary search
- The last occurrence can also be found using another modified binary search

Instead of scanning linearly, we use:

- Lower Bound → first index where value ≥ target
- Upper Bound → first index where value > target

Using these two boundaries, we can determine the exact range.

## Approach

### 1. Find Lower Bound

Lower bound returns:

```text
first index where nums[i] >= target
```

Example:

```text
nums = [5,7,7,8,8,10]
target = 8
```

Lower bound returns:

```text
3
```

because `nums[3] = 8`.

### 2. Verify Target Exists

If:

```text
lowerBound == len(nums)
```

or:

```text
nums[lowerBound] != target
```

then target does not exist.

Return:

```text
[-1, -1]
```

### 3. Find Upper Bound

Upper bound returns:

```text
first index where nums[i] > target
```

For:

```text
nums = [5,7,7,8,8,10]
target = 8
```

Upper bound returns:

```text
5
```

The actual last occurrence is:

```text
upperBound - 1
```

which becomes:

```text
4
```

### 4. Return Final Range

```text
[firstPosition, lastPosition]
```

## Dry Run

### Input

```text
nums = [5,7,7,8,8,10]
target = 8
```

### Lower Bound

```text
index = 3
```

### Upper Bound

```text
index = 5
```

### Final Answer

```text
[3, 4]
```

## Why Binary Search Works

Because the array is sorted:

- Left half contains smaller values
- Right half contains larger values

Binary search repeatedly halves the search space, giving:

```text
O(log n)
```

time complexity.

## Important Insight

Instead of searching for exact equality directly:

- Lower bound searches for first valid position
- Upper bound searches for first greater position

This technique is extremely useful in range-search problems.

## Complexity Analysis

### Time Complexity

Two binary searches:

```text
O(log n)
```

### Space Complexity

```text
O(1)
```

## Key Takeaways

- Sorted array → think Binary Search
- Use lower bound + upper bound
- Avoid linear scans
- Upper bound result must be decremented by 1
- Classic range-search binary search problem

## Notes

- Lower Bound:

```text
first index >= target
```

- Upper Bound:

```text
first index > target
```
