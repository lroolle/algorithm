"+++
title = "0701. Insert into a Binary Search Tree [C题解] 二叉树递归插入 "
author = ["justdoitno1"]
date = 2020-08-10T16:12:28+08:00
tags = ["Leetcode", "Algorithms", "C", "C++", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# [C题解] 二叉树递归插入

> [0701. Insert into a Binary Search Tree](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/)
> [[C题解] 二叉树递归插入](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/solution/cti-jie-er-cha-shu-di-gui-cha-ru-by-justdoitno1/) by [justdoitno1](https://leetcode-cn.com/u/justdoitno1/)

### 解题思路
此处撰写解题思路
LeetCode服务器不稳定，挑一次最好看的时空数据，心机boy
![image.png](https://pic.leetcode-cn.com/e6e283845c8bfd538d467b0d7d33b119c7096e4d396ad4f6a85ca3e23b6a31a7-image.png)
### 代码

```c
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
// 第一思路想通过像是链表的顺撸，是失败了。
这种方法到底是否可行，可能是需要考虑的边界过多。欢迎大佬帮帮我，奶农很难。
struct TreeNode* insertIntoBST(struct TreeNode* root, int val){
    if(root==NULL) return NULL;
    
    struct TreeNode* tmp = root;

    while(tmp!=NULL) {
    
        if(val < tmp->val) {  // 找到待插入的位置，一定为叶子节点
            if(tmp->left==NULL) {
                struct TreeNode* newNode = malloc(sizeof(struct TreeNode));
                newNode->val = val;
                newNode->left = NULL;
                newNode->right = NULL;
                tmp->left = newNode;
                break;
            }
            tmp = tmp->left;
        }
        if(val > tmp->val) {  // 找到待插入的位置，一定为叶子节点
            if(tmp->right==NULL) {
                struct TreeNode* newNode = malloc(sizeof(struct TreeNode));
                newNode->val = val;
                newNode->left = NULL;
                newNode->right = NULL;
                tmp->left = newNode;
                break;
            }                    
            tmp = tmp->right;
        }
    }
    return root;
}
*/
```
```
// 递归求解，二叉树题目确实第一思路应该是递归，这就好写多了。
struct TreeNode* insertIntoBST(struct TreeNode* root, int val){
    if(root==NULL) {
        struct TreeNode* node = malloc(sizeof(struct TreeNode));
        node->val = val;
        node->left = NULL;
        node->right = NULL;
        return node;
    }

    if(root->val > val) {
        root->left = insertIntoBST(root->left, val);
    }

    if(root->val < val) {
        root->right = insertIntoBST(root->right, val);
    }

    return root;

}
```