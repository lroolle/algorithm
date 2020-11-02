"+++
title = "0347. Top K Frequent Elements 比较直白的思路和简单的代码，借此复习了下qsort的用法 "
author = ["shaw-r"]
date = 2020-09-07T03:47:38+08:00
tags = ["Leetcode", "Algorithms", "C", "Sort"]
categories = ["leetcode"]
draft = false
+++

# 比较直白的思路和简单的代码，借此复习了下qsort的用法

> [0347. Top K Frequent Elements](https://leetcode-cn.com/problems/top-k-frequent-elements/)
> [比较直白的思路和简单的代码，借此复习了下qsort的用法](https://leetcode-cn.com/problems/top-k-frequent-elements/solution/bi-jiao-zhi-bai-de-si-lu-he-jian-dan-de-dai-ma-jie/) by [shaw-r](https://leetcode-cn.com/u/shaw-r/)

### 解题思路
1.将nums排序
2.一趟遍历计算各个数字出现的次数，用结构体类型记录
3.将结构体数组排序
4.返回前k个数字

### 代码

```c
//定义结构体，记录数字及其出现的次数
struct Times{
    int num;
    int cnt;
};
//用于初始数组nums的排序
int cmp(const void* a,const void* b){
    return *(int*)a - *(int*)b;
}
//用于结构体的一级排序
int cmps(const void* a,const void* b){
    return (*(struct Times*)b).cnt - (*(struct Times*)a).cnt;
}

int* topKFrequent(int* nums, int numsSize, int k, int* returnSize){
    *returnSize = k;
    int* res = (int*)malloc(sizeof(int)*k);
    struct Times time[numsSize];
    int index=0;
    //1.对初始数组nums进行排序
    qsort(nums,numsSize,sizeof(int),cmp);
    //2.遍历计算各个数字出现的次数
    time[index].num = nums[0];
    time[index].cnt = 1;
    for(int i=1;i<numsSize;i++){
        if(nums[i]==nums[i-1])time[index].cnt++;
        else{
            time[++index].num = nums[i];
            time[index].cnt = 1;
        }
    }
    //3.对结构体数组按出现次数降序排序
    qsort(time,index+1,sizeof(struct Times),cmps);
    //4.返回前k个数字
    for(int i=0;i<k;i++){
        res[i] = time[i].num;
    }

    return res;
}
```