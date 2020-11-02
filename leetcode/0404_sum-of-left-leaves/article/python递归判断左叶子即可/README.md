"+++
title = "0404. Sum of Left Leaves Python递归，判断左叶子即可 "
author = ["intwzt"]
date = 2020-09-03T17:33:50+08:00
tags = ["Leetcode", "Algorithms", "Python"]
categories = ["leetcode"]
draft = false
+++

# Python递归，判断左叶子即可

> [0404. Sum of Left Leaves](https://leetcode-cn.com/problems/sum-of-left-leaves/)
> [Python递归，判断左叶子即可](https://leetcode-cn.com/problems/sum-of-left-leaves/solution/pythondi-gui-pan-duan-zuo-xie-zi-ji-ke-by-intwzt/) by [intwzt](https://leetcode-cn.com/u/intwzt/)

**左叶子条件：**
*1. 左节点不为空。
2. 左节点的左节点为空。
3. 左节点的右节点为空。*

class Solution(object):
    def sumOfLeftLeaves(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        if not root:
            return 0
        if root.left != None and root.left.left == None and root.left.right == None:
            return root.left.val + self.sumOfLeftLeaves(root.right)
        return self.sumOfLeftLeaves(root.left) + self.sumOfLeftLeaves(root.right)