// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

// ProductListRequest
type ProductListRequest struct {
	Env_id int `json:"env_id"`
}

func (r ProductListRequest) Validate() *ValidateError {
	errs := []*FieldError{}
	if len(errs) > 0 {
		return &ValidateError{Errors: errs}
	}
	return nil
}
