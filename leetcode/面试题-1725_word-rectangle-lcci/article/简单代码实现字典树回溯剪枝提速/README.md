"+++
title = "面试题 17.25. Word Rectangle LCCI 简单代码实现字典树+回溯，剪枝提速 "
author = ["tian-ye"]
date = 2020-05-02T03:42:18+08:00
tags = ["Leetcode", "Algorithms", "Java", "Trie", "Backtracking"]
categories = ["leetcode"]
draft = false
+++

# 简单代码实现字典树+回溯，剪枝提速

> [面试题 17.25. Word Rectangle LCCI](https://leetcode-cn.com/problems/word-rectangle-lcci/)
> [简单代码实现字典树+回溯，剪枝提速](https://leetcode-cn.com/problems/word-rectangle-lcci/solution/jian-dan-dai-ma-shi-xian-zi-dian-shu-hui-su-jian-z/) by [tian-ye](https://leetcode-cn.com/u/tian-ye/)

回溯算法通常有一个“套路”，经典问题有[全排列](https://leetcode-cn.com/problems/permutations/)和[八皇后](https://leetcode-cn.com/problems/eight-queens-lcci/)。用伪代码表示出来：
```
function backTrack(供选择的列表，路径){
    if(终止条件){
        退出;
    }
    做选择;  //从列表选择一项加入路径
    backTrack(列表，路径);
    撤销选择;
}
```
这题的思路就是通过回溯来构成单词矩阵，构造字典树用来快速判断形成的单词是否在清单里。

现在这么做已经没问题了，但提交上去会有超时，需要给回溯做一个剪枝来优化一下。代码：

```
class Solution {
    class Trie{
        Trie[] childs;
        boolean isLeaf;
        public Trie(){
            childs = new Trie[26];
        }
    }

    Trie root;
    Map<Integer, Set<String>> map;  //把清单根据单词长度集合起来
    int maxArea;
    int maxLength;
    List<String> ans;   //别忘最后转换成数组输出
    public String[] maxRectangle(String[] words) {       
        root = new Trie();
        //构造字典树
        for(String str: words){
            Trie node = root;
            for(char c: str.toCharArray()){
                if(node.childs[c-'a'] == null){
                    node.childs[c-'a'] = new Trie();
                }
                node = node.childs[c-'a'];
            }
            node.isLeaf = true;
        }

        map = new HashMap<>();
        ans = new ArrayList<>();
        maxArea = 0;
        maxLength = 0;
        for(String w: words){
            int len = w.length();
            maxLength = Math.max(maxLength, len);
            Set<String> set = map.getOrDefault(len, new HashSet<>());
            set.add(w);
            map.put(len, set);
        }

        List<String> path = new ArrayList<>();
        for(int key: map.keySet()){
            path.clear();
            //回溯需要的参数是：相同长度单词的集合，存放路径的列表，当前单词的长度
            dfs(map.get(key), path, key);
        }

        String[] ultimate = new String[ans.size()];
        return ans.toArray(ultimate);
    }

    //回溯的“套路”
    public void dfs(Set<String> dic, List<String> path, int wordLen){
        //剪枝，dic里的情况不可能得到最优解，提前过滤掉不考虑
        if(wordLen*maxLength <= maxArea)   return;

        //终止条件：如果path矩阵的高度已经超过清单中最长单词长度，结束
        if(path.size() > maxLength) return;

        for(String str: dic){
            //做选择
            path.add(str);

            boolean[] res = isValid(path);
            if(res[0]){ //下面还可以再加
                int area = path.size()*path.get(0).length();
                if(res[1] && (area>maxArea)){   //每列都是单词的矩阵
                    maxArea = area;
                    //ans = path;   一定注意这里不能直接把path引用交给答案
                    ans = new ArrayList<>(path);
                }
                //回溯
                dfs(dic, path, wordLen);
            }
            //撤销选择
            path.remove(path.size()-1);
        }
    }

    /** 判断一个矩阵是否每一列形成的单词都在清单里
    *   存在两种情况：1.有的列中的字母不在字典树中，即这一列不可能构成单词，整个矩阵不合要求
    *   2.每列的所有字母都在字典树中但有的结尾不是leaf，也就是有的列目前还不是个单词
    *   所以需要一个boolean数组res[]来存放结果：
    *   res[0]表示是否有字母不在字典树中，true:都在，false:有不在的
    *   res[1]表示是否所有的列都构成了清单里的单词
    */
    public boolean[] isValid(List<String> input){
        boolean[] res = new boolean[2];
        boolean allLeaf = true;
        int m = input.size();
        int n = input.get(0).length();
        for(int j=0; j<n; j++){
            //按列来看单词是否在字典树
            Trie node = root;
            for(int i=0; i<m; i++){
                int c = input.get(i).charAt(j) - 'a';
                if(node.childs[c] == null)  return new boolean[]{false, false};
                node = node.childs[c];
            }
            if(!node.isLeaf)    allLeaf = false;
        }
        return new boolean[]{true, allLeaf};
    }
}
```
![1.png](https://pic.leetcode-cn.com/293fd9193322a08678dcd77a42699c3b081c83c731bbcd5994ad600733ec7eda-1.png)

执行用时不算特别快，但这种做法没问题的，相对应该也比较容易理解了。
如果对Trie字典树不了解，代码中Trie部分没看懂可以先快速浏览一下[面试题 17.13. 恢复空格](https://leetcode-cn.com/problems/re-space-lcci/solution/cong-bao-li-ru-shou-you-hua-yi-ji-triezi-dian-shu-/)的字典树解法。




