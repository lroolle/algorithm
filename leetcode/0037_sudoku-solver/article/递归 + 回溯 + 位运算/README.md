"+++
title = "0037. Sudoku Solver 递归 + 回溯 + 位运算 "
author = ["zoffer"]
date = 2020-06-23T05:33:15+08:00
tags = ["Leetcode", "Algorithms", "JavaScript", "Recursion", "BitManipulation", "Backtracking"]
categories = ["leetcode"]
draft = false
+++

# 递归 + 回溯 + 位运算

> [0037. Sudoku Solver](https://leetcode-cn.com/problems/sudoku-solver/)
> [递归 + 回溯 + 位运算](https://leetcode-cn.com/problems/sudoku-solver/solution/di-gui-hui-su-wei-yun-suan-by-zoffer-3/) by [zoffer](https://leetcode-cn.com/u/zoffer/)

### 递归
![sudoku.png](https://pic.leetcode-cn.com/743c22bc44e8911e6180e38e11a4b56ed420f4e8b6719ba3f2d99e00a3b9b4aa-sudoku.png)

- 每次选取**可填数字最少**的空格，试探次数更少，发现错误更快。
### 位运算
- 使用 9-bit 保存数字 1~9 的占用情况，通过位运算处理，更加轻松高效。

![250px-Sudoku-by-L2G-20050714.png](https://pic.leetcode-cn.com/c5a55e16c7bb150dd469e327d6f42872445068be666fe2f61df3ff0efa8b2d65-250px-Sudoku-by-L2G-20050714.png)

|  |                    |9|8|7|6|5|4|3|2|1
:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:
行|r = rows[5]          |0|0|1|1|0|0|0|1|0
列|c = columns[3]       |0|1|0|0|0|1|0|0|1
3x3方格|b = boxs[1][1]  |0|1|0|1|0|0|1|1|0
或运算|x = r \| c \| b  |0|1|1|1|0|1|1|1|1
取反 9-bit|p = x ^ 0b111111111 |1|0|0|0|1|0|0|0|0
截取最低位 1 | s = -p & p |0|0|0|0|1|0|0|0|0
### 代码

```javascript
/**
 * @param {character[][]} board
 * @return {void} Do not return anything, modify board in-place instead.
 */
var solveSudoku = function (board) {
    new Sudoku(board).solve();
};

class Sudoku {
    constructor(board) {
        this.board = board;
        //行
        this.rows = new Array(9).fill(0);
        //列
        this.columns = new Array(9).fill(0);
        //3x3方格
        this.boxs = Array.from({ length: 3 }, () => new Array(3).fill(0));
        //未填空格
        this.emptyCells = new Set();
    }
    solve() {
        //初始化已知的数字
        for (let i = 0; i < 9; i++) {
            for (let j = 0; j < 9; j++) {
                let num = this.board[i][j];
                if (num !== ".") {
                    //将数字转化为二进制标记
                    //1 -> 0b1, 2 -> 0b10, 3 -> 0b100, 4 -> 0b1000 ...
                    const sign = 1 << (Number(num) - 1);
                    this.rows[i] |= sign;
                    this.columns[j] |= sign;
                    this.boxs[Math.floor(i / 3)][Math.floor(j / 3)] |= sign;
                } else {
                    this.emptyCells.add((i << 4) | j);
                }
            }
        }
        //主逻辑
        return this.fillNext();
    }
    fillNext() {
        let cellInfo = this.getEmptyCell();
        if (cellInfo === null) {
            //没有空格，解题成功
            return true;
        }
        let [i, j, possible] = cellInfo;
        while (possible) {
            //截取其中一个可能性
            const sign = -possible & possible;
            //填入空格
            this.fillCell(i, j, sign);
            //继续下一个填充
            if (this.fillNext()) {
                //填充成功
                return true;
            } else {
                //排除当前数字
                possible ^= sign;
                //清空空格
                this.cleanCell(i, j, sign);
            }
        }
        //穷尽所有可能性，回溯
        return false;
    }
    getEmptyCell() {
        let min = 10;
        let cellInfo = null;
        for (const id of this.emptyCells) {
            const i = id >> 4, j = id & 0b1111;
            const possible = this.getCellPossible(i, j);
            const count = this.countPossible(possible);
            if (min > count) {
                //挑选可能性最少的格子，理论上可减少犯错回溯
                cellInfo = [i, j, possible];
                min = count;
            }
        }
        return cellInfo;
    }
    countPossible(possible) {
        //计算二进制 1 的数量
        let count = 0;
        while (possible) {
            possible &= (possible - 1);
            count++;
        }
        return count;
    }
    fillCell(i, j, sign) {
        //对应位变成1，标记占用
        this.rows[i] |= sign;
        this.columns[j] |= sign;
        this.boxs[Math.floor(i / 3)][Math.floor(j / 3)] |= sign;
        //填入空格
        this.emptyCells.delete((i << 4) | j);
        this.board[i][j] = String(Math.log2(sign) + 1);
    }
    cleanCell(i, j, sign) {
        //对应位变为0，清除占用
        this.rows[i] &= ~sign;
        this.columns[j] &= ~sign;
        this.boxs[Math.floor(i / 3)][Math.floor(j / 3)] &= ~sign;
        //清空格子
        this.emptyCells.add((i << 4) | j)
        this.board[i][j] = ".";
    }
    getCellPossible(i, j) {
        //获取格子可能的取值，二进制1表示可选
        return (this.rows[i] | this.columns[j] | this.boxs[Math.floor(i / 3)][Math.floor(j / 3)]) ^ 0b111111111;
    }
}
```
![截屏2020-06-27 12.54.31.png](https://pic.leetcode-cn.com/30e771b6e18ca3fb825c7a15764112ebfb3c83597469177dbac1709e82dc7e09-%E6%88%AA%E5%B1%8F2020-06-27%2012.54.31.png)