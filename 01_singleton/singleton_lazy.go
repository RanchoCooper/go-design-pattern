package singleton

import "sync"

var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

// GetLazyInstance 懒汉式
func GetLazyInstance() *Singleton {
	once.Do(func() {
		lazySingleton = &Singleton{a: "test"}
	})
	return lazySingleton
}
