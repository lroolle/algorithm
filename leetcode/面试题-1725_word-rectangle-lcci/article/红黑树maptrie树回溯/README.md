"+++
title = "面试题 17.25. Word Rectangle LCCI 红黑树map+Trie树+回溯 "
author = ["liousvious"]
date = 2020-08-04T12:08:51+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# 红黑树map+Trie树+回溯

> [面试题 17.25. Word Rectangle LCCI](https://leetcode-cn.com/problems/word-rectangle-lcci/)
> [红黑树map+Trie树+回溯](https://leetcode-cn.com/problems/word-rectangle-lcci/solution/hong-hei-shu-maptrieshu-hui-su-by-liousvious/) by [liousvious](https://leetcode-cn.com/u/liousvious/)

### 解题思路

### 代码

```cpp
/* 将所有单词插入字典树
*  将单词按单词长度进行分组，即构建hashmap
*  从单词长度最长的组开始遍历，对每组单词进行DFS搜索
*  利用Trie树检查是否合法，不合法进行回溯
*  代码需要进行优化
*/
// 构造Trie树，将单词插入到Trie树中
class Trie {
public:
    Trie() {}
    bool isEnd = false;
    Trie* next[26] = {nullptr};
    void insert(string& s) {
        Trie* cur = this;
        for (int i = 0; i < s.size(); ++i) {
            if (!cur->next[s[i] - 'a'])
                cur->next[s[i] - 'a'] = new Trie();
            cur = cur->next[s[i] - 'a'];
        }
        cur->isEnd = true;
    }
};

class Solution {
private:
    Trie* t;
    vector<string> ans;
    vector<string> temp;
public:
    vector<string> maxRectangle(vector<string>& words) {
        t = new Trie();
        map<int, vector<string>> m;
        int maxLen = 0, maxArea = 0, area;
        // 将单词表中的所有单词构成字典树Trie
        for (auto& w : words) {
            t->insert(w); // 单词插入Trie
            m[w.size()].push_back(w); // 单词按长度分组
            maxLen = max(maxLen, int(w.size())); // 单词最大的长度
        }

        for (map<int, vector<string>>::reverse_iterator it = m.rbegin(); it != m.rend(); it++) {
            // 反向遍历，按长度大的开始
            if (maxArea / (it->first) >= maxLen) break; // 最长的单词长度 * 宽小于最大面积时，跳出循环。
            temp.clear();
            area = 0;
            backTracking(it->second, maxArea, maxLen, area);
        }
        return ans;
    }

    void backTracking(vector<string>& word, int& maxArea, int& maxLen, int& area) {
        if (word[0].size() * maxLen <= maxArea) return; //找到的面积达到极限需要退出
        vector<bool> res;
        for (int i = 0; i < word.size(); i++) {
            temp.push_back(word[i]);
            // 检查是否合法
            res = isGood(temp); // 将单词列表中的单词按长度由大到小进行判断，对长度相同的单词按层叠加，当层与层之间的单词列构成的
            // 单词在单词表中时，结束且更新；否则返回错进行下一步的回溯迭代。
            if (res[0]) { // 可以往下加单词
                area = temp.size() * temp[0].size();
                if (res[1] && area > maxArea) { //都是结束单词
                    // 更新最大值
                    maxArea = area;
                    ans = temp;
                }
                backTracking(word, maxArea, maxLen, area); //单词没有结束时选择回溯迭代
            }
            temp.pop_back(); // 单词不满足要求将该单词移出结束回溯
        }
    }

    vector<bool> isGood(vector<string>& temp) {
        Trie* cur;
        bool allowed = true;
        int n = temp[0].size();
        for (int j = 0; j < n; ++j) { //将字母排成一列，对于每个字母的每一位进行逐项遍历，表示列
            cur = t;
            // 固定列，按行进行查找
            for (int i = 0; i < temp.size(); ++i) { // 对每个字母进行遍历，表示矩阵的行
                if (!cur->next[temp[i][j] - 'a']) // 不在Trie树上，返回错
                    return {false, false};
                cur = cur->next[temp[i][j] - 'a'];
            }
            allowed &= cur->isEnd; //用于判断每列字母组成的单词是否已经达到最后，即单词结束
        }
        return {true, allowed};
    }
};
```