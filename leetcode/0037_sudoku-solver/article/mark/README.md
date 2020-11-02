"+++
title = "0037. Sudoku Solver mark "
author = ["gOX1JKuQAs"]
date = 2020-09-01T03:16:19+08:00
tags = ["Leetcode", "Algorithms", "PHP"]
categories = ["leetcode"]
draft = false
+++

# mark

> [0037. Sudoku Solver](https://leetcode-cn.com/problems/sudoku-solver/)
> [mark](https://leetcode-cn.com/problems/sudoku-solver/solution/mark-by-gox1jkuqas-8/) by [gOX1JKuQAs](https://leetcode-cn.com/u/gOX1JKuQAs/)

### 解题思路
写的很乱，好歹自己做的，留个印记
执行用时：56 ms, 在所有 PHP 提交中击败了85.71%的用户
内存消耗：15.6 MB, 在所有 PHP 提交中击败了5.26%的用户
思路：数独未完成时，先把点的位置替换成可能出现的数字的数组，再回溯循环所有可能找出数独的解

### 代码

```php
class Solution {

    /**
     * @param String[][] $board
     * @return NULL
     */
    function solveSudoku(&$board) {
        $this->prev = '';
        $this->res = [];
        $break = false;
        while (!$this->is_ok($board)) {
           if(serialize($this->prev) == serialize($board)){
                $break = true;
                break;
           }
           $this->prev = $board;
           $board = $this->helper($board);
        }
        if ($break) {
            $this->dfs($board,0,0);
            $board = $this->res;
        }
        return $board;
    }

    function is_ok($board){
        $arr = [];
        for ($i=0; $i < 9; $i++) { 
            $map1 = $map2 = [];
            for ($j=0; $j < 9; $j++) { 
                if(!is_numeric($board[$i][$j]) || !is_numeric($board[$j][$i]))
                    return false;
                if(isset($map1[$board[$i][$j]]))
                    return false;
                if(isset($map2[$board[$j][$i]]))
                    return false;
                $map1[$board[$i][$j]] = $board[$i][$j];
                $map2[$board[$j][$i]] = $board[$j][$i];
                $arr[floor($j/3)+floor($i/3)*3][($i%3)*3+($j%3)] = $board[$i][$j];
            }
        }
        for ($i=0; $i < 9; $i++) { 
            $map3=[];
            for ($j=0; $j < 9; $j++) {
                if (isset($map3[$board[$i][$j]]))
                    return false;
                $map3[$board[$i][$j]] = $board[$i][$j];
            }
        }
        return true;
    }

    function numberIsNull($board){
        for ($i=0; $i < 9; $i++) { 
            for ($j=0; $j < 9; $j++) { 
                if(empty($board[$i][$j]))
                    return true;
            }
        }
        return false;
    }

    function helper($board){
        $arr=[];
        $map = [
            1=>1,2=>2,3=>3,4=>4,5=>5,6=>6,7=>7,8=>8,9=>9
        ];
        for ($j=0; $j < 9; $j++) { 
            $map1 = $map2 = $map;
            for ($i=0; $i < 9; $i++) {
                if(is_numeric($board[$j][$i])) unset($map1[$board[$j][$i]]);   
            }
            for ($i=0; $i < 9; $i++) { 
                if(is_array($board[$j][$i])){
                    $inter = array_intersect($board[$j][$i],$map1);
                    $board[$j][$i]=count($inter)==1? ((string) end($inter)):$inter;
                } 
                if($board[$j][$i]==='.') $board[$j][$i]=$map1;
            }
            for ($i=0; $i < 9; $i++) {
                if(is_numeric($board[$i][$j])) unset($map2[$board[$i][$j]]);   
            }
            for ($i=0; $i < 9; $i++) { 
                if(is_array($board[$i][$j])){
                    $inter=array_intersect($board[$i][$j],$map2);
                    $board[$i][$j]= count($inter)==1?((string) end($inter)):$inter;
                }
                if($board[$i][$j]==='.') $board[$i][$j]=$map2;
            }
            for ($i=0; $i < 9; $i++) { 
                $arr[floor($i/3)+floor($j/3)*3][($j%3)*3+($i%3)] =& $board[$j][$i];
            }
        }
        for ($j=0; $j < 9; $j++) { 
            $map3 = $map;
            for ($i=0; $i < 9; $i++) {
                if(is_numeric($arr[$j][$i]))  unset($map3[$arr[$j][$i]]); 
            }
            for ($i=0; $i < 9; $i++) { 
                if(!is_numeric($arr[$j][$i]))  {
                    $arr[$j][$i] = is_array($arr[$j][$i])?array_intersect($arr[$j][$i],$map3):$map3;
                    $board[floor($j/3)*3+floor($i/3)][($j%3)*3+($i%3)] = count($arr[$j][$i])==1?((string) end($arr[$j][$i])) :$arr[$j][$i];
                }
            }
        }
        return $board;
    }

    function dfs($board,$m,$n){
        if(!empty($this->res)) return;
        if($m>8 || $n>8) return;
        if ($this->numberIsNull($board)) return;
        if($this->is_ok($board)){
            $this->res = $board;
            return;
        }
        $start = $board;
        if (is_array($start[$m][$n])) {
            $tmpArr = $start[$m][$n];
            foreach ($tmpArr as $key => $value) {
                $board = $start;
                $board[$m][$n] = (string) $value;
                $prev = '';
                $break = false;
                while (!$this->is_ok($board)) {
                   if(serialize($prev) == serialize($board)){
                        $break = true;
                        break;
                   }
                   $prev = $board;
                   $board = $this->helper($board);
                }
                if ($break) {
                    if ($this->numberIsNull($board)) continue;                  
                    $this->dfs($board,($n==8 ? $m+1 : $m),($n==8 ? 0 : $n+1));
                }else{
                    $this->res = $board; 
                    return;                  
                }                
            }
        }else{
            $this->dfs($board,($n==8 ? $m+1 : $m),($n==8 ? 0 : $n+1));
        }
    }
}
```