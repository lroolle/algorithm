"+++
title = "0347. Top K Frequent Elements swift  双100%解题 "
author = ["Esdeath"]
date = 2020-09-07T03:18:41+08:00
tags = ["Leetcode", "Algorithms", "Swift"]
categories = ["leetcode"]
draft = false
+++

# swift  双100%解题

> [0347. Top K Frequent Elements](https://leetcode-cn.com/problems/top-k-frequent-elements/)
> [swift  双100%解题](https://leetcode-cn.com/problems/top-k-frequent-elements/solution/swift-shuang-100jie-ti-by-esdeath/) by [Esdeath](https://leetcode-cn.com/u/esdeath/)

### 解题思路
![image.png](https://pic.leetcode-cn.com/1599448717-VIpRiN-image.png)

### 代码

```swift
class Solution {
func topKFrequent(_ nums: [Int], _ k: Int) -> [Int] {
    return nums.reduce(into: [:]) { (parames, number) in
        parames[number,default: 0] += 1
    }.sorted(by: {$0.value > $1.value}).prefix(k).map{$0.key}
}
}
```