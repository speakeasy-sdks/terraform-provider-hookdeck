// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hashicups/internal/sdk/pkg/models/shared"
	"net/http"
)

type ArchiveRulesetRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type ArchiveRulesetResponse struct {
	// Not Found
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	// A single ruleset
	Ruleset     *shared.Ruleset
	StatusCode  int
	RawResponse *http.Response
}
