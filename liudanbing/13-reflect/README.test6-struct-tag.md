在这段 Go 代码中，主要涉及到结构体的标签（Tag）使用和通过反射（reflect）库来动态地访问这些标签。代码定义了一个 `resume` 结构体，并在 `main` 函数中创建了这个结构体的一个实例，然后通过反射查找并打印结构体字段的标签。下面是对这段代码的详细解释：

### 结构体定义和标签
首先，定义了一个名为 `resume` 的结构体，它有两个字段：`Name` 和 `Sex`，每个字段都附加了一些标签（Tag），这些标签在 Go 中以键值对的形式存在：

```go
type resume struct {
    Name string `info:"name" doc:"my name"`
    Sex  string `info:"sex"`
}
```
这里：
- `Name` 字段有两个标签：`info` 和 `doc`。
- `Sex` 字段有一个标签：`info`。

### 标签（Tag）
在 Go 中，标签通常用于提供关于结构体字段的元数据。它们可以被用于各种框架和库中，如序列化、数据库映射、配置解析等，这些库会通过反射来读取这些标签。

### 反射访问标签
在 `findTag` 函数中，我们使用反射来访问传入对象的标签：

```go
func findTag(str interface{}) {
    t := reflect.TypeOf(str).Elem()
    
    for i := 0; i < t.NumField(); i++ {
        tagInfo := t.Field(i).Tag.Get("info")
        tagDoc := t.Field(i).Tag.Get("doc")
        fmt.Println("info: ", tagInfo, ", doc: ", tagDoc)
    }
}
```
这里：
- `reflect.TypeOf(str).Elem()` 获取了 `str` 的类型的元素类型。因为 `str` 是一个指向 `resume` 的指针，`Elem()` 方法用于获取指针指向的元素类型，即 `resume`。
- `NumField()` 方法返回结构体中字段的数量。
- `Field(i)` 方法返回结构体的第 `i` 个字段的 `reflect.StructField` 对象。
- `Tag.Get("info")` 和 `Tag.Get("doc")` 方法分别用于获取字段的 `info` 和 `doc` 标签的值。

### 函数调用
在 `main` 函数中，创建了 `resume` 结构体的一个实例并取地址（生成一个指向 `resume` 的指针），然后将这个指针传递给 `findTag` 函数：

```go
var re resume
findTag(&re)
```

### 输出
当 `findTag` 函数被调用时，它会遍历 `resume` 结构体的所有字段，打印每个字段的 `info` 和 `doc` 标签值。对于上面的结构体定义，输出将是：

```
info: name , doc: my name
info: sex , doc: 
```

这段代码展示了如何结合使用结构体标签、反射和方法，来动态地获取和处理这些标签的信息，这在配置解析、ORM 映射、JSON/XML 解析等场景中非常有用。