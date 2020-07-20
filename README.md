### wxw-gocode

---

### 一、目录结构

```go
wxw-gocode  // go_project为GOPATH目录
  -- bin // 编译后生成的可执行文件（为了方便，可以把此目录加入到 $PATH 变量中，如果有多个gopath，那么使用${GOPATH//://bin:}/bin添加所有的bin目录）
     -- myApp1  // 编译生成
     -- myApp2  // 编译生成
     -- myApp3  // 编译生成
  -- pkg        // 编译时生成的中间文件（比如：.a）　　golang编译包时
  -- src        // 存放源代码(比如：.go .c .h .s等)   按照golang默认约定，go run，go install等命令的当前工作路径（即在此路径下执行上述命令）
     -- myApp1     // project1
        -- models
        -- controllers
        -- others
        -- main.go 
     -- myApp2     // project2
        -- models
        -- controllers
        -- others
        -- main.go 
     -- myApp3     // project3
        -- models
        -- controllers
        -- others
        -- main.go 
```

### 二、常用命令

```go
go env // 查看go当前环境
go get // 从远程下载需要用到的包、执行go install
go install // 会生成可执行文件直接放到bin目录下
```





























