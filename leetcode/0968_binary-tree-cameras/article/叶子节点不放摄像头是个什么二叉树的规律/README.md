"+++
title = "0968. Binary Tree Cameras 叶子节点不放摄像头是个什么二叉树的规律？ "
author = ["xiexb"]
date = 2020-09-22T02:58:26+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# 叶子节点不放摄像头是个什么二叉树的规律？

> [0968. Binary Tree Cameras](https://leetcode-cn.com/problems/binary-tree-cameras/)
> [叶子节点不放摄像头是个什么二叉树的规律？](https://leetcode-cn.com/problems/binary-tree-cameras/solution/gen-jie-dian-bu-fang-she-xiang-tou-shi-ge-shi-yao-/) by [xiexb](https://leetcode-cn.com/u/xiexb/)

### 解题思路
![图片.png](https://pic.leetcode-cn.com/1600738363-bcplVs-%E5%9B%BE%E7%89%87.png)
除根节点外，每一个节点都是父节点的子节点，所以：
1、根节点没有父节点，要考虑自己
2、非根节点只考虑子节点即可，因为自己会被作为父节点的子节点被考虑到

把节点值当作节点状态：
0：没拍别人也没被拍，需要被人拍
1：被拍了
2：拍别人

![图片.png](https://pic.leetcode-cn.com/1600742817-yIzwTU-%E5%9B%BE%E7%89%87.png)
### 代码

```python3
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def minCameraCover(self, root: TreeNode) -> int:
         
        self.res = 0
        def lrd(node):
            if node is None:
                return 1 #空节点不需要被人拍也不用拍别人，直接返回被拍了就好
            left = lrd(node.left)
            right = lrd(node.right)
            if left == 0 or right == 0: 
                #如果左儿子或者右儿子需要被拍，我就装个摄像机
                self.res += 1
                return 2        
            if left == 2 or right == 2:
                #如果左儿子或者右儿子装了摄像机，那我就被拍了
                return 1
            else:# left == 1 and right == 1:
                #如果左儿子和右儿子都是被拍的，都没有摄像机，那我就是需要被拍的状态
                return 0

        if lrd(root) == 0:
            ##看看根节点是不是需要被拍
            self.res += 1
        return self.res
```