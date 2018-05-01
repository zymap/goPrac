package _package

import (
	"sync"
	"fmt"
)

func LockUse()  {
	//访问一个互斥资源
	//mutex是一个互斥锁
	lock :=sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()
	fmt.Println("into lock...")
}

func ReadWriteLock()  {
	rwlock := sync.RWMutex{}
	rwlock.RLock()
	defer rwlock.RUnlock()
	fmt.Println("locker : ", rwlock.RLocker())

	rwlock.Lock()
	defer rwlock.Unlock()
	fmt.Println("write locker : ", rwlock.Unlock)
}