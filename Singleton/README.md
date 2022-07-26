### Singleton 獨體模式：
> 確保類別只會有一個物件實體存在, 並提供單一存取窗口

保證一個類別僅有一個實體, 並提供一個存取它的全域訪問點

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/Singleton.png?raw=true)

### Example
https://github.com/kimi0230/DesignPatternGolang/tree/master/Singleton 

```go
package singleton

import "sync"

type Singleton struct {
}

// 方法一: 用 sync.Mutex
var instance *Singleton
var mu sync.Mutex

func GetInstance() *Singleton {
	if instance != nil {
		mu.Lock()
		defer mu.Unlock()
		if instance != nil {
			instance = &Singleton{}
		}
	}
	return instance
}

// 方法二: 用 sync.Once
var once sync.Once

func GetInstanceOnce() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

```

### Benchmark
```sh
goos: darwin
goarch: amd64
pkg: DesignPatternGolang/Singleton
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkSingleton-4            401843040                2.971 ns/op           0 B/op          0 allocs/op
BenchmarkSingletonOnce-4        581861808                2.129 ns/op           0 B/op          0 allocs/op
PASS
ok      DesignPatternGolang/Singleton   2.960s
```