## package/cmd



### 前言

cmd 命名行工具可以使用` go tool` 进行查看：

```bash
wxw-go wxw$ go tool
addr2line
asm
buildid
cgo
compile
cover
dist
doc
fix
link
nm
objdump
pack
pprof
test2json
trace
vet

For more about each tool command, see 'go doc cmd/<command>'.
比如：go doc cmd/pprof
```

### pprof

pprof 是在做性能优化前的性能分析工具。

#### 1. 基本回顾

Go 语言自带的 pprof 库就可以分析程序的运行情况，并且提供可视化的功能。它包含两个相关的库：

- runtime/pprof：对于只跑一次的程序，例如每天只跑一次的离线预处理程序，调用 pprof 包提供的函数，手动开启性能数据采集。
- net/http/pprof：对于在线服务，对于一个 HTTP Server，访问 pprof 提供的 HTTP 接口，获得性能数据。当然，实际上这里底层也是调用的 runtime/pprof 提供的函数，封装成接口对外提供网络访问。

##### 1.1 pprof 的作用





















相关文档

1. [golang pprof 实战](https://blog.wolfogre.com/posts/go-ppof-practice/) 