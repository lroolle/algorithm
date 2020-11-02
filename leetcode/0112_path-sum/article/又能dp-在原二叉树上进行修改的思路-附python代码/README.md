"+++
title = "0112. Path Sum 又能DP: 在原二叉树上进行修改的思路; 附Python代码 "
author = ["xiao-yan-gou"]
date = 2020-07-06T17:01:10+08:00
tags = ["Leetcode", "Algorithms", "DynamicProgramming", "Python", "Python3", "C", "C++", "Java", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# 又能DP: 在原二叉树上进行修改的思路; 附Python代码

> [0112. Path Sum](https://leetcode-cn.com/problems/path-sum/)
> [又能DP: 在原二叉树上进行修改的思路; 附Python代码](https://leetcode-cn.com/problems/path-sum/solution/you-neng-dp-zai-yuan-er-cha-shu-shang-jin-xing-xiu/) by [xiao-yan-gou](https://leetcode-cn.com/u/xiao-yan-gou/)

这个题其实又能动态规划(不难的动态规划)，
我们一层一层把值加起来，到时候只看底部是否有叶子节点能和sum值相等即可
&ensp;
这么光说可能有点抽象，示例如图：
宗旨就是把左边这棵树里的值更新成右边这样
（其中两个数字的计算写了一下式子，其他位置以此类推）
然后拿底部叶子节点的值与sum比较
![gaitubao_屏幕快照 2020-07-07 上午12.49.48.png](https://pic.leetcode-cn.com/c31ad0d33274886bcfd1285f04b6a473b59625f8bfee022f70873d74d320af60-gaitubao_%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202020-07-07%20%E4%B8%8A%E5%8D%8812.49.48.png)

&ensp;
附Python代码，用一个函数的嵌套实现“进入新节点并更新节点的值”的操作
```python
class Solution:
    def hasPathSum(self, root: TreeNode, sum: int) -> bool:
        if not root:
            return False

        if not (root.left or root.right):
            return root.val == sum

        if root.left:
            root.left.val = root.val + root.left.val

        if root.right:
            root.right.val = root.val + root.right.val
        
        return self.hasPathSum(root.left,sum) or self.hasPathSum(root.right,sum)
```

如果不修改原二叉树，其实就是下面这个写法（改用减法）
```python
class Solution:
    def hasPathSum(self, root: TreeNode, sum: int) -> bool:
        if not root:
            return False
        if not root.left and not root.right:
            return sum == root.val
        return self.hasPathSum(root.left, sum - root.val) or self.hasPathSum(root.right, sum - root.val)
```

