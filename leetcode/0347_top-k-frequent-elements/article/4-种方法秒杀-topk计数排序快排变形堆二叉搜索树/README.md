"+++
title = "0347. Top K Frequent Elements 4 种方法秒杀 TopK（计数排序/快排变形/堆/二叉搜索树） "
author = ["sweetiee"]
date = 2020-09-07T05:37:16+08:00
tags = ["Leetcode", "Algorithms", "Java", "Sort", "Heap", "BinarySearchTree"]
categories = ["leetcode"]
draft = false
+++

# 4 种方法秒杀 TopK（计数排序/快排变形/堆/二叉搜索树）

> [0347. Top K Frequent Elements](https://leetcode-cn.com/problems/top-k-frequent-elements/)
> [4 种方法秒杀 TopK（计数排序/快排变形/堆/二叉搜索树）](https://leetcode-cn.com/problems/top-k-frequent-elements/solution/4-chong-fang-fa-miao-sha-topkji-shu-pai-xu-kuai-pa/) by [sweetiee](https://leetcode-cn.com/u/sweetiee/)

*每天起床第一句 先给自己打个气
每次多点一个赞 都要说声谢谢你
魔镜魔镜看看我 我的读者在哪里
努力我要努力 我要读者都满意 >3<*
#### 一、前言

众所周知，**TopK** 是面试必问的题目，本文提供 4 种解法，其中 **堆** 和 **二叉搜索树** 的时间复杂度是 $O(NlogK)$，**计数排序（桶排序）** 和 **快排变形** 的时间复杂度是 $O(N)$。
因为所有的 TopK 问题的解法都是一样的，所以具体讲解可以看我之前的一篇题解 [【Leetcode每日打卡】4种解法秒杀TopK（快排变形/堆/二叉搜索树/计数排序）](https://zhuanlan.zhihu.com/p/114699207)。

#### 二、代码实现
##### 方法一：小根堆
``` Java
class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        // 统计每个数字出现的次数
        Map<Integer, Integer> counter = IntStream.of(nums).boxed().collect(Collectors.toMap(e -> e, e -> 1, Integer::sum));
        // 定义小根堆，根据数字频率自小到大排序
        Queue<Integer> pq = new PriorityQueue<>((v1, v2) -> counter.get(v1) - counter.get(v2));
        // 遍历数组，维护一个大小为 k 的小根堆：
        // 不足 k 个直接将当前数字加入到堆中；否则判断堆中的最小次数是否小于当前数字的出现次数，若是，则删掉堆中出现次数最少的一个数字，将当前数字加入堆中。
        counter.forEach((num, cnt) -> {
            if (pq.size() < k) {
                pq.offer(num);
            } else if (counter.get(pq.peek()) < cnt) {
                pq.poll();
                pq.offer(num);
            }
        });
        // 构造返回结果
        int[] res = new int[k];
        int idx = 0;
        for (int num: pq) {
            res[idx++] = num;
        }
        return res;
    }
}
```

##### 方法二：二叉搜索树
本方法和上面的方法一一样，只是把小根堆换成了二叉搜索树。
``` Java
class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        // 统计每个数字出现的次数
        Map<Integer, Integer> counterMap = IntStream.of(nums).boxed().collect(Collectors.toMap(e -> e, e -> 1, Integer::sum));
        // 定义二叉搜索树：key 是数字出现的次数，value 是出现了key次的数字列表。
        TreeMap<Integer, List<Integer>> treeMap = new TreeMap<>();
        // 维护一个有 k 个数字的二叉搜索树：
        // 不足 k 个直接将当前数字加入到树中；否则判断当前树中的最小次数是否小于当前数字的出现次数，若是，则删掉树中出现次数最少的一个数字，将当前数字加入树中。
        int count = 0;
        for(Map.Entry<Integer, Integer> entry: counterMap.entrySet()) {
            int num = entry.getKey();
            int cnt = entry.getValue();
            if (count < k) {
                treeMap.computeIfAbsent(cnt, ArrayList::new).add(num);
                count++;
            } else {
                Map.Entry<Integer, List<Integer>> firstEntry = treeMap.firstEntry();
                if (cnt > firstEntry.getKey()) {
                    treeMap.computeIfAbsent(cnt, ArrayList::new).add(num);
                    List<Integer> list = firstEntry.getValue();
                    if (list.size() == 1) {
                        treeMap.pollFirstEntry();
                    } else {
                        list.remove(list.size() - 1);
                    }
                }
            }
        }
        // 构造返回结果
        int[] res = new int[k];
        int idx = 0;
        for (List<Integer> list: treeMap.values()) {
            for (int num: list) {
                res[idx++] = num;
            }
        }
        return res;
    }
}
```

##### 方法三：计数排序（桶排序）
``` Java
class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        // 统计每个数字出现的次数
        Map<Integer, Integer> counterMap = IntStream.of(nums).boxed().collect(Collectors.toMap(e -> e, e -> 1, Integer::sum));
        // 一个数字最多出现 nums.length 次，因此定义一个长度为 nums.length + 1 的数组，freqList[i] 中存储出现次数为 i 的所有数字。
        List<Integer>[] freqList = new List[nums.length + 1];
        for (int i = 0; i < freqList.length; i++) {
            freqList[i] = new ArrayList<>();
        }
        counterMap.forEach((num, freq) -> {
            freqList[freq].add(num);
        });
        // 按照出现频次，从大到小遍历频次数组，构造返回结果。
        int[] res = new int[k];
        int idx = 0;
        for (int freq = freqList.length - 1; freq > 0; freq--) {
            for (int num: freqList[freq]) {
                res[idx++] = num;
                if (idx == k) {
                    return res;
                }
            }
        }
        return res;
    }
}
```

##### 方法四：快排变形
``` Java
class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        // 统计每个数字出现的次数
        Map<Integer, Integer> counterMap = IntStream.of(nums).boxed().collect(Collectors.toMap(e -> e, e -> 1, Integer::sum));
        // 构造Pair数组，Pair.num 表示数值，Pair.freq 表示数字出现的次数
        Pair[] pairs = IntStream.of(nums).distinct().boxed().map(num -> new Pair(num, counterMap.get(num))).toArray(Pair[]::new);
        // 使用快排变形 O(N) 获取Pair数组中频次最大的 k 个元素（第 4 个参数是下标，因此是 k - 1）。
        Pair[] topKPairs = quickSelect(pairs, 0, pairs.length - 1, k - 1);
        
        // 构造返回结果
        int[] res = new int[k];
        int idx = 0;
        for (Pair pair: topKPairs) {
            res[idx++] = pair.num;
        }
        return res;
    }

    private Pair[] quickSelect(Pair[] pairs, int lo, int hi, int idx) {
        if (lo > hi) {
            return new Pair[0];
        }
        // 快排切分后，j 左边的数字出现的次数都大于等于 pairs[j].freq，j 右边的数字出现的次数都小于等于 pairs[j].freq。
        int j = partition(pairs, lo, hi);
        // 如果 j 正好等于目标idx，说明当前pairs数组中的[0, idx]就是出现次数最大的 K 个元素。
        if (j == idx) {
            return Arrays.copyOf(pairs, idx + 1);
        }
        // 否则根据 j 与 idx 的大小关系判断找左段还是右段
        return j < idx? quickSelect(pairs, j + 1, hi, idx): quickSelect(pairs, lo, j - 1, idx);
    }

    private int partition(Pair[] pairs, int lo, int hi) {
        Pair v = pairs[lo];
        int i = lo, j = hi + 1;
        while (true) {
            while(++i <= hi && pairs[i].freq > v.freq);
            while(--j >= lo && pairs[j].freq < v.freq);
            if (i >= j) {
                break;
            }
            Pair tmp = pairs[i];
            pairs[i] = pairs[j];
            pairs[j] = tmp;
        }
        pairs[lo] = pairs[j];
        pairs[j] = v;
        return j;
    }
}

class Pair {
    int num;
    int freq;

    public Pair(int num, int freq) {
        this.num = num;
        this.freq = freq;
    }
}
```
#### 三、关注我
如果觉得不错，那就快快戳我主页关注我吧，更多优质题解等你哦。