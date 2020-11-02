"+++
title = "0404. Sum of Left Leaves 左叶子之和，简单判断左叶子节点 "
author = ["zui-weng-jiu-xian"]
date = 2020-08-02T12:46:48+08:00
tags = ["Leetcode", "Algorithms", "Java", "Tree", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 左叶子之和，简单判断左叶子节点

> [0404. Sum of Left Leaves](https://leetcode-cn.com/problems/sum-of-left-leaves/)
> [左叶子之和，简单判断左叶子节点](https://leetcode-cn.com/problems/sum-of-left-leaves/solution/jian-dan-pan-duan-zuo-xie-zi-jie-dian-by-zui-weng-/) by [zui-weng-jiu-xian](https://leetcode-cn.com/u/zui-weng-jiu-xian/)

### 解题思路（递归遍历）
1. 遍历一次二叉树，每次对根节点进行判断它是否有左叶子节点即可。
2. 遍历方法可以采取前序遍历，中序遍历，后序遍历和层次遍历。
3. 这里采用前序遍历，问题的关键是怎么判断一个节点是否为左叶子节点。
4. 我们知道，当一个节点没有左右节点时，该节点是叶子结点，但是无法区分是左子节点还是右子节点。
5. 要区分左子节点和右子节点，要通过它们的父节点进行判断：当一个节点有左子节点，并且这个左子节点没有左右节点，说明这个左子节点是左叶子节点。

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
    int ans = 0;
    public int sumOfLeftLeaves(TreeNode root) {
        if (root == null) { // 递归出口
            return 0;
        }
        // 根节点有左子节点，并且左子节点没有左右节点，说明左子节点是左叶子节点
        if (root.left != null && root.left.left == null && root.left.right == null) {
            ans += root.left.val;
        }
        sumOfLeftLeaves(root.left);
        sumOfLeftLeaves(root.right);
        return ans;
    }

}
```
### 算法分析：
- 设二叉树节点个数为n.
- 对二叉树的所有节点访问一次且仅一次，时间复杂度为$O(n)$。
- 最坏情况下，二叉树退化为链表结构，递归需要$O(n)$的栈辅助空间。


**如果对你有帮助，请给个赞吧^_^**