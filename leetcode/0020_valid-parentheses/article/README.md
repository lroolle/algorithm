"+++
title = "0020. Valid Parentheses 数组栈 "
author = ["chenghua-3"]
date = 2020-08-27T09:13:37+08:00
tags = ["Leetcode", "Algorithms", "C"]
categories = ["leetcode"]
draft = false
+++

# 数组栈

> [0020. Valid Parentheses](https://leetcode-cn.com/problems/valid-parentheses/)
> [数组栈](https://leetcode-cn.com/problems/valid-parentheses/solution/shu-zu-zhan-by-chenghua-3/) by [chenghua-3](https://leetcode-cn.com/u/chenghua-3/)

### 解题思路
使用数组栈编写

### 代码

```c
bool isValid(char * s)
{
   if((*s)=='\0')   return true;
   int long_s=strlen(s);
   int *zu=(int *)malloc(sizeof(int)*long_s),count=-1;
   char *tmp=s;
   while((*tmp)!='\0'){
   if((*tmp)=='[' || (*tmp)=='(' || (*tmp)=='{')
   {
       count++;zu[count]=(*tmp);
   }
   else if(((*tmp)==']' || (*tmp)==')' || (*tmp)=='}'))
   {
       if(count==-1) return false;
       switch((*tmp))
       {
           case ']':
                if('['==zu[count])
                {
                    count--;
                }
                else return false;
                break;
            case '}':
                if('{'==zu[count])
                {
                    count--;
                }
                else return false;
                break;
            case ')':
                if('('==zu[count])
                {
                    count--;
                }
                else return false;
                break;
       }
   }
   tmp++;
   }
   if(count==-1)    return true;
   else return false;
}
```