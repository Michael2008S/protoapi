// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

// KeyValueListRequest
type KeyValueListRequest struct {
	Service_id int    `json:"service_id"`
	Keys       []*Key `json:"keys"`
}

func (r KeyValueListRequest) Validate() *ValidateError {
	errs := []*FieldError{}
	if len(errs) > 0 {
		return &ValidateError{Errors: errs}
	}
	return nil
}
