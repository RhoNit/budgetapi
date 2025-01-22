package custom_errors

type CustomError struct {
	Message string
}

func NewCustomError(msg string) *CustomError {
	return &CustomError{
		Message: msg,
	}
}

func (cusErr *CustomError) Error() string {
	return cusErr.Message
}
