// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// SourceIntegration - Integration object
type SourceIntegration struct {
	// List of enabled features
	Features []IntegrationFeature `json:"features"`
	// ID of the integration
	ID string `json:"id"`
	// Label of the integration
	Label string `json:"label"`
	// The provider name
	Provider IntegrationProvider `json:"provider"`
}

// Source - Associated [Source](#source-object) object
type Source struct {
	// List of allowed HTTP methods. Defaults to PUT, POST, PATCH, DELETE.
	AllowedHTTPMethods []SourceAllowedHTTPMethod `json:"allowed_http_methods,omitempty"`
	// Date the source was archived
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
	// Date the source was created
	CreatedAt time.Time `json:"created_at"`
	// Custom response object
	CustomResponse *SourceCustomResponse `json:"custom_response,omitempty"`
	// ID of the source
	ID string `json:"id"`
	// Integration object
	Integration *SourceIntegration `json:"integration,omitempty"`
	// ID of the integration
	IntegrationID *string `json:"integration_id,omitempty"`
	// Name for the source
	Name string `json:"name"`
	// ID of the workspace
	TeamID string `json:"team_id"`
	// Date the source was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// A unique URL that must be supplied to your webhook's provider
	URL string `json:"url"`
}
