"+++
title = "0347. Top K Frequent Elements 快排模板随便改下就好啦 "
author = ["tian-xin-yue-yuan-3"]
date = 2020-09-07T03:04:45+08:00
tags = ["Leetcode", "Algorithms", "cpp", "快速排序", "Recursion", "Sort", "TwoPointers"]
categories = ["leetcode"]
draft = false
+++

# 快排模板随便改下就好啦

> [0347. Top K Frequent Elements](https://leetcode-cn.com/problems/top-k-frequent-elements/)
> [快排模板随便改下就好啦](https://leetcode-cn.com/problems/top-k-frequent-elements/solution/kuai-pai-mo-ban-sui-bian-gai-xia-jiu-hao-la-by-tia/) by [tian-xin-yue-yuan-3](https://leetcode-cn.com/u/tian-xin-yue-yuan-3/)

从快排模板修改改。。。
快排模板。
```
void quick_sort(int q[], int l, int r) {  
  
    if (l >= r)  
        return;  
    int i = l - 1, j = r + 1, x = q[l + r >> 1];  
    while (i < j) {  
        do ++i; while (q[i] < x);  
        do --j; while (q[j] > x);  
        if (i < j)  
            swap(q[i], q[j]);  
    }  
    quick_sort(q, l, j);  
    quick_sort(q, j+1, r);  
}  

```
简单修改一下
```
using PII = pair<int,int>;
    vector<PII> arra;
    void quick_sort(int l , int r,vector<int>& ret,int  k){
        if( l > r)
            return ;
        if(l == r && k == 1){
            //l == r时且k ==1为原l>=r 中等于部分的递归基
            ret.push_back(arra[l].first);
            return ;
        }
        int i = l - 1 , j = r +1,mid = (l + r ) >> 1;
        PII x  = arra[mid];
        while( i < j){
            do ++i; while(arra[i].second > x.second);
            do --j; while(arra[j].second < x.second);
            if(i < j)
                swap(arra[i],arra[j]);
        }
        //左边数量够k个了
        if( j - l + 1 >= k)
            quick_sort(l,j,ret,k);
        else{
            //左边不够，把所有的左边的算上
            for(int ind = l; ind <=j;++ind){
                ret.push_back(arra[ind].first);
            }  
            //左边有 j - l +1 个，但是不够，从右边拿剩下的
            quick_sort(j + 1, r, ret, k - (j - l + 1));
        }

    }
    vector<int> topKFrequent(vector<int>& nums, int k) {
        unordered_map<int,int> mp;
        for(auto i :nums)
            mp[i]++;
        for(auto pii : mp)
            arra.push_back(pii);
        vector<int> ret;
        quick_sort(0,arra.size()-1,ret,k);
        return ret;
    }
```


