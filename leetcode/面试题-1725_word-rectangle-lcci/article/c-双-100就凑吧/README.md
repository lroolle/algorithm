"+++
title = "面试题 17.25. Word Rectangle LCCI C++ 双 100，就凑吧 "
author = ["sk-k"]
date = 2020-07-11T14:29:25+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# C++ 双 100，就凑吧

> [面试题 17.25. Word Rectangle LCCI](https://leetcode-cn.com/problems/word-rectangle-lcci/)
> [C++ 双 100，就凑吧](https://leetcode-cn.com/problems/word-rectangle-lcci/solution/c-shuang-100jiu-cou-ba-by-sk-k/) by [sk-k](https://leetcode-cn.com/u/sk-k/)

```C++
struct TrieNode {
  TrieNode(TrieNode* p = nullptr):parent(p) {}
  TrieNode* children[26]{};
  TrieNode* parent;
  bool word = false;
};
class Trie {
public:
    TrieNode root;
    Trie() {
      root.parent = &root;
    }
    void insert(const string& word) {
      auto cur = &root;
      for (auto n : word) {
        if (cur->children[n - 'a'] == nullptr) {
          cur->children[n - 'a'] = new TrieNode(cur);
        }
        cur = cur->children[n - 'a'];
      }
      cur->word = true;
    }
};

class Solution {
      vector<string_view> realres;
      int realsz = 0;
      Trie tree;
public:
    void dfs(vector<string_view>& res, vector<string_view>& w, vector<TrieNode*>& nodes) {
      auto l = w[0].size();
      //cout << l << " " << nodes.size() << endl;
      if (res.size()) {
        auto n = 0;
        for (int i = 0; i != l; ++i) {
          auto tmp = nodes[i];
          nodes[i] = nodes[i]->children[res.back()[i] - 'a'];
          if (!nodes[i]) {
            for (int k = 0; k != i; ++k) {
              nodes[k] = nodes[k]->parent;
            }
            nodes[i] = tmp;
            return;
          }
          if (nodes[i]->word) ++n;
        }
        if (n == l) {
          auto sz = res.size() * res[0].size();
          if (sz > realsz) {
            realres = res;
            realsz = sz;
          }
        }
      }
      for (int n = 0; n != w.size(); ++n) {
        res.emplace_back(w[n]);
        //cout << w[n].size() << endl;
        dfs(res, w, nodes);
        res.pop_back();
      }
      for (auto& n : nodes) {
        n = n->parent;
      }
    }
    vector<string> maxRectangle(vector<string>& words) {
      vector<vector<string_view>> w(101);
      for (auto& n : words) {
        w[n.size()].emplace_back(n);
        tree.insert(n);
      }
      vector<string_view> res;
      auto sz = 0;
      for (auto iter = w.rbegin(); iter != w.rend(); ++iter) {
        if (iter->empty() || (*iter)[0].size() * iter->size() <= realsz) continue;
        vector<string_view> tmp;
        vector<TrieNode*> nodes((*iter)[0].size(), &tree.root);
        dfs(tmp, *iter, nodes);
      }
      return {begin(realres), end(realres)};
    }
};
```
