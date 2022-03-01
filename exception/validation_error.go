package exception

type ValidationError struct {
	Message string
	Errors  interface{}
}

func (validationError ValidationError) Error() string {
	return validationError.Message
}
