// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

// RegisterServiceRequest
type RegisterServiceRequest struct {
    Env_id int `json:"env_id"`
    Product_id string `json:"product_id"`
    Service_name string `json:"service_name"`
    Tags []*Tag `json:"tags"`
    Desc string `json:"desc"`
}

func (r RegisterServiceRequest) Validate() *ValidateError {
    errs := []*FieldError{}
    if len(errs) > 0 {
		return &ValidateError{Errors: errs}
	}
	return nil
}
