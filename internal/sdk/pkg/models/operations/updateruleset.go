// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hashicups/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateRulesetRequestBody struct {
	IsTeamDefault *bool `json:"is_team_default,omitempty"`
	// Name for the ruleset
	Name *string `json:"name,omitempty"`
	// Array of rules to apply
	Rules []shared.Rule `json:"rules,omitempty"`
}

type UpdateRulesetRequest struct {
	RequestBody UpdateRulesetRequestBody `request:"mediaType=application/json"`
	ID          string                   `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateRulesetResponse struct {
	// Bad Request
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	// A single ruleset
	Ruleset     *shared.Ruleset
	StatusCode  int
	RawResponse *http.Response
}
