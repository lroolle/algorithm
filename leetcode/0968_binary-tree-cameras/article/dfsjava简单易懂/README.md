"+++
title = "0968. Binary Tree Cameras dfs，Java简单易懂 "
author = ["rorke76753"]
date = 2020-09-22T01:19:03+08:00
tags = ["Leetcode", "Algorithms", "Java", "DepthfirstSearch"]
categories = ["leetcode"]
draft = false
+++

# dfs，Java简单易懂

> [0968. Binary Tree Cameras](https://leetcode-cn.com/problems/binary-tree-cameras/)
> [dfs，Java简单易懂](https://leetcode-cn.com/problems/binary-tree-cameras/solution/dfsjavajian-dan-yi-dong-by-rorke76753/) by [rorke76753](https://leetcode-cn.com/u/rorke76753/)

根据题意可以确定一个节点可以有三个状态：0.没装摄像机但被观测到、1.没装摄像机但没被观测到、2.装了摄像机。

因为我们使用的是递归地进行深度优先搜索，因此本层的状态是由下一层的状态而决定的，而状态根据这样来分析得到的：一个摄像头最多影响3层（本层、下一层、上一层）；下一层如果能被观测到但没有摄像头，那么本层肯定是不被观测得到的（0->1）；如果下一层安装了摄像头，本层能够被观测到（2->0）；下一层如果不能被观测得到，那么本层一定安装了摄像头（1->2）。 ps：括号内容为 下一层状态为a 决定 本层状态为b

根据上面得想法，来讨论叶节点的子节点应该是什么状态：如果空节点的状态是不可被观测（1），那么上一层（叶节点层）必须安装摄像头（2），但是叶节点安装摄像机的话，那么这个摄像机贡献度为2（自己与叶节点的父节点），这样的贡献度是最大的吗？如果在叶节点的父节点安装摄像机，贡献度为4，这才是最大的贡献度，也就意味着能安装最少的监控。因此我们必须要在叶节点的父节点处安装摄像机才能确保总摄像机最少。根据这些分析，空节点状态为2无意义而且空节点的状态必不为1。

或者用这样一个例子来思考：
如果叶节点为2， 那么这台摄像机就覆盖了2个节点：叶节点本身以及其父节点。
如果叶节点不为2，因为所有节点都需要被覆盖到，因此只能是它的父节点装有摄像机，那么这台摄像机可以覆盖到最多4个节点。
很显然，在叶节点安装相机的总相机数不是最少的，举个例子方便说明：
>我们有这样一颗树
     0
    / \\
    0 0

如果在叶节点安装摄像机，那么需要2部才能将所有节点覆盖完，而在叶节点的父节点安装只需要1部。
因此我们的叶节点的状态应该是为1的。

这里对空节点的思考主要是指叶节点的子节点，对非叶节点的空子节点讨论没有意义：这是因为安装摄像机的位置不能确定，使得总摄像机数最少。

节点为空，根据上面的定义：状态为2没有意义，状态为1就意味着叶子节点需要安装摄像机得不到最少的答案；因此我们定义它的状态为0，那么意味着空节点的子节点（虚拟定义出来，并不存在的一个节点）为2。

根据上述分析不难得出状态的转移：
- 子节点为0，父节点为1（子节点无摄像机但被观测到了）
- 子节点为1，父节点为2（子节点不能被观测到，要么在子节点安装摄像机要么在父节点安装，上述分析已经知道应该在父节点安装）
- 子节点为2，父节点为0（子节点安装了摄像机，因此父节点能被观测到）

统计dfs时，`return 2`的节点的数目即可，但还需要注意的是如果`dfs(root)==1`，那么root节点还需要放多一部摄像机。


# 代码如下

```java
class Solution {
    int result = 0;
    public int minCameraCover(TreeNode root) {
        if(dfs(root)==1){
            result++;
        }
        return result;
    }
    //0:可被观测但无监控，上一层节点为1
    //1：不可被观测到，上一层节点为2
    //2：有摄像机，上一层节点为0
    private int dfs(TreeNode root){
        if(root==null){
            return 0;
        }
        int leftStatus = dfs(root.left),rightStatus = dfs(root.right);
        if(leftStatus==1||rightStatus==1){
            result++;
            return 2;
        }else if(leftStatus==2||rightStatus==2){
            return 0;
        }else{
            return 1;
        }
    }
}
```