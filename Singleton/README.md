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