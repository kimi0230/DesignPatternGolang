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
