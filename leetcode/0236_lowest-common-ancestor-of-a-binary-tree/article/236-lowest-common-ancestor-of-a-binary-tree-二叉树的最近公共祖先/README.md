"+++
title = "0236. Lowest Common Ancestor of a Binary Tree 236. Lowest Common Ancestor of a Binary Tree 二叉树的最近公共祖先 "
author = ["pengxurui"]
date = 2020-09-11T13:00:26+08:00
tags = ["Leetcode", "Algorithms", "Kotlin", "DFS"]
categories = ["leetcode"]
draft = false
+++

# 236. Lowest Common Ancestor of a Binary Tree 二叉树的最近公共祖先

> [0236. Lowest Common Ancestor of a Binary Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)
> [236. Lowest Common Ancestor of a Binary Tree 二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/solution/236-lowest-common-ancestor-of-a-binary-tree-er-cha/) by [pengxurui](https://leetcode-cn.com/u/pengxurui/)



#### 解法1：DFS遍历（递归方式） + 剪枝

```
class Solution {
    fun lowestCommonAncestor(root: TreeNode?, p: TreeNode?, q: TreeNode?): TreeNode? {
        
        var isFind = false
        fun search(node : TreeNode?) : TreeNode?{
            // 分析点 1
            if(null == node){
                return null
            }
            // 分析点 2
            if(node == p || node == q){
                // 如果当前点就是其中一个目标点，子树被剪枝
                return node
            }
            // 分析点 3
            val leftResult = search(node.left)

            if(isFind){
                // 左边已经找到了，右边被剪枝
                return leftResult
            }

            // 分析点 4
            val rightResult = search(node.right)

            if(null != leftResult && null != rightResult){
                // 左右都找到了，说明现在是最近父祖先
                isFind = true
                return node
            }
            // 分析点 5
            return if(null != leftResult) leftResult else rightResult
        }

        return search(root)
    }
}
```

###### 关键信息

1. 所有节点的值都是唯一的（重点）
2. p、q 为不同节点且均存在于给定的二叉树中（不是很重要，从编程技巧上容易兼容）

###### 测试用例：

1. root 为 null（此时，结果应为 null 或 root）
2. root 不为 null，p 或 q 至少一个为 null
3. 以上都不为 null，p 或 q 无祖先关系
4. 以上都不为 null，p 或 q 有祖先关系（例如 p 为 q 的子节点）

###### 代码解释：

我们实现了`search()`函数，实现了以下功能，其中参数为 node，对`search()`函数来说，它递归地处理以 node 为根节点的子树。因此 node 相当于这棵子树的 root 节点
1. root 为 null，返回 null（见测试用例 1）
2. root 等于 p 或 q，此时，没有必要继续递归左右子树了，直接返回root（即使遍历了，当前子树还是应该返回 root，所以应该剪去子树）
3. 递归搜索左子树，如果 isFind ，右子树没必要递归了（isFind，右子树上找不到目标了，所以应该剪去）
4. 否则，递归搜索右子树，如果左右子树同时搜索到，说明当前节点即为**最近共同祖先**，标志 isFind
5. 否则，返回左子树或右子树找到的节点（有可能两个都是 null）

![](https://pic.leetcode-cn.com/1599830631-VPDobD-image.png)

#### 复杂度
- 时间复杂度：$O(n)$
- 空间复杂度：平均$O(lgn)$，最差$O(n)$，最佳$O(lgn)$