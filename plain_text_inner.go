package ace

import "io"

// plainTextInner represents a plain text inner.
type plainTextInner struct {
	elementBase
}

// WriteTo writes data to w.
func (e *plainTextInner) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}

// CanHaveChildren returns false.
func (e *plainTextInner) CanHaveChildren() bool {
	return false
}

// newPlainTextInner creates and returns a plain text.
func newPlainTextInner(ln *line, rslt *result, parent element, opts *Options) *plainTextInner {
	return &plainTextInner{
		elementBase: newElementBase(ln, rslt, parent, opts),
	}
}