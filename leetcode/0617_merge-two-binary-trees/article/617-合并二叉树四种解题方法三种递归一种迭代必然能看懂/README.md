"+++
title = "0617. Merge Two Binary Trees 617. 合并二叉树：四种解题方法（三种递归，一种迭代）必然能看懂！ "
author = ["carlsun-2"]
date = 2020-07-31T04:17:26+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# 617. 合并二叉树：四种解题方法（三种递归，一种迭代）必然能看懂！

> [0617. Merge Two Binary Trees](https://leetcode-cn.com/problems/merge-two-binary-trees/)
> [617. 合并二叉树：四种解题方法（三种递归，一种迭代）必然能看懂！](https://leetcode-cn.com/problems/merge-two-binary-trees/solution/si-chong-jie-ti-fang-fa-san-chong-di-gui-yi-chong-/) by [carlsun-2](https://leetcode-cn.com/u/carlsun-2/)

## 思路 

四种写法，总有一款适合你，其实这道题目迭代法实现是比较困难的，大家可以试一试，是一道不错的面试进阶题目。

1. 第一种写法，递归修改了输入树的结构 
2. 第二种写法，递归不修改树的结构 
3. 第三种写法，一波指针的操作，自己写的野路子（可以用来深度理解一下C++的指针）
4. 第四种写法，迭代（这应该是最简单直观的迭代法代码了，一看就懂）

四种写法代码如下（详细注释）

## C++代码

### 递归

修改了输入树的结构
```
class Solution {
public:
    TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
        if (t1 == NULL) return t2;
        if (t2 == NULL) return t1;
        // 修改了t1的数值和结构
        t1->val += t2->val;
        t1->left = mergeTrees(t1->left, t2->left);
        t1->right = mergeTrees(t1->right, t2->right);
        return t1;
    }
};
```

不修改输入树的结构
```
class Solution {
public:
    TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
        if (t1 == NULL) return t2;
        if (t2 == NULL) return t1;
        // 重新定义新的节点，不修改原有两个树的结构
        TreeNode* root = new TreeNode(0);
        root->val = t1->val + t2->val;
        root->left = mergeTrees(t1->left, t2->left);
        root->right = mergeTrees(t1->right, t2->right);
        return root;
    }
};
```
一波指针的操作，自己写的野路子
想要更改二叉树的值，应该传入指向指针的指针。 
```
class Solution {
public:
    void process(TreeNode** t1, TreeNode** t2) {
        if ((*t1) == NULL && (*t2) == NULL) return;
        if ((*t1) != NULL && (*t2) != NULL) {
            (*t1)->val += (*t2)->val;
        }
        if ((*t1) == NULL && (*t2) != NULL) {
            *t1 = *t2;
            return;
        }
        if ((*t1) != NULL && (*t2) == NULL) {
            return;
        }
        process(&((*t1)->left), &((*t2)->left));
        process(&((*t1)->right), &((*t2)->right));
    }
    TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
        process(&t1, &t2);
        return t1;
    }
};
```
### 迭代

这应该是最简单直观的迭代法代码了
```
class Solution {
public:
    TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
        if (t1 == NULL) return t2;
        if (t2 == NULL) return t1;
        queue<TreeNode*> que;
        que.push(t1);
        que.push(t2);
        while(!que.empty()) {
            TreeNode* node1 = que.front(); que.pop();
            TreeNode* node2 = que.front(); que.pop();
            // 此时两个节点一定不为空，val相加
            node1->val += node2->val;
            // 如果左节点都不为空，加入队列
            if (node1->left != NULL && node2->left != NULL) {
                que.push(node1->left);
                que.push(node2->left);
            }
            // 如果右节点都不为空，加入队列
            if (node1->right != NULL && node2->right != NULL) {
                que.push(node1->right);
                que.push(node2->right);
            }
            // 当t1的左节点 为空 t2左节点不为空，就赋值过去
            if (node1->left == NULL && node2->left != NULL) {
                node1->left = node2->left;
            }
            // 当t1的右节点 为空 t2右节点不为空，就赋值过去
            if (node1->right == NULL && node2->right != NULL) {
                node1->right = node2->right;
            }
        }
        return t1;
    }
};
```

**我是程序员Carl，先后在腾讯和百度从事技术研发多年，利用工作之余重刷leetcode，更多精彩文章尽在公众号：[代码随想录](https://img-blog.csdnimg.cn/20200815195519696.png)，关注后，回复「Java」「C++」「python」「简历模板」等等，有我整理多年的学习资料，可以加我[微信](https://img-blog.csdnimg.cn/20200814140330894.png)，备注「组队刷题」，拉你进入刷题群，每天一道经典题目分析，我选的每一道题目都不是孤立的，而是由浅入深一脉相承的，如果跟住节奏每篇连续着看，定会融会贯通。本文  [https://github.com/youngyangyang04/leetcode-master](https://github.com/youngyangyang04/leetcode-master ) 已经收录，里面还有leetcode刷题攻略、各个类型经典题目刷题顺序、思维导图，看一看一定会有所收获！**

**以下资料希望对你有帮助：**
* [如何学习C++B站视频](https://www.bilibili.com/video/BV1rK4y1e7ed)
* [手把手带你读C++ primer](https://www.bilibili.com/video/BV1Z5411874t)
* [C++面试&C++学习指南知识点整理](https://github.com/youngyangyang04/TechCPP)
* [程序员应该如何写简历（附简历模板）](https://mp.weixin.qq.com/s/PkBpde0PV65dJjj9zZJYtg)
* [leetcode刷题攻略](https://github.com/youngyangyang04/leetcode-master)
* [究竟什么时候可以使用库函数，什么时候不要使用库函数，过来人来说一说](https://leetcode-cn.com/circle/article/E1Kjzn/)
* [究竟什么是时间复杂度，怎么求时间复杂度，看这一篇就够了](https://mp.weixin.qq.com/s/lYL9TSxLqCeFXIdjt4dcIw)
* [一文带你彻底理解程序为什么会超时](https://mp.weixin.qq.com/s/T-vcJSkq2-0s0bBB-itWbQ)
* [一场面试，带你彻底掌握递归算法的时间复杂度](https://mp.weixin.qq.com/s/Kt-Mvs8LeVqidLGUqySj1g)
* [算法分析中的空间复杂度，你真的会了么？](https://mp.weixin.qq.com/s/sXjjnOUEQ4Gf5F9QCRzy7g)

**如果感觉题解对你有帮助，不要吝啬给一个👍吧！**

