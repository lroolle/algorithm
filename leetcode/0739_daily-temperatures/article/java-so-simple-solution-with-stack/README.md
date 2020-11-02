"+++
title = "0739. Daily Temperatures [Java] So simple solution with stack "
author = ["lincs"]
date = 2020-08-28T08:53:28+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# [Java] So simple solution with stack

> [0739. Daily Temperatures](https://leetcode-cn.com/problems/daily-temperatures/)
> [[Java] So simple solution with stack](https://leetcode-cn.com/problems/daily-temperatures/solution/java-so-simple-solution-with-stack-by-lincs/) by [lincs](https://leetcode-cn.com/u/lincs/)

```java
class Solution {
    public int[] dailyTemperatures(int[] T) {
        int n = T.length;
        int[] ans = new int[n];
        Stack<Integer> stack = new Stack<>();
        for (int i = 0; i < n; i++) {
            while (!stack.isEmpty()) {
                int j = stack.peek();
                if (T[j] < T[i]) {
                    ans[j] = i - j;
                    stack.pop();
                }
                else {
                    break;
                }
            }
            stack.push(i);
        }
        return ans;
    }
}
```
