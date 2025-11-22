# QuickNote
整理 Grid169 內出現的經典題目出現過的算法筆記  
目的是幫助在面試之前可以快速複習必刷算法

# 刷題技巧
1. 最重要，除非你天生神力或看過類似題目，在看完 Grind 169 前千萬別做夢自己寫出來  
    寫不出來很正常
2. 類似題目可以嘗試寫看看，你會發現你不是不會，通常是忽略一些條件
3. 多參考幾個人的答案，要刷的有品質，通常會參考『最簡單』或『最有效率』的寫法  
    有時候思路是一樣的，但是別人的 code 比較乾淨
4. 該休息就休息，今天多刷一題，隔天太累少刷兩題不是很虧？
5. 刷題都說是刷了，重點就是要快，用看的看多了也會進步
6. 有空可以去打比賽鍛鍊一下，順便模擬面試那種緊張的氣氛

# 能刷題的時間不多了，推薦刷題類型
1. DP + Knapsack problem
2. Graph
3. Greedy
4. Backtrack

# 刷提提升 Runtime 小技巧與寫題好習慣
1. 盡量不要用 hash map，大部分題目也確實用不到，善加利用 array
2. 能預先 allocate array space 最好，在 Go 裡用 append 當空間不足時會重新分配空間並進行拷貝
3. 不要用神奇數字或名字命名，也讓你的 code 比較易懂
4. 手打的時候在 submit 前人工檢查有沒有 type
5. 準備三種 test case：超出左邊界、超出右邊界、符合邏輯同時又在邊界內
6. 寫之前估算一下 time & space complex，不要硬幹暴力解

# Data Struct & Algorithm

## String
### 核心概念
### 場景
1. 在 Go 中善用 `s:='A'` 把字串轉成 ASCII Code 來計算、分類、儲存
2. 要計算成數字，可以利用 ASCII 相減，例如 `i := 'a' - 'a'`

## Two Points
### 核心概念
1. 兩個指針從 0,len(arr)-1 開始像中間逼近進行交換或比較的一種方式
### 場景
```
func TwoPoints(arr []int){
    left,right:=0,0
    for left<right{
        ...
    }
}
```

## Tree
### 核心概念
1. 其實就是 Linked List 的變種，跟 Graph 也算親戚，都是加上某些條件形成的陣列結構
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
1. Preorder Traversal（前序遍歷）：形成樹的順序是中、左、右
2. Inorder Traversal（中序遍歷）：形成樹的順序是左、中、右
3. Postorder Traversal（後序遍歷）形成樹的順序是左、右、中
```
會有用某兩種 Traversal 方法產生的 Array 讓你還原成 Binary Tree 的問題
這要了解這幾種樹的特性才能解決
假設左樹長度為 n
1. Preorder Traversal：root：arr[0]，左樹範圍：arr[1:n]，右樹範圍：arr[n+1:]
2. Inorder Traversal：root：arr[n:n+1]，左樹範圍：arr[0:n]，右樹範圍：arr[n+1:]
```
4. Level-order Traversal（層序遍歷）：用 BFS 比較好解
```
func levelOrderBFS(root *TreeNode) []int {
        //初始化 BFS Array
	queue := []*TreeNode{root}
	
        //BFS
	for len(queue) > 0 {
	
	//計算長度最後要頂出，最重要的部分
        oldQLen:=len(queue)
        ...
        
        //頂出前面處理過的元素
        queue=queue[oldQLen:]
	}
}
```

#### 變種樹
1. Binary Search Tree：節點數值，左<中<右為特色
2. Height-Balanced Tree：意思指左右節點的樹高不會大於 1
3. Minimum Height Trees：同時也是 Graph 的樹中，找到最小高度的問題
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

## Hamming Weight
### 核心概念
1. 探究在一段二進制中1的密集程度
### 場景
1. 經典題裡沒有更進一步探索

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

## Backtrack
### 核心概念
1. DFS + 復原上一步（窮舉法的步驟）
2. 要有剪枝條件避免窮舉
3. 剪枝條件達成同時複製結果
### 場景
1. 在需要計算出所有結果時需要的演算法
```
func backtrackSample(nums []int) [][]int {
    var result [][]int
    var backtrack func(start int, path []int)
    
    //DFS
    backtrack = func(start int, path []int) {
        //達成減枝條件同時複製結果
        temp := make([]int, len(path))
        copy(temp, path)
        result = append(result, temp)
        
        //main logic
        for i := start; i < len(nums); i++ {
            path = append(path, nums[i])
            backtrack(i+1, path)
            
            //下列復原狀態為精華，不一定是擠出最後一個元素，能夠復原的方法都行
            path = path[:len(path)-1]
        }
    }
    backtrack(0, []int{})
    return result
}
```

## Graph
### 核心概念
1. 利用二維陣列紀錄頂點與邊
2. 或者利用 Hash Map 紀錄頂點與邊
### 場景
1. Activity On Vertex (AOV) Network，一種有向無環圖
```
作法通常是利用三個 Array 分別紀錄
1. 入邊次數，例如 arr[1,0]，從0>1，inDegree[arr[0]]++
2. 紀錄每個頂點的關係，建立 Graph，例如 adj[arr[1]]=append(adj[arr[1]],arr[0])
3. 從 Graph 中找到起始點，inDegree[i]=0 的點加入 Queue
4. 利用 BFS 從這些起點找到相鄰的點，在 inDegree 中去除次數
5. 如果 inDegree 此時為 0，就加入為下一次的起始點
6. 最後利用 inDegree中所有元素=0，或是其他條件判斷有沒有造成循環
```

## Least Recently Used（LRU）
### 核心概念
1. 利用 Linked List 易於改變位置的特性，時間複雜度O(1)，紀錄 Head,Tail 與元素
2. 最新被讀取到的放到 Head
3. 新增時如果達到長度上限則淘汰 Tail
4. 搭配 Hash Map 紀錄 Node 的位置，要刪除淘汰資料時讓不易於讀取的 Linked List  
    達到時間複雜度 O(1) 的刪除速度 
### 場景
1. 一種將最近時間內最少使用的資料替換掉的熱門搜尋算法，在 Redis 中的記憶體淘汰策略中有使用到
2. 類似的算法有 Least Frequency Used（LFU），但 LFU 是基於訪問次數來排除

## Longest Increasing Subsequence（LIS：最長遞增子序列）
### 核心概念
1. 將當前數比最常上升數列尾數的數字加入數列
2. 將當前數利用二分搜尋找尋能夠保持上升數列中能夠替換掉的最小數進行替換
### 場景
1. 在一個給定的數值序列中，找到一個子序列，使得這個子序列元素的數值依次遞增  
   並且這個子序列的長度儘可能地大

# 常用套件
## sort
```
sort.Strings(strs)
```

```
sort.Slice(points, func(i, j int) bool {}
```

```
sort.Ints(nums)
```

## strconv
```
strconv.Atoi(str)
```

## slice
```
slices.Sort(sortedRune)
```

```
sort.SliceStable(uniq, func(i, j int) bool {}
```

# 常用模版

## 比大小，找大的
* Go 1.21 有內建max/min函式
```
func min(a, b int) int {
	if a > b {
		return a
	}

	return b
}
```

## 比大小，找小的
* Go 1.21 有內建max/min函式
```
func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
```

## 轉換絕對值
```
func abs(x int)int{
    if x<0{
        return -x
    }
    return x
}
```