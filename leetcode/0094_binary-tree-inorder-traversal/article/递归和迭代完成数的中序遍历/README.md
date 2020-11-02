"+++
title = "0094. Binary Tree Inorder Traversal 递归和迭代完成数的中序遍历 "
author = ["50-hui"]
date = 2020-09-03T07:47:20+08:00
tags = ["Leetcode", "Algorithms", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# 递归和迭代完成数的中序遍历

> [0094. Binary Tree Inorder Traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)
> [递归和迭代完成数的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/di-gui-he-die-dai-wan-cheng-shu-de-zhong-xu-bian-l/) by [50-hui](https://leetcode-cn.com/u/50-hui/)

### 解题思路
1.递归较简单:
    直接上代码:
```javascript
var inorderTraversal = function(root) {
    let arr = [];
    if(!root) return [];

   let nums = [root];
    //递归方法实现
    MiddleOrder(root,arr);
    return arr;
};

function MiddleOrder(root,arr){
    if(root.left){
        MiddleOrder(root.left,arr);
    }
    arr.push(root.val);
    if(root.right){
        MiddleOrder(root.right,arr);
    }
}
```
### 解题思路
2.迭代完成，主要使用了辅助的栈来完成遍历

    将根节点入栈，
    定义一个flag(表示左子树有没有遍历)，初始值为false，表示没有经历遍历

    while循环 当栈的长度为0时跳出循环
        判断是否存在左节点左节点是否完成遍历,
            设置栈里面的第一个元素(最先出来的元素)为tmp
            判断tmp存在左节点,没有完成遍历:
                将左节点存放入栈中
            判断tmp存在左节点，并完成了遍历或者不存在左节点
                将值push到arr数组，
                并改变左子树为遍历完成，即flag = true;
                将该节点出栈
                判断是否存在右子树
                    存在：
                        修改flag(表示给节点的左子树没有遍历)
                        将该节点压入栈中，
                    不存在：不做任何操作
    

    返回 arr 即可。

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
 * @return {number[]}
 */
//递归遍历
var inorderTraversal = function(root) {
    let arr = [];
    if(!root) return [];
   
   let nums = [root];
   let flag = false;
   let tmp ;

    //递归方法实现
    // MiddleOrder(root,arr);
    
    //迭代方法实现
    while(nums.length){
        tmp = nums[0];
        if(tmp.left&&!flag){
            nums.unshift(tmp.left);
        }else{
            arr.push(tmp.val);
            nums.shift();
            flag = true;
            if(tmp.right){
                nums.unshift(tmp.right);
                flag = false;
            }
        }
        
    }
    return arr;
};

