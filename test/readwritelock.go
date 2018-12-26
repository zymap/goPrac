package main

import "sync"

func main() {
	readWriteLock()
}

func readWriteLock() {
	var rwlock sync.RWMutex
	go func() {
		rwlock.RLock()
		defer rwlock.RUnlock()

		rwlock.RLock()
		defer rwlock.RUnlock()
	}()

	rwlock.Lock()
	defer rwlock.Unlock()
}
