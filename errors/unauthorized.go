package errors

type Unuathorized string

func (u Unuathorized) Error() string {
	return string(u)
}
