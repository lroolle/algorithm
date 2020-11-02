"+++
title = "0037. Sudoku Solver 【解数独】回溯 + 状态压缩（使用 bitset） "
author = ["ikaruga"]
date = 2020-03-17T13:48:12+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# 【解数独】回溯 + 状态压缩（使用 bitset）

> [0037. Sudoku Solver](https://leetcode-cn.com/problems/sudoku-solver/)
> [【解数独】回溯 + 状态压缩（使用 bitset）](https://leetcode-cn.com/problems/sudoku-solver/solution/37-by-ikaruga/) by [ikaruga](https://leetcode-cn.com/u/ikaruga/)

### 思路
1. 状态压缩
    11. 使用 `bitset<9>` 来压缩存储每一行、每一列、每一个 `3x3` 宫格中 `1-9` 是否出现
    12. 这样每一个格子就可以计算出所有不能填的数字，然后得到所有**能填的数字** `getPossibleStatus()`
    13. 填入数字和回溯时，只需要更新存储信息
    14. 每个格子在使用时，会根据存储信息重新计算**能填的数字**

2. 回溯
    21. 每次都使用 `getNext()` 选择**能填的数字**最少的格子开始填，这样填错的概率最小，回溯次数也会变少
    22. 使用 `fillNum()` 在填入和回溯时负责更新存储信息
    23. 一旦全部填写成功，一路返回 true ，结束递归


### 图解
![图片.png](https://pic.leetcode-cn.com/1fb1c64cfddb5c66b61bd769224724a05027172d6486feb19b3a16d9473372ee-%E5%9B%BE%E7%89%87.png)


### 答题
```C++
class Solution {
public:
    bitset<9> getPossibleStatus(int x, int y)
    {
        return ~(rows[x] | cols[y] | cells[x / 3][y / 3]);
    }

    vector<int> getNext(vector<vector<char>>& board)
    {
        vector<int> ret;
        int minCnt = 10;
        for (int i = 0; i < board.size(); i++)
        {
            for (int j = 0; j < board[i].size(); j++)
            {
                if (board[i][j] != '.') continue;
                auto cur = getPossibleStatus(i, j);
                if (cur.count() >= minCnt) continue;
                ret = { i, j };
                minCnt = cur.count();
            }
        }
        return ret;
    }

    void fillNum(int x, int y, int n, bool fillFlag)
    {
        rows[x][n] = (fillFlag) ? 1 : 0;
        cols[y][n] = (fillFlag) ? 1 : 0;
        cells[x/3][y/3][n] = (fillFlag) ? 1: 0;
    }
    
    bool dfs(vector<vector<char>>& board, int cnt)
    {
        if (cnt == 0) return true;

        auto next = getNext(board);
        auto bits = getPossibleStatus(next[0], next[1]);
        for (int n = 0; n < bits.size(); n++)
        {
            if (!bits.test(n)) continue;
            fillNum(next[0], next[1], n, true);
            board[next[0]][next[1]] = n + '1';
            if (dfs(board, cnt - 1)) return true;
            board[next[0]][next[1]] = '.';
            fillNum(next[0], next[1], n, false);
        }
        return false;
    }

    void solveSudoku(vector<vector<char>>& board) 
    {
        rows = vector<bitset<9>>(9, bitset<9>());
        cols = vector<bitset<9>>(9, bitset<9>());
        cells = vector<vector<bitset<9>>>(3, vector<bitset<9>>(3, bitset<9>()));

        int cnt = 0;
        for (int i = 0; i < board.size(); i++)
        {
            for (int j = 0; j < board[i].size(); j++)
            {
                cnt += (board[i][j] == '.');
                if (board[i][j] == '.') continue;
                int n = board[i][j] - '1';
                rows[i] |= (1 << n);
                cols[j] |= (1 << n);
                cells[i / 3][j / 3] |= (1 << n);
            }
        }
        dfs(board, cnt);
    }

private:
    vector<bitset<9>> rows;
    vector<bitset<9>> cols;
    vector<vector<bitset<9>>> cells;
};
```

感谢 [@cocowy](/u/cocowy/) 提供的 bitset 风格的 `fillNum()` ，可读性得到质的提升，已更新到上面
下面是我原来位运算风格的 `fillNum()`
```C++
    void fillNum(int x, int y, int n, bool fillFlag)
    {
        bitset<9> pick(1 << n);
        rows[x] = (fillFlag) ? (rows[x] | pick) : (rows[x] ^ pick);
        cols[y] = (fillFlag) ? (cols[y] | pick) : (cols[y] ^ pick);
        cells[x / 3][y / 3] = (fillFlag) ? (cells[x / 3][y / 3] | pick) : (cells[x / 3][y / 3] ^ pick);
    }
```

### 执行时间

![图片.png](https://pic.leetcode-cn.com/a5e6450027aa59d98449f96ff4633613b49e2e51dc521085598667132e442c3f-%E5%9B%BE%E7%89%87.png)

对格子填入顺序的选择，使得搜索效率大大提升（加上运气好），击败 100%   
得益于状态压缩，内存的使用也击败 100% 
（别太较真儿，看个乐就好）



### 致谢

感谢您的观看，希望对您有帮助，欢迎热烈的交流！  

**如果感觉还不错就点个赞吧~**