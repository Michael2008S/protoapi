// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

// Error
type Error struct {
	Code    *ErrorCode `json:"code"`
	Message string     `json:"message"`
}

func (r Error) Validate() *ValidateError {
	errs := []*FieldError{}
	if len(errs) > 0 {
		return &ValidateError{Errors: errs}
	}
	return nil
}
