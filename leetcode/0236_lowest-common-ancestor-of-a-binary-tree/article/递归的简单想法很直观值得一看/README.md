"+++
title = "0236. Lowest Common Ancestor of a Binary Tree 递归的简单想法，很直观，值得一看 "
author = ["pang-pang-k"]
date = 2020-09-18T16:58:53+08:00
tags = ["Leetcode", "Algorithms", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# 递归的简单想法，很直观，值得一看

> [0236. Lowest Common Ancestor of a Binary Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)
> [递归的简单想法，很直观，值得一看](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/solution/di-gui-de-jian-dan-xiang-fa-hen-zhi-guan-zhi-de-yi/) by [pang-pang-k](https://leetcode-cn.com/u/pang-pang-k/)

### 解题思路
初看这个题，就觉得应该是这么干！
就是根据两个点的分布来决定向左子树递归还是右子树递归，或者直接返回根节点，于是，噔噔噔噔，做出来啦~

1、写一个find(root,target)函数，返回是否能在一棵树中找到这个节点。
2、开始主函数的递归。
3、如果根节点等于两个节点中的一个，直接返回根节点。为啥呢？因为这棵树一定得包括这两个节点，才会走到这一步，且看后面的递归条件。
4、root.left如果能够找到p,q，那么递归左子树，右子树同理，如果不能同时找到，说明两个点分布在两棵树的两边，那就返回根节点。
5、写完了，就是这么直观。
后面附上另外一种高智商的大佬才能想出来的高质量递归，我这辈子是想不出来的，人跟人的差距，就是注定的吧😭
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
 * @param {TreeNode} root
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {TreeNode}
 */

function find(root,target){
        if(root==null) return false;
        return root.val==target.val||find(root.left,target)||find(root.right,target);
    }

var lowestCommonAncestor = function(root, p, q) {

    if(root.val==p.val||root.val==q.val){
        return root;
    }else if(find(root.left,p)&&find(root.left,q)){
        return lowestCommonAncestor(root.left,p,q);
    }else if(find(root.right,p)&&find(root.right,q)){
        return lowestCommonAncestor(root.right,p,q);
    }else{
        return root;
    }
    }

//根据另一位大佬的思路，自己撸的代码，真心不直观但效率高啊，膜拜！！！

var lowestCommonAncestor = function(root, p, q) {
    if(!root||root.val==p.val||root.val==q.val) return root;
    let left=lowestCommonAncestor(root.left,p,q);
    let right=lowestCommonAncestor(root.right,p,q);

    if(!left&&!right) return null;
    if(!left) return right;
    if(!right) return left; 
    return root;
};

```