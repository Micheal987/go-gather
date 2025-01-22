package error

type WrapError struct {
	msg string
	err error
}

func (e *WrapError) Error() string {
	return e.msg
}

func (e *WrapError) Unwrap() error {
	return e.err
}

func NewWrapError(msg string, errs error) *WrapError {
	return &WrapError{msg: msg, err: errs}
}
