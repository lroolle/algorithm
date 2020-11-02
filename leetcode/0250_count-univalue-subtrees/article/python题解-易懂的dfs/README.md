"+++
title = "0250. Count Univalue Subtrees python题解--易懂的DFS "
author = ["xiao-xue-66"]
date = 2020-04-11T04:33:44+08:00
tags = ["Leetcode", "Algorithms", "Python", "DepthfirstSearch"]
categories = ["leetcode"]
draft = false
+++

# python题解--易懂的DFS

> [0250. Count Univalue Subtrees](https://leetcode-cn.com/problems/count-univalue-subtrees/)
> [python题解--易懂的DFS](https://leetcode-cn.com/problems/count-univalue-subtrees/solution/pythonti-jie-yi-dong-de-dfs-by-xiao-xue-66/) by [xiao-xue-66](https://leetcode-cn.com/u/xiao-xue-66/)

### 解题思路
![image.png](https://pic.leetcode-cn.com/aaaa82b36d399194ebb349cda987d45c301c21a75acac19a0d8782f8939225eb-image.png)

- 直接看代码比较好理解
- 每次需要判断当前的节点值是否是跟左右节点的值相等,在相等的情况下在看左右子树是否是一根同值树,如果都满足的话那么以当前节点为根节点的树满足同值树,否则不满足
- 使用DFS根据上面的规则遍历每个节点,记录结果数
### 代码

```python3
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def countUnivalSubtrees(self, root: TreeNode) -> int:
        result = 0
        def Solutions(root):
            nonlocal result
            if not root:
                return True
            left_ = Solutions(root.left)
            right_ = Solutions(root.right)
            if not root.right and not root.left: # 都为空
                result += 1
                return True
            elif root.right and root.left:#都不为空
                if root.right.val == root.val and root.val == root.left.val and left_ and right_:
                    result += 1
                    return True
                return False
            elif not root.left and root.right:#左子树为空,右子树不为空
                if root.right.val == root.val and right_:
                    result += 1
                    return True
                return False
            elif not root.right and root.left:#右子树为空,左子树不为空
                if root.left.val == root.val and left_:
                    result += 1
                    return True
                return False
                
        Solutions(root)
        return result
                        
        
```