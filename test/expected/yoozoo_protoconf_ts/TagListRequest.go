// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

// TagListRequest
type TagListRequest struct {
}

func (r TagListRequest) Validate() *ValidateError {
	errs := []*FieldError{}
	if len(errs) > 0 {
		return &ValidateError{Errors: errs}
	}
	return nil
}
