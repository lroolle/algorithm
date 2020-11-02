"+++
title = "面试题 17.25. Word Rectangle LCCI C++ 252ms超越了100%用户 "
author = ["ha-ha-310"]
date = 2020-09-02T10:01:47+08:00
tags = ["Leetcode", "Algorithms", "cpp", "DepthfirstSearch", "Trie", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# C++ 252ms超越了100%用户

> [面试题 17.25. Word Rectangle LCCI](https://leetcode-cn.com/problems/word-rectangle-lcci/)
> [C++ 252ms超越了100%用户](https://leetcode-cn.com/problems/word-rectangle-lcci/solution/c-252mschao-yue-liao-100yong-hu-by-ha-ha-310/) by [ha-ha-310](https://leetcode-cn.com/u/ha-ha-310/)

### 解题思路
![QQ图片20200902175129.png](https://pic.leetcode-cn.com/1599040900-rHnTnk-QQ%E5%9B%BE%E7%89%8720200902175129.png)
此处撰写解题思路

### 代码

```cpp
class Solution {
public:
    struct tire{
        struct tire *parent; //父节点指针
        struct tire *child[26]; //子节点指针数组
        //构造函数
        tire():parent(NULL)
        {
            memset(child,0,sizeof(struct tire *)*26);
        }
        //拷贝构造
        tire(const tire & tire):parent(tire.parent)
        {
            memcpy(child,tire.child,sizeof(struct tire *)*26);
            for(int i = 0; i< 26; i++)
            {
                if(child[i])
                {
                    child[i]->parent = this;
                }
            }
        }
        //添加函数
        void add(const string & word)
        {
            struct tire * temp=this;
            for(auto it:word)
            {
                if(!temp->child[it - 'a'])
                {
                    temp->child[it - 'a']=new tire();
                    temp->child[it - 'a']->parent = temp;
                }
                temp = temp->child[it -'a'];
            }
        }
        //释放函数
        void free()
        {
            for(int i  = 0; i < 26;i++)
            {
                if(child[i])
                {
                    child[i]->free();
                    delete child[i];
                }
            }
        }
    };

    vector<string> maxRectangle(vector<string>& words) {
        //寻找单词最大长度
        int max=0,len;
        for(auto & it:words)
        {
            len = it.size();
            if(len > max)
            {
                max = len;
            }
        }
        //将单词按长度分别放入list及字典树
        vector<list<string>> list(max+1);
        vector<tire> tire(max+1);
        for(auto & it:words)
        {
            len = it.size();
            list[len].push_back(it);
            tire[len].add(it);
        }
        //i确定矩形列长度，j确定矩形行长度
        for(int i = max;i>0;i--)
        {
            vector<string> ret(i);
            for(int j = i; j > 0; j--)
            {
                vector<struct tire*> tmp(j,&tire[i]);
                if(dfs(tmp,list[j],ret,0,i,j))
                {
//                    for(auto it:tire)
//                    {
//                        it.free();
//                    }
                    return ret;
                }
            }
        }
//        for(auto it:tire)
//        {
//            it.free();
//        }
        return {};
    }

    bool dfs(vector<struct tire*> & tire,list<string>& list,vector<string> & ret,int curX,int x,int y)
    {
        if(curX == x)
        {
            return true;
        }
        else
        {
            int i;
            //遍历所以行长度为y的list
            for(const auto & it:list)
            {
                //利用字典树剪枝
                for(i = 0; i < y;i++)
                {
                    if(!tire[i]->child[it[i] - 'a'])
                    {
                        break;
                    }
                }
                if(i == y)
                {
                    for(i = 0; i < y;i++)
                    {
                        tire[i] = tire[i]->child[it[i]-'a'];
                    }
                }
                else{
                    continue;
                }
                
                ret[curX] = it;
                //向下递归
                if(dfs(tire,list,ret,curX+1,x,y))
                {
                    return true;
                }
                else
                {
                    //递归失败，回溯字典树指针
                    for(i = 0; i < y;i++)
                    {
                        tire[i] = tire[i]->parent;
                    }
                }
            }
        }

        return false;
    }
};
```