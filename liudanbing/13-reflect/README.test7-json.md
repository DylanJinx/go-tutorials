定义了一个 `Movie` 结构体，并通过使用 `encoding/json` 包将其实例转换为 JSON 格式。这段代码展示了如何在 Go 中利用结构体字段标签来指定 JSON 键名称，并执行结构体到 JSON 字符串的序列化。下面是对这段代码详细的解释：

### 结构体定义与 JSON 标签

首先，定义了一个名为 `Movie` 的结构体，其中每个字段后面都附有一个 `json:"key"` 标签。这些标签指示 `encoding/json` 包在序列化结构体到 JSON 时应该使用的键名：

```go
type Movie struct {
    Title  string   `json:"title"`  // "title" 作为 JSON 中的键
    Year   int      `json:"year"`   // "year" 作为 JSON 中的键
    Price  int      `json:"dollar"` // "dollar" 作为 JSON 中的键
    Actors []string `json:"actors"` // "actors" 作为 JSON 中的键
}
```

这些标签告诉 `json.Marshal` 函数，当转换 `Movie` 结构体实例为 JSON 时，各个字段应该对应到哪个 JSON 键。例如，`Title` 字段的值应该在 JSON 对象中以键 `"title"` 来表示。

### JSON 序列化

在 `main` 函数中，创建了一个 `Movie` 结构体的实例，并填充了数据：

```go
movie := Movie{"喜剧之王", 2000, 10, []string{"周星驰", "莫文蔚"}}
```

接着，使用 `json.Marshal` 函数将这个 `Movie` 实例转换为 JSON 格式的字节数组：

```go
jsonStr, err := json.Marshal(movie)
```

- `json.Marshal` 返回两个值：一个字节切片（如果成功）和一个错误对象（如果序列化过程中发生错误）。
- 如果序列化成功，`jsonStr` 将包含 `Movie` 实例的 JSON 表示，例如：`{"title":"喜剧之王","year":2000,"dollar":10,"actors":["周星驰","莫文蔚"]}`。
- 如果发生错误，会打印错误消息并返回。

### 输出结果

最后，使用 `fmt.Printf` 打印出 JSON 字符串：

```go
fmt.Printf("jsonStr: %s\n", jsonStr)
```

这将输出：

```
jsonStr: {"title":"喜剧之王","year":2000,"dollar":10,"actors":["周星驰","莫文蔚"]}
```

### 小结

通过这段代码，可以看到 Go 如何利用结构体标签来控制如何将结构体字段映射到 JSON 对象的键上。这种映射非常灵活，使得数据在 Go 结构体与 JSON 之间的转换既直观又易于管理。这种方式广泛应用于构建 Web API、配置文件解析等需要数据序列化和反序列化的场景。