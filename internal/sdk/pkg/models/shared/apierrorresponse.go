// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type APIErrorResponseData struct {
}

// APIErrorResponse - Error response model
type APIErrorResponse struct {
	// Error code
	Code string                `json:"code"`
	Data *APIErrorResponseData `json:"data,omitempty"`
	// Error description
	Message string `json:"message"`
	// Status code
	Status float32 `json:"status"`
}
