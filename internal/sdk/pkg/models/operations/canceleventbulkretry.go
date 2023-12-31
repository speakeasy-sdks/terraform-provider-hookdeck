// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hashicups/internal/sdk/pkg/models/shared"
	"net/http"
)

type CancelEventBulkRetryRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type CancelEventBulkRetryResponse struct {
	// Not Found
	APIErrorResponse *shared.APIErrorResponse
	// A single events bulk retry
	BatchOperation *shared.BatchOperation
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
}
