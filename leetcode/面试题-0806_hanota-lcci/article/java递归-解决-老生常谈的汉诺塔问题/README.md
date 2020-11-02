"+++
title = "面试题 08.06. Hanota LCCI 【Java】递归 解决 老生常谈的“汉诺塔问题” "
author = ["leetcoder-youzg"]
date = 2020-08-23T01:44:54+08:00
tags = ["Leetcode", "Algorithms", "Java", "Recursion", "Stack"]
categories = ["leetcode"]
draft = false
+++

# 【Java】递归 解决 老生常谈的“汉诺塔问题”

> [面试题 08.06. Hanota LCCI](https://leetcode-cn.com/problems/hanota-lcci/)
> [【Java】递归 解决 老生常谈的“汉诺塔问题”](https://leetcode-cn.com/problems/hanota-lcci/solution/java-di-gui-jie-jue-lao-sheng-chang-tan-de-yi-nuo-/) by [leetcoder-youzg](https://leetcode-cn.com/u/leetcoder-youzg/)

### 解题思路
口头叙述有些麻烦，直接看代码，相信同学们一定会 豁然开朗

### 运行结果
![image.png](https://pic.leetcode-cn.com/1598146943-jjJfzb-image.png)

### 代码

```java
class Solution {
    /**
     * 将 A 上的所有盘子，借助 B，移动到C 上
     * @param A 原柱子
     * @param B 辅助柱子
     * @param C 目标柱子
     */
    public void hanota(List<Integer> A, List<Integer> B, List<Integer> C) {
        movePlate(A.size(), A, B, C);
    }

    private void movePlate(int num, List<Integer> original, List<Integer> auxiliary, List<Integer> target) {
        if (num == 1) {    // 只剩一个盘子时，直接移动即可
            target.add(original.remove(original.size() - 1));
            return;
        }

        movePlate(num - 1, original, target, auxiliary);   // 将 size-1 个盘子，从 original 移动到 auxiliary
        target.add(original.remove(original.size() - 1));   // 将 第size个盘子，从 original 移动到 target
        movePlate(num - 1, auxiliary, original, target);   // 将 size-1 个盘子，从 auxiliary 移动到 target
    }
}
```
打卡第36天，加油！！！