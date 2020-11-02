"+++
title = "0113. Path Sum II 帮助璐璐同学找思路 "
author = ["bad_bird"]
date = 2020-09-06T15:08:56+08:00
tags = ["Leetcode", "Algorithms", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# 帮助璐璐同学找思路

> [0113. Path Sum II](https://leetcode-cn.com/problems/path-sum-ii/)
> [帮助璐璐同学找思路](https://leetcode-cn.com/problems/path-sum-ii/solution/bang-zhu-lu-lu-tong-xue-zhao-si-lu-by-bad_bird/) by [bad_bird](https://leetcode-cn.com/u/bad_bird/)

### 解题思路
可以利用DFS的思路

- 当到达叶子结点，并且数组和等于sum，将数据拷贝一份到res。
- 因为题目没有给出节点大小是正数的条件，因此不能剪枝。
### 代码

```javascript
var pathSum = function(root, sum) {
    if(!root){
        return [];
    }
    let res = [];
    dfs(root, sum, root.val, [root.val], res );
    return res;
};

let dfs  = function(node, sum, now, nowarr, res){
    if(now == sum && !node.left && !node.right){
        res.push(nowarr.map(it=>it));
        return;
    }
    if(node.left){
        dfs(node.left, sum, now + node.left.val, nowarr.concat([node.left.val]), res);
    }
    if(node.right){
        dfs(node.right, sum, now + node.right.val, nowarr.concat([node.right.val]), res);
    }
}
```