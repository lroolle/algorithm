"+++
title = "0037. Sudoku Solver 递归回溯Golang+Bitmap实现 "
author = ["epic-hawkingqdf"]
date = 2020-08-24T15:11:24+08:00
tags = ["Leetcode", "Algorithms", "Go"]
categories = ["leetcode"]
draft = false
+++

# 递归回溯Golang+Bitmap实现

> [0037. Sudoku Solver](https://leetcode-cn.com/problems/sudoku-solver/)
> [递归回溯Golang+Bitmap实现](https://leetcode-cn.com/problems/sudoku-solver/solution/di-gui-hui-su-golangbitmapshi-xian-by-epic-hawking/) by [epic-hawkingqdf](https://leetcode-cn.com/u/epic-hawkingqdf/)

### 解题思路
此处撰写解题思路
![NL%%{WT8PT3\[%625\[M}7\[LY.png](https://pic.leetcode-cn.com/1598281875-buNuLx-NL%25%25%7BWT8PT3%5B%25625%5BM%7D7%5BLY.png)

### 代码

```golang
func leftmove (v int) int {
    return 0x1 << v
}

func byte2int(b byte) int {
    return int(b-'0')
}

func int2byte(i int) byte {
    return byte('0' + i)
}

var row, col, box [9]int // 9行 9列 9格
var success bool // 完成搜索状态标记

// 搜索回溯
func SearchAndRollback(step int, board [][]byte) {
    // 走完81格，完成搜索
    if step >= 81 {
        success = true
        return
    }

    var i, j = step/9, step%9

    // 非空格，直接查找下一步
    if board[i][j] != '.' {
        SearchAndRollback(step+1, board)
        return
    }

    // 1-9 总有一个数满足条件
    var idx = i/3*3 + j/3
    for k := 1; k < 10; k++ {
        var v = leftmove(k)
        // 筛除不符合条件的数，加速搜索
        if row[i] & v > 0 || col[j] & v > 0 || box[idx] & v > 0 {
            continue
        }

        // 标记
        board[i][j] = int2byte(k)
        row[i] |= v
        col[j] |= v
        box[idx] |= v

        SearchAndRollback(step+1, board)
        if success {
            break // 搜索结束，快速终止
        }

        // 不满足条件，回溯
        row[i] ^= v
        col[j] ^= v
        box[idx] ^= v
        board[i][j] = '.'
    }
}

func solveSudoku(board [][]byte)  {
    row, col, box, success = [9]int{}, [9]int{}, [9]int{}, false // 状态重置

    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            // 标记所有已给出的数字，加速搜索
            if board[i][j] != '.' {
                var idx = i/3*3 + j/3
                var v = leftmove(byte2int(board[i][j]))
                row[i] |= v
                col[j] |= v
                box[idx] |= v
            }
        }
    }

    // 从第一格开始查找
    SearchAndRollback(0, board)
}
```