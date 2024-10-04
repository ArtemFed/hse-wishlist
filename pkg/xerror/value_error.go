package xerror

type ValueError struct {
	message string
}

func NewValueError(message string) error {
	return ValueError{message: message}
}

func (r ValueError) Error() string {
	return r.message
}
