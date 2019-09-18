# ベンチーマーク

go test -bench .
参考[go標準のbenchmark機能の使い方](https://qiita.com/marnie_ms4/items/7014563083ca1d824905)
```go
func BenchmarkFunctionName(b *testing.B) {
	els := FunctionName(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}
```

