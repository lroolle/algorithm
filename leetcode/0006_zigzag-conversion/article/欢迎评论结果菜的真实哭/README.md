"+++
title = "0006. ZigZag Conversion 欢迎评论（结果菜的真实（哭）） "
author = ["elias-u"]
date = 2020-08-23T13:23:25+08:00
tags = ["Leetcode", "Algorithms", "C"]
categories = ["leetcode"]
draft = false
+++

# 欢迎评论（结果菜的真实（哭））

> [0006. ZigZag Conversion](https://leetcode-cn.com/problems/zigzag-conversion/)
> [欢迎评论（结果菜的真实（哭））](https://leetcode-cn.com/problems/zigzag-conversion/solution/huan-ying-ping-lun-jie-guo-cai-de-zhen-shi-ku-by-e/) by [elias-u](https://leetcode-cn.com/u/elias-u/)

### 解题思路
简单解释一下：找到每行上两个字符为一组时两字符在新的字符串中的位置规律，如果只有一行，直接返回字符串s。
             在第一行（numRows）时特殊处理（行数是倒着数的，第一行记为numRows，最后一行记为1）
             两个为一组
            1. numRows-0:从左到右对应新字符串中下标规律：(0、0+2*(numRows-0-1)*1)、(0+2*(numRows-0-1)*2、..)、...
            2. numRows-1:略
            3. ...
            4. numRows-i: *(i、t0={i+2*(numRows-i-1)})*、**(t0+2*(numRows-i)、t1={(t0+2*(numRows-i))+2*(numRows-i-1)})**、...//相邻组之间对应位置相差2*(numRows-i)+2*(numRows-i-1) //numRows-i为倒着数对应的行数 //一组中两下标之间差2*(numRows-i-1)
            5. ...
            6. 1:
            每一行结束条件:下标超过strlen(s);
            列结束条件：i=0；
            
### 代码

```c
char * convert(char * s, int numRows){
     int stairs=numRows;
     char* news=(char*)malloc(sizeof(char)*1000);
     int k=0;
     if(stairs==1)
     {
         return s;
     }
     for(int i=stairs;i>0;i--)//列结束判断
    {
         int cur=2*(i-1); //下标的差（部分）//组中两下标差
         int j=stairs-i;//每一行起始下标
         while(j<strlen(s))//行结束判断
         {
             news[k++]=s[j];
             j=j+cur;
             if(i!=stairs)//numRows对应行不执行以下语句
             {
                 if(j<strlen(s)&&cur!=0)
                 news[k++]=s[j];
                 j=j+(stairs-i)*2;//组之间第二大下标与第三大下标的差（右到左）
             }
         }
    }
    news[k]=0;
    return news;
}

```
欢迎评论，多多评论
败给现实，哎。。。
![image.png](https://pic.leetcode-cn.com/1598185447-YTdrmv-image.png)
