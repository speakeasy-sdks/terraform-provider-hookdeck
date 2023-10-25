// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hashicups/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateBookmarkRequestBody struct {
	// ID of the event data to bookmark
	EventDataID string `json:"event_data_id"`
	// Descriptive name of the bookmark
	Label string `json:"label"`
	// A unique, human-friendly name for the bookmark
	Name *string `json:"name,omitempty"`
	// ID of the associated connection
	WebhookID string `json:"webhook_id"`
}

func (o *CreateBookmarkRequestBody) GetEventDataID() string {
	if o == nil {
		return ""
	}
	return o.EventDataID
}

func (o *CreateBookmarkRequestBody) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *CreateBookmarkRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CreateBookmarkRequestBody) GetWebhookID() string {
	if o == nil {
		return ""
	}
	return o.WebhookID
}

type CreateBookmarkResponse struct {
	// Bad Request
	APIErrorResponse *shared.APIErrorResponse
	// A single bookmark
	Bookmark *shared.Bookmark
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateBookmarkResponse) GetAPIErrorResponse() *shared.APIErrorResponse {
	if o == nil {
		return nil
	}
	return o.APIErrorResponse
}

func (o *CreateBookmarkResponse) GetBookmark() *shared.Bookmark {
	if o == nil {
		return nil
	}
	return o.Bookmark
}

func (o *CreateBookmarkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateBookmarkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateBookmarkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
