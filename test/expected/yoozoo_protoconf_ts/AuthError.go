// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

// AuthError
type AuthError struct {
    Message string `json:"message"`
}

func (r AuthError) Validate() *ValidateError {
    errs := []*FieldError{}
    if len(errs) > 0 {
		return &ValidateError{Errors: errs}
	}
	return nil
}
