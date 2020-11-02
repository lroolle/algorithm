"+++
title = "0113. Path Sum II Java DFS递归 易懂 100.00% "
author = ["lonepwq"]
date = 2020-08-17T14:19:38+08:00
tags = ["Leetcode", "Algorithms", "Java", "DFS", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# Java DFS递归 易懂 100.00%

> [0113. Path Sum II](https://leetcode-cn.com/problems/path-sum-ii/)
> [Java DFS递归 易懂 100.00%](https://leetcode-cn.com/problems/path-sum-ii/solution/java-dfsdi-gui-yi-dong-10000-by-lonepwq/) by [lonepwq](https://leetcode-cn.com/u/lonepwq/)

执行用时：1 ms, 在所有 Java 提交中击败了100.00%的用户
内存消耗：40 MB, 在所有 Java 提交中击败了75.77%的用户
### 解题思路
完成此题之前，请先完成[112. 路径总和](https://leetcode-cn.com/problems/path-sum/);
此题在112题上，有以下两点变化：
- 要返回成功路径；
  此问题只需传入一个list，递归同时记录路径；
- 有多条成功路径；
  此问题只需全盘遍历即可，而不是有一个结果即返回；

下面上代码（代码中有注释）：

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
    List<List<Integer>> res = new ArrayList();//全局变量来获取所有的结果路径

    public List<List<Integer>> pathSum(TreeNode root, int sum) {
        LinkedList<Integer> temp = new LinkedList<>();// 定义一个有序list来存储路径
        helper(root, sum, temp);
        return res;
    }

    public void helper(TreeNode node, int sum, LinkedList<Integer> temp) {
        if (node == null) {
            return;
        }
        temp.addLast(node.val);// 记录路径
        if (node.left == null && node.right == null && sum == node.val) {
            res.add(new ArrayList(temp));
        }
        helper(node.left, sum - node.val, temp);
        helper(node.right, sum - node.val, temp);
        temp.removeLast();// 重点，遍历完后，需要把当前节点remove出去，因为用的是同一个list对象来存所有的路径
    }
}
```