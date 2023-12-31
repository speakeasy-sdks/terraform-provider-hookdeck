// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hashicups/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeleteSourceRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

// DeleteSource200ApplicationJSON - A single source
type DeleteSource200ApplicationJSON struct {
	// ID of the source
	ID string `json:"id"`
}

type DeleteSourceResponse struct {
	// Not Found
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	StatusCode       int
	RawResponse      *http.Response
	// A single source
	DeleteSource200ApplicationJSONObject *DeleteSource200ApplicationJSON
}
