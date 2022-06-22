package util

// RequestError struct stores the error that ocurrs and the httpcode for the response
type RequestError struct {
	StatusCode int
	Err        *WrappedError
}

// Error function overwrites the default output of error message so that we can see the context as well as the error
func (r *RequestError) Error() string {
	return r.Err.Error()
}

// WithoutContext function returns info about the error but without the context information on where did it happen
func (r *RequestError) ErrorWithoutContext() string {
	return r.Err.WithoutContext().Error()
}

// NewRequestError function receives the error and the status code, and return a pointer to a RequesteError with the received info
func NewRequestError(statusCode int, err *WrappedError) *RequestError {
	return &RequestError{
		StatusCode: statusCode,
		Err:        err,
	}
}
