// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// Ruleset - Associated [Ruleset](#ruleset-object) object
type Ruleset struct {
	// Date the ruleset was archived
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
	// Date the ruleset was created
	CreatedAt time.Time `json:"created_at"`
	// ID of the ruleset
	ID string `json:"id"`
	// Default ruleset of Workspace
	IsTeamDefault bool `json:"is_team_default"`
	// A unique name for the ruleset
	Name string `json:"name"`
	// Array of rules to apply
	Rules []Rule `json:"rules"`
	// ID of the workspace
	TeamID string `json:"team_id"`
	// Date the ruleset was last updated
	UpdatedAt time.Time `json:"updated_at"`
}
