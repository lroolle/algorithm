"+++
title = "0685. Redundant Connection II 分别讨论入度为2和入度为1的情况 "
author = ["jin-mu-yan-3"]
date = 2019-07-22T10:09:19+08:00
tags = ["Leetcode", "Algorithms", "C++", "UnionFind", "Graph"]
categories = ["leetcode"]
draft = false
+++

# 分别讨论入度为2和入度为1的情况

> [0685. Redundant Connection II](https://leetcode-cn.com/problems/redundant-connection-ii/)
> [分别讨论入度为2和入度为1的情况](https://leetcode-cn.com/problems/redundant-connection-ii/solution/fen-bie-tao-lun-ru-du-wei-2he-ru-du-wei-1de-qing-k/) by [jin-mu-yan-3](https://leetcode-cn.com/u/jin-mu-yan-3/)

- 入度为2代表这个点有两个父节点那么肯定要去点一个，那就都尝试一次从后往前看看一旦去除仍可连通就return
- 如果没有入度为2.那就从后往前考虑入度为1的点去掉。意思让这个点为根(根没有父节点入度为0)

```
class Solution {
public:
    int n;
    vector<int> fa;
    int find(int x) {
        return fa[x] == x ? x : fa[x] = find(fa[x]);
    }
    vector<int> findRedundantDirectedConnection(vector<vector<int>>& edges) {
        n = edges.size(); 
        fa.resize(n + 1);
        vector<int> degree(n + 1);
        for (auto& v : edges) {
            ++degree[v[1]];
        }
        for (int i = n - 1; i >= 0; --i) {
            if (degree[edges[i][1]] == 2) {
                if (helper(edges, i)) return edges[i];
            }
        }
        for (int i = n - 1; i >= 0; --i) {
            if (degree[edges[i][1]] == 1) {
                if (helper(edges, i)) return edges[i];
            }
        }
        return {};
    }
    bool helper(vector<vector<int>>& e, int except) {
        for (int i = 1; i <= n; ++i) fa[i] = i;
        int cnt = n;
        for (int i = 0; i < n; ++i) {
            if (i == except) continue;
            int fx = find(e[i][0]), fy = find(e[i][1]);
            if (fx != fy) {
                --cnt;
                fa[fy] = fx;
            }
        }
        return cnt == 1;
    }
};
```
