// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// TransformationExecution - A single transformation execution
type TransformationExecution struct {
	CreatedAt time.Time `json:"created_at"`
	ID        string    `json:"id"`
	IssueID   *string   `json:"issue_id,omitempty"`
	// The minimum log level to open the issue on
	LogLevel               TransformationExecutionLogLevel `json:"log_level"`
	Logs                   []ConsoleLine                   `json:"logs"`
	OriginalEventData      *ShortEventData                 `json:"original_event_data,omitempty"`
	OriginalEventDataID    string                          `json:"original_event_data_id"`
	TeamID                 string                          `json:"team_id"`
	TransformationID       string                          `json:"transformation_id"`
	TransformedEventData   *ShortEventData                 `json:"transformed_event_data,omitempty"`
	TransformedEventDataID string                          `json:"transformed_event_data_id"`
	UpdatedAt              time.Time                       `json:"updated_at"`
	WebhookID              string                          `json:"webhook_id"`
}
