# 102. Binary Tree Level Order Traversal

> Topics: Tree, Binary Tree, Breadth-First Search (BFS), Queue

## Problem Statement

Given the `root` of a binary tree, return the **level order traversal** of its nodes' values.

Level order traversal means visiting nodes level by level from **left to right**.

### Example

For the binary tree:

```text
        3
       / \
      9  20
         / \
        15  7
```

The traversal should be:

```text
[[3], [9, 20], [15, 7]]
```

Each inner array represents one level of the tree.

## Problem Overview

- **Input:** `root` of a binary tree  
- **Output:** List of lists containing node values level-by-level  
- **Goal:** Traverse the tree one level at a time

## Intuition

We need to process nodes **level by level**. This is exactly what **Breadth-First Search (BFS)** is designed for.

Using a queue:

1. Start from the root
2. Process all nodes at the current level
3. Add their children into the queue
4. Move to the next level

## Approach

### 1. Handle Empty Tree

If the root is `nil`, return `nil`.

### 2. Initialize Queue

Use a queue to process nodes in BFS order.  
Add the root node first.

### 3. Process Each Level

At every iteration:

- Determine the current level size
- Process exactly those nodes
- Collect their values
- Add their children into the queue

Using the current queue size ensures levels remain separated.

### 4. Traverse Nodes in Current Level

For each node:

- Remove from front of the queue
- Record its value
- Add left child if it exists
- Add right child if it exists

### 5. Store Level Result

After processing a level, append the collected values into the final result.

## Dry Run

### Input

```text
root = [3,9,20,null,null,15,7]
```

### Step-by-Step

#### Initial Queue

```text
[3]
```

#### Level 1

Process:

```text
3
```

Add children:

```text
9, 20
```

Result:

```text
[[3]]
```

Queue:

```text
[9, 20]
```

#### Level 2

Process:

```text
9, 20
```

Add children:

```text
15, 7
```

Result:

```text
[[3], [9,20]]
```

Queue:

```text
[15, 7]
```

#### Level 3

Process:

```text
15, 7
```

Result:

```text
[[3], [9,20], [15,7]]
```

Queue becomes empty.

## Final Answer

```text
[[3], [9,20], [15,7]]
```

## BFS Visualization

```text
Level 0:        3
               / \
Level 1:      9  20
                 / \
Level 2:       15  7
```

Traversal order:

```text
3 → 9 → 20 → 15 → 7
```

Grouped by levels:

```text
[[3], [9,20], [15,7]]
```

## Why BFS Works

BFS processes nodes in layers. Using a queue guarantees:

- Nodes are processed in insertion order
- One level is completely processed before moving deeper

This naturally produces level order traversal.

## Complexity Analysis

### Time Complexity

```text
O(n)
```

Each node is visited exactly once.

### Space Complexity

```text
O(n)
```

The queue may store an entire level of the tree.

## Edge Cases

### Empty Tree

```text
root = nil
```

Output:

```text
nil
```

### Single Node

```text
root = [1]
```

Output:

```text
[[1]]
```

### Unbalanced Tree

```text
    1
   /
  2
 /
3
```

Output:

```text
[[1], [2], [3]]
```

## Key Takeaways

- Classic BFS problem
- Queue is the core data structure
- Process nodes level-by-level
- Use current queue size to separate levels
- Each node is visited exactly once

## Notes

- BFS is generally preferred for level-order traversal problems
