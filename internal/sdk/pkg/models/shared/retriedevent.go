// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RetriedEvent - Retried event with event attempt
type RetriedEvent struct {
	Attempt EventAttempt `json:"attempt"`
	Event   Event        `json:"event"`
}
