"+++
title = "0404. Sum of Left Leaves 递归传入当前节点所在位置 "
author = ["developer-n"]
date = 2020-07-31T02:44:25+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Tree", "Python"]
categories = ["leetcode"]
draft = false
+++

# 递归传入当前节点所在位置

> [0404. Sum of Left Leaves](https://leetcode-cn.com/problems/sum-of-left-leaves/)
> [递归传入当前节点所在位置](https://leetcode-cn.com/problems/sum-of-left-leaves/solution/di-gui-chuan-ru-dang-qian-jie-dian-suo-zai-wei-zhi/) by [developer-n](https://leetcode-cn.com/u/developer-n/)

### 解题思路
1. 叶子节点中止
2. 递归传入当前节点所处的左右位置，0,1标识
### 代码

```python3
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def sumOfLeftLeaves(self, root: TreeNode) -> int:

        def collect_left_node(node,loc,leaf_val):
            if node:
                
                # 1. 叶子节点中止
                if not node.left and not node.right:
                    if loc == 0:
                        leaf_val.append(node.val)
                        # print(node.val,loc,leaf_val)
                    return 
                # 2. 递归传入当前节点所处的左右位置，0,1标识
                if node.left:
                    collect_left_node(node.left,0,leaf_val)
                if node.right:
                    collect_left_node(node.right,1,leaf_val)
            else:
                return 

        leaf_val = []
        collect_left_node(root,1,leaf_val)
        print(leaf_val)
        return sum(leaf_val)
            
                
```