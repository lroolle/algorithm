"+++
title = "0617. Merge Two Binary Trees C++ 广度优先遍历，单队列，比官方代码精简点 "
author = ["quickq"]
date = 2020-09-23T09:30:48+08:00
tags = ["Leetcode", "Algorithms", "cpp", "BreadthfirstSearch"]
categories = ["leetcode"]
draft = false
+++

# C++ 广度优先遍历，单队列，比官方代码精简点

> [0617. Merge Two Binary Trees](https://leetcode-cn.com/problems/merge-two-binary-trees/)
> [C++ 广度优先遍历，单队列，比官方代码精简点](https://leetcode-cn.com/problems/merge-two-binary-trees/solution/c-yan-du-you-xian-bian-li-dan-dui-lie-bi-guan-fang/) by [quickq](https://leetcode-cn.com/u/quickq/)

### 解题思路
不额外分配空间，原树调整
两棵树t1,t2：
1、以t1为基准，进行广度优先遍历；
2、当前遍历的节点都有值，则相加；
3、发现t1当前遍历节点的左子树或右子树和t2当前遍历节点长的不一样（t1左子树为空&t2左子树不为空 或 t1右子树为空&t2右子树不为空），把t2中不一样的子树挪到t1中，并且将t2挪走的节点置空；
bfs结束返回t1。
### 代码

```cpp
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
        if (t1 == nullptr) return t2;

		 queue<pair<TreeNode*,TreeNode*>> q;
		 q.emplace(make_pair(t1,t2));

		 while (!q.empty()) {
			 auto current = q.front();
			 if (current.first != nullptr) {
				 if (current.second != nullptr) {
					 current.first->val += current.second->val;
				 }

				 if (current.first->left == nullptr) {
					 if (current.second !=nullptr && current.second->left != nullptr) {
						 current.first->left = current.second->left;
						 current.second->left = nullptr;
					 }
				 }

				 if (current.first->right == nullptr) {
					 if (current.second != nullptr && current.second->right != nullptr) {
						 current.first->right = current.second->right;
						 current.second->right = nullptr;
					 }
				 }
			 }

			 q.pop();

			 if (current.first != nullptr && current.second != nullptr) {
				 q.emplace(make_pair(current.first->left, current.second->left));
				 q.emplace(make_pair(current.first->right, current.second->right));
			 }
		 }

		 return t1;
	 }
};
```