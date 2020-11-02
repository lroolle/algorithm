"+++
title = "0701. Insert into a Binary Search Tree 701. 二叉搜索树中的插入操作:【递归法】【迭代法】详解 "
author = ["carlsun-2"]
date = 2020-08-13T11:42:49+08:00
tags = ["Leetcode", "Algorithms", "cpp", "BinarySearchTree"]
categories = ["leetcode"]
draft = false
+++

# 701. 二叉搜索树中的插入操作:【递归法】【迭代法】详解

> [0701. Insert into a Binary Search Tree](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/)
> [701. 二叉搜索树中的插入操作:【递归法】【迭代法】详解](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/solution/701-er-cha-sou-suo-shu-zhong-de-cha-ru-cao-zuo-di-/) by [carlsun-2](https://leetcode-cn.com/u/carlsun-2/)

## 思路 

其实这道题目是一道很简单的题目，**但是题目中的提示：有多种有效的插入方式，还可以重构二叉搜索树，一下子吓退了不少人**，瞬间感觉题目复杂了很多。

其实**可以不考虑题目中提示所说的改变树的结构的插入方式。**

如下演示视频中可以看出：只要按照二叉搜索树的规则去遍历，遇到空节点就插入节点就可以了。

例如插入元素10 ，需要找到末尾节点插入便可，一样的道理来插入元素15，插入元素0，插入元素6，**需要调整二叉树的结构么？ 并不需要**。只需要遍历二叉搜索树，找到空节点 插入元素就可以了， 那么这道题其实就非常简单了。
![701.二叉搜索树中的插入操作.mp4](7beb48e6-3c87-4309-8241-80197c8514b0)
接下来就是遍历二叉搜索树的过程了。

代码如下：

## C++代码

### 递归 
```
class Solution {
public:
    TreeNode* insertIntoBST(TreeNode* root, int val) {
        if (root == NULL) {
            TreeNode* node = new TreeNode(val);
            return node;
        }
        if (root->val > val) root->left = insertIntoBST(root->left, val);
        if (root->val < val) root->right = insertIntoBST(root->right, val);
        return root;
    }
};
```

### 迭代

在来看看迭代法

在迭代法遍历的过程中，需要记录一下当前遍历的节点的父节点，这样才能做插入节点的操作。

```
class Solution {
public:
    TreeNode* insertIntoBST(TreeNode* root, int val) {
        if (root == NULL) {
            TreeNode* node = new TreeNode(val);
            return node;
        }
        TreeNode* cur = root;
        TreeNode* parent = root; // 这个很重要，需要记录上一个节点，否则无法赋值新节点
        while (cur != NULL) {
            parent = cur;
            if (cur->val > val) cur = cur->left;
            else cur = cur->right;
        }
        TreeNode* node = new TreeNode(val);
        if (val < parent->val) parent->left = node;// 此时是用parent节点的进行赋值
        else parent->right = node;
        return root;
    }
};
```

**我是程序员Carl，先后在腾讯和百度从事技术研发多年，利用工作之余重刷leetcode，更多精彩文章尽在公众号：[代码随想录](https://img-blog.csdnimg.cn/20200815195519696.png)，关注后，回复「Java」「C++」「python」「简历模板」等等，有我整理多年的学习资料，可以加我[微信](https://img-blog.csdnimg.cn/20200814140330894.png)，备注「组队刷题」，拉你进入刷题群（无任何广告，纯个人分享），每天一道经典题目分析，我选的每一道题目都不是孤立的，而是由浅入深一脉相承的，如果跟住节奏每篇连续着看，定会融会贯通。本文  [https://github.com/youngyangyang04/leetcode-master](https://github.com/youngyangyang04/leetcode-master ) 已经收录，里面还有leetcode刷题攻略、各个类型经典题目刷题顺序、思维导图，看一看一定会有所收获！**
**以下资料希望对你有帮助：**
* [手把手带你读C++ primer！B站](https://www.bilibili.com/video/BV1Z5411874t)
* [C++ primer 第一章，你要知道的知识点还有这些！B站](https://www.bilibili.com/video/BV1Kv41117Ya)
* [C++面试&C++学习指南知识点整理](https://github.com/youngyangyang04/TechCPP)
* [程序员应该如何写简历（附简历模板）](https://mp.weixin.qq.com/s/PkBpde0PV65dJjj9zZJYtg)
* [leetcode刷题攻略](https://github.com/youngyangyang04/leetcode-master)

**如果感觉题解对你有帮助，不要吝啬给一个👍吧！**
