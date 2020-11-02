"+++
title = "0739. Daily Temperatures java的5种解法，最后一种击败100%用户（图文详解） "
author = ["sdwwld"]
date = 2020-06-11T05:21:08+08:00
tags = ["Leetcode", "Algorithms", "Java", "Array", "Stack"]
categories = ["leetcode"]
draft = false
+++

# java的5种解法，最后一种击败100%用户（图文详解）

> [0739. Daily Temperatures](https://leetcode-cn.com/problems/daily-temperatures/)
> [java的5种解法，最后一种击败100%用户（图文详解）](https://leetcode-cn.com/problems/daily-temperatures/solution/java-by-sdwwld/) by [sdwwld](https://leetcode-cn.com/u/sdwwld/)


#### 1,最简单的方式是暴力求解，遍历每一个元素，然后再从当前元素往后找比它大的，找到之后记录下他俩位置的差值，然后停止内层循环，如果没找到默认为0。
![微信截图_20200611104024.png](https://pic.leetcode-cn.com/5123488ad7a1fee175480390025506fcc86524e8f0acf3d8b3a04f4bac0facbb-%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20200611104024.png)
![微信截图_20200611104034.png](https://pic.leetcode-cn.com/4ae7b3bd50fc5c53c07953fb39873032173d27c3d3e4767660402309eff46a53-%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20200611104034.png)


```
    public int[] dailyTemperatures(int[] T) {
        int length = T.length;
        int[] res = new int[length];
        for (int i = 0; i < length; i++) {
            for (int j = i + 1; j < length; j++) {
                if (T[j] > T[i]) {
                    res[i] = j - i;
                    break;
                }
            }
        }
        return res;
    }
```
#### 2，暴力求解毕竟效率不高，我们还可以只用栈来解决，这个栈中存放的是数组元素的下标，我们画个图看下
![微信截图_20200611131851.png](https://pic.leetcode-cn.com/7209ac1b4f5ec690c2bc5e019d88b87cc80e64c0ab3ab559e80bb303388a1482-%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20200611131851.png)
![微信截图_20200611131905.png](https://pic.leetcode-cn.com/4faa99b52a2f03d09a3836a82f5b68e9d1a483b0cd71e5bb22eebddf804d4b9a-%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20200611131905.png)
![微信截图_20200611131914.png](https://pic.leetcode-cn.com/b83a04163ceede3d45038a2d61e86f946d37e6d2c3aa82e8f7c08e5bea71badd-%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20200611131914.png)
![微信截图_20200611131924.png](https://pic.leetcode-cn.com/0aea854709433b259362dff1fc12fc80054881cc49ebc73bac6a8b42a3d908d0-%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20200611131924.png)
![微信截图_20200611131933.png](https://pic.leetcode-cn.com/13360afbcee7e2ab9ec6dd82cab1560fce9f4f2b336682b3017ee95695265306-%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20200611131933.png)
![微信截图_20200611131941.png](https://pic.leetcode-cn.com/1634ccfa6a688d91de575e20cf91583e7e733502461b5c6339b64a1702d709e9-%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20200611131941.png)
![微信截图_20200611131949.png](https://pic.leetcode-cn.com/17fec02033e9de5c0c07716e3c80bb22ce1e285cf95eba9ec33c2af698b2dbf7-%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20200611131949.png)

代码如下
```
    public int[] dailyTemperatures(int[] T) {
        Stack<Integer> stack = new Stack<>();
        int[] ret = new int[T.length];
        for (int i = 0; i < T.length; i++) {
            while (!stack.isEmpty() && T[i] > T[stack.peek()]) {
                int idx = stack.pop();
                ret[idx] = i - idx;
            }
            stack.push(i);
        }
        return ret;
    }
```

#### 3，我们还可以把栈改为数组的形式
```
    public int[] dailyTemperatures(int[] T) {
        int[] stack = new int[T.length];
        int top = -1;
        int[] res = new int[T.length];
        for (int i = 0; i < T.length; i++) {
            while (top >= 0 && T[i] > T[stack[top]]) {
                int idx = stack[top--];
                res[idx] = i - idx;
            }
            stack[++top] = i;
        }
        return res;
    }
```

#### 4，这题我们还可以参照第[84. 柱状图中最大的矩形](https://leetcode-cn.com/problems/largest-rectangle-in-histogram/solution/javade-5chong-jie-fa-xiao-lu-zui-gao-de-ji-bai-lia/)
来看一下代码
```
    public int[] dailyTemperatures(int[] T) {
        int length = T.length;
        Stack<Integer> stack = new Stack<>();
        int[] res = new int[length];
        for (int i = 0; i < length; i++) {
            int h = T[i];
            if (stack.isEmpty() || h <= T[stack.peek()]) {
                stack.push(i);
            } else {
                int top = stack.pop();
                res[top] = i - top;
                i--;
            }
        }
        return res;
    }
```

#### 5，最后一种解法，这种更厉害，从后面开始查找，效率更高，击败了100%的用户，代码中有注释，大家自己看
![微信截图_20200611134941.png](https://pic.leetcode-cn.com/b4cf88fd849b515a3f66786acbf25e72abe94193511ac9a7b9e81ee178a7269c-%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20200611134941.png)

```
    public int[] dailyTemperatures(int[] T) {
        int[] res = new int[T.length];
        //从后面开始查找
        for (int i = res.length - 1; i >= 0; i--) {
            int j = i + 1;
            while (j < res.length) {
                if (T[j] > T[i]) {
                    //如果找到就停止while循环
                    res[i] = j - i;
                    break;
                } else if (res[j] == 0) {
                    //如果没找到，并且res[j]==0。说明第j个元素后面没有
                    //比第j个元素大的值，因为这一步是第i个元素大于第j个元素的值，
                    //那么很明显这后面就更没有大于第i个元素的值。直接终止while循环。
                    break;
                } else {
                    //如果没找到，并且res[j]！=0说明第j个元素后面有比第j个元素大的值，
                    //然后我们让j往后挪res[j]个单位，找到那个值，再和第i个元素比较
                    j += res[j];
                }
            }
        }
        return res;
    }
```

### 如果想查看更多答案，可关注我