// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

// KeyValueResponse
type KeyValueResponse struct {
	Key_values []*KeyValue `json:"key_values"`
}

func (r KeyValueResponse) Validate() *ValidateError {
	errs := []*FieldError{}
	if len(errs) > 0 {
		return &ValidateError{Errors: errs}
	}
	return nil
}
