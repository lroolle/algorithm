"+++
title = "0047. Permutations II [Java] Backtrack solution using HashMap. "
author = ["lincs"]
date = 2020-09-18T03:16:49+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# [Java] Backtrack solution using HashMap.

> [0047. Permutations II](https://leetcode-cn.com/problems/permutations-ii/)
> [[Java] Backtrack solution using HashMap.](https://leetcode-cn.com/problems/permutations-ii/solution/java-backtrack-solution-using-hashmap-by-lincs/) by [lincs](https://leetcode-cn.com/u/lincs/)

```java
class Solution {
    private int N;
    private List<List<Integer>> premutations;
    public List<List<Integer>> permuteUnique(int[] nums) {
        N = nums.length;
        premutations = new LinkedList<>();
        Map<Integer, Integer> map = new HashMap<>();
        for (int i : nums) {
            map.put(i, map.getOrDefault(i, 0) + 1);
        }
        backtrack(map, new LinkedList<Integer>());
        return premutations;
    }
    private void backtrack(Map<Integer, Integer> map, LinkedList<Integer> premutation) {
        if (premutation.size() == N) {
            premutations.add(List.copyOf(premutation));
            return;
        }
        for (int i : map.keySet()) {
            if (map.get(i) == 0) continue;
            // place
            premutation.add(i);
            map.put(i, map.get(i) - 1);
            // backtrack
            backtrack(map, premutation);
            // recover
            map.put(i, map.get(i) + 1);
            premutation.removeLast();
        }
    }
}
```
