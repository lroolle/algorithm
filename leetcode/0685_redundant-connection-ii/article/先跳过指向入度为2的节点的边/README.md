"+++
title = "0685. Redundant Connection II 先跳过指向入度为2的节点的边 "
author = ["kobe24o"]
date = 2020-05-28T10:30:36+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# 先跳过指向入度为2的节点的边

> [0685. Redundant Connection II](https://leetcode-cn.com/problems/redundant-connection-ii/)
> [先跳过指向入度为2的节点的边](https://leetcode-cn.com/problems/redundant-connection-ii/solution/xian-tiao-guo-zhi-xiang-ru-du-wei-2de-jie-dian-de-/) by [kobe24o](https://leetcode-cn.com/u/kobe24o/)

- 参考 [数据结构--并查集（Disjoint-Set）](https://blog.csdn.net/qq_21201267/article/details/105294270)

类似题目：
[LeetCode 684. 冗余连接（并查集）](https://michael.blog.csdn.net/article/details/106396495)
[LeetCode 886. 可能的二分法（着色DFS/BFS/拓展并查集）](https://michael.blog.csdn.net/article/details/106401724)
[LeetCode 990. 等式方程的可满足性（并查集）](https://michael.blog.csdn.net/article/details/106354260)
[LeetCode 959. 由斜杠划分区域（并查集）](https://michael.blog.csdn.net/article/details/105739007)
[LeetCode 1202. 交换字符串中的元素（并查集）](https://michael.blog.csdn.net/article/details/106042356)
[LeetCode 1319. 连通网络的操作次数（BFS/DFS/并查集）](https://michael.blog.csdn.net/article/details/106087256)
[程序员面试金典 - 面试题 17.07. ***名字（并查集）](https://michael.blog.csdn.net/article/details/105300614)
本题是684题的升级版本

- 先找有没有入度为2的节点 node
- 然后遍历边的时候先跳过指向 node 的边
- 最后再遍历指向 node 的边
![在这里插入图片描述](https://pic.leetcode-cn.com/b338ad3ed9a804f4e591e55731211f597b75e1c4af1abab2b3f83dfb76840bb4.png)

```cpp
class dsu
{
	vector<int> f;
public:
	dsu(int n)
	{
		f.resize(n+1);
		for(int i = 0; i < n+1; ++i)
			f[i] = i;
	}
	void merge(int a, int b)
	{
		int fa = find(a), fb = find(b);
		f[fa] = fb;
	}
	int find(int a)//循环+路径压缩
	{
        int origin = a;
		while(a != f[a])
			a = f[a];
		return f[origin] = a;//路径压缩
	}
};
class Solution {
public:
    vector<int> findRedundantDirectedConnection(vector<vector<int>>& edges) {
        int N = edges.size();
        vector<int> indegree(N+1, 0);
        int node = -1;
        for(auto e : edges)
        {
        	indegree[e[1]]++;
        	if(indegree[e[1]] > 1)
        		node = e[1];
        }
        dsu u(N);
        int x, y;
        vector<vector<int>> E;
        for(auto e : edges)
        {
        	if(node == e[1])//边指向node，先跳过
        	{
        		E.push_back(e);
        		continue;
        	}
        	if(u.find(e[0]) != u.find(e[1]))//两个没有连接
        		u.merge(e[0], e[1]);//把边连接起来
        	else//已经连接了,有环
        		x = e[0], y = e[1];//记录下来
        }
        for(auto e : E)
        {
        	if(u.find(e[0]) != u.find(e[1]))//两个没有连接
        		u.merge(e[0], e[1]);//把边连接起来
        	else//已经连接了,有环
        		x = e[0], y = e[1];//记录下来
        }
        return {x, y};
    }
};
```
20 ms	9.9 MB