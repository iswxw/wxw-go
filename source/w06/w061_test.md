## 工程实践之测试



### 单元测试

Go语言拥有一套单元测试和性能测试系统，仅需要添加很少的代码就可以快速测试一段需求代码。

单元测试（unit testing），是指对软件中的最小可测试单元进行检查和验证。

#### 1. 单元测试规则

主要规则如下：

1. 单元测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头
2. 每个单元测试函数的参数必须为*testing.T，参数t用于报告测试是否失败以及日志信息。

```go
func TestAdd(t *testing.T){ ... }
func TestSum(t *testing.T){ ... }
func TestLog(t *testing.T){ ... }
```

`testing.T参数`的拥有的方法如下：

```go

func (c *T) Error(args ...interface{})
func (c *T) Errorf(format string, args ...interface{})
func (c *T) Fail()
func (c *T) FailNow()
func (c *T) Failed() bool
func (c *T) Fatal(args ...interface{})
func (c *T) Fatalf(format string, args ...interface{})
func (c *T) Log(args ...interface{})
func (c *T) Logf(format string, args ...interface{})
func (c *T) Name() string
func (t *T) Parallel()
func (t *T) Run(name string, f func(t *T)) bool
func (c *T) Skip(args ...interface{})
func (c *T) SkipNow()
func (c *T) Skipf(format string, args ...interface{})
func (c *T) Skipped() bool
```

#### 2. 测试方法

假设现在开发了1个单元（函数）， 该单元的功能是对string类型的变量进行split。

主要的源代码：

- `unit.go` 

  ```go
  package main
  
  import "strings"
  
  //NewSplit 切割字符串
  //example:
  //abc,b=>[ac]
  func NewSplit(str, sep string) (des []string) {
  	index := strings.Index(str, sep)
  	for index > -1 {
  		sectionBefor := str[:index]
  		des = append(des, sectionBefor)
  		str = str[index+1:]
  		index = strings.Index(str, sep)
  	}
  	//最后1
  	des = append(des, str)
  	return
  }
  ```

- `unit_test.go` 

  ```go
  
  //测试用例1：以字符分割
  func TestSplit(t *testing.T) {
  	got := NewSplit("123N456", "N")
  	want := []string{"123", "456"}
  	//DeepEqual比较底层数组
  	if !reflect.DeepEqual(got, want) {
  		//如果got和want不一致说明你写得代码有问题
  		t.Errorf("The values of %v is not %v\n", got, want)
  	}
  }
  
  //测试用例2：以标点符号分割
  func TestPunctuationSplit(t *testing.T) {
  	got := NewSplit("a:b:c", ":")
  	want := []string{"a", "b", "c"}
  	if !reflect.DeepEqual(got, want) {
  		t.FailNow()//出错就stop别往下测了！
  	}
  
  }
  ```

测试结果：

```bash
C:\Project\wxw-go\src\com.wxw\project_actual\actual_test\unit_test>go test -v
=== RUN   TestSplit
--- PASS: TestSplit (0.00s)
=== RUN   TestPunctuationSplit
--- PASS: TestPunctuationSplit (0.00s)
PASS
ok      src/com.wxw/04_project_actual/src/com.wxw/04_project_actual/w_test/unit_test     0.120s
```

当然，我们可以继续优化测试代码！利用结构体组织测试数据把多个测试用例合到一起，在1个函数内对1组测试用例进行统一测试。

**组测试**  

```go

package splitString
 
import (
    "reflect"
    "testing"
)
 
//测试组:在1个函数中写多个测试用例，切支持灵活扩展！
 
type testCase struct {
    str      string
    separate string
    want     []string
}
 
var testGroup = []testCase{
    //测试用例1:单个英文字母
    testCase{
        str:      "123N456",
        separate: "N",
        want:     []string{"123", "456"},
    },
    //测试用例2：符号
    testCase{
        str:      "a:b:c",
        separate: ":",
        want:     []string{"a", "b", "c"},
    },
    //测试用例3：多个英文字母
    testCase{
        str:      "hellowbsdjshdworld",
        separate: "bsdjshd",
        want:     []string{"hellow", "world"},
    },
    //测试用例4:单个汉字
    testCase{
        str:      "山西运煤车煤运西山",
        separate: "山",
        want:     []string{"西运煤车煤运西"},
    },
 
    //测试用例4：多个汉字
    testCase{
        str:      "京北北京之北",
        separate: "北京",
        want:     []string{"京北", "之北"},
    },
}
 
func TestSplit(t *testing.T) {
    for _, test := range testGroup {
        got := Newsplit(test.str, test.separate)
        if !reflect.DeepEqual(got, test.want) {
            t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
        }
    }
 
}
```

**子测试** 

基于测试组对测试代码再次进行优化，利用使用t *testing.T参数的run方法去执行测试用例。

这种方法可以针对测试组里的1个测试用例进行单独测试，所以也叫子测试。

```go

package splitString
 
import (
    "reflect"
    "testing"
)
 
//子测试
type testCase struct {
    str      string
    separate string
    want     []string
}
 
var testGroup = map[string]testCase{
    "punctuation": testCase{
        str:      "a:b:c",
        separate: ":",
        want:     []string{"a", "b", "c"},
    },
    "sigalLetter": testCase{
        str:      "123N456",
        separate: "N",
        want:     []string{"123", "456"},
    },
 
    "MultipleLetter": testCase{
        str:      "hellowbsdjshdworld",
        separate: "bsdjshd",
        want:     []string{"hellow", "world"},
    },
    "singalRune": testCase{
        str:      "山西运煤车煤运西山",
        separate: "山",
        want:     []string{"西运煤车煤运西"},
    },
    "multiplRune": testCase{
        str:      "京北北京之北",
        separate: "北京",
        want:     []string{"京北", "之北"},
    },
}
 
//测试用例函数
func TestSplit(t *testing.T) {
    for name, test := range testGroup {
        //使用t参数的run方法
        t.Run(name, func(t *testing.T) {
            got := Newsplit(test.str, test.separate)
            if !reflect.DeepEqual(got, test.want) {
                t.Fatalf("失败！got:%#v want:%#v\n", got, test.want)
            }
        })
    }
}
```

#### 3. 测试功能

测试覆盖率是你的代码被测试套件覆盖的百分比。

通常我们使用的都是语句的覆盖率，也就是在测试中至少被运行一次的代码占总代码的比例。

Go提供内置功能来检查你的代码覆盖率。我们可以使用 ` go test -cover` 

```bash
// 把测试覆盖率的详细信息输出到文件
go test -cover -coverprofile=test_report.out

// 把测试报告输出到文件，就是为了分析测试结果，go内置的工具支持以HTML的方式打开测试报告文件！
go tool cover -html=test_report.out
```

- 下图中每个用绿色标记的语句块表示被覆盖了，而红色的表示没有被覆盖。  

<img src="https://img2020.cnblogs.com/blog/1122865/202005/1122865-20200510162023356-1656223051.png" alt="img" style="zoom: 67%;" />   

- 覆盖率测试

  ```bash
  D:\goproject\src\LearningTest\splitString> go test -cover
  PASS
  coverage: 100.0% of statements
  ok      LearningTest/splitString        0.042s
   
  D:\goproject\src\LearningTest\splitString>
  ```

  



### 基准测试

当我们尝试去优化代码的性能时，首先得知道当前的性能怎么样。Go 语言标准库内置的 testing 测试框架提供了基准测试(benchmark)的能力，能让我们很容易地对某一段代码进行性能测试。

性能测试受环境的影响很大，为了保证测试的可重复性，在进行性能测试时，尽可能地保持测试环境的稳定。



在Go语言中，提供了测试函数性能（CPU和Memory）的测试方法，基准测试。

基准测试（`Benchmark Test`）主要用来测试CPU和内存的效率问题，来评估被测代码的性能。测试人员可以根据这些性能指标的反馈，来优化我们的代码，进而提高性能问题。

本节主要来介绍，基准测试的使用方法和性能指标的分析。

#### 1. 基准测试规则

主要规则如下：

```bash
1.基准测试的代码文件必须以_test.go结尾。
2.基准测试的函数必须以Benchmark开头。
3.基准测试函数必须接受一个指向testing.B类型的指针作为唯一参数。
4.在测试执行之前，需要调用b.ResetTimer（重置计时器）来重置时间，以便过滤掉测试之前代码所消耗的时间。
5.测试代码需要写在for循环中，并且循环中的最大值是b.N。
```

#### 2. 测试方法

##### 2.1 测试代码

文件结构为： benchmark.go, benchmark_test.go

<img src="https://static.studygolang.com/191014/92319c5203d5a0ca7d0acf65bb918d14.png" alt="benchmark" style="zoom: 50%;" /> 

- benchmark.go

  ```go
  package main
  
  import "fmt"
  
  // 基准测试 源代码
  
  // 测试方法 case_1
  func HandleType(flag int) string {
  	fmt.Printf("HandleType: %v",flag)
  	switch flag {
  	case 0:
  		return "Add"
  	case 1:
  		return "Sub"
  	case 2:
  		return "Multiply"
  	case 3:
  		return "Division"
  	}
  	return "NotExist"
  }
  ```

- benchmark_test.go

  ```go
  package main
  
  import (
  	"fmt"
  	"testing"
  )
  
  // 基准测试 测试代码 case_1
  func BenchmarkHandleType(b *testing.B) {
  	flag := 1
  	b.ResetTimer()
  	for i := 0; i < b.N; i++ {
  		fmt.Println("BenchMark Test:", HandleType(flag))
  	}
  }
  ```

##### 2.2 运行基准测试

**使用命令** 

```bash
$ go test -bench=. -benchtime=5s -benchmem -run=none

参数介绍

-bench=. ：表示的是运行所有的基准测试，. 表示全部。
-benchtime=5s:表示的是运行时间为5s，默认的时间是1s。
-benchmem:表示显示memory的指标。
-run=none:表示过滤掉单元测试，不去跑UT的cases。
```

**输出结果** 

<img src="https://static.studygolang.com/191014/867f9df11e1a8319c6cad3ea2679b380.png" alt="result" style="zoom:50%;" /> 

- 下面是实际值

```bash
HandleType: 1BenchMark Test: Sub
HandleType: 1BenchMark Test: Sub
   52947             98138 ns/op              96 B/op          3 allocs/op
PASS
ok      src/com.wxw/04_project_actual/src/com.wxw/04_project_actual/w_test/benchmark_test        6.476s
```

**输出结果分析** 

```bash
goos: darwin：表示的是操作系统是darwin。
goarch: amd64：表示目标平台的体系架构是amd64。
BenchmarkHandleWithType-4：BenchmarkHandleWithType表示运行的函数名称； 4表示的是，运行时对应的GOMAXPROCS的值。
10000000000：表示的是b.N的在5s内的值。
0.28 ns/op：表示执行一次这个函数，消耗的时间是0.28ns。
0 B/op：表示每次执行操作，分配0B的内存。
0 allocs/op：表示执行一次这个函数，分配内存的次数为1次。
```

#### 3. 基准测试原理

基准测试框架对一个测试用例的默认测试时间是 1 秒。开始测试时，当以 Benchmark 开头的基准测试用例函数返回时还不到 1 秒，那么 testing.B 中的 N 值将按 1、2、5、10、20、50……递增，同时以递增后的值重新调用基准测试用例函数。

#### 4. 基准测试功能

- 自定义测试时间 ` benchtime`

  ```bash
  $ go test -v -bench=. -benchtime=5s benchmark_test.go
  goos: linux
  goarch: amd64
  Benchmark_Add-4           10000000000                 0.33 ns/op
  PASS
  ok          command-line-arguments        3.380s
  ```

- 测试内存 ` benchmem` 

  ```bash
  $ go test -v -bench=Alloc -benchmem benchmark_test.go
  goos: linux
  goarch: amd64
  Benchmark_Alloc-4 20000000 109 ns/op 16 B/op 2 allocs/op
  PASS
  ok          command-line-arguments        2.311s
  ```

- 控制计时器

  ```go
  func Benchmark_Add_TimerControl(b *testing.B) {
  
      // 重置计时器
      b.ResetTimer()
  
      // 停止计时器
      b.StopTimer()
  
      // 开始计时器
      b.StartTimer()
  
      var n int
      for i := 0; i < b.N; i++ {
          n++
      }
  }
  ```

### 使用帮助

#### 1. 常用命令

```bash
这里介绍几个常用的参数：

-bench regexp 执行相应的 benchmarks，例如 -bench=.；
-cover 开启测试覆盖率；
-run regexp 只运行 regexp 匹配的函数，例如 -run=Array 那么就执行包含有 Array 开头的函数；
-v 显示测试的详细命令。
```



#### 2. 使用文档

根据测试维度可以把包内以`_test.go结尾的测试`文件中的函数划分为以下3种：

| 类型     | 格式                  | 作用                                     |
| -------- | --------------------- | ---------------------------------------- |
| 测试函数 | 函数名前缀为Test      | 测试程序的一些逻辑行为是否正确           |
| 基准函数 | 函数名前缀为Benchmark | 测试函数的性能（执行时间、内存申请情况） |
| 示例函数 | 函数名前缀为Example   | 为文档提供示例文档                       |

- **单元测试函数（**unit testing）：是指对软件中的最小可测试单元进行检查和验证。对于单元测试中单元的含义，一般来说，要根据实际情况去判定其具体含义，如C语言中单元指一个函数，Java里单元指一个类，图形化的软件中可以指一个窗口或一个菜单等。总的来说，单元就是人为规定的最小的被测功能模块。

- **基准测试函数：**测试程序执行时间复杂度、空间复杂度 

- **示例函数：**为调用该功能代码的人提供演示  

**具体指令** 

- ` go help test` 

  ```bash
  go help test
  
  usage: go test [build/test flags] [packages] [build/test flags & test binary flags]
  
  'Go test' automates testing the packages named by the import paths.
  It prints a summary of the test results in the format:
  
          ok   archive/tar   0.011s
          FAIL archive/zip   0.022s
          ok   compress/gzip 0.033s
          ...
  
  followed by detailed output for each failed package.
  
  'Go test' recompiles each package along with any files with names matching
  the file pattern "*_test.go".
  These additional files can contain test functions, benchmark functions, and
  example functions. See 'go help testfunc' for more.
  Each listed package causes the execution of a separate test binary.
  Files whose names begin with "_" (including "_test.go") or "." are ignored.
  
  Test files that declare a package with the suffix "_test" will be compiled as a
  separate package, and then linked and run with the main test binary.
  
  The go tool will ignore a directory named "testdata", making it available
  to hold ancillary data needed by the tests.
  
  As part of building a test binary, go test runs go vet on the package
  and its test source files to identify significant problems. If go vet
  finds any problems, go test reports those and does not run the test
  binary. Only a high-confidence subset of the default go vet checks are
  used. That subset is: 'atomic', 'bool', 'buildtags', 'errorsas',
  'ifaceassert', 'nilfunc', 'printf', and 'stringintconv'. You can see
  the documentation for these and other vet tests via "go doc cmd/vet".
  To disable the running of go vet, use the -vet=off flag.
  
  All test output and summary lines are printed to the go command's
  standard output, even if the test printed them to its own standard
  error. (The go command's standard error is reserved for printing
  errors building the tests.)
  
  Go test runs in two different modes:
  
  The first, called local directory mode, occurs when go test is
  invoked with no package arguments (for example, 'go test' or 'go
  test -v'). In this mode, go test compiles the package sources and
  tests found in the current directory and then runs the resulting
  test binary. In this mode, caching (discussed below) is disabled.
  After the package test finishes, go test prints a summary line
  showing the test status ('ok' or 'FAIL'), package name, and elapsed
  time.
  
  The second, called package list mode, occurs when go test is invoked
  with explicit package arguments (for example 'go test math', 'go
  test ./...', and even 'go test .'). In this mode, go test compiles
  and tests each of the packages listed on the command line. If a
  package test passes, go test prints only the final 'ok' summary
  line. If a package test fails, go test prints the full test output.
  If invoked with the -bench or -v flag, go test prints the full
  output even for passing package tests, in order to display the
  requested benchmark results or verbose logging. After the package
  tests for all of the listed packages finish, and their output is
  printed, go test prints a final 'FAIL' status if any package test
  has failed.
  
  In package list mode only, go test caches successful package test
  results to avoid unnecessary repeated running of tests. When the
  result of a test can be recovered from the cache, go test will
  redisplay the previous output instead of running the test binary
  again. When this happens, go test prints '(cached)' in place of the
  elapsed time in the summary line.
  
  The rule for a match in the cache is that the run involves the same
  test binary and the flags on the command line come entirely from a
  restricted set of 'cacheable' test flags, defined as -benchtime, -cpu,
  -list, -parallel, -run, -short, and -v. If a run of go test has any test
  or non-test flags outside this set, the result is not cached. To
  disable test caching, use any test flag or argument other than the
  cacheable flags. The idiomatic way to disable test caching explicitly
  is to use -count=1. Tests that open files within the package's source
  root (usually $GOPATH) or that consult environment variables only
  match future runs in which the files and environment variables are unchanged.
  A cached test result is treated as executing in no time at all,
  so a successful package test result will be cached and reused
  regardless of -timeout setting.
  
  In addition to the build flags, the flags handled by 'go test' itself are:
  
          -args
              Pass the remainder of the command line (everything after -args)
              to the test binary, uninterpreted and unchanged.
              Because this flag consumes the remainder of the command line,
              the package list (if present) must appear before this flag.
  
          -c
              Compile the test binary to 05_pkg.test but do not run it
              (where 05_pkg is the last element of the package's import path).
              The file name can be changed with the -o flag.
  
          -exec xprog
              Run the test binary using xprog. The behavior is the same as
              in 'go run'. See 'go help run' for details.
  
          -i
              Install packages that are dependencies of the test.
              Do not run the test.
              The -i flag is deprecated. Compiled packages are cached automatically.
  
          -json
              Convert test output to JSON suitable for automated processing.
              See 'go doc test2json' for the w_encoding details.
  
          -o file
              Compile the test binary to the named file.
              The test still runs (unless -c or -i is specified).
  
  The test binary also accepts flags that control execution of the test; these
  flags are also accessible by 'go test'. See 'go help testflag' for details.
  
  For more about build flags, see 'go help build'.
  For more about specifying packages, see 'go help packages'.
  
  See also: go build, go vet.
  ```

- ` go help testfunc`  

  ```bash
  The 'go test' command expects to find test, benchmark, and example functions
  in the "*_test.go" files corresponding to the package under test.
  
  A test function is one named TestXxx (where Xxx does not start with a
  lower case letter) and should have the signature,
  
          func TestXxx(t *testing.T) { ... }
  
  A benchmark function is one named BenchmarkXxx and should have the signature,
  
          func BenchmarkXxx(b *testing.B) { ... }
  
  An example function is similar to a test function but, instead of using
  *testing.T to report success or failure, prints output to os.Stdout.
  If the last comment in the function starts with "Output:" then the output
  is compared exactly against the comment (see examples below). If the last
  comment begins with "Unordered output:" then the output is compared to the
  comment, however the order of the lines is ignored. An example with no such
  comment is compiled but not executed. An example with no text after
  "Output:" is compiled, executed, and expected to produce no output.
  
  Godoc displays the body of ExampleXxx to demonstrate the use
  of the function, constant, or variable Xxx. An example of a method M with
  receiver type T or *T is named ExampleT_M. There may be multiple examples
  for a given function, constant, or variable, distinguished by a trailing _xxx,
  where xxx is a suffix not beginning with an upper case letter.
  
  Here is an example of an example:
  
          func ExamplePrintln() {
                  Println("The output of\nthis example.")
                  // Output: The output of
                  // this example.
          }
  
  Here is another example where the ordering of the output is ignored:
  
          func ExamplePerm() {
                  for _, value := range Perm(4) {
                          fmt.Println(value)
                  }
  
                  // Unordered output: 4
                  // 2
                  // 1
                  // 3
                  // 0
          }
  
  The entire test file is presented as the example when it contains a single
  example function, at least one other function, type, variable, or constant
  declaration, and no test or benchmark functions.
  
  See the documentation of the testing package for more information.
  
  ```

相关文章

1. https://studygolang.com/articles/23960
2. http://c.biancheng.net/view/124.html
3. https://www.cnblogs.com/sss4/p/12859027.html
4. https://geektutu.com/post/hpg-benchmark.html

