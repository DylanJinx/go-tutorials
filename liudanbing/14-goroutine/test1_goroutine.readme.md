从代码执行流程的角度来说，使用 `go` 关键字开启的新 goroutine 并不会阻塞 `main` 函数原有的执行流程。

1. **正常函数调用（无 go 关键字）**：  
   如果在 `main` 函数中直接调用 `newTask()`，那么程序的执行会顺序进行。也就是说，`main` 函数会一直在 `newTask()` 函数内运行，直到 `newTask()` 返回（但在这个例子中 `newTask()` 里的死循环不会结束，所以 `main` 会被“卡死”在 `newTask()` 内，永远也执行不到后面的代码）。

2. **使用 go 关键字调用函数（开启 goroutine）**：  
   当你在 `main` 中使用 `go newTask()` 时，程序会创建一个与 `main` 函数并行运行的 goroutine。这两个 goroutine（`main goroutine` 和 `newTask goroutine`）将同时执行。  
   - `main goroutine` 会继续往下执行自己的代码（即 `i++`、打印、`time.Sleep()`）。
   - 与此同时，`newTask goroutine` 也会独立执行自己的死循环（`i++`、打印、`time.Sleep()`），并不会阻塞 `main` 函数。  

换句话说，`go` 关键字启动的新 goroutine 会在后台并发运行，不会像传统函数调用那样占据当前 goroutine 的执行流，从而阻塞后续代码运行。这就是为什么在同一时间内，你能看到 `main goroutine` 和 `newTask goroutine` 交替打印输出，这表明它们在并行（并发）工作。

如果系统只有一个 CPU 核，那么 goroutine 的执行虽然是并发的，但由于只有一个执行单元，goroutine 之间实际上是被 Go runtime 的调度器轮换（time slicing）执行的，看起来像是“同时”执行，但本质上是快速切换，让每个 goroutine 都有机会运行。

如果系统有多个 CPU 核，那么在 Go 程序默认设置（GOMAXPROCS 不限于1）的情况下，多个 goroutine 可以分布在不同的 CPU 核上同时运行。这时就不仅仅是并发（concurrency），而是可以真正做到并行（parallelism）执行。
