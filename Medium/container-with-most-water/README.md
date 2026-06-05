# 11. Container With Most Water

> Topics: Array, Two Pointers, Greedy

## Problem Statement

LeetCode: [11. Container With Most Water](https://leetcode.com/problems/container-with-most-water/)

You are given an integer array `height` where:

```text
height[i]
```

represents the height of a vertical line drawn at position `i`.

Find two lines that together with the x-axis form a container that can hold the maximum amount of water.

Return the maximum amount of water the container can store.


## Visualization

Given:

```text
Index:   0 1 2 3 4 5 6 7 8
Height: [1,8,6,2,5,4,8,3,7]
```

The optimal container is formed by:

```text
Height 8 at index 1
Height 7 at index 8
```

```text
        |
        |               |
        |               |
        |               |
        |               |
        |               |
        |               |
        |               |
--------|---------------|--------
        <----- 7 ------->
```

Area:

```text
width     = 8 - 1 = 7
height    = min(8, 7) = 7

area = 7 × 7 = 49
```

Answer:

```text
49
```


## Key Observation

For any two lines:

```text
left, right
```

the container area is:

```text
width × height
```

where:

```text
width  = right - left
height = min(height[left], height[right])
```

Therefore:

```text
area = (right - left) *
       min(height[left], height[right])
```


## Brute Force Approach

Try every possible pair:

```go
for i := 0; i < n; i++ {
    for j := i + 1; j < n; j++ {
        area := (j - i) *
            min(height[i], height[j])
    }
}
```

Time Complexity:

```text
O(n²)
```

This becomes too slow for large inputs.


## Intuition

We need something better than checking all pairs.

Start with the widest possible container:

```text
left  = 0
right = n - 1
```

Since width is already maximum:

```text
width = right - left
```

the only way to potentially improve the area is by finding a taller boundary.


## Important Insight

Suppose:

```text
height[left] < height[right]
```

Example:

```text
left = 3
right = 8
```

Container height is:

```text
min(3, 8) = 3
```

Area:

```text
width × 3
```

Now consider moving the taller line:

```text
right--
```

The width becomes smaller:

```text
width - 1
```

but the limiting height remains at most:

```text
3
```

Therefore:

```text
area cannot increase
```

because:

```text
smaller width
same limiting height
```

So moving the taller line is never useful.

Instead, we move the shorter line hoping to find a taller one.


## Two-Pointer Strategy

Maintain:

```text
left  = 0
right = n - 1
```

At every step:

1. Compute area.
2. Update maximum area.
3. Move the shorter boundary inward.

Rule:

```text
if height[left] < height[right]
    left++
else
    right--
```

This guarantees we only explore candidates that might improve the result.


## Algorithm

### Step 1: Initialize Pointers

```go
left, right := 0, len(height)-1
```


### Step 2: Track Maximum Area

```go
maxArea := 0
```


### Step 3: Process While Pointers Have Not Crossed

```go
for left < right
```


### Step 4: Compute Current Container

Width:

```go
width := right - left
```

Height:

```go
minHeight := min(
    height[left],
    height[right],
)
```

Area:

```go
area := width * minHeight
```


### Step 5: Update Best Answer

```go
if area > maxArea {
    maxArea = area
}
```


### Step 6: Move Shorter Boundary

```go
if height[left] < height[right] {
    left++
} else {
    right--
}
```


### Step 7: Return Result

```go
return maxArea
```


## Dry Run

### Input

```text
[1,8,6,2,5,4,8,3,7]
```


### Initial State

```text
left  = 0
right = 8
```

```text
width  = 8
height = min(1,7) = 1

area = 8
```

```text
maxArea = 8
```

Move:

```text
left++
```

because:

```text
1 < 7
```


### State

```text
left  = 1
right = 8
```

```text
width  = 7
height = min(8,7) = 7

area = 49
```

```text
maxArea = 49
```

Move:

```text
right--
```

because:

```text
7 < 8
```


### Continue

The algorithm keeps shrinking the window while evaluating all useful candidates.

No later pair exceeds:

```text
49
```


### Final Answer

```text
49
```


## Why Moving the Shorter Line Works

Suppose:

```text
height[left] = 4
height[right] = 10
```

Current area:

```text
(right-left) × 4
```

If we move the taller line:

```text
right--
```

then:

```text
width decreases
```

and the limiting height can never exceed:

```text
4
```

Therefore:

```text
area cannot improve
```

The only chance to get a larger area is finding a taller line than:

```text
4
```

which means moving:

```text
left++
```

This observation eliminates unnecessary comparisons and reduces complexity from:

```text
O(n²)
```

to:

```text
O(n)
```


## Correctness Argument

For every step:

1. The current area is evaluated.
2. The shorter line determines the container height.
3. Moving the taller line cannot increase area because width decreases while the limiting height remains unchanged.
4. Therefore only moving the shorter line can possibly lead to a better solution.
5. Every useful candidate pair is considered exactly once.

Thus the algorithm always finds the maximum possible container area.


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
O(1)
```

Only a few variables are used.


## Example Walkthrough

### Input

```text
height = [4,3,2,1,4]
```

Initial:

```text
left  = 0
right = 4
```

```text
width  = 4
height = 4

area = 16
```

This is already optimal.

Output:

```text
16
```


## Two Pointers vs Brute Force

### Brute Force

```text
Check every pair
```

Complexity:

```text
Time  : O(n²)
Space : O(1)
```


### Two Pointers

```text
Start widest
Move shorter boundary
```

Complexity:

```text
Time  : O(n)
Space : O(1)
```


## Key Takeaways

* Area is determined by:

```text
width × min(leftHeight, rightHeight)
```

* Start with the maximum width.
* The shorter line limits the container.
* Moving the taller line can never improve the result.
* Move only the shorter line.
* This reduces the problem from:

```text
O(n²)
```

to:

```text
O(n)
```


## Notes

* Classic Two Pointers problem.
* One of the most important pointer-shrinking patterns.
* Core lesson:

```text
When width shrinks every step,
move the pointer that limits the answer.
```

For this problem:

```text
Move the shorter height.
```
