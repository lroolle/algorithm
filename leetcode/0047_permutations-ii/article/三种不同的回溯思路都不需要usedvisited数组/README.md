"+++
title = "0047. Permutations II 三种不同的回溯思路，都不需要used/visited数组 "
author = ["tai-yang-tian"]
date = 2020-06-27T09:49:59+08:00
tags = ["Leetcode", "Algorithms", "C++", "Python"]
categories = ["leetcode"]
draft = false
+++

# 三种不同的回溯思路，都不需要used/visited数组

> [0047. Permutations II](https://leetcode-cn.com/problems/permutations-ii/)
> [三种不同的回溯思路，都不需要used/visited数组](https://leetcode-cn.com/problems/permutations-ii/solution/shuang-bai-si-lu-chao-ji-jian-ji-by-tai-yang-tian/) by [tai-yang-tian](https://leetcode-cn.com/u/tai-yang-tian/)

# 解题思路
虽然都采用回溯算法，但是，可以按照不同的树生成全排列。同时，不同的树的剪枝方法都不一样，如果选取的树足够好，剪枝的方法会非常方便，不需要used/visited数组。（其实我觉得大家被46题的官解坑了，如果用那个官解的树，必须要使用used/visited数组。）
## 思路1
这个方法是我最早想到的一种思路，结果是双百。
![image.png](https://pic.leetcode-cn.com/0c918f199101f9ec0f03f03f88e2fc611ddcdbf932917bb5181c9946267cbcd4-image.png)
以[1, 2, 3]为例子，按照下面的树回溯出整个全排列：
[|1, 2, 3] -> [1|, 2, 3]
[1|, 2, 3] -> [1, 2|, 3], [2, 1|, 3]
[1, 2|, 3] -> [1, 2, 3|], [3, 2, 1|], [1, 3, 2|]
[2, 1|, 3] -> [2, 1, 3|], [3, 1, 2|], [2, 3, 1|]
ps：以[1, 2|, 3] -> [1, 2, 3|], [3, 2, 1|], [1, 3, 2|]为例，'|'往后移动一位，得到[1, 2, 3|]；[1, 2, 3|]中的3和1位置交换，得到[3, 2, 1|]；[1, 2, 3|]中的3和2交换位置，得到[1, 3, 2|]。
按照这个思路，可以写出46题的代码：
### 思路1对应的46题代码：
```cpp
class Solution {
public:
    vector<vector<int>> permuteUnique(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        int current_len = 0;
        vector<vector<int>> res;
        function<void(void)> backtrack = [&](void){
            if(current_len == nums.size()){
                res.push_back(nums);
                return;
            }
            current_len++;
            backtrack();
            current_len--;
            for(int i = 0; i < current_len; i++){
                swap(nums[i], nums[current_len]);
                current_len++;
                backtrack();
                current_len--;
                swap(nums[i], nums[current_len]);
            }
        };
        backtrack();
        return res;
    }
};
```
```python
class Solution(object):
    def permuteUnique(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        def backtrack(current_len = 0):
            if current_len == n:  
                res.append(nums[:])
                return
            backtrack(current_len + 1)
            for i in range(0, current_len):
                nums[current_len], nums[i] = nums[i], nums[current_len]
                backtrack(current_len + 1)
                nums[current_len], nums[i] = nums[i], nums[current_len]
        
        nums.sort()
        n = len(nums)
        res = []
        backtrack()
        return res
```
另外，考虑到存在重复元素，对于47题，上面的代码还需要额外添加一个判断句进行剪枝。47题代码如下：
### 思路1对应的47题代码：
```cpp
class Solution {
public:
    vector<vector<int>> permuteUnique(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        int current_len = 0;
        vector<vector<int>> res;
        function<void(void)> backtrack = [&](void){
            if(current_len == nums.size()){
                res.push_back(nums);
                return;
            }
            current_len++;
            backtrack();
            current_len--;
            for(int i = 0; i < current_len; i++){
                if(nums[i] == nums[current_len]) break;
                swap(nums[i], nums[current_len]);
                current_len++;
                backtrack();
                current_len--;
                swap(nums[i], nums[current_len]);
            }
        };
        backtrack();
        return res;
    }
};
```
```python
class Solution(object):
    def permuteUnique(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        def backtrack(current_len = 0):
            if current_len == n:  
                res.append(nums[:])
                return
            backtrack(current_len + 1)
            for i in range(0, current_len):
                if(nums[i] == nums[current_len]):
                    return
                nums[current_len], nums[i] = nums[i], nums[current_len]
                backtrack(current_len + 1)
                nums[current_len], nums[i] = nums[i], nums[current_len]
        
        nums.sort()
        n = len(nums)
        res = []
        backtrack()
        return res
```
为什么这么剪枝的原因放在附录，有兴趣的可以看。
## 思路2
以[1, 2, 3]为例，按照下面的树回溯出整个全排列:
[|1, 2, 3] -> [1|, 2, 3]
[1|, 2, 3] -> [1, 2|, 3], [2, 1|, 3]
[1, 2|, 3] -> [1, 2, 3|], [1, 3, 2|], [3, 1, 2|]
ps：以[1, 2|, 3] -> [1, 2, 3|], [1, 3, 2|], [3, 2, 1|]为例，'|'往后移动一位，得到[1, 2, 3|]；[1, 2, 3|]中的3和2位置交换，得到[1, 3, 2|]；[1, 3, 2|]中的1和3交换位置，得到[3, 1, 2|]。
按照这个思路，可以写出46题的代码：
### 思路2对应的46题代码：
```cpp
class Solution {
public:
    vector<vector<int>> permuteUnique(vector<int>& nums) {
        int current_len = 0;
        vector<vector<int>> res;
        function<void(void)> backtrack = [&](void){
            if(current_len == nums.size()){
                res.push_back(nums);
                return;
            }
            current_len++;
            backtrack();
            current_len--;
            int i = current_len;
            for(; i > 0; i--){
                swap(nums[i], nums[i-1]);
                current_len++;
                backtrack();
                current_len--;
            }
            for(i++; i <= current_len; i++){
                swap(nums[i], nums[i-1]);
            }
        };
        sort(nums.begin(), nums.end());
        backtrack();
        return res;
    }
};
```
类似的，加上一行判断句就可以剪枝解决47题，原因也放在附录。
### 思路2对应的47题代码：
```cpp
class Solution {
public:
    vector<vector<int>> permuteUnique(vector<int>& nums) {
        int current_len = 0;
        vector<vector<int>> res;
        function<void(void)> backtrack = [&](void){
            if(current_len == nums.size()){
                res.push_back(nums);
                return;
            }
            current_len++;
            backtrack();
            current_len--;
            int i = current_len;
            for(; i > 0; i--){
                if(nums[i] == nums[i-1]) break;
                swap(nums[i], nums[i-1]);
                current_len++;
                backtrack();
                current_len--;
            }
            for(i++; i <= current_len; i++){
                swap(nums[i], nums[i-1]);
            }
        };
        sort(nums.begin(), nums.end());
        backtrack();
        return res;
    }
};
```
## 思路3
以[1, 2, 3]为例，按照下面的树回溯出整个全排列:
[|1, 2, 3] -> [1|, 2, 3], [2|, 1, 3], [3|, 1, 2]
[1|, 2, 3] -> [1, 2|, 3], [1, 3|, 2]
[2|, 1, 3] -> [2, 1|, 3], [2, 3|, 1]
[3|, 1, 2] -> [3, 1|, 2], [3, 2|, 1]
[1, 2|, 3] -> [1, 2, 3|]
[1, 3|, 2] -> [1, 3, 2|]
...

以[|1, 2, 3] -> [1|, 2, 3], [2|, 1, 3], [3|, 1, 2]为例子，将'|'向右移动一位，得到[1|, 2, 3]；把[1|, 2, 3]中的1和2交换，得到[2|, 1, 3]；把[2|, 1, 3]中的2和3交换，得到[3|, 1, 2]。
对应的，可以写出46题的代码：
### 思路3对应的46题代码：
```cpp
class Solution {
public:
    vector<vector<int>> permuteUnique(vector<int>& nums) {
        int current_len = 0;
        vector<vector<int>> res;
        function<void(void)> backtrack = [&](void){
            if(current_len == nums.size()){
                res.push_back(nums);
                return;
            }
            current_len++;
            backtrack();
            current_len--;
            int i = current_len + 1;
            for(; i < nums.size(); i++){
                swap(nums[current_len], nums[i]);
                current_len++;
                backtrack();
                current_len--;
            }
            for(i--; i > current_len; i--){
                swap(nums[current_len], nums[i]);
            }
        };
        sort(nums.begin(), nums.end());
        backtrack();
        return res;
    }
};
```
类似的，对于47题，只要加上额外的判断条件即可，为什么这么剪枝的原因写在附录。
### 思路3对应的47题代码：
```cpp
class Solution {
public:
    vector<vector<int>> permuteUnique(vector<int>& nums) {
        int current_len = 0;
        vector<vector<int>> res;
        function<void(void)> backtrack = [&](void){
            if(current_len == nums.size()){
                res.push_back(nums);
                return;
            }
            current_len++;
            backtrack();
            current_len--;
            int i = current_len + 1;
            for(; i < nums.size(); i++){
                if(nums[current_len] == nums[i]) continue;
                swap(nums[current_len], nums[i]);
                current_len++;
                backtrack();
                current_len--;
            }
            for(i--; i > current_len; i--){
                swap(nums[current_len], nums[i]);
            }
        };
        sort(nums.begin(), nums.end());
        backtrack();
        return res;
    }
};
```
## 附录
思路3是
[@18838006233](/u/18838006233/)的算法。
值得一提的一些细节是，思路1和思路2中sort其实不是必要的，但是加上sort速度会更快一些，原因还不太清楚。思路3中，sort函数是必须的，并且思路3生成的其实是字典序。
下面是关于三个思路的剪枝条件，可能说的不是很清楚，凑合看吧。
### 思路1剪枝条件
考虑[1,1,2,2,2]的一个排列[2,2,2,1,1]。如果不加上面那一行判断，那么[2,2,2,1,1]可能由：[**1**,2,2,1,**2**]，[2,**1**,2,1,**2**]，[2,2,**1**,1,**2**]生成（粗体表示两个元素交换），这就会产生重复。仔细观察[**1**,2,2,1,**2**]，[2,**1**,2,1,**2**]，[2,2,**1**,1,**2**]，我们发现，[**1**,2,2,1,**2**]中的**1**前面没有2，[2,**1**,2,1,**2**]，[2,2,**1**,1,**2**]中的**1**前面有2，因此，只要前面出现过2，立马结束for循环，这就可以剪掉[2,**1**,2,1,**2**]，[2,2,**1**,1,**2**]生成[2,2,2,1,1]的可能。也就是只要添加上面一行代码就可以解决去重。
### 思路2剪枝条件
与思路1类似
### 思路3剪枝条件
只要注意，每次调用backtrack函数，nums[current_len]到nums[nums.size() - 1]都是一个有序的数列。