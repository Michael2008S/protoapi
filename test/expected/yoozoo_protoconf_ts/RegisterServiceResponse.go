// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

// RegisterServiceResponse
type RegisterServiceResponse struct {
	Env_id       int    `json:"env_id"`
	Product_id   string `json:"product_id"`
	Service_id   int    `json:"service_id"`
	Service_name string `json:"service_name"`
}

func (r RegisterServiceResponse) Validate() *ValidateError {
	errs := []*FieldError{}
	if len(errs) > 0 {
		return &ValidateError{Errors: errs}
	}
	return nil
}
