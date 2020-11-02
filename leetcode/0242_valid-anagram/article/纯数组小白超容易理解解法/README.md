"+++
title = "0242. Valid Anagram 纯数组小白超容易理解解法 "
author = ["duo-teng-shi-wo"]
date = 2020-08-21T16:22:33+08:00
tags = ["Leetcode", "Algorithms", "C++", "Array"]
categories = ["leetcode"]
draft = false
+++

# 纯数组小白超容易理解解法

> [0242. Valid Anagram](https://leetcode-cn.com/problems/valid-anagram/)
> [纯数组小白超容易理解解法](https://leetcode-cn.com/problems/valid-anagram/solution/chun-shu-zu-xiao-bai-chao-rong-yi-li-jie-jie-fa-by/) by [duo-teng-shi-wo](https://leetcode-cn.com/u/duo-teng-shi-wo/)

### 解题思路
纯数组解法,与哈希的方法其实原理是一样的,用一个26位的数组储存每一个字符串a-z 26个字母出现的次数(题目应该默认小写了,其他的转换一下也一样),然后再比较两个数组是否一致即可

### 代码
```cpp
class Solution {
public:
    bool isAnagram(string s, string t) {
        if(s.size() != t.size()) return false;
        int hash_s[26]={0};
        int hash_t[26]={0};
        for(auto i:s)
        {
            ++hash_s[i-'a'];
        }
        for(auto j:t)
        {
            ++hash_t[j-'a'];
        }
        for(int k=0;k<26;++k)
        {
            if(hash_s[k] != hash_t[k]) return false;
        }
        return true;
    }
};
```
![7354e835c8eac6597608f9159f80ac5.png](https://pic.leetcode-cn.com/1598026924-RNQdrJ-7354e835c8eac6597608f9159f80ac5.png)
