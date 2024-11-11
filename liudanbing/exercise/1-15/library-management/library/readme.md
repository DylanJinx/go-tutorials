## 1. `NextID` 是什么？

```go
NextID int
```

**作用**：`NextID` 是一个整数变量，用于为每本新添加的书籍生成一个**唯一的ID**。

**原因**：在图书管理系统中，每本书需要一个唯一的标识符（ID）来区分不同的书籍。使用 `NextID` 可以确保每次添加新书籍时，都能分配到一个独一无二的ID。

**工作原理**：

- **初始化**：当 `Library` 结构体被创建时，`NextID` 通常被初始化为 `1`。

- **添加书籍时**：

  - 使用当前的 `NextID` 作为新书籍的ID。
  - 将新书籍添加到 `Books` 映射（map）中，键为 `NextID`，值为书籍信息。
  - **自增 `NextID`**：将 `NextID` 的值加 `1`，为下一本书籍准备。

**示例代码**：

```go
func (lib *Library) AddBook(title, author string, category models.Category, price float64) int {
    lib.mu.Lock()
    defer lib.mu.Unlock()

    id := lib.NextID
    lib.Books[id] = models.Book{
        ID:       id,
        Title:    title,
        Author:   author,
        Category: category,
        Price:    price,
    }
    lib.NextID++ // 自增 NextID
    return id
}
```

在这个函数中：

- **获取当前的 `NextID`** 作为新书的ID。
- **添加新书** 到 `Books` 映射中。
- **自增 `NextID`**，为下次添加书籍准备。

---
**`NextID` 并不表示当前图书的总数量**，而是用于为每本新添加的书籍生成一个**唯一的ID**。它始终保持为下一个可用的ID，无论当前图书馆中实际有多少书籍。

**具体解释如下：**

### 1. `NextID` 的作用

- **唯一标识符**：`NextID` 用于确保每本新添加的书籍都有一个独一无二的ID。
- **自增特性**：每次添加一本新书，`NextID` 的值都会递增，即 `NextID++`，为下一个新书籍准备。
- **不受删除影响**：当你删除一本书时，`Books` 映射中的对应条目会被移除，但 `NextID` **不会减少**。它只在添加新书时递增。

### 2. 为什么 `NextID` 不等于图书总数

- **删除书籍的影响**：如果你删除了一些书籍，`Books` 中的条目数量会减少，但 `NextID` 不会因为删除操作而减少。
- **添加书籍的历史**：`NextID` 实际上反映了自系统初始化以来，**添加过的书籍数量加一**，而不是当前存在的书籍数量。

**举个例子：**

1. **初始状态**：`NextID = 1`，`Books` 为空。
2. **添加第一本书**：
   - 分配ID：`id = NextID`，即 `id = 1`。
   - 添加到 `Books`：`Books[1] = 新书籍信息`。
   - 更新 `NextID`：`NextID++`，现在 `NextID = 2`。
3. **添加第二本书**：
   - 分配ID：`id = NextID`，即 `id = 2`。
   - 添加到 `Books`：`Books[2] = 新书籍信息`。
   - 更新 `NextID`：`NextID++`，现在 `NextID = 3`。
4. **删除第一本书**：
   - 从 `Books` 中删除：`delete(Books, 1)`。
   - `NextID` **保持不变**，仍然是 `3`。
5. **当前状态**：
   - `Books` 中只有一本书（ID为2的书籍）。
   - `NextID = 3`。

此时，图书馆中有 **1 本书**，但 `NextID` 的值是 `3`。

### 3. 如何获取当前图书的总数量

要获取当前图书馆中实际存在的书籍数量，你需要检查 `Books` 映射的长度：

```go
totalBooks := len(lib.Books)
fmt.Printf("当前图书馆中有 %d 本书。\n", totalBooks)
```

这样，你就能准确地知道当前有多少书籍。

---

## 2. `sync.Mutex` 是什么？

```go
mu sync.Mutex // 保护 Books 和 NextID 的并发安全
```

**`sync.Mutex`**：

- **定义**：`sync.Mutex` 是 Go 语言标准库 `sync` 包中的一个互斥锁类型，用于**控制对共享资源的访问**，防止多个协程（goroutines）同时读写共享资源导致数据不一致的问题。

**作用**：

- **保护共享资源**：在我们的 `Library` 结构体中，`Books`（存储书籍的 `map`）和 `NextID` 是需要保护的共享资源。
  
- **防止数据竞争（Race Condition）**：当多个协程同时访问和修改 `Books` 或 `NextID` 时，可能会发生数据竞争，导致数据不一致或程序崩溃。

**工作原理**：

- **加锁（Lock）**：在访问或修改共享资源之前，调用 `mu.Lock()`，这会阻塞其他试图获取锁的协程，直到当前协程释放锁。

- **解锁（Unlock）**：在访问或修改完共享资源后，调用 `mu.Unlock()`，释放锁，允许其他被阻塞的协程获取锁并继续执行。

**示例代码**：

```go
func (lib *Library) AddBook(...) {
    lib.mu.Lock()         // 加锁，保护共享资源
    defer lib.mu.Unlock() // 解锁，确保函数结束前释放锁

    // 对 Books 和 NextID 的操作
}
```

在这个例子中：

- **`lib.mu.Lock()`**：在修改 `Books` 和 `NextID` 之前，获取互斥锁，确保只有当前协程能够访问这些资源。

- **`defer lib.mu.Unlock()`**：使用 `defer` 关键字，确保无论函数正常结束还是发生错误，都能在函数退出前释放锁，防止死锁。

