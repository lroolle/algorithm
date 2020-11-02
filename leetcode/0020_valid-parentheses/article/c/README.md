"+++
title = "0020. Valid Parentheses 【C语言实现】简明扼要分情况写即可 "
author = ["cao-mao-plasticine"]
date = 2020-08-06T05:39:15+08:00
tags = ["Leetcode", "Algorithms", "C", "Stack"]
categories = ["leetcode"]
draft = false
+++

# 【C语言实现】简明扼要分情况写即可

> [0020. Valid Parentheses](https://leetcode-cn.com/problems/valid-parentheses/)
> [【C语言实现】简明扼要分情况写即可](https://leetcode-cn.com/problems/valid-parentheses/solution/cyu-yan-shi-xian-jian-ming-e-yao-fen-qing-kuang-xi/) by [cao-mao-plasticine](https://leetcode-cn.com/u/cao-mao-plasticine/)

![QQ截图20200806133830.png](https://pic.leetcode-cn.com/8d3e638fe7387bb09a40e927cc7f6f11e97fb37c14639a45d7c5cf7f13619e74-QQ%E6%88%AA%E5%9B%BE20200806133830.png)

### 解题思路
① 先判断是否是左括号，是左括号则入栈；
② 不是左括号则先判断栈是否为空为空说明输入的第一个字符就是右括号或者输入的是空字符，非法，返回false；
③ 不是左括号且栈非空，进行匹配，如果匹配不成功直接返回false
④ 不是左括号且栈非空，进行匹配，如果匹配成功则出栈，即栈顶指针减1
⑤ 循环结束（即字符串遍历完毕），检查栈是否空，空则所有左括号匹配完毕
ps: 入栈的时候用左括号相应的右括号入栈而不是直接入栈左括号，这样在匹配的时候直接检查栈顶元素与当前遍历
    遍历的字符串是否相等即可完成匹配操作

### 代码

```c
bool isValid(char * s){
    int len = strlen(s);
    char stack[len + 1];                        // 栈
    int top = -1;                               // 栈顶指针

    for (int i = 0; i < len; i++) {
        // 左括号入栈 -- 用相应的右括号入栈 -- 方便匹配
        if (s[i] == '(') stack[++top] = ')';
        else if (s[i] == '[') stack[++top] = ']';
        else if (s[i] == '{') stack[++top] = '}';
        
        // 不是左括号那就只有右括号了，栈空或不匹配则非法
        else if (top == -1 || stack[top] != s[i])
            return false;
        // 栈非空匹配 -- 出栈 -- 即栈顶指针往下移
        else
            top--;
    }

    // 如果遍历完字符串s后栈非空 -- 还有左括号未匹配 -- 非法
    // 如果遍历完字符串s后栈为空 -- 所有左括号匹配完 -- 合法
    return (top == -1);
}
```