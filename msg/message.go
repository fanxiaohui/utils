package msg

import (
	"io"
)

// Reader 读消息方法
type Reader interface {
	ReadMessage() ([]byte, error)
}

// Writer 写消息方法
type Writer interface {
	WriteMessage(p []byte) error
}

// ReadCloser 读消息和关闭方法
type ReadCloser interface {
	Reader
	io.Closer
}

// WriteCloser 写消息和关闭方法
type WriteCloser interface {
	Writer
	io.Closer
}

// ReadWriter 读写方法
type ReadWriter interface {
	Reader
	Writer
}

// ReadWriteCloser 读写和关闭方法
type ReadWriteCloser interface {
	Reader
	Writer
	io.Closer
}
