"+++
title = "0141. Linked List Cycle Swift 双指针 94% "
author = ["gesen"]
date = 2020-10-09T02:48:23+08:00
tags = ["Leetcode", "Algorithms", "Swift"]
categories = ["leetcode"]
draft = false
+++

# Swift 双指针 94%

> [0141. Linked List Cycle](https://leetcode-cn.com/problems/linked-list-cycle/)
> [Swift 双指针 94%](https://leetcode-cn.com/problems/linked-list-cycle/solution/swift-shuang-zhi-zhen-94-by-gesen/) by [gesen](https://leetcode-cn.com/u/gesen/)

```swift
class Solution {
    func hasCycle(_ head: ListNode?) -> Bool {
        var slow = head
        var fast = head?.next
        
        while slow != nil && fast != nil {
            if slow === fast { return true }
            slow = slow?.next
            fast = fast?.next?.next
        }

        return false
    }
}
```