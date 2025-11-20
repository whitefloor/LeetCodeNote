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
5. Height-Balanced：意思指左右節點的樹高不會大於 1
#### 變種樹
1. Binary Search Tree：節點數值，左<中<右為特色
### 註記
1. 如果用遞迴（Recursion）來解題，要試著從樹底開始想 if condition 會比較好想

## Binary Search
### 核心概念
1. 利用搜尋法大幅減少時間複雜度，例如在遞增數列中可由O(N) > O(logN)
### 場景
1. 遞增數列中尋找某數
2. 尋找左邊界
```
func leftBoundary(n int) int {
    low:=1
    high:=n

    for low<high {
        mid:=(high-low)/2+low
        if ... {
            high=mid
        }else{
            low=mid+1
        }
    }

    return low
}
```
3. 尋找右邊界（找左右邊界概念是一樣的，但是要注意閉合條件）

## Flood Fill
### 核心概念
1. Breadth-First-Search（BFS）
2. Depth-First-Search（DFS）
### 場景
1. 處理 Matrix 問題時經常用到

## Fast and Slow Pointer
### 核心概念
1. 利用兩個指針從頭開始一快（走兩步）一慢（走一步）
### 場景
1. Linked List 尋找中間節點
2. Linked List 尋找有沒有循環
```
func findCycle(head *ListNode) bool {
    slow:=head
    fast:=head

    for fast!=nil && fast.Next!=nil{
        slow=slow.Next
        fast=fast.Next.Next
        if slow==fast{
            return true
        }
    }

    return false
}
```

## Reverse Linked List
### 核心概念
1. 每次都會忘記就記下來了
### 場景
1. 反轉 Linked List
```
func reverse(head *ListNode) *ListNode {
    //Declare a nil pointer
    var prevHead *ListNode

    for head!=nil {
        //記憶 Head 的下一個節點
        nextHead:=head.Next
        //把 Head 的下一個節點改為上一個節點
        head.Next=prevHead
        //把"上一個"節點挪到當前 Head，因為會成為"下一個"節點的"上一個"節點
        prevHead=head
        //把 Head 移到一開始記憶的"下一個"節點
        head=nextHead
    }

    return prevHead
}
```

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