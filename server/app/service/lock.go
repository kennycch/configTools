package service

import "sync"

// LockFunc 方法锁
func LockFunc(lockName string, f func()) {
	funcLock := getLock(lockName)
	funcLock.Lock()
	defer funcLock.Unlock()
	f()
}

// getLock 获取锁
func getLock(lockName string) *sync.Mutex {
	lock.Lock()
	defer lock.Unlock()
	funcLock, ok := locks[lockName]
	if !ok {
		funcLock = &sync.Mutex{}
		locks[lockName] = funcLock
	}
	return funcLock
}
