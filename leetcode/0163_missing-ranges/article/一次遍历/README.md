"+++
title = "0163. Missing Ranges 一次遍历 "
author = ["kobe24o"]
date = 2020-07-07T09:47:53+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# 一次遍历

> [0163. Missing Ranges](https://leetcode-cn.com/problems/missing-ranges/)
> [一次遍历](https://leetcode-cn.com/problems/missing-ranges/solution/yi-ci-bian-li-by-kobe24o-2/) by [kobe24o](https://leetcode-cn.com/u/kobe24o/)
```cpp
class Solution {
public:
    vector<string> findMissingRanges(vector<int>& nums, int lower, int upper) {
        long l = lower;
    	vector<string> ans;
    	for(int i = 0; i < nums.size(); ++i)
    	{
    		if(l == nums[i])
    			l++;//相等，我跳过你
    		else if(l < nums[i])
    		{	//有空缺
    			if(l < nums[i]-1)//大于1
    				ans.push_back(to_string(l)+"->"+to_string(nums[i]-1));
    			else if(l == nums[i]-1)//等于1
    				ans.push_back(to_string(l));
    			l = long(nums[i])+1;//更新l到nums[i]下一个数
                // [2147483647]
                // 0
                // 2147483647
    		}
    	}
    	if(l < upper)
    		ans.push_back(to_string(l)+"->"+to_string(upper));
    	else if(l==upper)
    		ans.push_back(to_string(l));
    	return ans;
    }
};
```
4 ms	7.2 MB	
---

我的CSDN[博客地址 https://michael.blog.csdn.net/](https://michael.blog.csdn.net/)

长按或扫码关注我的公众号（Michael阿明），一起加油、一起学习进步！
![Michael阿明](https://pic.leetcode-cn.com/e7842feff42eb6f4fc8125f5ca93bd7c5c39dfcd1f5665cc3dc00dda0649f844.png)