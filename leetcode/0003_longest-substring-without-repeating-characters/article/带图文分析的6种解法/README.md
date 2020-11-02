"+++
title = "0003. Longest Substring Without Repeating Characters 带图文分析的6种解法 "
author = ["sdwwld"]
date = 2020-06-05T05:51:41+08:00
tags = ["Leetcode", "Algorithms", "Java", "Map", "TwoPointers"]
categories = ["leetcode"]
draft = false
+++

# 带图文分析的6种解法

> [0003. Longest Substring Without Repeating Characters](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)
> [带图文分析的6种解法](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/javade-6chong-jie-fa-by-sdwwld/) by [sdwwld](https://leetcode-cn.com/u/sdwwld/)

1，我们使用两个指针，一个i一个j，最开始的时候i和j指向第一个元素，然后i往后移，把扫描过的元素都放到map中，如果i扫描过的元素没有重复的就一直往后移，顺便记录一下最大值max，如果i扫描过的元素有重复的，就改变j的位置，我们就以pwwkew为例画个图看一下
![image.png](https://pic.leetcode-cn.com/c759738650c417a23b1e9636f6f40a553d9258f1cc8907d8da3555b5d2f93e60-image.png)
![image.png](https://pic.leetcode-cn.com/bf4c4d4db704e5103f5394d8a790b8235ae160703ffa53b0700650287831b8c9-image.png)



代码如下
```
    public int lengthOfLongestSubstring(String s) {
        if (s.length() == 0)
            return 0;
        HashMap<Character, Integer> map = new HashMap<>();
        int max = 0;
        for (int i = 0, j = 0; i < s.length(); ++i) {
            if (map.containsKey(s.charAt(i))) {
                j = Math.max(j, map.get(s.charAt(i)) + 1);
            }
            map.put(s.charAt(i), i);
            max = Math.max(max, i - j + 1);
        }
        return max;
    }
```

2，我们还可以用一个队列，把元素不停的加入到队列中，如果有相同的元素，就把队首的元素移除，这样我们就可以保证队列中永远都没有重复的元素
![image.png](https://pic.leetcode-cn.com/e0d581509bc8bfa2cecc19a7e65da1629264bfe472ea7f9f83c286ba15b32be9-image.png)
```
    public int lengthOfLongestSubstring(String s) {
        //用链表实现队列，队列是先进先出的
        Queue<Character> queue = new LinkedList<>();
        int res = 0;
        for (char c : s.toCharArray()) {
            while (queue.contains(c)) {
                //如果有重复的，队头出队
                queue.poll();
            }
            //添加到队尾
            queue.add(c);
            res = Math.max(res, queue.size());
        }
        return res;
    }
```

3，由于都是字符，所以我们还可以使用数组来代替map
```
    public int lengthOfLongestSubstring(String s) {
        int[] map = new int[128];
        int start = 0, end = 0, len = 0;
        while (end < s.length()) {
            if (map[s.charAt(end++)]++ == 0) {
                len = Math.max(len, end - start);
            } else {
                while (map[s.charAt(end - 1)] > 1) {
                    map[s.charAt(start++)]--;
                }
            }
        }
        return len;
    }
```
4,我们还可以再优化一下减少一层while循环
```
    public int lengthOfLongestSubstring(String s) {
        int[] dict = new int[128];
        Arrays.fill(dict, -1);
        int longest = 0, m = 0;
        for (int i = 0; i < s.length(); i++) {
            m = Math.max(dict[s.charAt(i)] + 1, m);
            dict[s.charAt(i)] = i;
            longest = Math.max(longest, i - m + 1);
        }
        return longest;
    }
```


5，同样我们还可以使用集合set来代替队列，用两个指针，一个left一个right，如果有重复的就把left指向的给移除（left相当于队首，right相当于队尾）
```
    public int lengthOfLongestSubstring(String s) {
        int maxLen = 0;
        Set<Character> set = new HashSet<>();
        int left = 0, right = 0;
        while (right < s.length()) {
            while (set.contains(s.charAt(right)))
                set.remove(s.charAt(left++));
            set.add(s.charAt(right++));
            maxLen = Math.max(maxLen, right - left);
        }
        return maxLen;
    }
```
6,这里我们还可以改一下，把right-left改为set.size(),顺便再减少一层while循环
```
    public int lengthOfLongestSubstring(String s) {
        int left = 0, right = 0, max = 0;
        Set<Character> set = new HashSet<>();
        while (right < s.length()) {
            if (set.contains(s.charAt(right))) {
                set.remove(s.charAt(left++));
            } else {
                set.add(s.charAt(right++));
                max = Math.max(max, set.size());
            }
        }
        return max;
    }
```


**如果觉得有用就给个赞吧，你的赞是给我最大的鼓励，也是我写作的最大动力**

![image.png](https://pic.leetcode-cn.com/d56a80459005b444404d2ad6fbaabdabecd2b9ed3cb2cf432e570c315ae2fcf7-image.png)
##### 查看更多答案，可扫码关注我微信公众号“**[数据结构和算法](https://img-blog.csdnimg.cn/20190515124616751.jpg)**”