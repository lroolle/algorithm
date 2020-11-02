"+++
title = "0685. Redundant Connection II java 并查集 "
author = ["luma730"]
date = 2020-09-16T08:51:33+08:00
tags = ["Leetcode", "Algorithms", "Java", "UnionFind"]
categories = ["leetcode"]
draft = false
+++

# java 并查集

> [0685. Redundant Connection II](https://leetcode-cn.com/problems/redundant-connection-ii/)
> [java 并查集](https://leetcode-cn.com/problems/redundant-connection-ii/solution/java-bing-cha-ji-by-luma730-2/) by [luma730](https://leetcode-cn.com/u/luma730/)

### 解题思路
并查集：
547:朋友圈
684:冗余连接（无向图）
685:冗余连接（有向图）
737:句子相似性2
990:等式方程可满足性
1319:连通网络操作次数
952: 按公因数计算最大组件大小
### 代码

```java
class Solution {
   
   int[] result = new int[2];
    int doubleRoot = 0;
    int[] hadRoot;
    int[][] rootResult = new int[2][2];

    public int[] findRedundantDirectedConnection(int[][] edges) {
        hadRoot = new int[edges.length + 1];
        int[] father = new int[edges.length + 1];
        for (int i = 0; i < father.length; i++) {
            father[i] = i;
        }

        for (int[] edge : edges) {
            hadRoot[edge[1]]++;
            if (hadRoot[edge[1]] == 2) {
                doubleRoot = edge[1];
                rootResult[1] = edge;
            } else {
                union(father, edge[1], edge[0]);
            }
        }
        if (doubleRoot != 0) {
            for (int[] edge : edges) {
                if (edge[1] == doubleRoot) {
                    rootResult[0] = edge;
                    break;
                }
            }
            int root = 0;
            for (int i = 1; i < father.length; i++) {
                if (root == 0) {
                    root = findXFather(father,i);
                }
                if (root != findXFather(father,i)) {
                    return rootResult[0];
                }
            }
            return rootResult[1];
        }
        return result;
    }

    public int findXFather(int[] father, int x) {
        while (father[x] != x) {
            father[x] = father[father[x]];
            x = father[x];
        }
        return x;
    }

    public void union(int[] father, int x, int y) {
        int xFather = findXFather(father, x);
        int yFather = findXFather(father, y);
        if (xFather != yFather) {
            father[xFather] = yFather;
        } else {
            result[0] = y;
            result[1] = x;
        }
    }
}
```
![image.png](https://pic.leetcode-cn.com/1600246265-clxBIt-image.png)
