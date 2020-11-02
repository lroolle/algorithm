"+++
title = "0239. Sliding Window Maximum 滑动窗口+单调递增队列 "
author = ["user4358a"]
date = 2020-09-14T06:17:58+08:00
tags = ["Leetcode", "Algorithms", "Queue", "Java", "滑动窗口"]
categories = ["leetcode"]
draft = false
+++

# 滑动窗口+单调递增队列

> [0239. Sliding Window Maximum](https://leetcode-cn.com/problems/sliding-window-maximum/)
> [滑动窗口+单调递增队列](https://leetcode-cn.com/problems/sliding-window-maximum/solution/hua-dong-chuang-kou-dan-diao-di-zeng-dui-lie-by-us/) by [user4358a](https://leetcode-cn.com/u/user4358a/)

### 解题思路
    单调递增队列实现, 入队列的时候，因为队列从大到小排序，数字应该放到合适的位置。
    当前窗口的最大值就是队列的head。

    使用动态数组保存结果，返回的时候使用stream： res.stream().mapToInt(n -> n).toArray();

### 代码

```java
class Solution {
    /**
    单调递增队列实现, 入队列的时候，因为队列从大到小排序，数字应该放到合适的位置。
    当前窗口的最大值就是队列的head。

    使用列表保存结果
    **/
    public int[] maxSlidingWindow(int[] nums, int k) {
        List<Integer> res = new ArrayList<>(64);

        Deque<Integer> queue = new ArrayDeque<>();

        //j窗口左指针, i为窗口右边指针, i 直至一次遍历数组
        int i = 0, j = 0;
        //使用size表示窗口大小，简化逻辑，避免错误。也就是表示窗口有i，j和size
        int size = 0;
        while(i < nums.length) {
            while(size < k) {
                //维护单调递增队列
                while(!queue.isEmpty() && nums[i] > queue.getLast()) {
                    queue.removeLast();
                }
                queue.offer(nums[i]);
                size++;
                i++;
            }

            if (size == k) {
                res.add(queue.getFirst());
            }
            
            //收缩窗口
            if(!queue.isEmpty() && nums[j] == queue.peek()) {
                queue.poll();
            }
            //注意j++，收缩，向右走
            j++;
            size--;
        }
        return res.stream().mapToInt(n -> n).toArray();
    }
}
```