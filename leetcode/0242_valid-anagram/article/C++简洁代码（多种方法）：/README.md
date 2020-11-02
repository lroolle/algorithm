"+++
title = "0242. Valid Anagram C++简洁代码（多种方法）： "
author = ["OrangeMan"]
date = 2020-09-02T05:47:05+08:00
tags = ["Leetcode", "Algorithms", "cpp", "Sort", "Map"]
categories = ["leetcode"]
draft = false
+++

# C++简洁代码（多种方法）：

> [0242. Valid Anagram](https://leetcode-cn.com/problems/valid-anagram/)
> [C++简洁代码（多种方法）：](https://leetcode-cn.com/problems/valid-anagram/solution/cjian-ji-dai-ma-duo-chong-fang-fa-by-orangeman-10/) by [OrangeMan](https://leetcode-cn.com/u/orangeman/)

### 解题思路
解题思路：
方法一：`sort`
方法二：`map1`
方法三：`map2`
方法四：数组1(**推荐**)
方法五：数组2
### 代码

```sort
class Solution {
public:
    bool isAnagram(string s, string t) {
        sort(s.begin(), s.end());
        sort(t.begin(), t.end());
        return s == t;
    }
};
```
```map1
class Solution {
public:
    bool isAnagram(string s, string t) {
        if(s.size() != t.size()) return false;
        unordered_map<char, int> map;
        for(char c : s) map[c] ++;
        for(char c : t) 
            if(-- map[c] == -1) return false;
        return true;
    }
};
```
```map2
class Solution {
public:
    bool isAnagram(string s, string t) {
        unordered_map<char, int> map;
        for(int i = 0; i < max(s.size(), t.size()); i ++) 
            map[s[i]] ++, map[t[i]] --;
        for(auto it : map) 
            if(it.second != 0) return false;
        return true;
    }
};
```
```数组1
class Solution {
public:
    bool isAnagram(string s, string t) {
        if(s.size() != t.size()) return false;
        int hash[26] = {0};
        for(char c : s) hash[c - 'a'] ++;
        for(char c : t) 
            if(-- hash[c - 'a'] == -1) return false;
        return true;

    }
};
```
```数组2
class Solution {
public:
    bool isAnagram(string s, string t) {
        if(s.size() != t.size()) return false;
        int hash[26] = {0};
        for(int i = 0; i < s.size(); i ++) 
            hash[s[i] - 'a'] ++, hash[t[i] - 'a'] --;
        for(int i = 0; i < 26; i ++) 
            if(hash[i] != 0) return false;
        return true;

    }
};
```
