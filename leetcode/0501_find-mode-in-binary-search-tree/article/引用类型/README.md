"+++
title = "0501. Find Mode in Binary Search Tree &引用类型 "
author = ["zouyuhust"]
date = 2020-08-21T09:19:59+08:00
tags = ["Leetcode", "Algorithms", "C++", "Python3"]
categories = ["leetcode"]
draft = false
+++

# &引用类型

> [0501. Find Mode in Binary Search Tree](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/)
> [&引用类型](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/solution/yin-yong-lei-xing-by-zouyuhust/) by [zouyuhust](https://leetcode-cn.com/u/zouyuhust/)

for 个人学习：

注意迭代函数参数中pre加了&而root不用加&。&表示这是一个引用类型参数，而不是值参数。函数中对于引用参数的操作会在内存上修改原值，而函数中对值参数的操作并不会修改原变量的值。

值参数可以理解为，函数为该参数新开辟了一片内存来，然后把传进来的参数的值复制放到这片内存上，而对于值参数的修改是对新开辟的这块内存的值的修改。函数返回值时，若该值参数没有被直接返回或间接返回，则这片内存会在函数结束后当成垃圾回收了。而引用参数是在原变量的内存上操作，函数返回值时，原变量内存是不会被回收的。

在python中，判断一个变量或参数是引用类型还是值类型，是由其类型来决定的。若该变量或参数可哈希（hashable），则是值类型；若该变量或参数不可哈希（unhashable），则是引用类型。一般来说，hashable≈immutable，python对于immutable和hashable的解释，可在一下连接中搜索两个值查看：[https://docs.python.org/zh-cn/3/glossary.html](https://docs.python.org/zh-cn/3/glossary.html)。

如int、float、string等都是可哈希的，而list、dictionary、pandas都是不可哈希的。