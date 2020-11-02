"+++
title = "0006. ZigZag Conversion （详细）比较干净的解题思路，16行代码可破 "
author = ["trenlinhuang"]
date = 2020-09-12T07:28:33+08:00
tags = ["Leetcode", "Algorithms", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# （详细）比较干净的解题思路，16行代码可破

> [0006. ZigZag Conversion](https://leetcode-cn.com/problems/zigzag-conversion/)
> [（详细）比较干净的解题思路，16行代码可破](https://leetcode-cn.com/problems/zigzag-conversion/solution/ke-neng-shi-zui-hao-dong-de-na-chong-jie-ti-fang-s/) by [trenlinhuang](https://leetcode-cn.com/u/trenlinhuang/)

### 先上性能
执行用时：96 ms, 在所有 JavaScript 提交中击败了94.74%的用户
内存消耗：42.1 MB, 在所有 JavaScript 提交中击败了54.60%的用户

*这个垃圾算法的不垃圾的地方在于好理解*

### 解题思路
Z或者N型排列看着挺吓人的
```
L     D     R
E   O E   I I
E C   I H   N
T     S     G
```
对于题目想要的输出，我们其实可以转化成
```
L   D   R
E O E I I
E C I H N
T   S   G
```
这样一来问题就下降了一个难度，关键点其实是在于将这些字母进行分组，然后再按照一定的规律把它进行输出或者说拼接

将字母按`size = 2*nRows-2`个分为一组后（4行时就为6个一组）

```
L    0 
E O  1 5
E C  2 4
T    3 
```

可以发现，字母和他们的组号存在对应关系，其余字母相对于组内首字母存在一个偏移量
```
第一列：                第二列：
L -> size*0 + 0         
E -> size*0 + 1         O -> size*0 + size - 1
E -> size*0 + 2         C -> size*0 + size - 2
T -> size*0 + 3(size/2)
```
第二列比较特殊，逆序排列，可以利用和左侧字母组内序号加和刚好为`size`进行处理
找到规律后，拼接成待求字符串即可

我的遍历方式为按行遍历，取出各组对应位置进行拼接
其中首末行无第二列，判断处理下避免越界或者重复添加

### 代码

```javascript
/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */
var convert = function(s, numRows) {
    if(numRows<=1) return s;
    const length = s.length;
    const size = 2*numRows-2;
    let res = "";
    for(let j=0; j<=size/2; j++) {
        for(let i=0; i<=parseInt(length/size); i++) {
            let iCol1 = i*size + j;
            let iCol2 = (i+1)*size + -j;
            if(iCol1 < length) res += s[iCol1];
            if(j==0 || j==size/2) continue;
            if(iCol2 < length && iCol1!=iCol2) res += s[iCol2];
        }
    }
    return res;
};
```

不点个赞吗？🤐