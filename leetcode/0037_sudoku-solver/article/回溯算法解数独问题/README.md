"+++
title = "0037. Sudoku Solver 回溯算法解数独问题 "
author = ["yong-cai-gu-zi-xi"]
date = 2020-08-28T04:02:34+08:00
tags = ["Leetcode", "Algorithms", "Backtracking"]
categories = ["leetcode"]
draft = false
+++

# 回溯算法解数独问题

> [0037. Sudoku Solver](https://leetcode-cn.com/problems/sudoku-solver/)
> [回溯算法解数独问题](https://leetcode-cn.com/problems/sudoku-solver/solution/da-jia-bu-yao-kan-labuladong-de-jie-fa-fei-chang-w/) by [yong-cai-gu-zi-xi](https://leetcode-cn.com/u/yong-cai-gu-zi-xi/)

解数独的最一般的解法，类似于 $N$ 皇后问题的回溯算法，可以参考 [这里](https://leetcode-cn.com/problems/n-queens/solution/gen-ju-di-46-ti-quan-pai-lie-de-hui-su-suan-fa-si-/)，已经遍历过的数字需要记下来。

下面的代码来自 [题解](https://leetcode-cn.com/problems/sudoku-solver/solution/fan-zheng-wo-jue-de-wo-zhe-ge-cai-ji-de-dai-ma-hen/)，是思路最清晰的。我基于他的写法做了修改和格式上的调整和简单的注释。

```Java
public class Solution {

    private boolean[][] row;
    private boolean[][] col;
    private boolean[][] box;

    public void solveSudoku(char[][] board) {
        row = new boolean[9][9];
        col = new boolean[9][9];
        box = new boolean[9][9];

        // 题目说：给定数独永远是 9 x 9 形式的，因此不用做特殊判断
        // 步骤 1：先遍历棋盘一次，然后每一行，每一列在 row col cell 里占住位置
        for (int i = 0; i < 9; i++) {
            for (int j = 0; j < 9; j++) {
                if (board[i][j] != '.') {
                    // 计算在布尔数值中的下标
                    int index = board[i][j] - '1';

                    // 在下标为 i 这一行，记录为 true
                    row[i][index] = true;
                    // 在下标为 j 这一列，记录为 true
                    col[j][index] = true;
                    // 在下标为 i / 3 * 3 + j / 3 的地方，看到数值 board[i][j] - '1'
                    box[i / 3 * 3 + j / 3][index] = true;
                }
            }
        }
        // 步骤 2：进行一次深度优先遍历，尝试所有的可能性
        dfs(board, 0, 0);
    }

    private boolean dfs(char[][] board, int i, int j) {
        if (i > 8) {
            return true;
        }

        // 对 '.' 尝试从 1 填到 9
        if (board[i][j] == '.') {
            for (char ch = '1'; ch <= '9'; ch++) {
                // 如果行、列、box 已经填了 ch - '1' 则尝试下一个数字
                int index = ch - '1';
                if (row[i][index] || col[j][index] || box[i / 3 * 3 + j / 3][index]) {
                    continue;
                }

                // 填写当前字符，并且对应 row、col、box 占位
                board[i][j] = ch;
                row[i][index] = true;
                col[j][index] = true;
                box[i / 3 * 3 + j / 3][index] = true;

                // ① 只需找到一个解即可，继续填写下一格 i + (j + 1) / 9 表示如果 j 已经在一列的末尾，跳转到下一行
                if (dfs(board, i + (j + 1) / 9, (j + 1) % 9)) {
                    return true;
                }

                // 重置变量
                board[i][j] = '.';
                row[i][index] = false;
                col[j][index] = false;
                box[i / 3 * 3 + j / 3][index] = false;
            }
        } else {
            // 填写下一格和 ① 一样
            return dfs(board, i + (j + 1) / 9, (j + 1) % 9);
        }
        // 全部尝试过以后，返回 false
        return false;
    }
}
```

大家还可以参考 [《状压、选择分支少的进行优化、预处理》](https://leetcode-cn.com/problems/sudoku-solver/solution/zhuang-ya-xuan-ze-fen-zhi-shao-de-jin-xing-you-hua/)，他的解法更优。