"+++
title = "面试题 17.25. Word Rectangle LCCI trie+暴力搜索 "
author = ["lddlinan"]
date = 2020-07-16T05:38:47+08:00
tags = ["Leetcode", "Algorithms", "Python", "Trie"]
categories = ["leetcode"]
draft = false
+++

# trie+暴力搜索

> [面试题 17.25. Word Rectangle LCCI](https://leetcode-cn.com/problems/word-rectangle-lcci/)
> [trie+暴力搜索](https://leetcode-cn.com/problems/word-rectangle-lcci/solution/triebao-li-sou-suo-by-lddlinan/) by [lddlinan](https://leetcode-cn.com/u/lddlinan/)

### 解题思路
遍历矩形的长和寛，2956ms勉强过了

### 代码

```python
class Solution(object):
    def maxRectangle(self, words):
        """
        :type words: List[str]
        :rtype: List[str]
        """
        trie = [None]*101
        mlen = 0
        for w in words:
            n = len(w)
            if mlen<n: mlen=n
            if trie[n] is None: trie[n] = [None]*26
            r = trie[n]
            for c in w:
                cc = ord(c)-ord('a')
                if r[cc] is None:
                    r[cc] = [None]*26
                r = r[cc]
        rr, cr = [None]*101, [None]*101
        rss = [["x"]*101 for _ in range(101)]
        def dfs(c, r, w, h):
            if c==w: c, r = 0, r+1
            if r==h: return True
            for i in range(26):
                if rr[c][i] and cr[r][i]:
                    bc, br = rr[c], cr[r]
                    rr[c], cr[r] = rr[c][i], cr[r][i]
                    rss[r][c]=chr(i+ord('a'))
                    if dfs(c+1, r, w, h): return True
                    rr[c], cr[r] = bc, br
            return False
        w, rs, rw, rh = mlen, 0, 0, 0
        while w>0:
            if w*mlen<=rs: break
            cw, w = w, w-1
            if trie[cw] is None: continue
            h = mlen
            while h>0:
                ch, h = h, h-1
                if trie[ch] is None: continue
                for i in range(cw): rr[i] = trie[ch]
                for i in range(ch): cr[i] = trie[cw]
                if dfs(0, 0, cw, ch):
                    cs = cw*ch
                    if rs<cs: rs, rw, rh = cs, cw, ch
                    break
        # print rs, rw, rh
        if rs==0: return [""]
        for i in range(rw): rr[i] = trie[rh]
        for i in range(rh): cr[i] = trie[rw]
        dfs(0, 0, rw, rh)
        return ["".join(r[:rw]) for r in rss[:rh]]

        
```