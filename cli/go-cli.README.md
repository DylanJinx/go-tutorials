## 1. `go run`

**用途**：编译并运行 Go 程序。

**示例**：

```bash
go run main.go
```

**说明**：此命令会临时编译 `main.go` 并执行生成的可执行文件。适用于快速测试代码。

## 2. `go build`

**用途**：编译 Go 源代码，生成可执行文件。

**示例**：

```bash
go build main.go
```

**说明**：编译 `main.go` 并在当前目录下生成可执行文件（在 macOS 上为 `main`）。如果是包级别的编译，执行 `go build` 会生成包的可执行文件。

## 3. `go install`

**用途**：编译并安装可执行文件到 `$GOPATH/bin` 或模块模式下到 `$GOBIN`。

**示例**：

```bash
go install ./...
```

**说明**：编译当前模块下的所有包并将可执行文件安装到指定目录，方便全局调用。

## 4. `go fmt` / `gofmt`

**用途**：格式化 Go 源代码。

**示例**：

```bash
go fmt ./...
gofmt -w main.go
```

**说明**：

- `go fmt ./...`：格式化当前模块及所有子目录中的 Go 文件。
- `gofmt -w main.go`：使用 `gofmt` 工具格式化 `main.go` 文件，并直接覆盖写入。

## 5. `go vet`

**用途**：静态分析代码，检查潜在的问题。

**示例**：

```bash
go vet ./...
```

**说明**：分析当前模块及所有子目录中的 Go 文件，报告潜在的错误和不规范的代码。

## 6. `go test`

**用途**：运行测试。

**示例**：

```bash
go test ./...
go test -v
go test -cover
```

**说明**：

- `go test ./...`：运行当前模块及所有子目录中的测试。
- `-v`：详细模式，显示每个测试的运行结果。
- `-cover`：生成测试覆盖率报告。

## 7. `go get`

**用途**：获取并安装远程包及其依赖。

**示例**：

```bash
go get github.com/gin-gonic/gin
```

**说明**：下载并安装 `gin` 框架到模块的依赖中。

> **注意**：在使用 Go 模块（`go.mod`）时，`go get` 也用于更新模块的依赖版本。

## 8. `go mod` 相关命令

### 8.1 `go mod init`

**用途**：初始化 Go 模块。

**示例**：

```bash
go mod init github.com/yourusername/yourproject
```

**说明**：在项目根目录下创建 `go.mod` 文件，初始化模块。

### 8.2 `go mod tidy`

**用途**：整理模块依赖，添加缺失的依赖并移除不需要的依赖。

**示例**：

```bash
go mod tidy
```

**说明**：确保 `go.mod` 和 `go.sum` 文件的依赖是最新且干净的。

### 8.3 `go mod vendor`

**用途**：将依赖包复制到 `vendor` 目录。

**示例**：

```bash
go mod vendor
```

**说明**：将所有依赖包复制到 `vendor` 目录，便于在没有网络的环境下构建项目。

## 9. `go env`

**用途**：显示 Go 环境的配置信息。

**示例**：

```bash
go env
go env GOPATH
```

**说明**：

- `go env`：显示所有 Go 环境变量。
- `go env GOPATH`：只显示 `GOPATH` 的值。

## 10. `go list`

**用途**：列出包的信息。

**示例**：

```bash
go list ./...
go list -json ./...
```

**说明**：

- `go list ./...`：列出当前模块及所有子目录中的包。
- `-json`：以 JSON 格式输出包的信息。

## 11. `go doc`

**用途**：查看包或符号的文档。

**示例**：

```bash
go doc fmt.Println
go doc net/http
```

**说明**：查看 `fmt.Println` 函数或 `net/http` 包的文档说明。

## 12. `go generate`

**用途**：运行代码生成工具。

**示例**：

```bash
go generate ./...
```

**说明**：根据源代码中的 `//go:generate` 指令运行相应的生成命令。

## 13. `go clean`

**用途**：清理构建缓存和生成的文件。

**示例**：

```bash
go clean
go clean -cache
go clean -modcache
```

**说明**：

- `go clean`：删除生成的对象文件。
- `-cache`：删除构建缓存。
- `-modcache`：删除模块缓存。

## 14. `go tool`

**用途**：访问 Go 工具链的低级工具。

**示例**：

```bash
go tool pprof
go tool compile
```

**说明**：`go tool` 提供访问编译器、链接器、性能分析等工具的接口。

## 15. `go version`

**用途**：显示当前 Go 版本信息。

**示例**：

```bash
go version
```

**说明**：输出类似 `go version go1.20.3 darwin/amd64` 的信息。

## 16. `go install`（模块模式）

**用途**：安装指定版本的模块。

**示例**：

```bash
go install github.com/some/module@v1.2.3
```

**说明**：安装 `some/module` 的指定版本 `v1.2.3`。

## 17. 其他有用的工具

### 17.1 `golint`

**用途**：检查代码是否符合 Go 语言的风格规范。

**安装**：

```bash
go install golang.org/x/lint/golint@latest
```

**使用**：

```bash
golint ./...
```

### 17.2 `staticcheck`

**用途**：静态分析工具，检查代码中的潜在问题和改进建议。

**安装**：

```bash
go install honnef.co/go/tools/cmd/staticcheck@latest
```

**使用**：

```bash
staticcheck ./...
```

### 17.3 `gopls`

**用途**：Go 语言的语言服务器，提供智能补全、跳转到定义等功能，常用于编辑器集成。

**安装**：

```bash
go install golang.org/x/tools/gopls@latest
```
