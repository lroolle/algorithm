"+++
title = "0617. Merge Two Binary Trees 递归解法 java "
author = ["wallee"]
date = 2019-06-12T07:26:05+08:00
tags = ["Leetcode", "Algorithms", ]
categories = ["leetcode"]
draft = false
+++

# 递归解法 java

> [0617. Merge Two Binary Trees](https://leetcode-cn.com/problems/merge-two-binary-trees/)
> [递归解法 java](https://leetcode-cn.com/problems/merge-two-binary-trees/solution/di-gui-jie-fa-java-by-wallee/) by [wallee](https://leetcode-cn.com/u/wallee/)

```java
class Solution {
    public TreeNode mergeTrees(TreeNode t1, TreeNode t2) {
        if(t1 == null){
            return t2;
        }
        if(t2 == null){
            return t1;
        }
        TreeNode result = new TreeNode(t1.val + t2.val);
        result.left = mergeTrees(t1.left,t2.left);
        result.right = mergeTrees(t1.right,t2.right);
        return result;
    }
}
```