"+++
title = "0037. Sudoku Solver 解数独 "
author = ["LeetCode"]
date = 2019-05-17T11:37:08+08:00
tags = ["Leetcode", "Algorithms", "Backtracking", "Java", "Python"]
categories = ["leetcode"]
draft = false
+++

# 解数独

> [0037. Sudoku Solver](https://leetcode-cn.com/problems/sudoku-solver/) > [解数独](https://leetcode-cn.com/problems/sudoku-solver/solution/jie-shu-du-by-leetcode/) by [LeetCode](https://leetcode-cn.com/u/leetcode/)

#### 方法 0：蛮力法

首先的想法是通过蛮力法来生成所有可能用`1` 到 `9`填充空白格的解，
并且检查合法从而保留解。这意味着共有 $9^{81}$ 个操作需要进行。
其中 $9$ 是可行的数字个数，$81$ 是需要填充的格子数目。
因此我们必须考虑进一步优化。

---

#### 方法 1：回溯法

**使用的概念**

了解两个编程概念会对接下来的分析有帮助。

> 第一个叫做 _约束编程_。

基本的意思是在放置每个数字时都设置约束。在数独上放置一个数字后立即
排除当前 _行_， _列_ 和 _子方块_ 对该数字的使用。这会传播 _约束条件_ 并有利于减少需要考虑组合的个数。

![37_const3.png](https://pic.leetcode-cn.com/ce5d93149307f8c4b5e3bf5d0924985f39721905ccc6537cbdf92030ad79b789-image.png){:width="500"}

> 第二个叫做 _回溯_。

让我们想象一下已经成功放置了几个数字
在数独上。
但是该组合不是最优的并且不能继续放置数字了。该怎么办？ _回溯_。
意思是回退，来改变之前放置的数字并且继续尝试。如果还是不行，再次 _回溯_。

![37_backtrack2.png](https://pic.leetcode-cn.com/575afd37ae93cd906adc1cd46e939bfc951af5aa1fe411c693d09ce5140f8822-image.png){:width="500"}

**如何枚举子方块**

> 一种枚举子方块的提示是：

使用 `方块索引= (行 / 3) * 3 + 列 / 3`
其中 `/` 表示整数除法。

![36_boxes_2.png](https://pic.leetcode-cn.com/5a7856c3c2a2185600b7cb5cd3fd50101281af7391a70a63293d82d62873aadd-36_boxes_2.png){:width="370"}

**算法**

现在准备好写回溯函数了
`backtrack(row = 0, col = 0)`。

- 从最左上角的方格开始 `row = 0, col = 0`。直到到达一个空方格。
- 从`1` 到 `9` 迭代循环数组，尝试放置数字 `d` 进入 `(row, col)` 的格子。

  - 如果数字 `d` 还没有出现在当前行，列和子方块中：
    - 将 `d` 放入 `(row, col)` 格子中。
    - 记录下 `d` 已经出现在当前行，列和子方块中。
    - 如果这是最后一个格子`row == 8, col == 8` ：
      - 意味着已经找出了数独的解。
    - 否则
      - 放置接下来的数字。
    - 如果数独的解还没找到：
      将最后的数从 `(row, col)` 移除。

**代码**

```Java
class Solution {
  // box size
  int n = 3;
  // row size
  int N = n * n;

  int [][] rows = new int[N][N + 1];
  int [][] columns = new int[N][N + 1];
  int [][] boxes = new int[N][N + 1];

  char[][] board;

  boolean sudokuSolved = false;

  public boolean couldPlace(int d, int row, int col) {
    /*
    Check if one could place a number d in (row, col) cell
    */
    int idx = (row / n ) * n + col / n;
    return rows[row][d] + columns[col][d] + boxes[idx][d] == 0;
  }

  public void placeNumber(int d, int row, int col) {
    /*
    Place a number d in (row, col) cell
    */
    int idx = (row / n ) * n + col / n;

    rows[row][d]++;
    columns[col][d]++;
    boxes[idx][d]++;
    board[row][col] = (char)(d + '0');
  }

  public void removeNumber(int d, int row, int col) {
    /*
    Remove a number which didn't lead to a solution
    */
    int idx = (row / n ) * n + col / n;
    rows[row][d]--;
    columns[col][d]--;
    boxes[idx][d]--;
    board[row][col] = '.';
  }

  public void placeNextNumbers(int row, int col) {
    /*
    Call backtrack function in recursion
    to continue to place numbers
    till the moment we have a solution
    */
    // if we're in the last cell
    // that means we have the solution
    if ((col == N - 1) && (row == N - 1)) {
      sudokuSolved = true;
    }
    // if not yet
    else {
      // if we're in the end of the row
      // go to the next row
      if (col == N - 1) backtrack(row + 1, 0);
        // go to the next column
      else backtrack(row, col + 1);
    }
  }

  public void backtrack(int row, int col) {
    /*
    Backtracking
    */
    // if the cell is empty
    if (board[row][col] == '.') {
      // iterate over all numbers from 1 to 9
      for (int d = 1; d < 10; d++) {
        if (couldPlace(d, row, col)) {
          placeNumber(d, row, col);
          placeNextNumbers(row, col);
          // if sudoku is solved, there is no need to backtrack
          // since the single unique solution is promised
          if (!sudokuSolved) removeNumber(d, row, col);
        }
      }
    }
    else placeNextNumbers(row, col);
  }

  public void solveSudoku(char[][] board) {
    this.board = board;

    // init rows, columns and boxes
    for (int i = 0; i < N; i++) {
      for (int j = 0; j < N; j++) {
        char num = board[i][j];
        if (num != '.') {
          int d = Character.getNumericValue(num);
          placeNumber(d, i, j);
        }
      }
    }
    backtrack(0, 0);
  }
}
```

```Python
from collections import defaultdict
class Solution:
    def solveSudoku(self, board):
        """
        :type board: List[List[str]]
        :rtype: void Do not return anything, modify board in-place instead.
        """
        def could_place(d, row, col):
            """
            Check if one could place a number d in (row, col) cell
            """
            return not (d in rows[row] or d in columns[col] or \
                    d in boxes[box_index(row, col)])

        def place_number(d, row, col):
            """
            Place a number d in (row, col) cell
            """
            rows[row][d] += 1
            columns[col][d] += 1
            boxes[box_index(row, col)][d] += 1
            board[row][col] = str(d)

        def remove_number(d, row, col):
            """
            Remove a number which didn't lead
            to a solution
            """
            del rows[row][d]
            del columns[col][d]
            del boxes[box_index(row, col)][d]
            board[row][col] = '.'

        def place_next_numbers(row, col):
            """
            Call backtrack function in recursion
            to continue to place numbers
            till the moment we have a solution
            """
            # if we're in the last cell
            # that means we have the solution
            if col == N - 1 and row == N - 1:
                nonlocal sudoku_solved
                sudoku_solved = True
            #if not yet
            else:
                # if we're in the end of the row
                # go to the next row
                if col == N - 1:
                    backtrack(row + 1, 0)
                # go to the next column
                else:
                    backtrack(row, col + 1)


        def backtrack(row = 0, col = 0):
            """
            Backtracking
            """
            # if the cell is empty
            if board[row][col] == '.':
                # iterate over all numbers from 1 to 9
                for d in range(1, 10):
                    if could_place(d, row, col):
                        place_number(d, row, col)
                        place_next_numbers(row, col)
                        # if sudoku is solved, there is no need to backtrack
                        # since the single unique solution is promised
                        if not sudoku_solved:
                            remove_number(d, row, col)
            else:
                place_next_numbers(row, col)

        # box size
        n = 3
        # row size
        N = n * n
        # lambda function to compute box index
        box_index = lambda row, col: (row // n ) * n + col // n

        # init rows, columns and boxes
        rows = [defaultdict(int) for i in range(N)]
        columns = [defaultdict(int) for i in range(N)]
        boxes = [defaultdict(int) for i in range(N)]
        for i in range(N):
            for j in range(N):
                if board[i][j] != '.':
                    d = int(board[i][j])
                    place_number(d, i, j)

        sudoku_solved = False
        backtrack()
```

**复杂性分析**

- 这里的时间复杂性是常数由于数独的大小是固定的，因此没有 N 变量来衡量。
  但是我们可以计算需要操作的次数：$(9!)^9$ 。
  我们考虑一行，即不多于 $9$ 个格子需要填。
  第一个格子的数字不会多于 $9$ 种情况，
  两个格子不会多于 $9 \times 8$ 种情况，
  三个格子不会多于 $9 \times 8 \times 7$ 种情况等等。
  总之一行可能的情况不会多于 $9!$ 种可能，
  所有行不会多于 $(9!)^9$ 种情况。比较一下：

  - $9^{81}$ = 196627050475552913618075908526912116283103450944214766927315415537966391196809
    为蛮力法，
  - 而 $(9!)^{9}$ = 109110688415571316480344899355894085582848000000000
    为回溯法，
    即数字的操作次数减少了 $10^{27}$ 倍！

- 空间复杂性：数独大小固定，空间用来存储数独，行，列和子方块的结构，每个有 `81` 个元素。
