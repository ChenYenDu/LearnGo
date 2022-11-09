# Function

## what is a function?

* 可重複執行的一個程式碼區
* 減少 main function 內的複雜度


## How to declare a function

* Function 基本定義方式
    ```go
    func name(parameter-list) (result-list){
        body
    }

    e.g.
    func hypot(x, y float64) float64 {
            return math.Sqrt(x*x + y*y)
    }

    ```
* parameter-list: 帶入function的參數, 必須指定型態
* result-list: function 回傳的結果，同樣必須指定型態


* Function 的參數結果都可以複數定義
    ```go
    /* 這兩個定義是一樣的 */
    func f(i, j, k int, s, t string){/*...*/}
    func f(i int, j int, k int, s string, t string){/*...*/}
    ```

* Function 必須以return結尾，除非function本身不回傳任何結果或是只想要操作其他變數


* 一些我覺得golang很特別的部分：
  ```go
  func f(p1 string, p2 int)(r1 int, r2 string) {
    r2 := p1
    r1 := p2
    return
  }

  f("1", 2)
  /* 跟python不一樣，這裡會回傳 2, "1" */
  ```

* Go 是一個 100% copy value language，這表示所有的參數傳入function後都會有一份copy, main function caller 裡的variable不會被function內影響，需要用return 結果覆蓋
除非，你傳入function的參數是 pointer, slice, map, function 或 channel

* Go 是辦法用參數名傳入的，一定要按定義的順序給參數

### Error
Error 跟 exception 不一樣， error 偏向已知的預期的錯誤，golang也有exception

Go 的 error handling strategy:
* _propagate_ error: 把 error 作爲return的一部份，並在caller裡面進形診斷
    ```go
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }

    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
    }

    /* fmt.Errorf 代表用 fmt.Sprintf 的格式回傳一個error message */
    ```

* _retry_
  ```go
    func WaitForServer(url string) error {
 
        const timeout = 1 * time.Minute
        deadline := time.Now().Add(timeout)
        for tries := 0; time.Now().Before(deadline); tries++ {
            _, err := http.Head(url)
            if err == nil {
                return nil // success
            }
            log.Printf("server not responding (%s); retrying...", err)
            time.Sleep(time.Second << uint(tries)) // exponential back-off
        }
        return fmt.Errorf("server %s failed to respond after %s", url, timeout)
    }
  ```

* ＿print error and stop the program_
  ```go
    // (In function main.)
    if err := WaitForServer(url); err != nil {
        fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
        os.Exit(1)
    }
    
    // or more convenience
    if err := WaitForServer(url); err != nil {
        log.Fatalf("Site is down: %v\n", err)
    }
  ```

* _log error and continue_
  ```go
    if err := Ping(); err != nil {
        log.Printf("ping failed: %v; network is disabled", err)
    }
  ```

* _ignore error_
  ```go
    dir, err := ioutil.TempDir("", "scratch")
    if err != nil {
        return fmt.Errorf("failed to create temp dir: %v", err)
    }
    // ...use temp dir...
    os.RemoveAll(dir) //ignore errors; $TEMPDIR is cleaned periodically
    // 這裡表示我們故意不去處理 os.RemoveAll()也沒關係，因為我們本來就要移除所以的temp dir, 他沒有建成也沒關係
  ```

* End of File (EOF)
io package 針對讀取到file end的時候有特定的error
```go
in := bufio.NewReader(os.Stdin)
for {
    r, _, err := in.ReadRune()
    if err == io.EOF {
        break // finished reading
    }
    if err != nil {
        return fmt.Errorf("read failed: %v", err)
    }
    // ...use r...
}
```

## Function Values
function 是 Go 的 first-class 之一，這表示function value 是有型別的，這讓可以被指地成變數或者是其他function回傳的結果
```go
func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }
f := square
fmt.Println(f(3)) // "9"
f = negative
fmt.Println(f(3))     // "-3"
fmt.Printf("%T\n", f) // "func(int) int"
f = product // compile error: can't assign f(int, int) int to f(int) int
```

* zero value for function type is: __nil__
* Calling a nil function value can caused a panic
```go
var f func(int) int
f(3) // panic: call of nil funciton
```
* function value 可以跟nil進行確認
```go
var f func(int) int
if f != nil {
    f(3)
}
```

## Anonymous Functions
```go
// this method is also called closures
func squares() func() int {
    var x int
    return func() int {
        x++
        return x * x
    }
}

func main() {
    f := squares()
    fmt.Println(f()) // "1"
    fmt.Println(f()) // "4"
    fmt.Println(f()) // "9"
    fmt.Println(f()) // "16"
}
```


## variadic functions
```go
func sum(vals ...int) int {
    total := 0
    for _, val := range vals {
        total += val
    }
    return total
}
```
實際用法：
```go
fmt.Println(sum(1,2,3,4)) // 10

vales := []int{1, 2, 3, 4}
fmt.Println(sum(values...)) // 10
```