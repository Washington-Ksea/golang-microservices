# ベンチーマーク

go test -bench=.
参考[go標準のbenchmark機能の使い方](https://qiita.com/marnie_ms4/items/7014563083ca1d824905)
```go
func BenchmarkFunctionName(b *testing.B) {
	els := FunctionName(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}
```

# 各パッケージ内でファイルごとにメソッドを分ける

パッケージ内でpackageName.A.method()という形でアクセスできるようにする。

```go
type A struct {}
var (
	a A
)

func (f *A) method() {}
```

# テストのisolationの確保に関して

Aのパッケージを読み込むBを隔離してテストする(Unitテスト)

Aパッケージ

```go
//A
package Apackage

var (
	Aglob Ainterface
)

type Ainterface interface{
	Amethod() int64
}

type A struct{}

func (m *A) Amethod() int64 {
	//something
	return 1
}

func init() {
	Aglob = &A
}
```

Bパッケージ

```go
//B
package Bpackage

var Bglob B

type B struct{}

func (m *B) Bmethod() int64 {
	return Apackage.Aglob.Amethod() //ApackageのAmethodを実行s
	
}

```

もし、BパッケージをUnitテストする場合、Aパッケージから隔離する。


```go
//B_test.go
package Bpackage

//Mock
var(
	mock AMock
	getAmethod func() int64
) 

type AMock struct{}

func (m *AMock) Amethod() int64{
	return getAmethod()
}

func init(){
	Apackage.Agolb = &AMock
}

//Test
func TestBpackage(){
	getAmethod = func(){
		return 2
	}
	//test... Bglob.Bmethod内で呼び出されるAmethodの返り値が2になる。

}
```

# パッケージを調べる

例えば、ginのNewを調べる場合、

```
go doc github.com/gin-gonic/gin New
```

# Testの初期化

下記の関するを各テストファイルに記載することで、各テスト関数に共通して実行されるメソッドを記載できる。

```go
func TestMain(m *testing.M) {
	//Before method
	os.Exit(m.Run())
	//After method
}
```

# limit concurrency

go runtimeの数を制限する。また、すべての処理完了までプログラム終了を待つ。
競合状態の処理(Mutex)

```go

limitN = make(chan int, 5)
var wg sync.WaitGroup
var lock sync.Mutex

for i:=0; i< 100; i++ {
	c <- true
	wg.Add(1)
	go func (c chan int, wg *sync.WaitGroup){
		lock.Lock()
		defer lock.Unlock()
		//something doing
		defer wg.Done()
		<-c
	}(limitN, wg)
}

wg.Wait()


```