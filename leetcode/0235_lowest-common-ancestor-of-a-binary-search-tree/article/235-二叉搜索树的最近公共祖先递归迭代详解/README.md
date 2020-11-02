"+++
title = "0235. Lowest Common Ancestor of a Binary Search Tree 235. 二叉搜索树的最近公共祖先:【递归】【迭代】详解 "
author = ["carlsun-2"]
date = 2020-09-27T01:48:45+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# 235. 二叉搜索树的最近公共祖先:【递归】【迭代】详解

> [0235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
> [235. 二叉搜索树的最近公共祖先:【递归】【迭代】详解](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/235-er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu--24/) by [carlsun-2](https://leetcode-cn.com/u/carlsun-2/)

## 思路 
遇到这个题目首先想的是要是能自底向上查找就好了，这样就可以找到公共祖先了，可惜二叉树只能自上向低。

那么自上相下查找的话，如何记录祖先呢？

做过[236. 二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)题目的同学，应该知道，只要判断一个节点的左子树里有p，右子树里有q，那么当前节点就是最近公共祖先。

那么本题是二叉搜索树，二叉搜索树是有序的，那得好好利用一下这个特点。

在有序树里，如果判断一个节点的左子树里有p，右子树里有q呢？

其实**只要从上到下遍历的时候，如果 (p->val <=  cur->val && cur->val <= q->val)则说明该节点cur就是最近公共祖先了。**

理解这一点，本题就很好解了。

如图所示

![235.二叉搜索树的最近公共祖先.png](https://pic.leetcode-cn.com/1601171298-QLvqHD-235.%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91%E7%9A%84%E6%9C%80%E8%BF%91%E5%85%AC%E5%85%B1%E7%A5%96%E5%85%88.png)
在遍历二叉搜索树的时候就是寻找区间[p->val, q->val]（注意这里是左闭右闭）

那么如果 cur->val 大于 p->val，同时 cur->val 大于q->val，那么就应该向左遍历。（因为我们此时不知道p和q谁大，所以两个都要判断）

代码如下：

```
        if (cur->val > p->val && cur->val > q->val) {
            return traversal(cur->left, p, q);
        }
```

如果 cur->val 小于 p->val，同时 cur->val 小于 q->val，那么就应该向右遍历。

```
        } else if (cur->val < p->val && cur->val < q->val) {
            return traversal(cur->right, p, q);
        }
```

剩下的情况，我们就找到了区间使（p->val <=  cur->val && cur->val <= q->val）或者是 （q->val <=  cur->val && cur->val <= p->val）

代码如下：
```
        else {
            return cur;
        }

```

那么整体递归代码如下:

## C++递归代码

（我这里特意把递归的过程抽出一个函数traversal，这样代码更清晰，有助于读者理解。）

```
class Solution {
private:
    TreeNode* traversal(TreeNode* cur, TreeNode* p, TreeNode* q) {
        if (cur->val > p->val && cur->val > q->val) {
            return traversal(cur->left, p, q);
        } else if (cur->val < p->val && cur->val < q->val) {
            return traversal(cur->right, p, q);
        } else return cur;
    }
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {

        return traversal(root, p, q);
    }
};
```
## C++迭代法代码

同时给出一个迭代的版本，思想是一样的，代码如下：

```
class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        while(root) {
            if (root->val > p->val && root->val > q->val) {
                root = root->left;
            } else if (root->val < p->val && root->val < q->val) {
                root = root->right;
            } else return root;
        }
        return NULL;
    }
};
```
**我是程序员Carl，先后在腾讯和百度从事技术研发多年，利用工作之余重刷leetcode，更多[精彩算法文章](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzUxNjY5NTYxNA==&action=getalbum&album_id=1485825793120387074&scene=173#wechat_redirect)尽在公众号：[代码随想录](https://img-blog.csdnimg.cn/20200815195519696.png)，关注后，回复「Java」「C++」「python」「简历模板」等等，有我整理多年的学习资料，可以加我[微信](https://img-blog.csdnimg.cn/20200814140330894.png)，备注「简单自我介绍」+「组队刷题」，拉你进入刷题群（无任何广告，纯个人分享），每天一道经典题目分析，我选的每一道题目都不是孤立的，而是由浅入深一脉相承的，如果跟住节奏每篇连续着看，定会融会贯通。本文  [https://github.com/youngyangyang04/leetcode-master](https://github.com/youngyangyang04/leetcode-master ) 已经收录，里面还有leetcode刷题攻略、各个类型经典题目刷题顺序、思维导图，可以fork到自己仓库，有空看一看一定会有所收获，如果对你有帮助也给一个star支持一下吧！**
**以下资料希望对你有帮助：**

* [手把手带你读C++ primer！B站](https://www.bilibili.com/video/BV1Z5411874t)
* [如何学习C++？ B站](https://www.bilibili.com/video/BV1rK4y1e7ed)
* [帮你把KMP算法学个通透！（理论篇）B站](https://www.bilibili.com/video/BV1PD4y1o7nd/)
* [帮你把KMP算法学个通透！（求next数组代码篇）B站](https://www.bilibili.com/video/BV1M5411j7Xx/)
* [leetcode刷题攻略](https://github.com/youngyangyang04/leetcode-master)
* [程序员应该如何写简历（附简历模板）](https://mp.weixin.qq.com/s/PkBpde0PV65dJjj9zZJYtg)
* [一线互联网公司技术面试的流程以及注意事项](https://mp.weixin.qq.com/s/1VMvQ_6HbVpEn85CNilTiw)
* [C++面试&C++学习指南知识点整理](https://github.com/youngyangyang04/TechCPP)

**如果感觉题解对你有帮助，不要吝啬给一个👍吧！**
