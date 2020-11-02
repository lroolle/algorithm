"+++
title = "0617. Merge Two Binary Trees 这思路简直不要太清晰 "
author = ["xmblgt"]
date = 2020-09-21T14:21:49+08:00
tags = ["Leetcode", "Algorithms", "Java", "二叉树"]
categories = ["leetcode"]
draft = false
+++

# 这思路简直不要太清晰

> [0617. Merge Two Binary Trees](https://leetcode-cn.com/problems/merge-two-binary-trees/)
> [这思路简直不要太清晰](https://leetcode-cn.com/problems/merge-two-binary-trees/solution/zhe-si-lu-jian-zhi-bu-yao-tai-qing-xi-by-xmblgt/) by [xmblgt](https://leetcode-cn.com/u/xmblgt/)

**思路**
假设我们要把`t1`和`r2`两棵树都合并到`t1`上。
1. 特判：
    - 如果`t1`为空，直接返回`t2`
    - 如果`t2`为空，直接返回`t1`
2. 递归调用，合并`t1.left`和`t2.left`，结果放在`t1.left`上；
3. 递归调用，合并`t1.right`和`t2.right`，结果放在`t1.right`上；
4. 返回`t1`;

```java
class Solution {
    public TreeNode mergeTrees(TreeNode t1, TreeNode t2) {
        if (t1 == null) return t2;
        if (t2 == null) return t1;
        t1.left = mergeTrees(t1.left, t2.left);
        t1.right = mergeTrees(t1.right, t2.right);
        t1.val += t2.val;
        return t1;
    }
}
```

![微信图片_20200921212827.jpg](https://pic.leetcode-cn.com/1600698103-Ahnrri-%E5%BE%AE%E4%BF%A1%E5%9B%BE%E7%89%87_20200921212827.jpg)
