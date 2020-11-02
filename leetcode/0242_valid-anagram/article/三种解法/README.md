"+++
title = "0242. Valid Anagram 三种解法 "
author = ["chenlq"]
date = 2020-09-06T02:49:55+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# 三种解法

> [0242. Valid Anagram](https://leetcode-cn.com/problems/valid-anagram/)
> [三种解法](https://leetcode-cn.com/problems/valid-anagram/solution/san-chong-jie-fa-by-chenlq/) by [chenlq](https://leetcode-cn.com/u/chenlq/)

```
import java.util.Arrays;
import java.util.HashMap;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public boolean isAnagram(String s, String t) {
        /* 解法一：HashMap
        HashMap<Character,Integer> smap = new HashMap<Character, Integer> ();
        HashMap<Character,Integer> tmap = new HashMap<Character,Integer>();
        for (int i = 0; i < s.length(); i++){
            if (smap.containsKey(s.charAt(i))){
                //如果map中存在，则将value+1
                smap.put(s.charAt(i),smap.get(s.charAt(i))+1);
            }else{
                smap.put(s.charAt(i),1);
            }
        }

        for (int i = 0; i < t.length(); i++){
            if (tmap.containsKey(t.charAt(i))){
                //如果map中存在，则将value+1
                tmap.put(t.charAt(i),tmap.get(t.charAt(i))+1);
            }else{
                tmap.put(t.charAt(i),1);
            }
        }

        return smap.equals(tmap);
         */
        /* 方法二：通过数组排序
        //思路：如果是字母异位词，则排序后的字符数组应该完全相同
        if (s.length() != t.length()){
            return false;
        }
        //转换为char数组
        char[] sArr = s.toCharArray();
        char[] tArr = t.toCharArray();
        //对字符数组进行排序
        Arrays.sort(sArr);
        Arrays.sort(tArr);

        //判断字符数组是否相同
        for (int i = 0; i < sArr.length; i++) {
            if (sArr[i] != tArr[i]){
                return false;
            }
        }
        return true;
         */
        /*
        方法三：制作建议哈希表
        //思路：制作建议哈希表，统计每个字母出现的次数，s中出现一次，对应位置的值+1，t中出现一次，对应位置的值-1
        if (s.length() != t.length()){
            return false;
        }

        char[] sArr = s.toCharArray();
        char[] tArr = t.toCharArray();
        int[] map = new int[26];//简易哈希表

        for (int i = 0; i < sArr.length; i++) {
            map[s.charAt(i) - 'a'] += 1;
            map[t.charAt(i) - 'a'] -= 1;
        }

        for (int i = 0; i < map.length; i++) {
            if (map[i] != 0){
                return false;
            }
        }
        return true;
         */

    }
}
```
