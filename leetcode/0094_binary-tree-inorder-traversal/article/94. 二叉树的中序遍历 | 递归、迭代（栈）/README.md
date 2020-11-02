"+++
title = "0094. Binary Tree Inorder Traversal 94. 二叉树的中序遍历 | 递归、迭代（栈） "
author = ["yiluolion"]
date = 2020-09-14T10:08:36+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Recursion", "迭代", "Stack", "Python"]
categories = ["leetcode"]
draft = false
+++

# 94. 二叉树的中序遍历 | 递归、迭代（栈）

> [0094. Binary Tree Inorder Traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)
> [94. 二叉树的中序遍历 | 递归、迭代（栈）](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/94-er-cha-shu-de-zhong-xu-bian-li-di-gui-die-dai-z/) by [yiluolion](https://leetcode-cn.com/u/yiluolion/)

# 解题思路

---

## 思路：递归、迭代（借助栈）

### 递归

关于中序遍历，我们了解到的说法可能是这样的：**先遍历左子树，然后访问根节点，最后遍历左子树**。不能说错，但是其中还是存在些许的歧义，下面两篇文章篇幅涉及遍历方面存在歧义的情况，有兴趣可以点击下方链接进行了解，（**文章来源：** [笨猪爆破组](https://leetcode-cn.com/u/xiao_ben_zhu/)）

[「手画图解」用栈模拟中序遍历，怎么做以及为什么 | 递归与迭代](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/shou-hua-tu-jie-yong-zhan-mo-ni-zhong-xu-bian-li-z/)

[谈谈别的，前、中、后序遍历的区别只有一点](https://leetcode-cn.com/problems/binary-tree-paths/solution/tu-jie-er-cha-shu-de-suo-you-lu-jing-by-xiao_ben_z/)

先说中序遍历的输出顺序，这一点我们可能会记住类似【左根右】的口诀，也就是左子树、根、左子树。

二叉树的中序遍历，在遍历时会先访问根节点，然后是左子树，再是右子树，**只不过处理节点值是放在访问完左子树之后**（这就是区别于先序、后序遍历的地方）。而在访问左子树或右子树的时候，还是会以同样的方式去遍历。那么这里我们就可以考虑使用递归去解决。

根据上面的分析，这里直接放代码：

```python
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def inorderTraversal(self, root: TreeNode) -> List[int]:
        def inorder(root):
            if not root:
                return
            
            inorder(root.left)
            # 递归完左子树后，处理节点值
            ans.append(root.val)
            inorder(root.right)
        
        ans = []
        inorder(root)
        return ans
```

### 迭代（栈）

进阶中提及，是否能够用迭代的方法实现？

这里我们需要借助栈模拟递归来实现迭代。

这里直接用图示来说明，如何去用栈模拟递归，如下：

![图示](https://pic.leetcode-cn.com/1600077029-qHwyxn-94_%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E4%B8%AD%E5%BA%8F%E9%81%8D%E5%8E%86.gif)


根据上面图示思路，实现代码如下。

```python
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def inorderTraversal(self, root: TreeNode) -> List[int]:
        ans = []
        stack = []

        while stack or root:
            if root:
                # 这部分将左子树压入栈中
                stack.append(root)
                root = root.left
            else:
                # 进入右子树前处理值
                tmp = stack.pop()
                ans.append(tmp.val)
                # 进入右子树，继续循环
                root = tmp.right
            
        return ans
```

## 欢迎关注

---

公众号 【[书所集录](https://i.loli.net/2020/07/09/sNEGeV8g6fmW5Ub.jpg)】

---

 **如有错误，烦请指正，欢迎交流指点。**