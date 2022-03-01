package exception

type ExceptionError struct {
	Code    int
	Message string
	Data    interface{}
}

func (exceptionError ExceptionError) Error() string {
	return exceptionError.Message
}
