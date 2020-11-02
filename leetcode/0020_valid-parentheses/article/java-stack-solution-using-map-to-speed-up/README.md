"+++
title = "0020. Valid Parentheses [Java] Stack solution using map to speed up! "
author = ["lincs"]
date = 2020-08-28T08:37:13+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# [Java] Stack solution using map to speed up!

> [0020. Valid Parentheses](https://leetcode-cn.com/problems/valid-parentheses/)
> [[Java] Stack solution using map to speed up!](https://leetcode-cn.com/problems/valid-parentheses/solution/java-stack-solution-using-map-to-speed-up-by-lincs/) by [lincs](https://leetcode-cn.com/u/lincs/)

```java
class Solution {
    public boolean isValid(String s) {
        Map<Character, Character> map = new HashMap<>();
        map.put(')', '(');
        map.put(']', '[');
        map.put('}', '{');
        
        Stack<Character> stack = new Stack<>();
        for (char ch : s.toCharArray()) {
            if (ch == '(' || ch == '[' || ch == '{') {
                stack.push(ch);
            }
            else if (stack.isEmpty() || stack.pop() != map.get(ch)){
                return false;
            }
        }
        return stack.isEmpty();
    }
}
```
