"+++
title = "0685. Redundant Connection II 最简单方法，不用分类，一看就懂 "
author = ["leck"]
date = 2020-04-11T07:04:59+08:00
tags = ["Leetcode", "Algorithms", "C++", "UnionFind"]
categories = ["leetcode"]
draft = false
+++

# 最简单方法，不用分类，一看就懂

> [0685. Redundant Connection II](https://leetcode-cn.com/problems/redundant-connection-ii/)
> [最简单方法，不用分类，一看就懂](https://leetcode-cn.com/problems/redundant-connection-ii/solution/zui-jian-dan-fang-fa-bu-yong-fen-lei-yi-kan-jiu-do/) by [leck](https://leetcode-cn.com/u/leck/)

emmm，做完之后发现我和评论区大神的方法都不太一样，大神们都在分类讨论，只有我傻傻的归为一种情况。

思路如下：
1. 首先遍历第一次找到入度为2的点，要删的边肯定和这个点有关。这个点记为*root*
2. 遍历第二次，把非指向 root 的边用并查集合并；指向root的边用个数组记录下来 另一点。
3. 遍历数组记录的点，判断这个点和root是否指向同一点，如果不是的话，说明这俩没有联通，那么合并。指向同一点，那么这条边[这点，root]，完全可以删去，返回这条边就行了。(其实就两个点，因为只有一条边是多余的，所以也可以简单的用if判断 )
4. 特判当入度全为1的情况，这样返回环中最后出现的边就可以
```
class Solution {
public:
    int pre[1005],ind[1005];
    int find(int x){
        return pre[x]==x?x:pre[x]=find(pre[x]);
    }
    vector<int> findRedundantDirectedConnection(vector<vector<int>>& edges) {
        int n=edges.size(),root=-1;
        vector<int> v;
        for(int i=1;i<=n;++i) pre[i]=i;
        for(auto i:edges){
            ind[i[1]]++;
            if(ind[i[1]]==2)
                root=i[1];
        }
        if(root == -1) {
            for(auto i : edges) {
                int a = find(i[0]), b = find(i[1]);
                if(a != b) pre[b] = a;
                else return i;
            }
        }
        for(auto i:edges){
            if(i[1]!=root){
                pre[find(i[1])]=find(i[0]);
            }
            else
                v.emplace_back(i[0]);
        }
        if(find(root)!=find(v[0]))
            return {v[1],root};
        else
            return {v[0],root};
    }
};
```
