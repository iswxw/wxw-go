## 基础知识之流程控制

---

Go语言中最常用的流程控制有`if`和`for`，而`switch`和`goto`主要是为了简化代码、降低重复代码而生的结构，属于扩展类的流程控制。

### 分支结构`if else` 

**` if else`基本写法** 

Go语言中`if`条件判断的格式如下：

```go
if 表达式1 {
    分支1
} else if 表达式2 {
    分支2
} else{
    分支3
}
```

当表达式1的结果为`true`时，执行分支1，否则判断表达式2，如果满足则执行分支2，都不满足时，则执行分支3。 if判断中的`else if`和`else`都是可选的，可以根据实际需要进行选择。

Go语言规定与`if`匹配的左括号`{`必须与`if和表达式`放在同一行，`{`放在其他位置会触发编译错误。 同理，与`else`匹配的`{`也必须与`else`写在同一行，`else`也必须与上一个`if`或`else if`右边的大括号在同一行。

举个例子：

```go
func ifDemo1() {
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
```

**if else 特殊写法** 

if条件判断还有一种特殊的写法，可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断，举个例子：

```go
func ifDemo2() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
```

### 循环结构 ` for` 

Go 语言中的所有循环类型均可以使用`for`关键字来完成。

for循环的基本格式如下：

```go
for 初始语句;条件表达式;结束语句{
    循环体语句
}
```

条件表达式返回`true`时循环体不停地进行循环，直到条件表达式返回`false`时自动退出循环。

```go
func forDemo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
```

for循环的初始语句可以被忽略，但是初始语句后的分号必须要写，例如：

```go
func forDemo2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}
```

for循环的初始语句和结束语句都可以省略，例如：

```go
func forDemo3() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
```

这种写法类似于其他编程语言中的`while`，在`while`后添加一个条件表达式，满足条件表达式时持续循环，否则结束循环。

#### 1. 无限循环

```go
for {
    循环体语句
}
```

for循环可以通过`break`、`goto`、`return`、`panic`语句强制退出循环。

#### 2. 键值循环（for range） 

Go语言中可以使用`for range`遍历数组、切片、字符串、map 及通道（channel）。 通过`for range`遍历的返回值有以下规律：

1. 数组、切片、字符串返回索引和值。
2. map返回键和值。
3. 通道（channel）只返回通道内的值。

### ` switch case` 

用`switch`语句可方便地对大量的值进行条件判断。

```go
func switchDemo1() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}
```

Go语言规定每个`switch`只能有一个`default`分支。

一个分支可以有多个值，多个case值中间使用英文逗号分隔。

```go
func testSwitch3() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}
```

分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量。例如：

```go
func switchDemo4() {
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}
```

**`fallthrough`语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。** 

```go
func switchDemo5() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
```

输出：

```bash
a
b
```

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

### 跳出循环 break

`break`语句可以结束`for`、`switch`和`select`的代码块。

`break`语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的`for`、`switch`和 `select`的代码块上。 举个例子：

```go
func breakDemo1() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}
```

### 继续下次循环continue

`continue`语句可以结束当前循环，开始下一次的循环迭代过程，仅限在`for`循环内使用。

在 `continue`语句后添加标签时，表示开始标签对应的循环。例如：

```go
func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}
```

**案例：打印9*9乘法表** 

```
package main

import (
	"fmt"
)
// 编写代码打印9*9乘法表。
func main() {
    for i := 1; i < 10; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%v*%v=%v\t", j, i, i*j)
        }
        fmt.Println()
    }
}
```









