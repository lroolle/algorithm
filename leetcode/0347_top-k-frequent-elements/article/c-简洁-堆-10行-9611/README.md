"+++
title = "0347. Top K Frequent Elements C++ 简洁 堆 10行 96.11% "
author = ["xyzufo"]
date = 2020-10-09T08:45:34+08:00
tags = ["Leetcode", "Algorithms", "cpp", "Heap", "优先队列"]
categories = ["leetcode"]
draft = false
+++

# C++ 简洁 堆 10行 96.11%

> [0347. Top K Frequent Elements](https://leetcode-cn.com/problems/top-k-frequent-elements/)
> [C++ 简洁 堆 10行 96.11%](https://leetcode-cn.com/problems/top-k-frequent-elements/solution/c-jian-ji-dui-10xing-9611-by-xyzufo/) by [xyzufo](https://leetcode-cn.com/u/xyzufo/)

### 解题思路
![Snipaste_2020-10-09_16-43-50.png](https://pic.leetcode-cn.com/1602233041-AUjFRa-Snipaste_2020-10-09_16-43-50.png)

make_pair的用法
无需写出型别， 就可以生成一个pair对象

### 代码

```cpp
class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        unordered_map<int,int> map;         //统计频次
        for(auto i:nums) map[i]++;
        
        priority_queue<pair<int,int>> que;  
        for(auto m:map){
            que.push(make_pair(m.second, m.first));     //优先队列（堆）按照第一项进行排序
        }

        vector<int> vec;
        for(int i=0;i<k;i++){
            vec.push_back(que.top().second);
            que.pop();
        }
        return vec;
    }
};
```