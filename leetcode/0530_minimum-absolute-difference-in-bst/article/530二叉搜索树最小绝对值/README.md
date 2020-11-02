"+++
title = "0530. Minimum Absolute Difference in BST 530.二叉搜索树最小绝对值 "
author = ["mystery-z"]
date = 2020-05-29T09:57:51+08:00
tags = ["Leetcode", "Algorithms", "Java", "二叉树"]
categories = ["leetcode"]
draft = false
+++

# 530.二叉搜索树最小绝对值

> [0530. Minimum Absolute Difference in BST](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/)
> [530.二叉搜索树最小绝对值](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/solution/530er-cha-sou-suo-shu-zui-xiao-jue-dui-zhi-by-myst/) by [mystery-z](https://leetcode-cn.com/u/mystery-z/)

### 解题思路
这道题主要还是利用搜索二叉树的特性，中序遍历是一个升序数组，而最小值的产生一定是在数组中相邻两个元素的差之中，因此，在中序遍历时候抓住前一个数，和当前数字的差 于最小值作比较
我给出了编译三次，不同效率的代码，希望有所帮助。
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
//第三次修改后的代码 ： 0ms 100% 39.7MB 25%
    int last = -1;
    int min = Integer.MAX_VALUE;
    public int getMinimumDifference(TreeNode root) {
        inorde(root);
        return min;
    }
    public void inorde (TreeNode root){
        if(root ==null)return;
        inorde(root.left);
        if(last == -1){
            last = root.val;
        }else{
            min= (root.val-last)<min? root.val-last:min;
            last = root.val;
        }
        inorde(root.right);
    }
//第二次的代码  1ms 40.3MB
   int min = Integer.MAX_VALUE;
    ArrayList<Integer>arrayList = new ArrayList<>();
    public int getMinimumDifference(TreeNode root) {
        inorde(root);
        return min;
    }
    public void inorde (TreeNode root){
        if(root ==null)return;
        inorde(root.left);
        if(!arrayList.isEmpty())
        if(root.val-arrayList.get(arrayList.size()-1)<min)min = root.val-arrayList.get(arrayList.size()-1);
        arrayList.add(root.val);
        inorde(root.right);
    }
//第一次的代码 2ms 40.3MB
    ArrayList<Integer> arrayList = new ArrayList<>();
    public int getMinimumDifference(TreeNode root) {
        inorde(root);
        int now = arrayList.size()-1;
        int min = Integer.MAX_VALUE;
        while(now>0){
            int A = arrayList.get(now);
            int B = arrayList.get(now-1);
            if(A-B < min)
                min = A-B;
            now--;
        }
        return min;
    }
    public void inorde (TreeNode root){
        if(root ==null)return;
        inorde(root.left);
        arrayList.add(root.val);
        inorde(root.right);
    }
}
```