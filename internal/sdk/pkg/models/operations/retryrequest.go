// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hashicups/internal/sdk/pkg/models/shared"
	"net/http"
)

type RetryRequestRequestBody struct {
	// Subset of webhook_ids to re-run the event logic on. Useful to retry only specific ignored_events
	WebhookIds []string `json:"webhook_ids"`
}

type RetryRequestRequest struct {
	RequestBody RetryRequestRequestBody `request:"mediaType=application/json"`
	ID          string                  `pathParam:"style=simple,explode=false,name=id"`
}

type RetryRequestResponse struct {
	// Bad Request
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	// Retry request operation result
	RetryRequest *shared.RetryRequest
	StatusCode   int
	RawResponse  *http.Response
}
