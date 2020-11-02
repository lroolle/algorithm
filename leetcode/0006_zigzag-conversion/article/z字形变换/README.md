"+++
title = "0006. ZigZag Conversion Z字形变换 "
author = ["a-mao-da-ma"]
date = 2020-08-19T14:43:00+08:00
tags = ["Leetcode", "Algorithms", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# Z字形变换

> [0006. ZigZag Conversion](https://leetcode-cn.com/problems/zigzag-conversion/)
> [Z字形变换](https://leetcode-cn.com/problems/zigzag-conversion/solution/zzi-xing-bian-huan-by-a-mao-da-ma/) by [a-mao-da-ma](https://leetcode-cn.com/u/a-mao-da-ma/)

### 解题思路
看题意大概理解，就是边上的字符顺着遍历，而中间斜线的字符是倒着遍历的（除边上的头和尾）。所以中间的遍历是要比头尾的遍历要多的，两个用不同的遍历来记录遍历次数。
用二维数组来存放z字形变换，其中用front来记录头尾的遍历次数，用mid来记录中间的遍历次数。
那么就是一次竖着遍历走numRows个数组，front、mid各+1，然后再倒着遍历除numRows-1和0的数组（即中间），mid +1。
遍历完之后获得的二维数组再拼接成字符串就是答案了。
为了方便理解，我画了个草图
![image.png](https://pic.leetcode-cn.com/6c0eb917b2c1b4fe7afa6657b602fe054cad4b1df521ecec6561ffc9241b519a-image.png)
### 代码

```javascript
/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */
var convert = function(s, numRows) {
    let map = new Array(numRows); //存放Z字形的二维数组
    for(let i=0;i<map.length;i++){
        map[i] = new Array();
    }
    let front = 0; //记录头尾的遍历次数
    let mid = 0; //记录中间的遍历次数
    let num = 0; //记录当前在s字符串中的位置

    while(num<s.length){
        for(let i=0;i<map.length;i++){
            if(num >= s.length) break; //每次遍历都可能超出字符串，所以要及时退出
            //分头尾和中间的情况
            if(i ==0 || i== map.length - 1){
                map[i][front] = s[num++]; 
            } 
            else{
                map[i][mid] = s[num++];
            }
        }
        front += 1;
        mid += 1;
        //斜线上的遍历
        for(let i= map.length - 2;i>=1;i--){
            if(num >= s.length) break;
            map[i][mid] = s[num++];
        }
        mid += 1;
    }
    //二维数组转字符串
    let res = '';
    for(let i=0;i<map.length;i++){
        res += map[i].join('');
    }
    return res

};
```