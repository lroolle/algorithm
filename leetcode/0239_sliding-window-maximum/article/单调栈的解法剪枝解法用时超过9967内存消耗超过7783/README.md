"+++
title = "0239. Sliding Window Maximum 单调栈的解法+剪枝解法用时超过99.67%，内存消耗超过77.83% "
author = ["ni-shui-qian-fan"]
date = 2020-09-01T13:59:07+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# 单调栈的解法+剪枝解法用时超过99.67%，内存消耗超过77.83%

> [0239. Sliding Window Maximum](https://leetcode-cn.com/problems/sliding-window-maximum/)
> [单调栈的解法+剪枝解法用时超过99.67%，内存消耗超过77.83%](https://leetcode-cn.com/problems/sliding-window-maximum/solution/dan-diao-zhan-de-si-xiang-by-ni-shui-qian-fan/) by [ni-shui-qian-fan](https://leetcode-cn.com/u/ni-shui-qian-fan/)

![截图.PNG](https://pic.leetcode-cn.com/1599046723-POWEDv-%E6%88%AA%E5%9B%BE.PNG)
方法一：暴力解法
不用多想，凡是任何题，首先考虑暴力解法，其他算法都是暴利解法的优化。

方法二：单调栈
由于题目要求在线性时间复杂度内，暴力解法的时间复杂度是O(（n-k + 1）*k），显然不满足要求。

此题需要求窗口中的最大值，滑动窗口对应一个栈，不同的是栈里面保存的是窗口中从最大元素开始单调递减的元素的下标，和单调栈的思想一致，栈底永远是最大的。需要注意的是，解题思想用的是单调栈的思想，但是数据结构需要选用双端队列，因为栈是FILO，不能优先取到栈底元素，恰好Java ArrayDeque满足这种要求，以上思路想通了代码很简单，一个for循环搞定。
```
class Solution {
    public int[] maxSlidingWindow(int[] nums, int k) {
        int len = nums.length;
        int[] outArr = new int[len - k + 1];
        Deque<Integer> deque = new ArrayDeque<>();

        for (int i = 0; i < len; i++) {
            // 当队列 i < k的时候，就添加，当 i>=k的时候，就要考虑移除
            while (!deque.isEmpty() && nums[deque.peekLast()] < nums[i]) {
//引深：合理利用单调栈的性质，如果题目明确窗口中可能出现多个最大元素，则当前代码输出下表最小的那个；加上=，则输出下标最大的那个
                deque.pollLast();
            }
            deque.addLast(i);


            if (i >= k - 1) {
                outArr[i - k + 1] = nums[deque.peekFirst()];
                if (!deque.isEmpty() && deque.peekFirst() == i - k + 1) {
                    deque.pollFirst();
                }
            }
        }

        return outArr;
    }

}
```
方法三：我也不知道叫什么算法，算是剪枝吧（暴力解法的优化，减少重复计算），执行用时超过99.67%，内存消耗超过77.83%。

1、初始化一个index = -1，用来表示当前窗口中最大值的下标
2、确定index是否存在于当前窗口中，如果不存在，则重新求index
3、如果存在，则把当前元素添加到输出数组中
4、通过和将要进入窗口的元素比较，确定下一个窗口的最大值

代码很简单：
```
    public static int[] maxSlidingWindow1(int[] nums, int k) {
        int len = nums.length;
        int[] outArr = new int[len - k + 1];
        int index = -1;

        // 当 i 在 index的最大管辖范围 [index - k + 1, index, index + k - 1] 内，就是index
        for (int i = 0; i < len; i++) {
            // index最大管辖范围 [index - k + 1, index, index + k - 1]，当i = index + k的时候失效
            if (i == index + k) {
                int max = Integer.MIN_VALUE;
                for (int j = i - k + 1; j <= i; j++) {
                    if (nums[j] > max) {
                        max = nums[j];
                        index = j;
                    }
                }
            }

            if (index != -1) {
                outArr[i - k + 1] = nums[index];

                if (i < len - 1 && nums[index] < nums[i + 1]) {
                    index = i + 1;
                }
            }
        }
        return outArr;
    }
```
