package writer

import "testing"
import "bytes"
import "io"

var _ io.Writer = new(SyncWriter)

func Test(t *testing.T) {
	b := bytes.NewBuffer(nil)
	w := New(b)
	w.Write([]byte("test"))
}
