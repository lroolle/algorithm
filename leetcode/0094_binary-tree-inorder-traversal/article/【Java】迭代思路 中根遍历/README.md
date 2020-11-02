"+++
title = "0094. Binary Tree Inorder Traversal 【Java】迭代思路 中根遍历 "
author = ["leetcoder-youzg"]
date = 2020-09-14T01:36:29+08:00
tags = ["Leetcode", "Algorithms", "Java", "迭代", "DepthfirstSearch", "搜索"]
categories = ["leetcode"]
draft = false
+++

# 【Java】迭代思路 中根遍历

> [0094. Binary Tree Inorder Traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)
> [【Java】迭代思路 中根遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/java-die-dai-si-lu-zhong-gen-bian-li-by-leetcoder-/) by [leetcoder-youzg](https://leetcode-cn.com/u/leetcoder-youzg/)

### 解题思路
递归的解决方案，想必大家已经很熟悉了
那么，在本题解中，本人来展示下 **`迭代`** 的解决方案

首先，本人来介绍下 什么是 **`迭代`**，它与 **`递归`** 的**区别**是什么：
> **`迭代`** 就是利用 **变量的原值** 推算出变量的一个新值
> （我的理解就是 —— 利用变量的**原值**推出**新值**）
> 若 **`递归`** 是自己 **调用自己** 的话，**`迭代`** 就是自己不停地 **调用别人**

代码如下：

### 运行结果
![image.png](https://pic.leetcode-cn.com/1600047117-TUwaPa-image.png)
### 代码

```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> result = new ArrayList<>();
        if (root == null) {
            return result;
        }

        Deque<TreeNode> queue = new ArrayDeque<>();
        TreeNode p = root;
        TreeNode curNode = null;
        while (!queue.isEmpty() || p != null) {
            while (p != null) { // 将 当前p 的 左孩子路径 全部录入queue
                queue.push(p);  // 将 当前p 加入 队列
                p = p.left; // 让 当前p 指向 左孩子
            }
            curNode = queue.pop();
            result.add(curNode.val);    // 录入顺序为 左->根 (因为在查找 左孩子们之前，将根节点提前录入过了)
            p = curNode.right;  // 让 p 指向 右孩子(达成 左->根->右 的顺序)
        }
        
        return result;
    }
}
```
打卡第55天，加油🦆