package writer

import (
	"io"
	"sync"
)

// SyncWriter wraps the underlying writer in a mutex.
type SyncWriter struct {
	io.Writer
	sync.Mutex
}

// New sync writer.
func New(w io.Writer) *SyncWriter {
	return &SyncWriter{
		Writer: w,
	}
}

// Write implements the io.Writer interface.
func (w *SyncWriter) Write(b []byte) (int, error) {
	w.Lock()
	defer w.Unlock()
	return w.Writer.Write(b)
}
