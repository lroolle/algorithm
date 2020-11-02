"+++
title = "0404. Sum of Left Leaves 炒鸡容易理解，其实和求全路径是一样的思想~ "
author = ["Omooo"]
date = 2019-12-12T01:22:13+08:00
tags = ["Leetcode", "Algorithms", "Java", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 炒鸡容易理解，其实和求全路径是一样的思想~

> [0404. Sum of Left Leaves](https://leetcode-cn.com/problems/sum-of-left-leaves/)
> [炒鸡容易理解，其实和求全路径是一样的思想~](https://leetcode-cn.com/problems/sum-of-left-leaves/solution/chao-ji-rong-yi-li-jie-qi-shi-he-qiu-quan-lu-jing-/) by [Omooo](https://leetcode-cn.com/u/omooo/)

```
class Solution {
    int sum = 0;
    private void getSum(TreeNode root, boolean isLeft){
        if(root != null){
            if(isLeft && root.left == null && root.right == null){
                sum += root.val;
            }else{
                getSum(root.left, true);
                getSum(root.right, false);
            }
        }
    }
    public int sumOfLeftLeaves(TreeNode root) {
        getSum(root, false);
        return sum;
    }
}
```
