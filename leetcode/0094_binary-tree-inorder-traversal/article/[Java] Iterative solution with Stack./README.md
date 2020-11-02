"+++
title = "0094. Binary Tree Inorder Traversal [Java] Iterative solution with Stack. "
author = ["lincs"]
date = 2020-08-28T12:38:13+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# [Java] Iterative solution with Stack.

> [0094. Binary Tree Inorder Traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)
> [[Java] Iterative solution with Stack.](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/java-iterative-solution-with-stack-by-lincs/) by [lincs](https://leetcode-cn.com/u/lincs/)

```java
class Solution {
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> list = new LinkedList<>();
        Stack<TreeNode> stack = new Stack<>();
        TreeNode node = root;
        while (!stack.isEmpty() || node != null) {
            while (node != null) {
                stack.push(node);
                node = node.left;
            }
            node = stack.pop();
            list.add(node.val);
            node = node.right;
        }
        return list;
    }
}
```
