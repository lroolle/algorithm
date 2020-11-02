"+++
title = "0047. Permutations II 仅需增加三行代码：从无重复数字到有重复数字的全排列解法 "
author = ["toheng"]
date = 2020-08-02T13:41:06+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# 仅需增加三行代码：从无重复数字到有重复数字的全排列解法

> [0047. Permutations II](https://leetcode-cn.com/problems/permutations-ii/)
> [仅需增加三行代码：从无重复数字到有重复数字的全排列解法](https://leetcode-cn.com/problems/permutations-ii/solution/jin-xu-zeng-jia-san-xing-dai-ma-cong-wu-zhong-fu-s/) by [toheng](https://leetcode-cn.com/u/toheng/)

# 1、无重复数字全排列
核心思想是"交换"，即依次将数组后面的元素交换到数组的start索引处，完成更深层的backtrack后再次执行交换还原数组，并将list中末尾元素弹出，然后执行下一轮循环。
```
public class solution {
    List<List<Integer>> lists = new ArrayList<>();
    public List<List<Integer>> permute(int[] nums) {
        List<Integer> list = new ArrayList<>();
        if (nums == null || nums.length == 0) return lists;
        backtrack(nums, 0, list);
        return lists;
    }
    void backtrack(int[] nums, int start, List<Integer> list){
        if (list.size() == nums.length){
            lists.add(new ArrayList<>(list));
            return;
        }

        for (int i = start; i < nums.length; i++){
            swap(nums, start, i);
            list.add(nums[start]);
            backtrack(nums, start+1, list);
            swap(nums, start, i);
            list.remove(list.size()-1);
        }
    }
    void swap(int[] nums, int left, int right){
        int temp = nums[left];
        nums[left] = nums[right];
        nums[right] = temp;
    }
```

# 2、有重复数字全排列
在无重复数字全排列的基础上，backtrack方法中增加了一个set来去重。具体来说如果set中存在了当前要交换的数字，说明当前这个数字已经放到过索引start的位置，如果将当前数字再次交换到start位置，则会产生重复的排列，因此跳过这一轮循环进入下一轮。
```
public class solution {
    List<List<Integer>> lists = new ArrayList<>();
    public List<List<Integer>> permuteUnique(int[] nums) {
        if (nums == null || nums.length == 0) return lists;
        List<Integer> list = new ArrayList<>();
        backtrack(nums, 0, list);
        return lists;
    }

    void backtrack(int[] nums, int start, List<Integer> list){
        if (list.size() == nums.length){
            lists.add(new ArrayList<>(list));
            return;
        }

        HashSet<Integer> set = new HashSet<>();
        for (int i = start; i < nums.length; i++){
            
            // 与无重复数字全排列的唯一不同之处
            if (set.contains(nums[i])){
                continue;
            }
            set.add(nums[i]);

            swap(nums, i, start);
            list.add(nums[start]);
            backtrack(nums, start + 1, list);
            list.remove(list.size() - 1);
            swap(nums, i, start);
        }
    }

    void swap(int[] digit, int left, int right){
        int tmp = digit[left];
        digit[left] = digit[right];
        digit[right] = tmp;
    }
}
```
