package _interface

import . "bytes"

//接口嵌套
//接口file包含了readwrite,lock,和close
type ReadWrite interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
}

type Lock interface {
	Lock()
	Unlock()
}

type File interface {
	ReadWrite
	Lock
	Close()
}