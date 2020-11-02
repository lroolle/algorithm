"+++
title = "0020. Valid Parentheses  20. 有效的括号：想清楚三种情况不匹配，代码就很好写了（详细讲解 & C++极简代码） "
author = ["carlsun-2"]
date = 2020-08-09T08:07:59+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

#  20. 有效的括号：想清楚三种情况不匹配，代码就很好写了（详细讲解 & C++极简代码）

> [0020. Valid Parentheses](https://leetcode-cn.com/problems/valid-parentheses/)
> [ 20. 有效的括号：想清楚三种情况不匹配，代码就很好写了（详细讲解 & C++极简代码）](https://leetcode-cn.com/problems/valid-parentheses/solution/20-you-xiao-de-gua-hao-xiang-qing-chu-san-chong-qi/) by [carlsun-2](https://leetcode-cn.com/u/carlsun-2/)

## 思路 

### 题外话
括号匹配是使用栈解决的经典问题。

题意其实就像我们在写代码的过程中，要求括号的顺序是一样的，有左括号，响应的位置必须要有右括号，如果还记得编译原理的话，编译器在 词法分析的过程中处理 括号，花括号等这个符号的逻辑，也是使用了栈这种数据结构。

所以栈在计算机领域中应用是非常广泛的。 有的同学可以经常会想学的这些数据结构有什么用，也开发不了什么软件，大多数同学说的软件应该都是可视化的软件例如APP之类的，那都是非常上层的应用了，底层很多功能的实现都是基础的数据结构和算法。 这里我就不过多展开了，我们先来看题。

### 进入正题

由于栈结构的特殊性，非常适合做对称匹配类的题目。首先我们要弄清楚，字符串里的括号不匹配有几种情况。

一些同学，在面试中看到这种题目上来就开始写代码，建议要写代码之前要分析好有哪几种不匹配的情况，如果不动手之前分析好，写出的代码也会有很多问题。 会给面试官留下不好的印象。

我们先来分析一下 这里有三种不匹配的情况，

1. 第一种情况，字符串里左方向的括号多余了 ，所以不匹配。
![image.png](https://pic.leetcode-cn.com/a8f6eab320dbcacaded82af7a0b14eda752fa38dc6e28366d7fb319627e86c67-image.png)
2. 第二种情况，括号没有多余，但是 括号的类型没有匹配上。

![image.png](https://pic.leetcode-cn.com/b98abe29244f1db4c3e45b39e7dabd9a2bfb7193826e9db618e9fbb032780e7a-image.png)
3. 第三种情况，字符串里右方向的括号多余了，所以不匹配。

![image.png](https://pic.leetcode-cn.com/ef62ae33b112b52cd984c2eb6a4f73a51ad815a9a71b7e9aab7473611ba15300-image.png)

我们的代码只要覆盖了这三种不匹配的情况，基本就不会出问题，可以看出 动手之前分析好题目的重要性。

动画如下：
![20.有效括号.mp4](f746c057-3314-48ff-98a9-80a2d93b5060)
第一种情况：已经遍历完了字符串，但是栈不为空，说明有相应的左括号没有右括号来匹配，所以return false

第二种情况：遍历字符串匹配的过程中，发现栈里没有要匹配的字符。所以return false

第三种情况：遍历字符串匹配的过程中，栈已经为空了，没有匹配的字符了，说明右括号没有找到对应的左括号return false

那么什么时候说明左括号和右括号全都匹配了呢，就是字符串遍历完之后，栈也是空的，就说明全都匹配了。

接下来我们来看一下代码。

## C++代码
```
class Solution {
public:
    bool isValid(string s) {
        stack<int> st;
        for (int i = 0; i < s.size(); i++) {
            if (s[i] == '(') st.push(')');
            else if (s[i] == '{') st.push('}');
            else if (s[i] == '[') st.push(']');
            // 第三种情况 是遍历字符串匹配的过程中，栈已经为空了，没有匹配的字符了，说明右括号没有找到对应的左括号 return false
            // 第二种情况 遍历字符串匹配的过程中，发现栈里没有我们要匹配的字符。所以return false
            else if (st.empty() || st.top() != s[i]) return false;
            else st.pop(); // st.top() == s[i]
        }
        // 第一种情况 此时我们已经遍历完了字符串，但是栈不为空，说明有相应的左括号没有右括号来匹配，所以return false，否则就return true
        return st.empty();
    }
};
```

**我是程序员Carl，先后在腾讯和百度从事技术研发多年，利用工作之余重刷leetcode，更多精彩文章尽在公众号：[代码随想录](https://img-blog.csdnimg.cn/20200815195519696.png)，关注后，回复「Java」「C++」「python」「简历模板」等等，有我整理多年的学习资料，可以加我[微信](https://img-blog.csdnimg.cn/20200814140330894.png)，备注「组队刷题」，拉你进入刷题群（无任何广告，纯个人分享），每天一道经典题目分析，我选的每一道题目都不是孤立的，而是由浅入深一脉相承的，如果跟住节奏每篇连续着看，定会融会贯通。本文  [https://github.com/youngyangyang04/leetcode-master](https://github.com/youngyangyang04/leetcode-master ) 已经收录，里面还有leetcode刷题攻略、各个类型经典题目刷题顺序、思维导图，看一看一定会有所收获！**
**以下资料希望对你有帮助：**
* [手把手带你读C++ primer！B站](https://www.bilibili.com/video/BV1Z5411874t)
* [C++ primer 第一章，你要知道的知识点还有这些！B站](https://www.bilibili.com/video/BV1Kv41117Ya)
* [C++面试&C++学习指南知识点整理](https://github.com/youngyangyang04/TechCPP)
* [程序员应该如何写简历（附简历模板）](https://mp.weixin.qq.com/s/PkBpde0PV65dJjj9zZJYtg)
* [leetcode刷题攻略](https://github.com/youngyangyang04/leetcode-master)

**如果感觉题解对你有帮助，不要吝啬给一个👍吧！**

