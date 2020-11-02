"+++
title = "0003. Longest Substring Without Repeating Characters 最长不重复子串解题分享 "
author = ["orange_go"]
date = 2020-09-11T04:01:43+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# 最长不重复子串解题分享

> [0003. Longest Substring Without Repeating Characters](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)
> [最长不重复子串解题分享](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/zui-chang-bu-zhong-fu-zi-chuan-jie-ti-fen-xiang-by/) by [orange_go](https://leetcode-cn.com/u/orange_go/)

### 解题思路
先归纳有字符串出现的情况
 * （1）整个子串不重复：abcdef
 *  (2) 两个子串衔接:abcaabcda
 *  (3) 两个子串交差：abcaefgb
 *  (4) 两段型：asdderty
 *  该题的问题的关键在于：如果找出每一个子串。我们可以考虑记录一个子串，当新加加入的元素和子串中的元素有重复时，当前子串不重复元素个数达到大；
 *  同时下一个子串的开始位置为子串的重复元素的后一个元素（不是新加入的元素，需要考虑情况3）
 * 
 * 1、记录游标start，要求：从start开始到当前位置i之内没有重复元素，目的：用于记录从当前i位置到最近的不重复子段的启始点；
 *    记录size,用于表示当前子串的长度；记录start用于表示当前子串的启始位置。
 * 2、遍历s中的每一个元素：
 *     检查从star开始到i-1元素中是否有元素与i元素一致。如果有则返回重复元素下标index，否则返回-1
 *     如果存在重复元素：
 *         start加一为作为下一个子串的起始点
 *         记录当前最大长度：历史最大值以及当前子串长度中最大中
 *         下一个字段的长度size=当前位置-start
 *     如果不存在重复元素
 *         size++
 *
 *  3、如果start==0,说明是情况1，返回字符串长度
 *  4、返回记录的最大长度和size最大值

### 代码

```java
class Solution {
    public int lengthOfLongestSubstring(String s) {
        if (s == null || s.equals("")) {
            return 0;
        }
        int start = 0;
        int index = 0;
        int size = 0;
        int maxSize = 0;
        int tempSize = 0;
        for (int i = 0; i < s.length(); i++) {
            index = find(start, i, s, s.charAt(i));
            if (index != -1) {
                size=i - index;
                start=index+1;
                maxSize = Math.max(maxSize, size);
            }else{
                size++;
            }
        }
        if(start==0){
            return s.length();
        }
        return Math.max(maxSize, size);

    }

    private int find(int start, int end, String sub, char c) {
        for (int i = start; i < end; i++) {
            if (sub.charAt(i) == c) {
                return i;
            }
        }
        return -1;
    }
}
```