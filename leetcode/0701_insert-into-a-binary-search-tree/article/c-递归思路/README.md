"+++
title = "0701. Insert into a Binary Search Tree C++ 递归+思路 "
author = ["youlookdeliciousc"]
date = 2020-05-11T16:16:09+08:00
tags = ["Leetcode", "Algorithms", "C++", "Recursion", "BinarySearchTree"]
categories = ["leetcode"]
draft = false
+++

# C++ 递归+思路

> [0701. Insert into a Binary Search Tree](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/)
> [C++ 递归+思路](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/solution/c-di-gui-si-lu-by-youlookdeliciousc/) by [youlookdeliciousc](https://leetcode-cn.com/u/youlookdeliciousc/)

### 思路
与搜索操作类似，对于每个节点，我们将：

1. 根据节点值与目标节点值的关系，搜索左子树或右子树；
2. 重复步骤 1 直到到达外部节点；
3. 根据节点的值与目标节点的值的关系，将新节点添加为其左侧或右侧的子节点。

这样，我们就可以添加一个新的节点并依旧维持二叉搜索树的性质。
```cpp
class Solution {
public:
    TreeNode* insertIntoBST(TreeNode* root, int val) {
        if(!root)   return new TreeNode(val);
        if(val > root -> val)
            root -> right = insertIntoBST(root -> right, val);
        else
            root -> left = insertIntoBST(root -> left, val);
        return root;
    }
};
```
