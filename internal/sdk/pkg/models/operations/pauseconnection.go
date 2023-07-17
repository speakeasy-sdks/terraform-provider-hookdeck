// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hashicups/internal/sdk/pkg/models/shared"
	"net/http"
)

type PauseConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type PauseConnectionResponse struct {
	// Not Found
	APIErrorResponse *shared.APIErrorResponse
	// A single connection
	Connection  *shared.Connection
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
