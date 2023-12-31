// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hashicups/internal/sdk/pkg/models/shared"
	"net/http"
)

type AddCustomDomainRequest struct {
	AddCustomHostname shared.AddCustomHostname `request:"mediaType=application/json"`
	TeamID            string                   `pathParam:"style=simple,explode=false,name=team_id"`
}

type AddCustomDomainResponse struct {
	// Custom domain successfuly added
	AddCustomHostname *shared.AddCustomHostname
	ContentType       string
	StatusCode        int
	RawResponse       *http.Response
}
