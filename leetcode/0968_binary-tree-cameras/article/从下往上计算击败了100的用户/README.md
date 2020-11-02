"+++
title = "0968. Binary Tree Cameras 从下往上计算，击败了100%的用户 "
author = ["sdwwld"]
date = 2020-09-22T06:42:29+08:00
tags = ["Leetcode", "Algorithms", "Java", "Tree"]
categories = ["leetcode"]
draft = false
+++

# 从下往上计算，击败了100%的用户

> [0968. Binary Tree Cameras](https://leetcode-cn.com/problems/binary-tree-cameras/)
> [从下往上计算，击败了100%的用户](https://leetcode-cn.com/problems/binary-tree-cameras/solution/cong-xia-wang-shang-ji-suan-ji-bai-liao-100de-yong/) by [sdwwld](https://leetcode-cn.com/u/sdwwld/)


这题正常的逻辑是从上往下，但我们还可以逆向思维，从下往上来统计。那么一个节点就会有3种情况
- 1，当前节点有相机
- 2，当前节点不需要相机（子节点有相机把它给覆盖了）
- 3，当前节点没有相机并且也没有被子节点给覆盖（那么他只能等他的父节点把它给覆盖了）

```
    //NO_CAMERA表示的是子节点没有相机，当前节点也没放相机
    private final int NO_CAMERA = 0;
    //HAS_CAMERA表示当前节点有一个相机
    private final int HAS_CAMERA = 1;
    //NO_NEEDED表示当前节点没有相机，但他的子节点有一个相机，把它给
    //覆盖了，所以它不需要了。或者他是一个空的节点也是不需要相机的
    private final int NO_NEEDED = 2;

    //全局的，统计有多少相机
    int res = 0;

    public int minCameraCover(TreeNode root) {
        //边界条件判断
        if (root == null)
            return 0;
        //如果最后返回的是NO_CAMERA，表示root节点的子节点也没有相机，
        //所以root节点要添加一个相机
        if (dfs(root) == NO_CAMERA)
            res++;
        //返回结果
        return res;
    }

    public int dfs(TreeNode root) {
        //如果是空的，就不需要相机了
        if (root == null)
            return NO_NEEDED;
        int left = dfs(root.left), right = dfs(root.right);
        //如果左右子节点有一个是NO_CAMERA，表示的是子节点既没相机，也没相机覆盖它，
        //所以当前节点需要有一个相机
        if (left == NO_CAMERA || right == NO_CAMERA) {
            //在当前节点放一个相机，统计相机的个数
            res++;
            return HAS_CAMERA;
        }
        //如果左右子节点只要有一个有相机，那么当前节点就不需要相机了，否则返回一个没有相机的标记
        return left == HAS_CAMERA || right == HAS_CAMERA ? NO_NEEDED : NO_CAMERA;
    }

```

看一下运行结果
![image.png](https://pic.leetcode-cn.com/1600756304-RNkShA-image.png)


<br>

**如果觉得有用就给个赞吧，你的赞是给我最大的鼓励，也是我写作的最大动力**

![image.png](https://pic.leetcode-cn.com/d56a80459005b444404d2ad6fbaabdabecd2b9ed3cb2cf432e570c315ae2fcf7-image.png)
##### 查看更多答案，可扫码关注我微信公众号“**[数据结构和算法](https://img-blog.csdnimg.cn/20200807155236311.png)**”