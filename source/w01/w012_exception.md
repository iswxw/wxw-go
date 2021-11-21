## 基础知识之异常处理

---

### 延迟语句：`defer` 

Go 中有几个比较特殊的关键字，如 defer，尤其 `defer+panic+recover`的组合可以发挥出 java 中 `try...catch...fanilly` 的作用，功能非常强大，值的去深入学习。同时他们每个又有自身的特性，这一节我们先来理解一下 `defer`。

**`defer`** 关键字在 go 中的使用率算是非常高的，类似于 `finally` 与 析构函数的作用，用来做方法的善后工作。

**先抛砖引玉 defer的延迟调用：** 

defer特性

1. 关键字 defer 用于注册延迟调用。
2. 这些调用直到 return 前才被执性。因此，可以用来做资源清理。
3. 多个defer语句，按先进后出的方式执行（从后往前执行）。
4. defer语句中的变量，在defer声明时就决定了。

defer用途

1. 关闭文件句柄
2. 锁资源释放
3. 数据库连接释放

#### 1. 怎么使用

Go语言中的`defer`语句会将其后面跟随的语句进行延迟处理。在`defer`归属的函数即将返回时，将延迟处理的语句按`defer`定义的逆序进行执行，也就是说，先被`defer`的语句最后被执行，最后被`defer`的语句，最先被执行。

举个例子：

```go
func main() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}
```

输出结果：

```go
start
end
3
2
1
```

由于`defer`语句延迟调用的特性，所以`defer`语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等。

#### 2. 理解defer的执行时机

在Go语言的函数中`return`语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。而`defer`语句执行的时机就在返回值赋值操作后，RET指令执行前。具体如下图所示：

<img src="https://www.liwenzhou.com/images/Go/func/defer.png" alt="defer执行时机" style="zoom: 80%;" />  

下面有一段相对复杂的 defer 定义，这个方法里面有定义 4 个 defer , 分别以不同的方式来定义 defer 后的函数

```go
func deferTest() {
	fish := 0
	
	defer func() {
		fmt.Println("d1: ", fish)	
	}()
	
	defer fmt.Println("d2: ", fish) 
	
	defer func(fish int) {
		fish += 2						
		fmt.Println("d3: ", fish)	
	}(fish)							    
	
	defer func() {
		fmt.Println("d4: ", fish)	
		fish +=2						
	}()
  
	fish++
}
```

这里可以先不要往下看答案，或者直接本地运行看结果，可以先猜猜每个结果的值，然后去跟真实结果做对比

如果是新手的话，可能会这么认为：d1: 1, d2: 1, d3: 3, d4: 3 , 这肯定是错的啦，要不然我还讲什么呢？哈哈

首先要明白一点：

- **一个方法中声明了多个defer, 那么 defer 是按顺序从后往前执行的** 

根据这个规则，我们再猜想一次结果：d4: 1, d3:5, d2: 5, d1: 5， 真的对么？ 我们来看看真正的结果

```bash
d4:  1
d3:  2
d2:  0
d1:  3
```

这个结果是不是很奇怪，无论你怎么想都想不出来，因为它牵涉到 defer 声明时的一些特殊规则，defer 后面的函数对外部参数有两种引用方式：

- **参数传递：在defer声明时，即将值传递给defer,并缓存起来，调用defer的时候使用缓存值进行计算**
- **直接引用：根据上下文确定当前值，作用域与外部变量相同** 

`d1,d4` 是在闭包里面直接引用，而 `d2,d3` 是经过参数传递来饮用的，因此 `d2,d3` 传递进去的初始值都是 0

```go
func deferTest() {
  // 1. 声明 fish=0
	fish := 0
  
	// 直接引用(闭包)
  //
	defer func() {
    // 8. 打印 fish=3
		fmt.Println("d1: ", fish)			// 3
	}()
  
	// 参数传递
  // 2. 传递 fish=0
  // 7. 打印内部 fish 0
	defer fmt.Println("d2: ", fish) // 0
  
	// 参数传递(闭包)
  // 3. 传递 fish=0
	defer func(fish int) {
    // 6. 内部 fish 0 + 2 =2
		fish += 2											// fish 只作用于内部
    // 7. 打印内部 fish
		fmt.Println("d3: ", fish)			// 2
	}(fish)							    				// 声明时传递, fish = 0
  
	// 直接引用(闭包)
	defer func() {
    // 4. 打印 fish = 1
		fmt.Println("d4: ", fish)			// 1
    // 5. fish = 3
		fish +=2											// fish = 3, 作用域与外部的相同
	}()
  
  // 3. fish=1
	fish++
}
```

#### 3. defer与return的关系

再看一段代码，猜猜下面3个方法的返回值各是多少

```go
func defer1() (res int) {
	defer func() {
		res ++
	}()
	return 10
}

func defer2() (res int) {
	sb := 10
	defer func() {
		sb += 5
	}()
	return sb
}

func defer3() (res int) {
	res = 2
	defer func(res int) {
		res += 2
		fmt.Println("内部 res ", res)  
	}(res)		
	return 10
}
```

接下来我们一个个看，也许你会认为`defer1` 的结果是10，其实是11，在这里我们先明白一个概念：

**return 语句并不是一条原子指令**，有没有被震慑到！！！

**有返回值的且带有 defer 函数的方法中， return 语句执行顺序：**  

```go
1. 返回值赋值
2. 调用 defer 函数 (在这里是可以修改返回值的)
3. return 返回值
```

那 `defer1` 方法可以这样解析：

```go
// 不是 10 , 是 11
func defer1() (res int) {
	defer func() {
		// 2. res = 10 + 1 = 11
		res ++
	}()
	// 1. res = 10
  // 3. return res
	return 10
}
```

也就是说最后的值为 11

`defer2` 返回值不是15，而是10，这样解析

```go
// 不是15, 是10
func defer2() (res int) {
	// 1. sb = 10
	sb := 10
	defer func() {
		// 3. sb = 15, 但是 res = 10
		sb += 5
	}()
	// 2. res = sb = 10
  // 4. return res(10)
	return sb
}
```

`defer3` 的返回值不是12， 而是10, 解析如下：

```go
// 不是12，是10
func defer3() (res int) {
  // 1. res=2
	res = 2
	defer func(res int) {
    // 3. 内部res为形参，不影响外边的值 res=2+2=4
		res += 2
		fmt.Println("内部 res ", res)   // 4
	}(res)		// defer 参数的值是在声明的时候确定的，也就是只有 defer 之前的语句会影响这个值
	// 2. res = 10
  // 4. return res(10)
	return 10
}
```

#### 4. 深入理解defer



通过实例加深理解，我们先看看一段代码

- **defer 执行顺序可以理解为 先进后出**   

  ```go
  package main
  
  import "fmt"
  
  func main() {
      var users [5]struct{}
      for i := range users {
          defer fmt.Println(i)
      }
  }
  
  ----
  输出：4 3 2 1 0 ,defer 是先进后出,这个输出没啥好说的
  ```

- **defer后面的语句在执行的时候，函数调用的参数会被保存起来，但是不执行。也就是复制了一份。但是并没有说struct这里的*指针如何处理** 

  我们把上面的代码改下：

  **[defer 换上闭包]**：  

  ```go
  package main
  
  import "fmt"
  
  func main() {
      var users [5]struct{}
      for i := range users {
          defer func() { fmt.Println(i) }()
      }
  }
  ```

  输出：4 4 4 4 4，很多人也包括我。预期的结果不是 4 3 2 1 0 吗？

  官网对defer 闭包的使用大致是这个意思：

  - 函数正常执行,由于闭包用到的变量 i 在执行的时候已经变成4,所以输出全都是4。那么 如何正常输出预期的 4 3 2 1 0 呢？

  **[不用闭包，换成函数：]** 

  ```
  package main
  
  import "fmt"
  
  func main() {
      var users [5]struct{}
      for i := range users {
          defer Print(i)
      }
  }
  func Print(i int) {
      fmt.Println(i)
  }
  ```

   函数正常延迟输出：4 3 2 1 0。

  我们再举一个可能一不小心会犯错的例子：

  **[defer调用引用结构体函数]** 

  ```
  package main
  
  import "fmt"
  
  type Users struct {
      name string
  }
  
  func (t *Users) GetName() { // 注意这里是 * 传地址 引用Users
      fmt.Println(t.name)
  }
  func main() {
      list := []Users{{"乔峰"}, {"慕容复"}, {"清风扬"}}
      for _, t := range list {
          defer t.GetName()
      }
  }
  ```

  输出：清风扬 清风扬 清风扬。

  这个输出并不会像我们预计的输出：清风扬 慕容复 乔峰

  可是按照前面的go defer函数中的使用说明,应该输出清风扬 慕容复 乔峰才对啊？

  那我们换一种方式来调用一下

  ```
  package main
  
  import "fmt"
  
  type Users struct {
      name string
  }
  
  func (t *Users) GetName() { // 注意这里是 * 传地址 引用Users
      fmt.Println(t.name)
  }
  func GetName(t Users) { // 定义一个函数，名称自定义
      t.GetName() // 调用结构体USers的方法GetName
  }
  func main() {
      list := []Users{{"乔峰"}, {"慕容复"}, {"清风扬"}}
      for _, t := range list {
          defer GetName(t)
      }
  }
  ```

  输出：清风扬 慕容复 乔峰。

  这个时候输出的就是所谓"预期"滴了

  当然,如果你不想多写一个函数,也很简单,可以像下面这样（改2处）,同样会输出清风扬 慕容复 乔峰

  ```
  package main
  
  import "fmt"
  
  type Users struct {
      name string
  }
  
  func (t *Users) GetName() { // 注意这里是 * 传地址 引用Users
      fmt.Println(t.name)
  }
  func GetName(t Users) { // 定义一个函数，名称自定义
      t.GetName() // 调用结构体USers的方法GetName
  }
  func main() {
      list := []Users{{"乔峰"}, {"慕容复"}, {"清风扬"}}
      for _, t := range list {
          t2 := t // 定义新变量t2 t赋值给t2
          defer t2.GetName()
      }
  }
  ```

  输出：清风扬 慕容复 乔峰。

  通过以上例子

  我们可以得出下面的结论：

  - defer后面的语句在执行的时候，函数调用的参数会被保存起来，但是不执行。也就是复制了一份。但是并没有说struct这里的*指针如何处理，

  通过这个例子可以看出go语言并没有把这个明确写出来的this指针(比如这里的* Users)当作参数来看待。

  到这里有滴朋友会说。看似多此一举的声明，直接去掉指针调用 t *Users改成 t Users 不就行了？

  ```
package main
  
  import "fmt"
  
  type Users struct {
      name string
  }
  
  func (t Users) GetName() { // 注意这里是 * 传地址 引用Users
      fmt.Println(t.name)
  }
  
  func main() {
      list := []Users{{"乔峰"}, {"慕容复"}, {"清风扬"}}
      for _, t := range list {
          defer t.GetName()
      }
  }
  ```
  
  输出：清风扬 慕容复 乔峰。这就回归到上面的 defer 函数非引用调用的示例了。

  所以这里我们要注意：

  **defer后面的指针函数和普通函数的调用区别** ,很容易混淆出错。

- **多个 defer 注册，按 FILO 次序执行 ( 先进后出 )。 哪怕函数或某个延迟调用发生错误，这些调用依旧会被执行**，

  我们看看这一段代码：

  ```
  package main
  
  func users(i int) {
      defer println("北丐")
      defer println("南帝")
  
      defer func() {
          println("西毒")
          println(10 / i) // 异常未被捕获，逐步往外传递，最终终止进程。
      }()
  
      defer println("东邪")
  }
  
  func main() {
      users(0)
      println("武林排行榜,这里不会被输出哦")
  }
  
  ----
  东邪
  西毒
  南帝
  北丐
  panic: runtime error: integer divide by zero
  goroutine 1 [running]:
  main.users.func1(0x0)
  ```

  我们发现函数中异常，最后才捕获输出，但是一旦捕获了异常，后面就不会再执行了，即终止了程序。

- **延迟调用参数在求值或复制，指针或闭包会 "延迟" 读取。** 

  ```
  package main
  
  func test() {
      x, y := "乔峰", "慕容复"
  
      defer func(s string) {
          println("defer:", s, y) // y 闭包引用 输出延迟和的值，即y+= 后的值=慕容复第二
          
      }(x) // 匿名函数调用，传送参数x 被复制,注意这里的x 是 乔峰,而不是下面的 x+= 后的值
  
      x += "第一"
      y += "第二"
      println("x =", x, "y =", y)
  }
  
  func main() {
      test()
  }
  
  ---
  x = 乔峰第一 
  y = 慕容复第二
  defer: 乔峰 慕容复第二
  ```

- **defer与return** 

  ```
  package main
  
  import "fmt"
  
  func Users() (s string) {
  
      s = "乔峰"
      defer func() {
          fmt.Println("延迟执行后:"+s)
      }()
  
      return "清风扬"
  }
  
  func main() {
      Users() // 输出:延迟执行后:清风扬
  }
  ```

  解释：

  - 在有命名返回值的函数中（这里命名返回值为 s），执行 return "风清扬" 的时候实际上已经将s 的值重新赋值为 风清扬。

    所以defer 匿名函数 输出结果为 风清扬 而不是 乔峰。

- **在错误的位置使用 defer,来一段不严谨滴代码：** 

  ```
  package main
  
  import "net/http"
  
  func request() error {
      res, err := http.Get("http://www.google.com") // 不FQ的情况下。是无法访问滴
      defer res.Body.Close()
      if err != nil {
          return err
      }
  
      // ..继续业务code...
  
      return nil
  }
  
  func main() {
      request()
  }
  
  输出
  -------- 
  panic: runtime error: invalid memory address or nil pointer dereference
  [signal 0xc0000005 code=0x0 addr=0x40 pc=0x5e553e]
  ```

  Why？

  - 因为在这里我们并没有检查我们的请求是否成功执行，当它失败的时候，我们访问了 Body 中的空变量 res ，所以会抛出异常。

  怎么优化呢？

  - 我们应该总是在一次成功的资源分配下面使用 defer ，简单点说就是：当且仅当 http.Get 成功执行时才使用 defer.

  ```
  package main
  
  import "net/http"
  
  func request() error {
      res, err := http.Get("http://www.google.com")
      if res != nil {
          defer res.Body.Close()
      }
  
      if err != nil {
          return err
      }
  
      // ..继续业务code...
  
      return nil
  }
  
  func main() {
      request()
  }
  ```

  这样，当有错误的时候，err 会被返回，否则当整个函数返回的时候，会关闭 res.Body 。

  解释：

  - 在这里，同样需要检查 res 的值是否为 nil ，这是 http.Get 中的一个警告。

  通常情况下，出错的时候，返回的内容应为空并且错误会被返回，可当你获得的是一个重定向 error 时， res 的值并不会为 nil ，

  但其又会将错误返回。所以上面的代码保证了无论如何 Body 都会被关闭。

另外我们再聊下关于文件的defer close。我们看一段代码：

- **在这里，f.Close() 可能会返回一个错误，可这个错误会被我们忽略掉** 

  ```
  package main
  
  import "os"
  
  func open() error {
      f, err := os.Open("result.json") // 确保文件名存在
      if err != nil {
          return err
      }
  
      if f != nil {
          defer f.Close()
      }
  
      // ..code...
  
      return nil
  }
  
  func main() {
      open()
  }
  ```

  表面上看似没问题，其实f.Close可能关闭文件失败，我们优化下：

  ```
  package main
  
  import "os"
  
  func open() error {
      f, err := os.Open("result.json")
      if err != nil {
          return err
      }
  
      if f != nil {
          defer func() {
              if err := f.Close(); err != nil {
                  return
              }
          }()
      }
  
      // ..code...
  
      return nil
  }
  
  func main() {
      open()
  }
  ```

  如果有代码洁癖优化强迫症滴，哈哈。这里我们还可以优化下，可以通过命名的返回变量来返回 defer 内的错误。 如下：

  ```
  package main
  
  import "os"
  
  func open() (err error) {
      f, err := os.Open("result.json")
      if err != nil {
          return err
      }
  
      if f != nil {
          defer func() {
              if ferr := f.Close(); ferr != nil {
                  err = ferr //这里 通过命名的返回变量ferr赋值给err 来返回 defer 内的错误
              }
          }()
      }
  
      // ..code...
  
      return nil
  }
  
  func main() {
      open()
  }
  ```

最后一个容易忽视的问题：**如果你尝试使用相同的变量释放不同的资源，那么这个操作可能无法正常执行**

神马意思？继续看：

```
package main

import (
    "fmt"
    "os"
)

func open() error {
    f, err := os.Open("result.json")
    if err != nil {
        return err
    }
    if f != nil {
        defer func() {
            if err := f.Close(); err != nil {
                fmt.Printf("延迟关闭文件result.json 错误 %v\n", err)
            }
        }()
    }

    // ..code...

    f, err = os.Open("result2.json")
    if err != nil {
        return err
    }
    if f != nil {
        defer func() {
            if err := f.Close(); err != nil {
                fmt.Printf("延迟关闭文件result2.json 错误 %v\n", err)
            }
        }()
    }

    return nil
}

func main() {
    open()
}

输出
----
延迟关闭文件result.json 错误 close result2.json: file already closed
```

**结论：** 

当延迟函数执行时，只有最后一个变量会被用到，因此，f 变量 会成为最后那个资源 (result2.json)。

而且两个 defer 都会将这个资源作为最后的资源来关闭,也就是优先关闭了result2.json后，再执行第一个defer Close result1.json的时候，

其实还是在关闭result2.json.这样重复关闭同一个文件导致错误异常。肿么解决？很好办？用io.Closer属性

```
package main

import (
    "fmt"
    "io"
    "os"
)

func open() error {
    f, err := os.Open("result.json")
    if err != nil {
        return err
    }
    if f != nil {
        defer func(f io.Closer) { // 注意修改滴地方
            if err := f.Close(); err != nil {
                fmt.Printf("延迟关闭文件result.json 错误 %v\n", err)
            }
        }(f) // 注意修改滴地方
    }

    // ..code...

    f, err = os.Open("result2.json")
    if err != nil {
        return err
    }
    if f != nil {
        defer func(f io.Closer) {// 注意修改滴地方
            if err := f.Close(); err != nil {
                fmt.Printf("延迟关闭文件result2.json 错误 %v\n", err)
            }
        }(f)// 注意修改滴地方
    }

    return nil
}

func main() {
    open()
}
```

### 捕获异常：recover

`For example`  

```go
func recoverDemo1() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Runtime error caught: %v", r)
		}
	}()
    
    // 调用foo() 方法
    foo()
}

func foo(){
    panic("test panic")
}
```

无论`  foo()` 中是否触发了错误处理流程，该匿名defer函数都将在函数退出时得到执行。假 如 `  foo()` 中触发了错误处理流程，recover()函数执行将使得该错误处理过程终止。如果错误处 理流程被触发时，程序传给panic函数的参数不为nil，则该函数还会打印详细的错误信息。

`Golang`  异常的抛出与捕获，依赖两个内置函数：

- panic：抛出异常，使程序崩溃
- recover：捕获异常，恢复程序或做收尾工作

revocer 调用后，抛出的 panic 将会在此处终结，不会再外抛，但是 recover，并不能任意使用，它有 **强制要求，必须得在 defer 下才能发挥用途** 。

#### 1.  recover 仅在延迟函数 defer 中有效

先来看个例子：

```go
package main

import "fmt"

func main() {
	recover()     // 无任何作用
	panic("停止运行")
	recover()     // 不会执行到
	fmt.Println("结束")
	
	// 输出
	// panic: 停止运行
	// goroutine 1 [running]:exit status 2
}
```

修改下代码：

```go
package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("捕获到异常:", recover())
	}()
	panic("手动抛出异常")
	
	// 输出
	// 捕获到异常: 手动抛出异常
}
```

**结论：**

 recover 仅在延迟函数 defer 中有效，在正常的执行过程中，调用 recover 会返回 nil 并且没有其他任何效果.

**重要的事再说一遍：仅当在一个defer函数中被完成时，调用recover()才生效。** 

#### 2. recover在defer中直接调用才生效

举个例子

```go
package main

import "fmt"

func doRecover() {
	fmt.Println("捕获到异常 =>", recover()) //输出: 捕获到异常 => <nil>
}
func main() {
	defer func() {
		doRecover() //注意：这里间接使用函数，在函数中调用了recover()函数，
		// panic 没有恢复,没有捕获到错误信息
	}()
	panic("手动抛出异常")
}
```

输出

```bash
捕获到异常 => <nil>
panic: 手动抛出异常

goroutine 1 [running]:
main.main()
```

**总结：panic配合recover使用，recover要在defer函数中直接调用才生效。** 

#### 3. 案例解析

`panic`可以在任何地方引发，但`recover`只有在`defer`调用的函数中有效。 首先来看一个例子：

```go
func funcA() {
	fmt.Println("func A")
}

func funcB() {
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}
func main() {
	funcA()
	funcB()
	funcC()
}
```

输出：

```bash
func A
panic: panic in B

goroutine 1 [running]:
main.funcB(...)
        .../code/func/main.go:12
main.main()
        .../code/func/main.go:20 +0x98
```

程序运行期间`funcB`中引发了`panic`导致程序崩溃，异常退出了。

这个时候我们就可以通过`recover`将程序恢复回来，继续往后执行。

```go
func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}
func main() {
	funcA()
	funcB()
	funcC()
}
```

**注意：**

1. `recover()`必须搭配`defer`使用。
2. `defer`一定要在可能引发`panic`的语句之前定义。





















