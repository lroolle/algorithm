"+++
title = "0047. Permutations II 【Java】DFS + 剪枝 解决 “全排列”问题 "
author = ["leetcoder-youzg"]
date = 2020-09-18T01:08:46+08:00
tags = ["Leetcode", "Algorithms", "Java", "DepthfirstSearch", "搜索"]
categories = ["leetcode"]
draft = false
+++

# 【Java】DFS + 剪枝 解决 “全排列”问题

> [0047. Permutations II](https://leetcode-cn.com/problems/permutations-ii/)
> [【Java】DFS + 剪枝 解决 “全排列”问题](https://leetcode-cn.com/problems/permutations-ii/solution/java-dfs-jian-zhi-jie-jue-quan-pai-lie-wen-ti-by-l/) by [leetcoder-youzg](https://leetcode-cn.com/u/leetcoder-youzg/)

### 解题思路
本题解 **`剪枝`**处 借鉴了 题解区大佬的思想：
> 1、当前元素 被使用过
> 2、当前元素 和 前一个元素相同(排序保证相同元素排在一起)，
>    且 前一个元素 没有被使用过(防止 重复排列)

### 运行结果
![image.png](https://pic.leetcode-cn.com/1600391138-PJQFGf-image.png)

### 代码

```java
class Solution {
    public List<List<Integer>> permuteUnique(int[] nums) {
        List<List<Integer>> results = new ArrayList<List<Integer>>();
        if (nums == null || nums.length <= 0) {
            return results;
        }

        Arrays.sort(nums);  // 剪枝 的必要条件
        int length = nums.length;
        Deque<Integer> curPath = new ArrayDeque<>();
        boolean[] used = new boolean[length];
        dfs(nums, 0, curPath, length, used, results);
        return results;
    }

    private void dfs(int[] nums, int index, Deque<Integer> curPath, int length, boolean[] used, List<List<Integer>> results) {
        if (index >= length) {  // 递归 结束标志
            results.add(new ArrayList<>(curPath));
            return;
        }

        for (int i = 0; i < length; i++) {
            /*
                剪枝:
                    1、当前元素 被使用过
                    2、当前元素 和 前一个元素相同(排序保证相同元素排在一起)，
                        且 前一个元素 没有被使用过(保证重复元素只按照一个顺序被遍历，防止 重复排列)
             */
            if (used[i] || (i > 0 && nums[i] == nums[i - 1] && !used[i - 1])) {
                continue;
            }

            /*
                递归
             */
            curPath.addLast(nums[i]);
            used[i] = true;
            dfs(nums, index + 1, curPath, length, used, results);

            /*
                复位
             */
            used[i] = false;
            curPath.removeLast();
        }
    }
}
```
这种题做多了，在再次遇到时，心中就会有一个模板，大致都是如上格式
只有剪枝处会有些许差别
打卡第59天，加油！！！