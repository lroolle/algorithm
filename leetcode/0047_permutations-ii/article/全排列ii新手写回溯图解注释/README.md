"+++
title = "0047. Permutations II 全排列II，新手写回溯，图解+注释 "
author = ["zui-weng-jiu-xian"]
date = 2020-09-18T04:08:26+08:00
tags = ["Leetcode", "Algorithms", "Java", "DepthfirstSearch", "Backtracking"]
categories = ["leetcode"]
draft = false
+++

# 全排列II，新手写回溯，图解+注释

> [0047. Permutations II](https://leetcode-cn.com/problems/permutations-ii/)
> [全排列II，新手写回溯，图解+注释](https://leetcode-cn.com/problems/permutations-ii/solution/quan-pai-lie-iixin-shou-xie-hui-su-tu-jie-zhu-shi-/) by [zui-weng-jiu-xian](https://leetcode-cn.com/u/zui-weng-jiu-xian/)

## 思路
### 中心思想：要找出所有可能的组合，回溯法应该是可行的。

### 回溯法四要素
![image.png](https://pic.leetcode-cn.com/1600399890-NLmVTc-image.png)
决策树共`nums.length`层，图中是`3`层。
<br>
1. 明确选择：在决策树的当前层，有多少种选择？把所有选择统称为选择列表，选择受约束条件制约。
 
2. 明确路径，在决策树的第`i`层，前面走过的所有层选择的数构成一条路径。

3. 明确结束条件：当选了`nums.length`个数后，说明找到了一个全排列，加入答案中。

4. 明确约束条件：在当前层，我们不能选哪些数？
    - 被选过的不能选：因为有**重复数字**，所以是否被选过的不能用数值判断，**要用它的下标判断**！
    
    - 解释图中标号`1`和标号`2`：第二棵决策树从开始就不应该出现，为什么？因为他的开头和前一棵决策树是一样的！这说明，**以这个数为开头的所有全排列，前一棵树已经都找到了**，你这一棵没必要再找一次，砍了吧！同理那个分支也砍了吧！**即同一层不能开头相同**！
## 代码

```java
class Solution {
    List<List<Integer>> permutations = new LinkedList<>();  // 结果
    List<Integer> path = new LinkedList<>();                // 路径
    boolean[] mark;     // 标记是否用过
    int length;         // 数组长度
    public List<List<Integer>> permuteUnique(int[] nums) {
        length = nums.length;
        if (length == 0) { return permutations; }   // 空数组
        Arrays.sort(nums);                          // 排序，可有可无，原因见后文注释
        mark = new boolean[length];                 
        backtrack(nums);
        return permutations;
    }

    public void backtrack(int[] nums) {
        // 结束条件
        if (path.size() == length) {
            permutations.add(new LinkedList(path));
            return;
        }
        /* 
            用set去重，无需排序
            牺牲空间换时间
        */
        //Set<Integer> set = new HashSet<>();
        for (int i = 0; i < length; i++) {
            // 被用过了
            if (mark[i]) {
                continue;
            }
            /*
            // 如果这个开头用过
            if (set.contains(nums[i])) {
                continue;
            }
            set.add(nums[i]);
            */
            // 需要排序，相同的开头是相邻的
            if (i > 0 && nums[i] == nums[i - 1] && mark[i - 1]) {
                continue;
            }

            mark[i] = true;         // 标记被用过
            path.add(nums[i]);      // 加入路径
            backtrack(nums);        // 深度搜索路径
            path.remove(path.size() - 1);   // 从路径中去掉刚才加的
            mark[i] = false;                // 恢复没用过
        }
    }
}
```
## 算法分析
设`n`是数组元素个数。
你看啊，这个，这个，如果没有重复数字，就老老实实要列出`n!`个全排列，好吧我不会，大佬教我。
不过我觉得至少有一个$O(n!)$在里面。
### 如果本文对你有帮助，可以给个大拇指呀！
### 如果你有什么建议或疑问，可以评论留言呀！