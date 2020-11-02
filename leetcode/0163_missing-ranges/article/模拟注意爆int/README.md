"+++
title = "0163. Missing Ranges 模拟、注意爆int "
author = ["wangdh15"]
date = 2020-07-09T17:58:02+08:00
tags = ["Leetcode", "Algorithms", "C++", "模拟"]
categories = ["leetcode"]
draft = false
+++

# 模拟、注意爆int

> [0163. Missing Ranges](https://leetcode-cn.com/problems/missing-ranges/)
> [模拟、注意爆int](https://leetcode-cn.com/problems/missing-ranges/solution/mo-ni-zhu-yi-bao-int-by-acw_wangdh15/) by [wangdh15](https://leetcode-cn.com/u/wangdh15/)

### 解题思路

模拟即可。注意转longlong.....

[个人博客](http://wangdh15.github.io)

### 代码

```cpp
class Solution {
public:
    vector<string> findMissingRanges(vector<int>& nums, int lower, int upper) {
        int n = nums.size();
        vector<string> ans;
        if (lower > upper) swap(lower, upper);
        if (n == 0) {
            if (lower != upper) return vector<string>({to_string(lower) + "->" + to_string(upper)});
            else return vector<string>({to_string(lower)});
        }
        if (lower > nums[n-1]) return vector<string>({to_string(lower) + "->" + to_string(upper)});
        if (upper < nums[0]) return vector<string>({to_string(lower) + "->" + to_string(upper)});
        
        int l = 0, r = n - 1;
        while (l < r) {
            int mid = l + r >> 1;
            if (nums[mid] < lower) l = mid + 1;
            else r = mid;
        }
        // cout << l << endl;
        if ((long long)nums[l] - lower >= 2) ans.push_back(to_string(lower) + "->" + to_string((long long)nums[l] - 1));
        else if ((long long)nums[l] - lower == 1) ans.push_back(to_string(lower));
        // cout << ans.size() << endl;
        int last = nums[l];
        for (int i = l + 1; i < n; i ++) {
            if (nums[i] > upper) break;
            if (nums[i] > (long long)last + 1) {
                if ((long long)nums[i] - last == 2) ans.push_back(to_string((long long)last + 1));
                else ans.push_back(to_string((long long)last + 1) + "->" + to_string((long long)nums[i] - 1));
            }
            last = nums[i];
        }
        if (last != upper) {
            if ((long long)upper - last == 1) ans.push_back(to_string((long long)last + 1));
            else ans.push_back(to_string((long long)last + 1) + "->" + to_string(upper));
        }
        return ans;
        
    }
};
```