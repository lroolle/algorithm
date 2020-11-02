"+++
title = "0617. Merge Two Binary Trees 617. 合并二叉树-都是老套路了 "
author = ["xiao-jian-feng"]
date = 2020-07-07T17:31:23+08:00
tags = ["Leetcode", "Algorithms", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# 617. 合并二叉树-都是老套路了

> [0617. Merge Two Binary Trees](https://leetcode-cn.com/problems/merge-two-binary-trees/)
> [617. 合并二叉树-都是老套路了](https://leetcode-cn.com/problems/merge-two-binary-trees/solution/617-he-bing-er-cha-shu-du-shi-lao-tao-lu-liao-by-x/) by [xiao-jian-feng](https://leetcode-cn.com/u/xiao-jian-feng/)

### 解题思路
看到这题，立马想到这肯定要遍历树啊，二叉树遍历？上模板！！
算了，二叉树三种递归方式的模板已经不想上了，自己默写去。
这里挑个前序遍历吧

其实这道题跟 [226. 翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree/)比较类似，亲爱的你细细体味下（噢严格来说二叉树遍历相关的题都比较类似）

代码很短，看下注释吧

### 代码

```javascript
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} t1
 * @param {TreeNode} t2
 * @return {TreeNode}
 */
var mergeTrees = function(t1, t2) {
    if(!t1) return t2 //若t1节点为空，那直接返回t2节点，不管t2是否为空
    if(!t2) return t1 //若t2为空，那肯定t1肯定不为空，返回t1节点
    t1.val = t1.val + t2.val //能执行到这里证明t1与t2节点均不为空，那就两值相加，替换t1原来的值
    t1.left = mergeTrees(t1.left, t2.left ) //递归遍历两者的左子树
    t1.right = mergeTrees(t1.right, t2.right) ////递归遍历两者的右左子树
    return t1 //t1必然是返回的根节点，为啥？因为都拼到t1树上了啊
};
```