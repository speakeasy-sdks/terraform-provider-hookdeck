// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hashicups/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeleteCustomDomainRequest struct {
	DomainID string `pathParam:"style=simple,explode=false,name=domain_id"`
	TeamID   string `pathParam:"style=simple,explode=false,name=team_id"`
}

func (o *DeleteCustomDomainRequest) GetDomainID() string {
	if o == nil {
		return ""
	}
	return o.DomainID
}

func (o *DeleteCustomDomainRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

type DeleteCustomDomainResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Custom domain successfuly removed
	DeleteCustomDomainSchema *shared.DeleteCustomDomainSchema
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteCustomDomainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteCustomDomainResponse) GetDeleteCustomDomainSchema() *shared.DeleteCustomDomainSchema {
	if o == nil {
		return nil
	}
	return o.DeleteCustomDomainSchema
}

func (o *DeleteCustomDomainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteCustomDomainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
