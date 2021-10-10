## 流程控制

---

### 无条件跳转：`goto` 

Go 居然会保留 goto，因为很多人不建议使用 goto，所以在一些编程语言中甚至直接取消了 goto。

Go 语言的 goto 语句可以无条件地转移到过程中指定的行。

goto 语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。

但是，在结构化程序设计中一般不主张使用 goto 语句， 以免造成程序流程的混乱，使理解和调试程序都产生困难。

#### 1. goto 语法

goto 语法格式如下：

```
goto 标签;
...
...
标签: 表达式;
```

#### 2. 为何golang支持goto语句

其实goto不是洪水猛兽，不推荐使用是因为这东西太灵活，容易造成额外的调试复杂度。只要用法得当，goto有的时候反而更加干脆便捷。

最常见的两种情况：

1. **golang里跳出多级循环**，除非使用单独变量，且在每一级循环中增加额外判断，否则想跳出多级循环就只能用goto（其他语言里的break标签形式，可以理解为goto的语法糖版本）。
2. **golang的错误处理**，就是那个if err!=nil，通常的写法是逐级处理，当需要忽略后续判断时，使用goto更加简洁，比一层一层的判断标志变量更清晰。

**举个例子** 

- 跳出多层循环

  - 传统编码：

  ```go
  package main
   
  import "fmt"
   
  func main() {
   
      var breakAgain bool
   
      // 外循环
      for x := 0; x < 10; x++ {
          // 内循环
          for y := 0; y < 10; y++ {
              // 满足某个条件时, 退出循环
              if y == 2 {
                  // 设置退出标记
                  breakAgain = true
                  // 退出本次循环
                  break
              }
          }
   
          // 根据标记, 还需要退出一次循环
          if breakAgain {
                  break
          }
      }
   
      fmt.Println("done")
  }
  ```

  - 如果使用goto 跳出多层循环优化

  ```go
  package main
   
  import "fmt"
   
  func main() {
   
      for x := 0; x < 10; x++ {
          for y := 0; y < 10; y++ {
              if y == 2 {
                  // 跳转到标签
                  goto breakHere
              }
          }
      }
      // 手动返回, 避免执行进入标签
      return
   
      // 标签
  breakHere:
      fmt.Println("done")
  }
  ```

  使用 goto 语句后，无须额外的变量就可以快速退出所有的循环。

- 统一错误处理

  - 多处错误处理存在代码重复时是非常棘手的,比如：

    ```go
    err := firstCheckError()
    if err != nil {
        fmt.Println(err)
        exitProcess()
        return
    }
     
    err = secondCheckError()
     
    if err != nil {
        fmt.Println(err)
        exitProcess()
        return
    }
     
    fmt.Println("done")
    ```

    加粗部分都是重复的错误处理代码。后期陆续在这些代码中如果添加更多的判断，就需要在每一块雷同代码中依次修改，极易造成疏忽和错误。

  - 如果使用 goto 语句来实现同样的逻辑：

    ```go
        err := firstCheckError()
        if err != nil {
            goto onExit
        }
        err = secondCheckError()
     
        if err != nil {
            goto onExit
        }
        fmt.Println("done")
        return
     
    onExit:
        fmt.Println(err)
        exitProcess()
    ```

#### 3. 使用goto注意事项

goto语句与标签之间不能有变量声明，否则编译错误。

```
import "fmt"

func main() {
    fmt.Println("start")
    goto flag
    var say = "hello oldboy"
    fmt.Println(say)
flag:
    fmt.Println("end")
```

编译错误

```
.\main.go:7:7: goto flag jumps over declaration of say at .\main.go:8:6
```

