package msg

type Reader interface {
	ReadMessage() ([]byte, error)
}

type Writer interface {
	WriteMessage(p []byte) error
}

type Closer interface {
	Close() error
}

type ReadCloser interface {
	Reader
	Closer
}

type WriteCloser interface {
	Writer
	Closer
}

type ReadWriter interface {
	Reader
	Writer
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}
