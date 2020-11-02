"+++
title = "0037. Sudoku Solver 反正我觉得我这个菜鸡的代码很好理解 "
author = ["amer-3"]
date = 2020-05-05T13:15:51+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# 反正我觉得我这个菜鸡的代码很好理解

> [0037. Sudoku Solver](https://leetcode-cn.com/problems/sudoku-solver/)
> [反正我觉得我这个菜鸡的代码很好理解](https://leetcode-cn.com/problems/sudoku-solver/solution/fan-zheng-wo-jue-de-wo-zhe-ge-cai-ji-de-dai-ma-hen/) by [amer-3](https://leetcode-cn.com/u/amer-3/)

### 解题思路
![image.png](https://pic.leetcode-cn.com/6578a7fccf1302e082e9cfe3641ad75c47ee12670d6b6de5f5c98615f821f541-image.png)

### 代码

```cpp
class Solution {
public:
    int row[9][9] = {0};
    int col[9][9] = {0};
    int l[9][9] = {0};
    void solveSudoku(vector<vector<char>>& board) {
        for(int i=0;i<9;++i){
            for(int j=0;j<9;++j){
                if(board[i][j]!='.'){
                    row[i][board[i][j]-'1'] = 1;
                    col[j][board[i][j]-'1'] = 1;
                    l[i/3*3+j/3][board[i][j]-'1'] = 1;
                }
            }
        }
        backtrack(board,0,0);
    }

    bool backtrack(vector<vector<char>>& board,int r,int c){
        //递归边界
        if(r>8) return true;
        if(board[r][c]=='.'){
            for(char ch='1';ch<='9';++ch){
                //符合要求
                if(row[r][ch-'1']==1||col[c][ch-'1']==1||l[r/3*3+c/3][ch-'1']==1) continue;
                board[r][c] = ch;
                row[r][ch-'1'] = 1;
                col[c][ch-'1'] = 1;
                l[r/3*3+c/3][ch-'1'] = 1;

                //只需找到一个解即可
                if(backtrack(board,r+(c+1)/9,(c+1)%9)) return true;

                //回溯
                board[r][c] = '.';
                row[r][ch-'1'] = 0;
                col[c][ch-'1'] = 0;
                l[r/3*3+c/3][ch-'1'] = 0;
            }
        }
        else return backtrack(board,r+(c+1)/9,(c+1)%9);

        //对空格处操作：没有得到合适的解就会跳到此处 
        return false;
    }
};
```