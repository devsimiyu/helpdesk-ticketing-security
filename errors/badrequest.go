package errors

type BadRequest string

func (b BadRequest) Error() string {
	return string(b)
}
