// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// IssueTrigger - A single issue trigger
type IssueTrigger struct {
	// Notification channels object for the specific channel type
	Channels *IssueTriggerChannels `json:"channels,omitempty"`
	// Configuration object for the specific issue type selected
	Configs interface{} `json:"configs"`
	// ISO timestamp for when the issue trigger was created
	CreatedAt time.Time `json:"created_at"`
	// ISO timestamp for when the issue trigger was deleted
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// ISO timestamp for when the issue trigger was disabled
	DisabledAt *time.Time `json:"disabled_at,omitempty"`
	// ID of the issue trigger
	ID string `json:"id"`
	// Optional unique name to use as reference when using the API
	Name *string `json:"name,omitempty"`
	// ID of the workspace
	TeamID *string `json:"team_id,omitempty"`
	// Issue type
	Type IssueType `json:"type"`
	// ISO timestamp for when the issue trigger was last updated
	UpdatedAt time.Time `json:"updated_at"`
}
