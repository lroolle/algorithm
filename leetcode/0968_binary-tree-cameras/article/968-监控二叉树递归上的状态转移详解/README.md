"+++
title = "0968. Binary Tree Cameras 968. 监控二叉树:【递归上的状态转移】详解 "
author = ["carlsun-2"]
date = 2020-09-22T03:18:15+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# 968. 监控二叉树:【递归上的状态转移】详解

> [0968. Binary Tree Cameras](https://leetcode-cn.com/problems/binary-tree-cameras/)
> [968. 监控二叉树:【递归上的状态转移】详解](https://leetcode-cn.com/problems/binary-tree-cameras/solution/968-jian-kong-er-cha-shu-di-gui-shang-de-zhuang-ta/) by [carlsun-2](https://leetcode-cn.com/u/carlsun-2/)

## 思路 

这道题目其实不是那么好理解的，题目举的示例不是很典型，会误以为摄像头必须要放在中间，其实放哪里都可以只要覆盖了就行。

这道题目难在两点：

1. 需要确定遍历方式 
2. 需要状态转移的方程 

我们之前做动态规划的时候，只要最难的地方在于确定状态转移方程，至于遍历方式无非就是在数组或者二维数组上。

**而本题，不仅要确定状态转移方式，而且要在树上进行推导，所以难度就上来了，一些同学知道这道题目难，但其实说不上难点究竟在哪。** 

1. 需要确定遍历方式

首先先确定遍历方式，才能确定转移方程，那么该如何遍历呢？

在安排选择摄像头的位置的时候，**我们要从底向上进行推导，因为尽量让叶子节点的父节点安装摄像头，这样摄像头的数量才是最少的** 

如何从低向上推导呢？ 

就是后序遍历也就是左右中的顺序，这样就可以从下到上进行推导了。

后序遍历代码如下：

```
    int traversal(TreeNode* cur) {

        // 空节点，该节点有覆盖
        if (终止条件) return ;

        int left = traversal(cur->left);    // 左
        int right = traversal(cur->right);  // 右

        逻辑处理                            // 中

        return ;
    }
```

**注意在以上代码中我们取了左孩子的返回值，右孩子的返回值，即left 和 right， 以后推导中间节点的状态**

2. 需要状态转移的方程 

确定了遍历顺序，再看看这个状态应该如何转移，先来看看每个节点可能有几种状态：

可以说有如下三种：

* 该节点无覆盖 
* 本节点有摄像头
* 本节点有覆盖

我们分别有三个数字来表示：

* 0：该节点无覆盖 
* 1：本节点有摄像头
* 2：本节点有覆盖

大家应该找不出第四个节点的状态了。

**那么问题来了，空节点究竟是哪一种状态呢？ 空节点表示无覆盖？ 表示有摄像头？还是有覆盖呢？ **

回归本质，为了让摄像头数量最少，我们要尽量让叶子节点的父节点安装摄像头，这样才能摄像头的数量最少。

那么空节点不能是无覆盖的状态，这样叶子节点就可以放摄像头了，空节点也不能是有摄像头的状态，这样叶子节点的父节点就没有必要放摄像头了，而是可以把摄像头放在叶子节点的爷爷节点上。

**所以空节点的状态只能是有覆盖，这样就可以在叶子节点的父节点放摄像头了** 

接下来就是递推关系。

那么递归的终止条件应该是遇到了空节点，此时应该返回2（有覆盖），原因上面已经解释过了。

代码如下：

```
        // 空节点，该节点有覆盖
        if (cur == NULL) return 2;
```

递归的函数，以及终止条件已经确定了，再来看单层逻辑处理。

主要有如下四类情况：

1. 情况1：左右节点都有覆盖

左孩子有覆盖，右孩子有覆盖，那么此时中间节点应该就是无覆盖的状态了。

如图：

![968.监控二叉树2.png](https://pic.leetcode-cn.com/1600744629-AgRMTO-968.%E7%9B%91%E6%8E%A7%E4%BA%8C%E5%8F%89%E6%A0%912.png)
代码如下：

```
        // 左右节点都有覆盖
        if (left == 2 && right == 2) return 0;
```

2. 情况2：左右节点至少有一个无覆盖的情况

如果是以下情况，则中间节点（父节点）应该放摄像头：

left == 0 && right == 0 左右节点无覆盖
left == 1 && right == 0 左节点有摄像头，右节点无覆盖
left == 0 && right == 1 左节点有无覆盖，右节点摄像头
left == 0 && right == 2 左节点无覆盖，右节点覆盖
left == 2 && right == 0 左节点覆盖，右节点无覆盖

这个不难理解，毕竟有一个孩子没有覆盖，父节点就应该放摄像头。

此时摄像头的数量要加一，并且return 1，代表中间节点放摄像头。

代码如下：
```
        if (left == 0 || right == 0) {
            result++;
            return 1;
        }
```

3. 情况3：左右节点至少有一个有摄像头

如果是以下情况，其实就是 左右孩子节点有一个有摄像头了，那么其父节点就应该是2（覆盖的状态）

left == 1 && right == 2 左节点有摄像头，右节点有覆盖
left == 2 && right == 1 左节点有覆盖，右节点有摄像头
left == 1 && right == 1 左右节点都有摄像头

代码如下：

```
        if (left == 1 || right == 1) return 2;
```

**从这个代码中，可以看出，如果left == 1, right == 0 怎么办？其实这种条件在情况2中已经判断过了**，如图：

![968.监控二叉树1.png](https://pic.leetcode-cn.com/1600744648-dlgGvO-968.%E7%9B%91%E6%8E%A7%E4%BA%8C%E5%8F%89%E6%A0%911.png)
这种情况也是大多数同学容易迷惑的情况。

4. 情况4：头结点没有覆盖

以上都处理完了，递归结束之后，可能头结点 还有一个无覆盖的情况，如图：

![968.监控二叉树3.png](https://pic.leetcode-cn.com/1600744659-SFiBRU-968.%E7%9B%91%E6%8E%A7%E4%BA%8C%E5%8F%89%E6%A0%913.png)
所以递归结束之后，还要判断根节点，如果没有覆盖，result++，代码如下：

```
    int minCameraCover(TreeNode* root) {
        result = 0;
        if (traversal(root) == 0) { // root 无覆盖
            result++;
        }
        return result;
    }
```

以上四种情况我们分析完了，代码也差不多了，整体代码如下：

（**以下我的代码是可以精简的，但是我是为了把情况说清楚，特别把每种情况列出来，因为精简之后的代码读者不好理解。**）

## C++代码 

```
class Solution {
private:
    int result;
    int traversal(TreeNode* cur) {

        // 空节点，该节点有覆盖
        if (cur == NULL) return 2;

        int left = traversal(cur->left);    // 左
        int right = traversal(cur->right);  // 右

        // 情况1
        // 左右节点都有覆盖
        if (left == 2 && right == 2) return 0;

        // 情况2
        // left == 0 && right == 0 左右节点无覆盖
        // left == 1 && right == 0 左节点有摄像头，右节点无覆盖
        // left == 0 && right == 1 左节点有无覆盖，右节点摄像头
        // left == 0 && right == 2 左节点无覆盖，右节点覆盖
        // left == 2 && right == 0 左节点覆盖，右节点无覆盖
        if (left == 0 || right == 0) {
            result++;
            return 1;
        }

        // 情况3
        // left == 1 && right == 2 左节点有摄像头，右节点有覆盖
        // left == 2 && right == 1 左节点有覆盖，右节点有摄像头
        // left == 1 && right == 1 左右节点都有摄像头
        // 其他情况前段代码均已覆盖
        if (left == 1 || right == 1) return 2;

        // 以上代码我没有使用else，主要是为了把各个分支条件展现出来，这样代码有助于读者理解
        // 这个 return -1 逻辑不会走到这里。
        return -1;
    }

public:
    int minCameraCover(TreeNode* root) {
        result = 0;
        // 情况4
        if (traversal(root) == 0) { // root 无覆盖
            result++;
        }
        return result;
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
