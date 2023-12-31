// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hashicups/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetEventRawBodyRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetEventRawBodyResponse struct {
	// Not Found
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	// A request raw body data
	RawBody     *shared.RawBody
	StatusCode  int
	RawResponse *http.Response
}
