# 94. Binary Tree Inorder Traversal

> Topics: Tree, Depth-First Search, Stack, Binary Tree

## Problem Statement

LeetCode: [94. Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/)

Given the root of a binary tree, return the **inorder traversal** of its nodes' values.

Inorder traversal visits nodes in the following order:

```text
Left -> Root -> Right
```

### Example 1

```text
Input:

    1
     \
      2
     /
    3

Output:

[1,3,2]
```

### Example 2

```text
Input:

    1

Output:

[1]
```

### Example 3

```text
Input:

    1
   /
  2

Output:

[2,1]
```

## Visualization

Consider the tree:

```text
        4
      /   \
     2     6
    / \   / \
   1   3 5   7
```

Inorder traversal means:

```text
Left -> Root -> Right
```

Therefore the visit order becomes:

```text
1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7
```

Result:

```text
[1,2,3,4,5,6,7]
```

## Key Observation

For every node:

```text
1. Visit the entire left subtree
2. Visit the node itself
3. Visit the entire right subtree
```

The challenge is remembering where to return after exploring a left subtree.

A stack solves this naturally.

## Recursive Idea

The recursive definition is straightforward:

```go
func dfs(node *TreeNode) {
    if node == nil {
        return
    }

    dfs(node.Left)
    visit(node)
    dfs(node.Right)
}
```

However, recursion uses the call stack internally.

We can simulate the same behavior explicitly using our own stack.

## Intuition

Whenever we move left:

```text
current = current.Left
```

we are postponing the visit of the current node until its entire left subtree has been processed.

Therefore we push the node onto a stack before moving left.

```text
        4
       /
      2
     /
    1
```

Traversal:

```text
push(4)
push(2)
push(1)
```

Once we can no longer go left:

```text
current == nil
```

the node at the top of the stack is the next node that should be visited.

## Stack Strategy

Maintain:

```text
current = root
stack = []
```

At every step:

### Go as far left as possible

```text
while current != nil
```

Push nodes:

```text
stack.push(current)
current = current.Left
```

### Process node

When no left child remains:

```text
current = stack.pop()
```

Visit:

```text
result.append(current.Val)
```

### Explore right subtree

Move to:

```text
current = current.Right
```

and repeat.

## Algorithm

### Step 1: Initialize

```go
var result []int
var stack []*TreeNode

current := root
```

### Step 2: Continue While Work Exists

```go
for current != nil || len(stack) > 0
```

Either:

```text
there is a current node
```

or

```text
there are nodes waiting in the stack
```

### Step 3: Traverse Left

```go
for current != nil {
    stack = append(stack, current)
    current = current.Left
}
```

### Step 4: Pop Node

```go
current = stack[len(stack)-1]
stack = stack[:len(stack)-1]
```

### Step 5: Visit Node

```go
result = append(result, current.Val)
```

### Step 6: Traverse Right

```go
current = current.Right
```

### Step 7: Return Result

```go
return result
```

## Dry Run

### Input

```text
    1
     \
      2
     /
    3
```

Initial state:

```text
current = 1
stack   = []
result  = []
```

### Go Left

Push:

```text
[1]
```

Move:

```text
current = nil
```

### Pop

Visit:

```text
1
```

Result:

```text
[1]
```

Move right:

```text
current = 2
```

---

Go left:

```text
push(2)
push(3)
```

Stack:

```text
[2,3]
```

Current:

```text
nil
```

---

Pop:

```text
3
```

Result:

```text
[1,3]
```

Move right:

```text
nil
```

---

Pop:

```text
2
```

Result:

```text
[1,3,2]
```

Traversal complete.

## Why This Works

Whenever a node is pushed onto the stack:

```text
its visit is postponed
```

until its entire left subtree has been processed.

The stack therefore stores:

```text
nodes waiting to be visited
```

in exactly the same order recursion would.

By:

```text
1. Going left
2. Visiting node
3. Going right
```

the algorithm follows inorder traversal precisely.

## Correctness Argument

For every node:

1. The node is pushed before exploring its left subtree.
2. The node is popped only after all left descendants have been processed.
3. The node is visited exactly once.
4. The right subtree is explored only after the node itself is visited.

Therefore each node is visited in:

```text
Left -> Root -> Right
```

order.

Hence the algorithm always returns the correct inorder traversal.

## Complexity Analysis

Let:

```text
n = number of nodes
```

### Time Complexity

```text
O(n)
```

Each node is:

```text
pushed once
popped once
visited once
```

### Space Complexity

```text
O(h)
```

where:

```text
h = height of the tree
```

because the stack stores at most one path from root to leaf.

Worst case:

```text
O(n)
```

for a completely skewed tree.

Balanced tree:

```text
O(log n)
```

## Example Walkthrough

### Balanced Tree

```text
        4
      /   \
     2     6
    / \   / \
   1   3 5   7
```

Visit order:

```text
1
2
3
4
5
6
7
```

Result:

```text
[1,2,3,4,5,6,7]
```

## Recursive vs Iterative

### Recursive DFS

```text
Uses call stack
Easy to write
```

Complexity:

```text
Time  : O(n)
Space : O(h)
```

### Iterative Stack

```text
Uses explicit stack
Avoids recursion
```

Complexity:

```text
Time  : O(n)
Space : O(h)
```

Both approaches are equivalent.

The iterative version makes the traversal process more explicit.

## Key Takeaways

* Inorder traversal means:

```text
Left -> Root -> Right
```

* A stack can simulate recursive DFS.
* Push nodes while moving left.
* Pop when no more left children exist.
* Visit node after left subtree is processed.
* Then explore the right subtree.
* Every node is visited exactly once.

## Notes

* Classic Tree + Stack problem.
* Fundamental binary tree traversal pattern.
* Excellent example of converting recursion into iteration.

Core lesson:

```text
Whenever recursion needs to remember where to return,
an explicit stack can store that information.
```

For inorder traversal:

```text
Push while moving left,
visit when popping,
then move right.
```
