"+++
title = "0047. Permutations II 47 全排列 Ⅱ (较为详细易懂的说法) "
author = ["feng-zui-zhu-xin-sui-i"]
date = 2020-09-09T14:33:43+08:00
tags = ["Leetcode", "Algorithms", "Java", "Backtracking"]
categories = ["leetcode"]
draft = false
+++

# 47 全排列 Ⅱ (较为详细易懂的说法)

> [0047. Permutations II](https://leetcode-cn.com/problems/permutations-ii/)
> [47 全排列 Ⅱ (较为详细易懂的说法)](https://leetcode-cn.com/problems/permutations-ii/solution/47-quan-pai-lie-ii-jiao-wei-xiang-xi-yi-dong-de-sh/) by [feng-zui-zhu-xin-sui-i](https://leetcode-cn.com/u/feng-zui-zhu-xin-sui-i/)

这种全排列的问题，可以用回溯来解决，
46比较简单，看注释应该能明白
47在46的基础上难一点，由于出现了重复的变量，需先排序(让同一层相同变量挨在一起)再同层去重
建议自己画图理解
![微信图片_20200909223118.jpg](https://pic.leetcode-cn.com/1599661889-pqdLkW-%E5%BE%AE%E4%BF%A1%E5%9B%BE%E7%89%87_20200909223118.jpg)

# **1.无重复变量的全排列** 46题


class Solution {

    private List<List<Integer>> result = new ArrayList<>();
    public List<List<Integer>> permute(int[] nums) {
        int[] visited = new int [nums.length];  //标记已有,默认为0
        List<Integer> path = new ArrayList<>();

        fun(path,visited,nums);

        return result;
    }

    private void fun(List<Integer> path,int [] visited,int []  nums)
    {
        if(path.size()==nums.length)
        {
            result.add(new ArrayList<>(path));
            return;
        }

        for(int i = 0;i<nums.length;i++ )
        {
            if(visited[i]==1)continue;  //判断是否用过
            path.add(nums[i]);           //添加path并是标记visited标志置为1
            visited[i]=1;
            fun(path,visited,nums);     
                                        //回溯
            visited[i]=0;
            path.remove(path.size()-1);
        }
    }
}


# **2.有重复变量的全排列**  47题

class Solution {

    
    private List<List<Integer>> result = new ArrayList<>();
    public List<List<Integer>> permuteUnique(int[] nums) {
        
    
        int[] visited = new int [nums.length];  //标记已有,默认为0
        List<Integer> path = new ArrayList<>();

        //和46差不多，需先排序，然后同层去重
        Arrays.sort(nums);
        fun(path,visited,nums);

        return result;
    }

    private void fun(List<Integer> path,int [] visited,int []  nums)
    {
        if(path.size()==nums.length)
        {
            result.add(new ArrayList<>(path));
            return;
        }

        for(int i = 0;i<nums.length;i++ )
        {

            if( i>0 && nums[i]==nums[i-1])  //同层去重
            {
                if(visited[i-1]==1);        //判断上次是否用过，上次没用过，则这次这个可以用
                else continue;
            }

            if(visited[i]==1)continue;     //判断这次是否用过
            path.add(nums[i]);           //添加path并是标记visited标志置为1
            visited[i]=1;
            fun(path,visited,nums);     
                                        //回溯
            visited[i]=0;
            path.remove(path.size()-1);
        }
    
    }
}