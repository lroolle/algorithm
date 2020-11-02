"+++
title = "0617. Merge Two Binary Trees 617. 合并二叉树:【三种递归】【一种迭代】详解 "
author = ["carlsun-2"]
date = 2020-09-23T01:20:01+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# 617. 合并二叉树:【三种递归】【一种迭代】详解

> [0617. Merge Two Binary Trees](https://leetcode-cn.com/problems/merge-two-binary-trees/)
> [617. 合并二叉树:【三种递归】【一种迭代】详解](https://leetcode-cn.com/problems/merge-two-binary-trees/solution/617-he-bing-er-cha-shu-san-chong-di-gui-yi-chong-d/) by [carlsun-2](https://leetcode-cn.com/u/carlsun-2/)

几个月前写了一个题解，写的不是清晰，乘着每日一天，重新写一篇！
## 思路 

相信这道题目很多同学疑惑的点是如何同时遍历两个二叉树呢？ 

其实和遍历一个树逻辑是一样的，只不过传入两个树的节点，同时操作。

那么前中后序应该使用哪种遍历呢？ 

**本题使用哪种遍历都是可以的！**

我们下面以前序遍历为例。

动画如下：
![617.合并二叉树.gif](https://pic.leetcode-cn.com/1600824986-KMKaJm-617.%E5%90%88%E5%B9%B6%E4%BA%8C%E5%8F%89%E6%A0%91.gif)
那么我们来按照递归三部曲来解决：

1. **确定递归函数的参数和返回值：**
首先那么要合入两个二叉树，那么参数至少是要传入两个二叉树的根节点，返回值就是合并之后二叉树的根节点。

代码如下：

```
TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
```

2. **确定终止条件：** 

因为是传入了两个树，那么就有两个树遍历的节点t1 和 t2，如果t1 == NULL 了，两个树合并就应该是 t2 了啊（如果t2也为NULL也无所谓）。

反过来如果t2 == NULL，那么两个数合并就是t1（如果t1也为NULL也无所谓）。

代码如下：

```
if (t1 == NULL) return t2; // 如果t1为空，合并之后就应该是t2
if (t2 == NULL) return t1; // 如果t2为空，合并之后就应该是t1
```
3. **确定单层递归的逻辑：**

单层递归的逻辑就比较好些了，这里我们用重复利用一下t1这个树，t1就是合并之后树的根节点（所谓的修改了元数据的结构）。

那么单层递归中，就要把两棵树的元素加到一起。
```
t1->val += t2->val;
```

那么此时t1 的左子树 应该是 合并 t1左子树 t2左子树之后的左子树，t1 的右子树 应该是 合并 t1右子树 t2右子树之后的右子树。

代码如下：

```
        t1->left = mergeTrees(t1->left, t2->left);
        t1->right = mergeTrees(t1->right, t2->right);
        return t1;
```

此时前序遍历，修改原输入树结构的完整代码就写出来了，如下：

```
class Solution {
public:
    TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
        if (t1 == NULL) return t2; // 如果t1为空，合并之后就应该是t2
        if (t2 == NULL) return t1; // 如果t2为空，合并之后就应该是t1
        // 修改了t1的数值和结构
        t1->val += t2->val;                             // 中
        t1->left = mergeTrees(t1->left, t2->left);      // 左
        t1->right = mergeTrees(t1->right, t2->right);   // 右
        return t1;
    }
};
```

那么中序遍历可不可以呢，也是可以的，代码如下：

```
class Solution {
public:
    TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
        if (t1 == NULL) return t2; // 如果t1为空，合并之后就应该是t2
        if (t2 == NULL) return t1; // 如果t2为空，合并之后就应该是t1
        // 修改了t1的数值和结构
        t1->left = mergeTrees(t1->left, t2->left);      // 左
        t1->val += t2->val;                             // 中
        t1->right = mergeTrees(t1->right, t2->right);   // 右
        return t1;
    }
};
```

后序遍历呢，依然可以，代码如下：

```
class Solution {
public:
    TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
        if (t1 == NULL) return t2; // 如果t1为空，合并之后就应该是t2
        if (t2 == NULL) return t1; // 如果t2为空，合并之后就应该是t1
        // 修改了t1的数值和结构
        t1->left = mergeTrees(t1->left, t2->left);      // 左
        t1->right = mergeTrees(t1->right, t2->right);   // 右
        t1->val += t2->val;                             // 中
        return t1;
    }
};
```

**但是前序遍历是最好理解的，我建议大家用前序遍历来做就OK。**

**那么如下还总结了四种方法，递归的方式均使用了前序遍历，此时大家应该知道了，以下每一种递归的方法都可以换成中序和后序遍历，所以本题的解法是很多的。**

**其实这道题目迭代法实现是比较困难的，大家可以试一试，是一道不错的面试进阶题目。**

四种写法如下：

1. 递归修改了输入树的结构 
2. 递归不修改树的结构 
3. 递归，一波指针的操作，自己写的野路子（可以用来深度理解一下C++的指针）
4. 迭代（这应该是最简单直观的迭代法代码了，一看就懂）

## C++代码

### 递归

修改了输入树的结构，前序遍历
```
class Solution {
public:
    TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
        if (t1 == NULL) return t2; // 如果t1为空，合并之后就应该是t2
        if (t2 == NULL) return t1; // 如果t2为空，合并之后就应该是t1
        // 修改了t1的数值和结构
        t1->val += t2->val;
        t1->left = mergeTrees(t1->left, t2->left);
        t1->right = mergeTrees(t1->right, t2->right);
        return t1;
    }
};
```

不修改输入树的结构，前序遍历
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
想要更改二叉树的值，应该传入指向指针的指针， 如果process(t1, t2);这么写的话，其实只是传入的一个int型的指针，并没有传入地址，要传入指向指针的指针才能完成对t1的修改。
（前序遍历）
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

这应该是最简单直观的迭代法了，模拟的层序遍历。
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
**我是程序员Carl，先后在腾讯和百度从事技术研发多年，利用工作之余重刷leetcode，更多[精彩算法文章](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzUxNjY5NTYxNA==&action=getalbum&album_id=1485825793120387074&scene=173#wechat_redirect)尽在公众号：[代码随想录](https://img-blog.csdnimg.cn/20200815195519696.png)，关注后，回复「Java」「C++」「python」「简历模板」等等，有我整理多年的学习资料，可以加我[微信](https://img-blog.csdnimg.cn/20200814140330894.png)，备注「简单自我介绍」+「组队刷题」，拉你进入刷题群（无任何广告，纯个人分享），每天一道经典题目分析，我选的每一道题目都不是孤立的，而是由浅入深一脉相承的，如果跟住节奏每篇连续着看，定会融会贯通。本文  [https://github.com/youngyangyang04/leetcode-master](https://github.com/youngyangyang04/leetcode-master ) 已经收录，里面还有leetcode刷题攻略、各个类型经典题目刷题顺序、思维导图，可以fork到自己仓库，有空看一看一定会有所收获，如果对你有帮助也给一个star支持一下吧！**
**以下资料希望对你有帮助：**

* [帮你把KMP算法学个通透！（理论篇）B站](https://www.bilibili.com/video/BV1PD4y1o7nd/)
* [帮你把KMP算法学个通透！（求next数组代码篇）B站](https://www.bilibili.com/video/BV1M5411j7Xx/)
* [如何学习C++？ B站](https://www.bilibili.com/video/BV1rK4y1e7ed)
* [手把手带你读C++ primer！B站](https://www.bilibili.com/video/BV1Z5411874t)
* [leetcode刷题攻略](https://github.com/youngyangyang04/leetcode-master)
* [程序员应该如何写简历（附简历模板）](https://mp.weixin.qq.com/s/PkBpde0PV65dJjj9zZJYtg)
* [一线互联网公司技术面试的流程以及注意事项](https://mp.weixin.qq.com/s/1VMvQ_6HbVpEn85CNilTiw)
* [C++面试&C++学习指南知识点整理](https://github.com/youngyangyang04/TechCPP)

**如果感觉题解对你有帮助，不要吝啬给一个👍吧！**
