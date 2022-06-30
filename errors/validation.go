package errors

import "encoding/json"

type ValidationError map[string]interface{}

func (v *ValidationError) Error() string {
	message, _ := json.Marshal(v)
	return string(message)
}
