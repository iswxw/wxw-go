### Go语言

---

#### **项目目录结构**
```go
wxw-go // go_project为GOPATH目录
  -- bin // 编译后生成的可执行文件（为了方便，可以把此目录加入到 $PATH 变量中，如果有多个gopath，那么使用${GOPATH//://bin:}/bin添加所有的bin目录）
     -- myApp1  // 编译生成
     -- myApp2  // 编译生成
  -- pkg        // 编译时生成的中间文件（比如：.a）golang编译包时
  -- src        // 存放源代码(比如：.go .c .h .s等) 按照golang默认约定，go run，go install等命令的当前工作路径（即在此路径下执行上述命令）
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

#### **常用命令**

```go
go env       // 查看go当前环境
go get       // 从远程下载需要用到的包、执行go install
go install   // 会生成可执行文件直接放到bin目录下
go run       // 在编译后直接运行Go语言程序，编译过程中会产生一个临时文件，但不会生成可执行文件
go build     // 将Go语言程序代码编译成二进制的可执行文件，但是需要我们手动运行该二进制文件；
```
  
#### **学习资源**   
1. Go入门
   - [学习视频](https://www.bilibili.com/video/BV1h7411x7JB?p=14)        
  
2. GO进阶+实战
   - [学习视频](https://www.bilibili.com/video/bv1FV411r7m8/?spm_id_from=333.788.b_636f6d6d656e74.26)      
   
   