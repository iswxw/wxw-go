## package/sync

> 来源：https://golang.google.cn/pkg/sync



### 概览

包同步提供了基本的同步原语，例如互斥锁。除了Once 和WaitGroup 类型之外，大多数类型都供低级库例程使用。更高级别的同步最好通过通道和通信来完成。

不应复制包含在此包中定义的类型的值。

- 文章出处：https://golang.google.cn/pkg/sync/

### sync.Map 并发安全

### sync.Mutex 互斥锁

### sync.RWMutex 读写锁

### sync.Pool 复用对象

**一句话总结：保存和复用临时对象，减少内存分配，降低 GC 压力。** 

#### 1. 使用示例

json 的反序列化在文本解析和网络通信过程中非常常见，当程序并发度非常高的情况下，短时间内需要创建大量的临时对象。而这些对象是都是分配在堆上的，会给 GC 造成很大压力，严重影响程序的性能。【[GC工作原理](https://geektutu.com/post/qa-golang-2.html#Q5-%E7%AE%80%E8%BF%B0-Go-%E8%AF%AD%E8%A8%80GC-%E5%9E%83%E5%9C%BE%E5%9B%9E%E6%94%B6-%E7%9A%84%E5%B7%A5%E4%BD%9C%E5%8E%9F%E7%90%86)】 

```go
type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})

func unmarsh() {
	stu := &Student{}
	json.Unmarshal(buf, stu)
}
```

Go 语言从 1.3 版本开始提供了对象重用的机制，即 sync.Pool。sync.Pool 是可伸缩的，同时也是并发安全的，其大小仅受限于内存的大小。sync.Pool 用于存储那些被分配了但是没有被使用，而未来可能会使用的值。这样就可以不用再次经过内存分配，可直接复用已有对象，减轻 GC 的压力，从而提升系统的性能。

sync.Pool 的大小是可伸缩的，高负载时会动态扩容，存放在池中的对象如果不活跃了会被自动清理。

#### 2. 如何使用

##### 2.1 声明对象池

只需要实现 New 函数即可。对象池中没有对象时，将会调用 New 函数创建。

```go
var studentPool = sync.Pool{
    New: func() interface{} { 
        return new(Student) 
    },
}
```

##### 2.2 Get & Put

```go
stu := studentPool.Get().(*Student)

json.Unmarshal(buf, stu)
studentPool.Put(stu)
```

- `Get()` 用于从对象池中获取对象，因为返回值是 `interface{}`，因此需要类型转换。
- `Put()` 则是在对象使用完毕后，返回对象池。

#### 3. 性能测试

##### 3.1 struct 反序列化

```go
func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}
```

测试结果如下：

```bash
$ go test -bench . -benchmem
goos: darwin
goarch: amd64
05_pkg: example/hpg-sync-pool
BenchmarkUnmarshal-8           1993   559768 ns/op   5096 B/op 7 allocs/op
BenchmarkUnmarshalWithPool-8   1976   550223 ns/op    234 B/op 6 allocs/op
PASS
ok      example/hpg-sync-pool   2.334s
```

在这个例子中，因为 Student 结构体内存占用较小，内存分配几乎不耗时间。而标准库 json 反序列化时利用了反射，效率是比较低的，占据了大部分时间，因此两种方式最终的执行时间几乎没什么变化。但是内存占用差了一个数量级，使用了 `sync.Pool` 后，内存占用仅为未使用的 234/5096 = 1/22，对 GC 的影响就很大了。

##### 3.2 bytes.Buffer

```go
var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

var data = make([]byte, 10000)

func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Write(data)
		buf.Reset()
		bufferPool.Put(buf)
	}
}

func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		buf.Write(data)
	}
}
```

测试结果如下：

```bash
BenchmarkBufferWithPool-8    8778160    133 ns/op       0 B/op   0 allocs/op
BenchmarkBuffer-8             906572   1299 ns/op   10240 B/op   1 allocs/op
```

这个例子创建了一个 `bytes.Buffer` 对象池，而且每次只执行一个简单的 `Write` 操作，存粹的内存搬运工，耗时几乎可以忽略。而内存分配和回收的耗时占比较多，因此对程序整体的性能影响更大。

#### 4. 标准库中使用

##### 4.1 fmt.Printf

Go 语言标准库也大量使用了 `sync.Pool`，例如 `fmt` 和 `encoding/json`。

以下是 `fmt.Printf` 的源代码(go/src/fmt/print.go)：

```go
// go 1.13.6

// pp is used to store a printer's state and is reused with sync.Pool to avoid allocations.
type pp struct {
    buf buffer
    ...
}

var ppFree = sync.Pool{
	New: func() interface{} { return new(pp) },
}

// newPrinter allocates a new pp struct or grabs a cached one.
func newPrinter() *pp {
	p := ppFree.Get().(*pp)
	p.panicking = false
	p.erroring = false
	p.wrapErrs = false
	p.fmt.init(&p.buf)
	return p
}

// free saves used pp structs in ppFree; avoids an allocation per invocation.
func (p *pp) free() {
	if cap(p.buf) > 64<<10 {
		return
	}

	p.buf = p.buf[:0]
	p.arg = nil
	p.value = reflect.Value{}
	p.wrappedErr = nil
	ppFree.Put(p)
}

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintf(format, a)
	n, err = w.Write(p.buf)
	p.free()
	return
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}
```

`fmt.Printf` 的调用是非常频繁的，利用 `sync.Pool` 复用 pp 对象能够极大地提升性能，减少内存占用，同时降低 GC 压力。

相关文章

1. https://www.cnblogs.com/qcrao-2018/p/12736031.html
2. https://geektutu.com/post/hpg-sync-pool.html

### sync.Cond 条件变量

#### 1. 使用场景

**一句话总结：`sync.Cond` 条件变量用来协调想要访问共享资源的那些 goroutine，当共享资源的状态发生变化的时候，它可以用来通知被互斥锁阻塞的 goroutine。** 

`sync.Cond` 基于互斥锁/读写锁，它和互斥锁的区别是什么呢？

- 互斥锁 `sync.Mutex` 通常用来保护临界区和共享资源，条件变量 `sync.Cond` 用来协调想要访问共享资源的 goroutine。

`sync.Cond` 经常用在多个 goroutine 等待，一个 goroutine 通知（事件发生）的场景。如果是一个通知，一个等待，使用互斥锁或 channel 就能搞定了。

**我们想象一个非常简单的场景：**  

```bash
有一个协程在异步地接收数据，剩下的多个协程必须等待这个协程接收完数据，才能读取到正确的数据。在这种情况下，如果单纯使用 chan 或互斥锁，那么只能有一个协程可以等待，并读取到数据，没办法通知其他的协程也读取数据。
```

这个时候，就需要有个全局的变量来标志第一个协程数据是否接受完毕，剩下的协程，反复检查该变量的值，直到满足要求。或者创建多个 channel，每个协程阻塞在一个 channel 上，由接收数据的协程在数据接收完毕后，逐个通知。总之，需要额外的复杂度来完成这件事。

Go 语言在标准库 sync 中内置一个 `sync.Cond` 用来解决这类问题。

#### 2. 四个核心方法

sync.Cond 的定义如下：

```go
// Cond implements a condition variable, a rendezvous point
// for goroutines waiting for or announcing the occurrence
// of an event.
//
// Each Cond has an associated Locker L (often a *Mutex or *RWMutex),
// which must be held when changing the condition and
// when calling the Wait method.
//
// A Cond must not be copied after first use.
type Cond struct {
	noCopy noCopy

	// L is held while observing or changing the condition
	L Locker

	notify  notifyList
	checker copyChecker
}
```

每个 Cond 实例都会关联一个锁 L（互斥锁 *Mutex，或读写锁 *RWMutex），当修改条件或者调用 Wait 方法时，必须加锁。

和 sync.Cond 相关的有如下几个方法：

##### 2.1 NewCond 创建实例

```go
// NewCond 创建 Cond 实例时，需要关联一个锁。
// NewCond returns a new Cond with Locker l.
func NewCond(l Locker) *Cond {
	return &Cond{L: l}
}
```

##### 2.2 Broadcast 广播唤醒所有协程

```go
// Broadcast 唤醒所有等待条件变量 c 的 goroutine，无需锁保护。
// Broadcast wakes all goroutines waiting on c.
// 
// It is allowed but not required for the caller to hold c.L
// during the call.
func (c *Cond) Broadcast() {
	c.checker.check()
	runtime_notifyListNotifyAll(&c.notify)
}
```

##### 2.3 Signal 唤醒一个协程

```go
// Signal 只唤醒任意 1 个等待条件变量 c 的 goroutine，无需锁保护。
// Signal wakes one goroutine waiting on c, if there is any.
//
// It is allowed but not required for the caller to hold c.L
// during the call.
func (c *Cond) Signal() {
	c.checker.check()
	runtime_notifyListNotifyOne(&c.notify)
}
```

##### 2.4 Wait 等待

```go
// Wait atomically unlocks c.L and suspends execution
// of the calling goroutine. After later resuming execution,
// Wait locks c.L before returning. Unlike in other systems,
// Wait cannot return unless awoken by Broadcast or Signal.
//
// Because c.L is not locked when Wait first resumes, the caller
// typically cannot assume that the condition is true when
// Wait returns. Instead, the caller should Wait in a loop:
//
//    c.L.Lock()
//    for !condition() {
//        c.Wait()
//    }
//    ... make use of condition ...
//    c.L.Unlock()
//
func (c *Cond) Wait() {
	c.checker.check()
	t := runtime_notifyListAdd(&c.notify)
	c.L.Unlock()
	runtime_notifyListWait(&c.notify, t)
	c.L.Lock()
}
```

调用 Wait 会自动释放锁 c.L，并挂起调用者所在的 goroutine，因此当前协程会阻塞在 Wait 方法调用的地方。如果其他协程调用了 Signal 或 Broadcast 唤醒了该协程，那么 Wait 方法在结束阻塞时，会重新给 c.L 加锁，并且继续执行 Wait 后面的代码。

对条件的检查，使用了 `for !condition()` 而非 `if`，是因为当前协程被唤醒时，条件不一定符合要求，需要再次 Wait 等待下次被唤醒。为了保险起见，使用 `for` 能够确保条件符合要求后，再执行后续的代码。

```go
c.L.Lock()
for !condition() {
    c.Wait()
}
... make use of condition ...
c.L.Unlock()
```

#### 3. 使用示例

接下来我们实现一个简单的例子，三个协程调用 `Wait()` 等待，另一个协程调用 `Broadcast()` 唤醒所有等待的协程。

```go
package main

import (
	"log"
	"sync"
	"time"
)

// 条件变量
var done = false

func main() {

	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	write("writer", cond)

	time.Sleep(time.Second * 3)
}

// 读
func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		c.Wait()
	}
	log.Println(name, "starts reading")
	c.L.Unlock()
}

// 写
func write(name string, c *sync.Cond) {
	log.Println(name, "starts writing")
	time.Sleep(time.Second)
	c.L.Lock()
	done = true
	c.L.Unlock()
    // 唤醒所有协程
	log.Println(name, "wakes all")
	c.Broadcast()
}
```

- `done` 即互斥锁需要保护的条件变量。
- `read()` 调用 `Wait()` 等待通知，直到 done 为 true。
- `write()` 接收数据，接收完成后，将 done 置为 true，调用 `Broadcast()` 通知所有等待的协程。
- `write()` 中的暂停了 1s，一方面是模拟耗时，另一方面是确保前面的 3 个 read 协程都执行到 `Wait()`，处于等待状态。main 函数最后暂停了 3s，确保所有操作执行完毕。

运行结果如下：

```bash
$ go run main.go

2021/12/05 21:55:18 writer starts writing
2021/12/05 21:55:19 writer wakes all
2021/12/05 21:55:19 reader3 starts reading
2021/12/05 21:55:19 reader1 starts reading
2021/12/05 21:55:19 reader2 starts reading
```

writer 接收数据花费了 1s，同步通知所有等待的协程。

相关资料

1. https://geektutu.com/post/hpg-sync-cond.html

### sync.Once 如何提升性能

#### 1. sync.Once 使用场景

`sync.Once` 是 Go 标准库提供的使函数只执行一次的实现，常应用于单例模式，例如初始化配置、保持数据库连接等。作用与 `init` 函数类似，但有区别。

- init 函数是当所在的 package 首次被加载时执行，若迟迟未被使用，则既浪费了内存，又延长了程序加载时间。
- sync.Once 可以在代码的任意位置初始化和调用，因此可以延迟到使用时再执行，并发场景下是线程安全的。

在多数情况下，`sync.Once` 被用于控制变量的初始化，这个变量的读写满足如下三个条件：

- 当且仅当第一次访问某个变量时，进行初始化（写）；
- 变量初始化过程中，所有读都被阻塞，直到初始化完成；
- 变量仅初始化一次，初始化完成后驻留在内存里。

`sync.Once` 仅提供了一个方法 `Do`，参数 f 是对象初始化函数。

```go
// Do calls the function f if and only if Do is being called for the
// first time for this instance of Once. In other words, given
// 	var once Once
// if once.Do(f) is called multiple times, only the first call will invoke f,
// even if f has a different value in each invocation. A new instance of
// Once is required for each function to execute.
//
// Do is intended for initialization that must be run exactly once. Since f
// is niladic, it may be necessary to use a function literal to capture the
// arguments to a function to be invoked by Do:
// 	config.once.Do(func() { config.init(filename) })
//
// Because no call to Do returns until the one call to f returns, if f causes
// Do to be called, it will deadlock.
//
// If f panics, Do considers it to have returned; future calls of Do return
// without calling f.
//
func (o *Once) Do(f func()) {
	// Note: Here is an incorrect implementation of Do:
	//
	//	if atomic.CompareAndSwapUint32(&o.done, 0, 1) {
	//		f()
	//	}
	//
	// Do guarantees that when it returns, f has finished.
	// This implementation would not implement that guarantee:
	// given two simultaneous calls, the winner of the cas would
	// call f, and the second would return immediately, without
	// waiting for the first's call to f to complete.
	// This is why the slow path falls back to a mutex, and why
	// the atomic.StoreUint32 must be delayed until after f returns.

	if atomic.LoadUint32(&o.done) == 0 {
		// Outlined slow-path to allow inlining of the fast-path.
		o.doSlow(f)
	}
}
```

#### 2. 使用示例

##### 2.1 sync.Once 读取配置

考虑一个简单的场景，函数 ReadConfig 需要读取环境变量，并转换为对应的配置。环境变量在程序执行前已经确定，执行过程中不会发生改变。ReadConfig 可能会被多个协程并发调用，为了提升性能（减少执行时间和内存占用），使用 `sync.Once` 是一个比较好的方式。

```go
package main

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

type Config struct {
	Server string
	Port int64
}

var (
	once sync.Once
	config *Config
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			_ = ReadConfig()
		}()
	}
	time.Sleep(time.Second)
}

// 读取配置
func ReadConfig() *Config{
	once.Do(func() {
		var err error
		config = &Config{Server: os.Getenv("TT_SERVER_URL")}
		config.Port,err = strconv.ParseInt(os.Getenv("TT_ROOT"),10,0)
		if err != nil {
			config.Port = 8080 // default port
		}
		log.Printf("init config: %+v",config)
	})
	return config
}
```

- 在这个例子中，声明了 2 个全局变量，once 和 config；
- config 是需要在 ReadConfig 函数中初始化的(将环境变量转换为 Config 结构体)，ReadConfig 可能会被并发调用。

如果 ReadConfig 每次都构造出一个新的 Config 结构体，既浪费内存，又浪费初始化时间。如果 ReadConfig 中不加锁，初始化全局变量 config 就可能出现并发冲突。这种情况下，使用 sync.Once 既能够保证全局变量初始化时是线程安全的，又能节省内存和初始化时间。

运行结果如下：

```bash
$ go run .
2021/01/07 23:51:49 init config
```

**`init config` 仅打印了一次，即 sync.Once 中的初始化函数仅执行了一次。**  

##### 2.2 标准库 sync.Once 的使用

`sync.Once` 在 Go 语言标准库中被广泛使用，我们可以简单地搜索一下：

```bash
$ grep -nr "sync\.Once" "$(dirname $(which go))/../src"
/usr/local/go/bin/../src/cmd/go/internal/cache/default.go:25:   defaultOnce  sync.Once
/usr/local/go/bin/../src/cmd/go/internal/cache/default.go:63:   defaultDirOnce sync.Once
/usr/local/go/bin/../src/cmd/go/internal/auth/netrc.go:23:      netrcOnce sync.Once
/usr/local/go/bin/../src/cmd/go/internal/modfetch/sumdb.go:53:  dbOnce sync.Once
/usr/local/go/bin/../src/cmd/go/internal/modfetch/sumdb.go:117: once    sync.Once
/usr/local/go/bin/../src/cmd/go/internal/modfetch/codehost/git.go:126:  refsOnce sync.Once
/usr/local/go/bin/../src/cmd/go/internal/modfetch/codehost/git.go:132:  localTagsOnce sync.Once
...
$ grep -nr "sync\.Once" "$(dirname $(which go))/../src" | wc -l
111
```

在 go1.13.6 版本的源码目录下，可以 grep 到 111 处使用。

比如 package `html` 中，对象 entity 只被初始化一次：

```go
var populateMapsOnce sync.Once
var entity           map[string]rune

func populateMaps() {
    entity = map[string]rune{
        "AElig;":                           '\U000000C6',
        "AMP;":                             '\U00000026',
        "Aacute;":                          '\U000000C1',
        "Abreve;":                          '\U00000102',
        "Acirc;":                           '\U000000C2',
        // 省略 2000 项
    }
}

func UnescapeString(s string) string {
    populateMapsOnce.Do(populateMaps)
    i := strings.IndexByte(s, '&')

    if i < 0 {
            return s
    }
    // 省略后续的实现
}
```

- 字典 `entity` 包含 2005 个键值对，若使用 init 在包加载时初始化，若不被使用，将会浪费大量内存。
- `html.UnescapeString(s)` 函数是线程安全的，可能会被用户程序在并发场景下调用，因此对 entity 的初始化需要加锁，使用 `sync.Once` 能保证这一点。

#### 3. sync.Once 原理

- 第一：保证变量仅被初始化一次，需要有个标志来判断变量是否已初始化过，若没有则需要初始化。

- 第二：线程安全，支持并发，无疑需要互斥锁来实现。

##### 3.1 源码实现

以下是 `sync.Once` 的源码实现，代码位于 `$(dirname $(which go))/../src/sync/once.go`：

```go
package sync

import (
    "sync/atomic"
)

type Once struct {
    done uint32
    m    Mutex
}

func (o *Once) Do(f func()) {
    if atomic.LoadUint32(&o.done) == 0 {
        o.doSlow(f)
    }
}

func (o *Once) doSlow(f func()) {
    o.m.Lock()
    defer o.m.Unlock()
    if o.done == 0 {
        defer atomic.StoreUint32(&o.done, 1)
        f()
    }
}
```

`sync.Once` 的实现与一开始的猜测是一样的，使用 `done` 标记是否已经初始化，使用锁 `m Mutex` 实现线程安全。

##### 3.2 done 为什么是第一个字段

字段 `done` 的注释也非常值得一看：

```go
type Once struct {
    // done indicates whether the action has been performed.
    // It is first in the struct because it is used in the hot path.
    // The hot path is inlined at every call site.
    // Placing done first allows more compact instructions on some architectures (amd64/x86),
    // and fewer instructions (to calculate offset) on other architectures.
    done uint32
    m    Mutex
}
```

其中解释了为什么将 done 置为 Once 的第一个字段：done 在热路径中，done 放在第一个字段，能够减少 CPU 指令，也就是说，这样做能够提升性能。

简单解释下这句话：

1. 热路径(hot path)是程序非常频繁执行的一系列指令，sync.Once 绝大部分场景都会访问 `o.done`，在热路径上是比较好理解的，如果 hot path 编译后的机器码指令更少，更直接，必然是能够提升性能的。
2. 为什么放在第一个字段就能够减少指令呢？因为结构体第一个字段的地址和结构体的指针是相同的，如果是第一个字段，直接对结构体的指针解引用即可。如果是其他的字段，除了结构体指针外，还需要计算与第一个值的偏移(calculate offset)。在机器码中，偏移量是随指令传递的附加值，CPU 需要做一次偏移值与指针的加法运算，才能获取要访问的值的地址。因为，访问第一个字段的机器代码更紧凑，速度更快。

相关资料

1. https://geektutu.com/post/hpg-sync-once.html

### sync.WaitGroup

#### 1. 场景引入

经常会看到以下代码：

```go
package main

import (
    "fmt"
    "time"
)

func main(){
    for i := 0; i < 100 ; i++{
        go fmt.Println(i)
    }
    time.Sleep(time.Second)
}
```

主线程为了等待goroutine都运行完毕，不得不在程序的末尾使用`time.Sleep()` 来睡眠一段时间，等待其他线程充分运行。

对于简单的代码，100个for循环可以在1秒之内运行完毕，`time.Sleep()` 也可以达到想要的效果。

但是对于实际生活的大多数场景来说，1秒是不够的，并且大部分时候我们都无法预知for循环内代码运行时间的长短。

这时候就不能使用`time.Sleep()` 来完成等待操作了。

可以考虑使用管道来完成上述操作：

```go
func main() {
    c := make(chan bool, 100)
    for i := 0; i < 100; i++ {
        go func(i int) {
            fmt.Println(i)
            c <- true
        }(i)
    }

    for i := 0; i < 100; i++ {
        <-c
    }
}
```

首先可以肯定的是使用管道是能达到我们的目的的，而且不但能达到目的，还能十分完美的达到目的。

但是管道在这里显得有些大材小用，因为它被设计出来不仅仅只是在这里用作简单的同步处理，在这里使用管道实际上是不合适的。而且假设我们有一万、十万甚至更多的`for`循环，也要申请同样数量大小的管道出来，对内存也是不小的开销。

对于这种情况，go语言中有一个其他的工具`sync.WaitGroup` 能更加方便的帮助我们达到这个目的。

`WaitGroup` 对象内部有一个计数器，最初从0开始，它有三个方法：`Add(), Done(), Wait()` 用来控制计数器的数量。

- `Add(n)` 把计数器设置为`n` 
- `Done()` 每次把计数器`-1` 
- `wait()` 会阻塞代码的运行，直到计数器地值减为`0`。

使用`WaitGroup` 将上述代码可以修改为：

```go
func main() {
    wg := sync.WaitGroup{}
    wg.Add(100)
    for i := 0; i < 100; i++ {
        go func(i int) {
            fmt.Println(i)
            wg.Done()
        }(i)
    }
    wg.Wait()
}
```

这里首先把`wg` 计数设置为`100`， 每个`for`循环运行完毕都把计数器减一，主函数中使用`Wait()` 一直阻塞，直到`wg`为零——也就是所有的`100`个`for`循环都运行完毕。相对于使用管道来说，`WaitGroup` 轻巧了许多。

#### 2. 注意事项

- 计数器不能为负值
- WaitGroup 对象不是一个引用类型

（1）**计数器不能为负值** 

我们不能使用`Add()` 给`wg` 设置一个负值，否则代码将会报错：

```bash
panic: sync: negative WaitGroup counter

goroutine 1 [running]:
sync.(*WaitGroup).Add(0x0, 0x0)
	D:/Installed/go1.17.1/src/sync/waitgroup.go:74 +0x105
main.main()
	D:/Project/wxw-go/src/com.wxw/01_basic_grammar/14_waitgroup/demo.go:15 +0x36
```

同样使用`Done()` 也要特别注意不要把计数器设置成负数了。

（2）**WaitGroup 对象不是一个引用类型** 

WaitGroup对象不是一个引用类型，在通过函数传值的时候需要使用地址：

```go
func main() {
    wg := sync.WaitGroup{}
    
    wg.Add(100)
    for i := 0; i < 100; i++ {
        go f(i, &wg)
    }
    wg.Wait()
}


// 一定要通过指针传值，不然进程会进入死锁状态
func f(i int, wg *sync.WaitGroup) { 
    fmt.Println(i)
    wg.Done()
}
```

相关资料

1. https://golang.google.cn/pkg/sync/#WaitGroup





