"+++
title = "0685. Redundant Connection II 685	 冗余连接 II  C语言 并查集 "
author = ["xingmian"]
date = 2020-08-02T07:14:10+08:00
tags = ["Leetcode", "Algorithms", "C", "UnionFind"]
categories = ["leetcode"]
draft = false
+++

# 685	 冗余连接 II  C语言 并查集

> [0685. Redundant Connection II](https://leetcode-cn.com/problems/redundant-connection-ii/)
> [685	 冗余连接 II  C语言 并查集](https://leetcode-cn.com/problems/redundant-connection-ii/solution/685-rong-yu-lian-jie-ii-cyu-yan-bing-cha-ji-by-xin/) by [xingmian](https://leetcode-cn.com/u/xingmian/)

### 解题思路
比684题复杂的地方在于：
如果没有入度为2的点，那么直接按照684题处理
如果有入度为2的点，找到这两条边，先不考虑last这条边：
如果能顺利merge，那么last就是要去掉的。如果不能顺利merge，那么first就是要去掉的

### 代码

```c

 int find_parents(int x, int *parents)
{
    while (x != parents[x]) {
        parents[x] = parents[parents[x]];
        x = parents[x];
    }
    return x;
}

void init(int n, int **parents, int **size)
{
    *parents = (int *)malloc(sizeof(int) * n);
    *size = (int *)malloc(sizeof(int) * n);

    for (int i = 0; i < n; i++) {
        (*parents)[i] = i;
        (*size)[i] = 1;
    }
}

bool merge(int *edge, int *parents, int *size)
{
    int p1 = find_parents(edge[0], parents);
    int p2 = find_parents(edge[1], parents);
    
    if (p1 == p2) {
        return false;
    } else {
        if (size[p1] > size[p2]) {
            parents[p2] = p1;
            size[p1] += size[p2];
        } else {
            parents[p1] = parents[p2];
            size[p2] = size[p1];
        }
    }
    return true;
}

int* findRedundantDirectedConnection(int** edges, int edgesSize, int* edgesColSize, int* returnSize){
    *returnSize = 0;
    if (edges == NULL) {
        return NULL;
    }
    
    int *parents = NULL;
    int *size = NULL;
    init(edgesSize + 1, &parents, &size);
    
    //统计入度
    int *in_degree = (int *)calloc(edgesSize + 1, sizeof(int));
    int node = -1;//入度为2的点
    for (int i = 0; i < edgesSize; i++) {
        in_degree[edges[i][1]]++;
        if (in_degree[edges[i][1]] == 2) {
            node = edges[i][1];
            break;
        }
    }
    
    int *res = (int *)malloc(2 * sizeof(int));
    if (node == -1) {   //说明只是成环了，按照684题即可
        for (int i = 0; i < edgesSize; i++) {
            if (merge(edges[i], parents, size) == false) {
                res[0] = edges[i][0];
                res[1] = edges[i][1];
                *returnSize = 2;
                return res;
            }
        }
    } else {  // 有入度为2的点，去掉这个边，如果还能成环，就该删掉成环的这个边，否则，删除去掉的这个边
        int first = -1; //入度为2的点所在的第一条边
        int last = -1;  //入度为2的点所在的第二条边
        for (int i = 0; i < edgesSize; i++) {
            if (edges[i][1] == node) {
                if (first == -1) {
                    first = i;
                } else {
                    last = i;
                    break;
                }
            }
        }
        printf("first = %d, last = %d\n", first, last);
        
        // 在合并的时候，不合并last这条边
        for (int i = 0; i < edgesSize; i++) {
            if (i == last) {
                continue;
            }
            如果出现环了，出现环的这条就是要删除的
            if (merge(edges[i], parents, size) == false) {
                res[0] = edges[first][0];
                res[1] = edges[first][1];
                *returnSize = 2;
                return res;
            }
        }
        //如果没有成环的，那么last这条就是要删除的
        res[0] = edges[last][0];
        res[1] = edges[last][1];
        *returnSize = 2;
        return res;
        
    }
    
    
    
    return NULL;
}
```