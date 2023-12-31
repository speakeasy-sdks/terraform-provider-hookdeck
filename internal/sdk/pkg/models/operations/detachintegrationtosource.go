// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hashicups/internal/sdk/pkg/models/shared"
	"net/http"
)

type DetachIntegrationToSourceRequest struct {
	ID       string `pathParam:"style=simple,explode=false,name=id"`
	SourceID string `pathParam:"style=simple,explode=false,name=source_id"`
}

type DetachIntegrationToSourceResponse struct {
	// Bad Request
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	// Detach operation success status
	DetachedIntegrationFromSource *shared.DetachedIntegrationFromSource
	StatusCode                    int
	RawResponse                   *http.Response
}
