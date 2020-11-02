"+++
title = "面试题 17.25. Word Rectangle LCCI [Java]字典树+回溯+dfs "
author = ["tlzxsun"]
date = 2020-08-13T02:35:25+08:00
tags = ["Leetcode", "Algorithms", "Java", "Trie", "Backtracking"]
categories = ["leetcode"]
draft = false
+++

# [Java]字典树+回溯+dfs

> [面试题 17.25. Word Rectangle LCCI](https://leetcode-cn.com/problems/word-rectangle-lcci/)
> [[Java]字典树+回溯+dfs](https://leetcode-cn.com/problems/word-rectangle-lcci/solution/javazi-dian-shu-hui-su-dfs-by-tlzxsun/) by [tlzxsun](https://leetcode-cn.com/u/tlzxsun/)

一顿瞎写居然AC了，具体的还是看注释吧
```
class Solution {

    int max = 0, maxLength = 0;
    List<String> result = null;

    public String[] maxRectangle(String[] words) {
        //System.out.println(Arrays.stream(words).map(it -> "\"" + it + "\"").collect(Collectors.toList()));
        //用来保存每个长度的字符串
        Map<Integer, List<String>> map = new HashMap<>();
        //字典树根节点
        Node root = new Node();
        //遍历字符串
        for (String word: words) {
            //记下最大长度
            maxLength = Math.max(maxLength, word.length());
            //加入对应长度的字符串列表
            map.putIfAbsent(word.length(), new ArrayList<>());
            map.get(word.length()).add(word);
            //当前默认为根节点
            Node current = root;
            for (int i = 0; i < word.length(); i++) {
                //如果根节点对应的子节点为null，那么赋值
                if (current.child[word.charAt(i) - 'a'] == null) {
                    current.child[word.charAt(i) - 'a'] = new Node();
                }
                //更改当前节点指向
                current = current.child[word.charAt(i) - 'a'];
                //如果是最后一个字母，标记为叶子
                if (i == word.length() - 1) {
                    current.leaf = true;
                }
            }
        }
        //遍历每个长度的字符串列表
        for (Map.Entry<Integer, List<String>> entry: map.entrySet()) {
            //剪枝，只有新的长度可能的组合面积大于已有最大值才需要查找
            if (entry.getKey() * entry.getValue().size() > max) {
                dfs(entry.getValue(), new ArrayList<>(), new Node[entry.getKey()], root, maxLength);
            }
        }
        System.out.println(result);
        return result.toArray(new String[result.size()]);
    }

    /**
     *
     * @param list 当前可选字符串
     * @param current 当前已经选中字符串
     * @param last 上一轮选中的节点
     * @param root 根节点
     * @param maxLength 最大长度
     */
    void dfs(List<String> list, List<String> current, Node[] last, Node root, int maxLength) {
        //如果选中的已经达到最大长度了，再选肯定不符合要求直接返回
        if (current.size() == maxLength) {
            return;
        }
        Node[] newLast = new Node[last.length];
        //遍历所有可选字符串
        for (String word: list) {
            Arrays.fill(newLast, null);
            //分别标记选中当前后是否可以作为结果，当前选中的是否不能选中
            boolean ends = true, skip = false;
            for (int i = 0; i < word.length(); i++) {
                //得到上一个节点
                Node node = last[i]==null?root:last[i];
                //如果上一个节点对应的节点为null，说明字典中不存在这个序列，那么这个字符串不能加到结果中，跳过
                if (node.child[word.charAt(i) - 'a'] == null) {
                    skip = true;
                    break;
                } else if (!node.child[word.charAt(i) - 'a'].leaf) {
                    //如果上一个节点对应的节点不是叶子，说明加上当前字符串后不是完整的结果
                    ends = false;
                }
                //将上一节点赋值成当前的节点
                newLast[i] = node.child[word.charAt(i) - 'a'];
            }
            //如果需要跳过
            if (skip) {
                continue;
            }
            //将当前字符串加入选中的列表中
            current.add(word);
            //如果当前是一个完整的结果并且面积大于最大值，那么记下结果
            if (ends) {
                if (current.size() * current.get(0).length() > max) {
                    max = current.size() * current.get(0).length();
                    result = new ArrayList<>(current);
                }
            }
            //继续递归
            dfs(list, current, newLast, root, maxLength);
            //回溯
            current.remove(current.size() - 1);
        }
    }

    //字典数结构
    class Node {
        //当前节点是否是一个词的结尾
        boolean leaf = false;
        //当前节点的子节点，代码26个小写字母
        Node[] child = new Node[26];
    }
}
```