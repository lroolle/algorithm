"+++
title = "0141. Linked List Cycle 巧妙利用javascript的特点 "
author = ["finnwu"]
date = 2020-08-23T00:35:46+08:00
tags = ["Leetcode", "Algorithms", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# 巧妙利用javascript的特点

> [0141. Linked List Cycle](https://leetcode-cn.com/problems/linked-list-cycle/)
> [巧妙利用javascript的特点](https://leetcode-cn.com/problems/linked-list-cycle/solution/qiao-miao-li-yong-javascriptde-te-dian-by-finnwu/) by [finnwu](https://leetcode-cn.com/u/finnwu/)

### 巧妙利用javascript的特点

```javascript
var hasCycle = function(head) {
    while(head) {
        if (head.cycle) {
            return true
        } else {
            head.cycle = 1
            head = head.next
        }
    }
    return false
};
```
