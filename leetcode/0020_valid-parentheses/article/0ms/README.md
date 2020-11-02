"+++
title = "0020. Valid Parentheses 0ms，用空间换时间 "
author = ["zumin"]
date = 2020-09-06T02:47:52+08:00
tags = ["Leetcode", "Algorithms", "cpp", "Stack"]
categories = ["leetcode"]
draft = false
+++

# 0ms，用空间换时间

> [0020. Valid Parentheses](https://leetcode-cn.com/problems/valid-parentheses/)
> [0ms，用空间换时间](https://leetcode-cn.com/problems/valid-parentheses/solution/0msyong-kong-jian-huan-shi-jian-by-zumin/) by [zumin](https://leetcode-cn.com/u/zumin/)

### 解题思路
- 利用数组模拟Hash，以省去比较的时间
- 在右括号位置上存放对应的左括号
- 在左括号位置上存放一个标识符

### 代码

```cpp
class Solution {
public:
    bool isValid(string s) {
      stack<char> st;
      short bucket[126];
      //右括号位置上存放对应的左括号
      bucket[')'] = '(';
      bucket[']'] = '[';
      bucket['}'] = '{';
      //左括号位置上存放一个标识符
      bucket['('] = bucket['['] = bucket['{'] = -1;

      for (const char &c : s) {
        //若为左括号，则将其压入栈中
        if (bucket[c] == -1) {
          st.emplace(c);
          //若为右括号且栈顶是对应的左括号，则将左括号弹出
        } else if (!st.empty() && st.top() == bucket[c]) {
          st.pop();
          //否则该括号不合法
        } else {
          return false;
        }
      }
      //若栈中还有左括号，则也不合法
      return st.empty();
    }
};
```