# 881. Boats to Save People

> Topics: Greedy, Two Pointers, Sorting

## Problem Statement

LeetCode: [881. Boats to Save People](https://leetcode.com/problems/boats-to-save-people/)

## Problem Overview

You are given:

* An array `people`

  * `people[i]` represents the weight of the `i-th` person
* An integer `limit`

  * Maximum weight a boat can carry

Rules:

* Each boat can carry at most **2 people**
* The total weight on a boat cannot exceed `limit`

Return the **minimum number of boats** required to carry everyone.

## Intuition

We want to minimize the number of boats.

A greedy strategy works best here:

* Always try to pair the **heaviest person**
  with the **lightest possible person**
* If they can share a boat:

```text
lightest + heaviest <= limit
```

then we save one boat.

Otherwise:

* the heaviest person must go alone.

This naturally leads to:

* Sorting
* Two pointers

## Approach

### 1. Sort the Array

Sorting helps us efficiently pair:

* the lightest person
* the heaviest person

Example:

```text
people = [3,2,2,1]

After sorting:
[1,2,2,3]
```

### 2. Use Two Pointers

Maintain:

```text
left  → lightest person
right → heaviest person
```

Initially:

```text
left = 0
right = n - 1
```

### 3. Try to Pair People

If:

```text
people[left] + people[right] <= limit
```

then:

* both can share a boat
* move both pointers

```text
left++
right--
```

Otherwise:

* the heaviest person goes alone

```text
right--
```

In both cases:

* one boat is used

```text
boats++
```

## Greedy Insight

Why pair the heaviest with the lightest?

Because:

* the heaviest person is the hardest to fit
* if the lightest cannot fit with them,
  nobody else can

So sending the heaviest alone is always optimal.

## Dry Run

### Input

```text
people = [3,2,2,1]
limit = 3
```

### After Sorting

```text
[1,2,2,3]
```

### Step 1

```text
1 + 3 = 4 > 3
```

3 goes alone.

```text
boats = 1
```

### Step 2

```text
1 + 2 = 3 <= 3
```

Pair them.

```text
boats = 2
```

### Step 3

Remaining:

```text
2
```

One final boat.

```text
boats = 3
```

## Final Answer

```text
3
```

## Why This Greedy Approach Works

At every step:

* we optimally place the heaviest remaining person
* pairing with the lightest maximizes boat utilization

This guarantees:

* minimum boats used

## Complexity Analysis

### Time Complexity

```text
O(n log n)
```

Due to sorting.

Two-pointer traversal is:

```text
O(n)
```

### Space Complexity

```text
O(1)
```

Ignoring sorting space.

## Key Takeaways

* This is a Greedy + Two Pointer problem
* Sort the array first
* Pair:

```text
lightest + heaviest
```

* If pairing fails:

  * heaviest must go alone
* Greedy works because the heaviest person is the limiting factor

## Notes

* Classic greedy pairing problem
* Common pattern in array and sequence problems
* Very common two-pointer problem-solving pattern
* Standard approach for pairing or scanning problems
