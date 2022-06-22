package util

import "fmt"

// WrappedError struct stores the error that ocurrs and a context about where did it happen
type WrappedError struct {
	Context string
	Err     error
}

// Wrap function receives the error and the context info and return a pointer to a WrappedError with the received info
func Wrap(err error, info string) *WrappedError {
	return &WrappedError{
		Context: info,
		Err:     err,
	}
}

// Error function overwrites the default output of error message so that we can see the context as well as the error
func (w *WrappedError) Error() string {
	return fmt.Sprintf("%s: %v", w.Context, w.Err)
}

// WithoutContext function returns info about the error but without the context information on where did it happen
func (w *WrappedError) WithoutContext() error {
	return w.Err
}
