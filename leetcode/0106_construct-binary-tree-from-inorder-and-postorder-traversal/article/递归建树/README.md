"+++
title = "0106. Construct Binary Tree from Inorder and Postorder Traversal 递归建树 "
author = ["yi-wen-statistics"]
date = 2020-09-16T09:29:01+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# 递归建树

> [0106. Construct Binary Tree from Inorder and Postorder Traversal](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)
> [递归建树](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/solution/di-gui-jian-shu-by-yi-wen-statistics/) by [yi-wen-statistics](https://leetcode-cn.com/u/yi-wen-statistics/)

## 方法1——递归
和前面那个题一个道理，后序遍历【左子树，右子树，根节点】，中序遍历【左子树，根节点，右子树】，所以根据后序遍历可以得知根节点的位置，再根据根节点的位置从中序遍历中找到左右子树的位置，然后重复这一过程，自底向上构建唯一地二叉树。

### 代码

```python3
class Solution:
    def buildTree(self, inorder: List[int], postorder: List[int]) -> TreeNode:
        def dfs(in_l, in_r, post_l, post_r):
            if post_l < post_r:
                return None
            inorder_root = Hash[postorder[post_l]]

            root = TreeNode(postorder[post_l])
            subtree_size = in_r - inorder_root
            root.right = dfs(inorder_root+1, in_r, post_l-1, post_l-subtree_size)
            root.left = dfs(in_l, inorder_root-1, post_l-subtree_size-1, post_r)
            return root

        Hash = {key:value for value, key in enumerate(inorder)}
        return dfs(0, len(inorder)-1, len(postorder)-1, 0)
```