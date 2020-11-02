"+++
title = "0003. Longest Substring Without Repeating Characters ghs "
author = ["9Y3dLqLt7F"]
date = 2020-09-13T02:15:58+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# ghs

> [0003. Longest Substring Without Repeating Characters](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)
> [ghs](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/ghs-by-9y3dlqlt7f/) by [9Y3dLqLt7F](https://leetcode-cn.com/u/9Y3dLqLt7F/)

### 解题思路
此处撰写解题思路

### 代码

```cpp

//最长无重复子串 On 滑动窗口+set
class Solution {
public:
    int lengthOfLongestSubstring(string s) 
    {
        int MAX=0;
        set<char>m;
        int left=0;
        for(int right=0;right<s.size();right++)
        {
                while(m.find(s[right])!=m.end())
                {
                    m.erase(s[left]);
                    left++;
                }

            MAX=max(MAX,right-left+1);
                  m.insert(s[right]);
        }
        return MAX;
    }
};
```