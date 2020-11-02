"+++
title = "0037. Sudoku Solver 37. 解数独 + 递归判断每行每列每格子 "
author = ["Jasion_han"]
date = 2020-08-15T11:35:44+08:00
tags = ["Leetcode", "Algorithms", "Java", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 37. 解数独 + 递归判断每行每列每格子

> [0037. Sudoku Solver](https://leetcode-cn.com/problems/sudoku-solver/)
> [37. 解数独 + 递归判断每行每列每格子](https://leetcode-cn.com/problems/sudoku-solver/solution/37-jie-shu-du-di-gui-pan-duan-mei-xing-mei-lie-mei/) by [Jasion_han](https://leetcode-cn.com/u/jasion_han/)

### 解题思路
此题和 `36` 题很相似，但是要简单一点，对比着学习才能熟练掌握该类题型！

### 代码

```java
class Solution {
    public void solveSudoku(char[][] board) {
        if (board == null || board.length == 0 || board[0].length == 0) {
            return;
        }
        solve(board);
    }
    
    private boolean solve(char[][] board) {
        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[0].length; j++) {
                // 如果是空格
                if (board[i][j] == '.') {
                    // 就寻找剩余能放的数字
                    for (char num = '1'; num <= '9'; num++) {
                        // 先判断当前所指数字是否有效，也就是之前没有放过该数字
                        if (isValid(board, i, j, num)) {
                            // 放入该有效位置
                            board[i][j] = num;
                            // 接着递归判断剩余空格是否也可以满足
                            if (solve(board)) {
                                return true;
                            // 不满足就恢复原状进行下一轮循环
                            } else {
                                board[i][j] = '.';
                            }
                        }
                    }
                    return false;
                }
            }
        }
        return true;
    }
    
    private boolean isValid(char[][] board, int row, int col, char num) {
        for (int i = 0; i < 9; i++) {
            // 判断当前行，当前列，当前格子里是否出现过现在正遍历到的数字
            if (board[i][col] == num) {
                return false;
            }
            if (board[row][i] == num) {
                return false;
            }
            if (board[3 * (row / 3) + i / 3][3 * (col / 3) + i % 3] == num) {
                return false;
            }
        }
        return true;
    }
}
```