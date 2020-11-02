"+++
title = "0006. ZigZag Conversion Z 字形变换 - 消费补偿算法 "
author = ["banana798"]
date = 2020-09-15T01:13:38+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# Z 字形变换 - 消费补偿算法

> [0006. ZigZag Conversion](https://leetcode-cn.com/problems/zigzag-conversion/)
> [Z 字形变换 - 消费补偿算法](https://leetcode-cn.com/problems/zigzag-conversion/solution/z-zi-xing-bian-huan-xiao-fei-bu-chang-suan-fa-by-b/) by [banana798](https://leetcode-cn.com/u/banana798/)

解决这题的关键是找到每一行的字符位置规律，根据分析可以知道：
第一行和最后一行的字符位置相差 `2*numRows-2`，如下图所示：

![规律1](https://pic.leetcode-cn.com/1600132308-cYeVLD-2020091421520727.png)

而对于其他行来说，字符位置加4会略过当前行的一个字符。如下图所示：

![202009142155307.png](https://pic.leetcode-cn.com/1600132339-CrXwXb-202009142155307.png)

先定义一个money作为这个跨度，即`int money = 2*numRows-2;`，定义 j 为要访问的s字符串的字符下标。

然后，我打算对当前行做一个"补偿"，即让他倒退2个位置。即`j = j + money - (2*i - 2)`：，这样一来，T的位置就可以在E的基础上算出来（j=1, money=4）： 1+4-(2*2-2)=3。

你可以理解为原来我要消费money块钱，现在我补偿回一些钱给你。那么在下次消费时，就要恰好把money消费完。即：`money -= money - (2*i - 2)`，即`money=4-(2*2-2)=2`，这样就可以访问到O字符了。

这个算法，我称之为`消费补偿算法`，还没有看题解，不知道其他大佬的思想是不是类似。下面贴我的代码：

```cpp
#include <iostream>
#include <string>
using namespace std;

/**
 * LeetCode
 * 6. Z 字形变换
 * https://space.bilibili.com/54183978
 */

class Solution {
public:
    string convert(string s, int numRows) {
         int size = s.size();
         string ans = "";

         if(numRows == 1){
             return s;
         }

         for(int i = 1; i <= numRows; i++){
             // 加第i行字符
             int j = i - 1;
             // 每行第一个字符
             ans += s[j];
             int money = 2*numRows-2;

             while(j < size){

                 if(money == 2*numRows-2 && i!=numRows){
                     // 消费补偿
                     j = j + money - (2*i - 2);
                     if(j >= size){
                         break;
                     }
                     ans += s[j];
                     money -= money - (2*i - 2);
                 } else{
                     // 消费剩余
                     j = j + money;
                     if(j >= size){
                         break;
                     }
                     ans += s[j];
                     money = 0;
                 }

                 if(money == 0){
                     money = 2*numRows-2;
                 }
             }
         }
         return ans;
    }
};

int main(){
    Solution solution;
    cout << solution.convert("LEETCODEISHIRING", 3);
}
```
下面是代码评测结果：
![20200914220619159.png.jpeg](https://pic.leetcode-cn.com/1600132386-ZGehfL-20200914220619159.png.jpeg)
