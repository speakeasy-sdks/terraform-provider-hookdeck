// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hashicups/internal/sdk/pkg/models/shared"
	"net/http"
)

type UnpauseConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UnpauseConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UnpauseConnectionResponse struct {
	// Not Found
	APIErrorResponse *shared.APIErrorResponse
	// A single connection
	Connection *shared.Connection
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UnpauseConnectionResponse) GetAPIErrorResponse() *shared.APIErrorResponse {
	if o == nil {
		return nil
	}
	return o.APIErrorResponse
}

func (o *UnpauseConnectionResponse) GetConnection() *shared.Connection {
	if o == nil {
		return nil
	}
	return o.Connection
}

func (o *UnpauseConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UnpauseConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UnpauseConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
