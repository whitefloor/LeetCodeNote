# QuickNote
整理 Grid169 內出現的經典題目出現過的算法筆記  
目的是幫助在面試之前可以快速複習必刷算法

# Data Struct & Algorithm

## Binary Tree
### 核心概念
1. 其實就是 Linked List 的變種，變成加上某些條件的樹狀結構
```
//資料結構
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
```
### 場景
#### Algo
1. Preorder Traversal（前序遍歷）
2. Inorder Traversal（中序遍歷）
3. Postorder Traversal（後序遍歷）
4. Level-order Traversal（層序遍歷）
#### 變種樹
1. Binary Search Tree：節點數值，左<中<右為特色 

## Binary Search
### 核心概念
1. 利用搜尋法大幅減少時間複雜度，例如在遞增數列中可由O(N) > O(logN)
### 場景
1. 遞增數列

## Flood Fill
### 核心概念
1. Breadth-First-Search（BFS）
2. Depth-First-Search（DFS）
### 場景
1. 處理 Matrix 問題時經常用到

## Union Find
### 核心概念
1. 利用 Hash Map 的 Key-Value 來記憶父集
### 場景
1. 找子集的共同父集時壓縮路徑時會使算法時間複雜度大幅減少

## Polish notation（波蘭表示法、前綴表示法）
### 核心概念
1. Last In, First Out  
2. 由 Array 實現
### 場景
1. 計算加減乘除及括號會使用到，十分重要