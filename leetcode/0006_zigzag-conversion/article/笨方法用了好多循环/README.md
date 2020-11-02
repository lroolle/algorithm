"+++
title = "0006. ZigZag Conversion 笨方法，用了好多循环 "
author = ["denfel"]
date = 2020-08-28T04:40:48+08:00
tags = ["Leetcode", "Algorithms", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# 笨方法，用了好多循环

> [0006. ZigZag Conversion](https://leetcode-cn.com/problems/zigzag-conversion/)
> [笨方法，用了好多循环](https://leetcode-cn.com/problems/zigzag-conversion/solution/ben-fang-fa-yong-liao-hao-duo-xun-huan-by-denfel/) by [denfel](https://leetcode-cn.com/u/denfel/)

### 解题思路
>翻转一下，用数组做，例如s='ABCD',  numRows = 3;
则arrZ = [['ABC'],['CD']],
再把arrZ转换成[['A','B','C'],['','D']]以对应位置，
最后循环遍历一下就ok了。
边错边改的代码，条理不是很清晰，继续学习！

### 代码

```javascript
/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */
var convert = function(s, numRows) {
    if(numRows == 1) return s;
    var arrZ = [];
    var i = 0;
    while(i < s.length){
        arrZ.push(s.slice(i,i+numRows));
        i+=numRows-1;
    }
    for(let j = 0; j < arrZ.length; j++){
        if(j%2==0){
            arrZ[j] = arrZ[j].split('');
            
        }else{
            arrZ[j] = arrZ[j].split('');
            if(arrZ[j].length<numRows){
                //z字斜着的转换成['',...,'char',...,''],中间不足用''补齐
                for(let k = arrZ[j].length; k<numRows; k++){
                    arrZ[j][k] = '';
                }
            }
            arrZ[j][0] = '';
            var length = arrZ[j].length == numRows?numRows-2:arrZ[j].length-1;
            arrZ[j].splice(1,length,...arrZ[j].slice(1,length+1).reverse());
            if(arrZ[j].length==numRows) arrZ[j][arrZ[j].length-1]='';
        }
    }
    var str = '';
    for(let j = 0; j < numRows; j++){
        for(let k = 0; k < arrZ.length; k++){
            if(arrZ[k][j]) str += arrZ[k][j];
        }
    }
    return str;
};
```