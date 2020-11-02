"+++
title = "面试题 08.06. Hanota LCCI JAVA 简洁代码， 思路 "
author = ["acw_weian"]
date = 2020-06-20T02:41:33+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# JAVA 简洁代码， 思路

> [面试题 08.06. Hanota LCCI](https://leetcode-cn.com/problems/hanota-lcci/)
> [JAVA 简洁代码， 思路](https://leetcode-cn.com/problems/hanota-lcci/solution/java-jian-ji-dai-ma-si-lu-by-acw_weian/) by [acw_weian](https://leetcode-cn.com/u/acw_weian/)

    /**
    分治思想： 拿只有A只有两个盘子的时候做模拟，发现步骤如下:
    第一步， 把A上面的盘子放到B
    第二步， 把A上最后一个盘子放到C
    第三步， 把B上的盘子移动到C

    推广到A上有N个盘子：
    第一步， 把A上的N - 1个盘子放到B
    第二步， 把A上的最后一个盘子放到C
    第三步， 把B上的N - 1个盘子放到C

    边界情况： 当N == 1时候， 把A移动到C
    */
```
    public void hanota(List<Integer> A, List<Integer> B, List<Integer> C) {
        hanoi(A.size(), A, B, C);
    }

    public void hanoi(int n, List<Integer> A, List<Integer> B, List<Integer> C){
        
        if(n == 1){
            C.add(A.get(A.size() - 1));
            A.remove(A.size() - 1);
        }else{
            //把A经过辅助C放到B上
            hanoi(n - 1, A, C, B);
            //把A放到C上
            C.add(A.get(A.size() - 1));
            A.remove(A.size() - 1);
            //把B经过辅助A放到C上
            hanoi(n - 1, B, A, C);
        }
    }
```
