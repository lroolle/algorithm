"+++
title = "0404. Sum of Left Leaves Java递归与迭代实现思路 "
author = ["lastwhispers"]
date = 2020-01-26T05:54:50+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# Java递归与迭代实现思路

> [0404. Sum of Left Leaves](https://leetcode-cn.com/problems/sum-of-left-leaves/)
> [Java递归与迭代实现思路](https://leetcode-cn.com/problems/sum-of-left-leaves/solution/javadi-gui-yu-die-dai-shi-xian-si-lu-by-ggb2312/) by [lastwhispers](https://leetcode-cn.com/u/lastwhispers/)

# 方法一：递归

直接让我们求左叶子之和，可能一下想不到，我们先求节点之和。
```java
// 先序遍历求所有节点值之和
    public int sumOfTrees(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int leave = root.val;
        int left = sumOfTrees(root.left);
        int right = sumOfTrees(root.right);
        return left + right + leave;
    }
```
在求节点之和的基础上，我们再来求所有叶子节点之和。
```java
// 先序遍历求所有叶子节点值之和
    public int sumOfLeaves(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int leave = 0;
        // 叶子节点
        if (root.left == null && root.right == null) {
            leave = root.val;
        }
        int left = sumOfLeaves(root.left);
        int right = sumOfLeaves(root.right);
        return left + right + leave;
    }
```
在求叶子节点之和的基础上，我们再来求所有左叶子节点之和。

当前节点是不是左节点，可以通过父节点标识。
```java
    public int sumOfLeftLeaves(TreeNode root) {
        return sumOfLeftLeavesHelper(root, false);
    }

    // 先序遍历求所有左叶子节点值之和
    public int sumOfLeftLeavesHelper(TreeNode root, boolean flag) {
        if (root == null) {
            return 0;
        }
        int leave = 0;
        // 左叶子节点
        if (flag && root.left == null && root.right == null) {
            leave = root.val;
        }
        int left = sumOfLeftLeavesHelper(root.left, true);
        int right = sumOfLeftLeavesHelper(root.right, false);
        return left + right + leave;
    }

```

评论区的一种写法。
```java

public int sumOfLeftLeaves(TreeNode root) {
    return helper(root);
}

public int helper(TreeNode node) {
    int sum = 0;
    if (node == null) {
        return 0;
    }
    if (node.left != null && (node.left.left == null && node.left.right == null)) {
        sum += node.left.val;
    }
    sum += helper(node.left) + helper(node.right);
    return sum;
}

```

# 方法二：迭代
将先序遍历递归改为迭代
```java
public int sumOfLeftLeaves(TreeNode root) {
        if (root == null) {
            return 0;
        }
        LinkedList<Pair<TreeNode, Boolean>> stack = new LinkedList<>();
        stack.push(new Pair<>(root, false));

        int sum = 0;
        Boolean flag;
        while (!stack.isEmpty()) {
            Pair<TreeNode, Boolean> pair = stack.pop();
            root = pair.getKey();
            flag = pair.getValue();
            if (flag && root.left == null && root.right == null) {
                sum += root.val;
            }
            if (root.left != null) {
                stack.push(new Pair<>(root.left, true));
            }
            if (root.right != null) {
                stack.push(new Pair<>(root.right, false));
            }
        }
        return sum;
    }
```
记得Pair需要导包