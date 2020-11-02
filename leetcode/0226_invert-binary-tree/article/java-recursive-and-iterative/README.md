"+++
title = "0226. Invert Binary Tree [Java] Recursive and Iterative. "
author = ["lincs"]
date = 2020-09-16T02:46:29+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# [Java] Recursive and Iterative.

> [0226. Invert Binary Tree](https://leetcode-cn.com/problems/invert-binary-tree/)
> [[Java] Recursive and Iterative.](https://leetcode-cn.com/problems/invert-binary-tree/solution/java-recursive-solution-by-lincs/) by [lincs](https://leetcode-cn.com/u/lincs/)

- Recursive
```java
class Solution {
    public TreeNode invertTree(TreeNode root) {
        if (root == null) return root;
        TreeNode left = invertTree(root.left);
        TreeNode right = invertTree(root.right);
        root.left = right;
        root.right = left;
        return root;
    }
}
```
- Iterative
```java
class Solution {
    public TreeNode invertTree(TreeNode root) {
        Deque<TreeNode> queue = new LinkedList<>();
        queue.offer(root);
        while (!queue.isEmpty()) {
            for (int i = queue.size(); i > 0; i--) {
                TreeNode p = queue.poll();
                if (p == null) continue;
                TreeNode left = p.left;
                p.left = p.right;
                p.right = left;
                queue.offer(p.right);
                queue.offer(p.left);
            }
        }
        return root;
    }
}
```

