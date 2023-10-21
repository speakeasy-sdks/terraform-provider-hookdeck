// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"hashicups/internal/sdk/pkg/models/shared"
	"net/http"
)

// TriggerBookmarkRequestBodyTarget - Bookmark target
type TriggerBookmarkRequestBodyTarget string

const (
	TriggerBookmarkRequestBodyTargetHTTP TriggerBookmarkRequestBodyTarget = "http"
	TriggerBookmarkRequestBodyTargetCli  TriggerBookmarkRequestBodyTarget = "cli"
)

func (e TriggerBookmarkRequestBodyTarget) ToPointer() *TriggerBookmarkRequestBodyTarget {
	return &e
}

func (e *TriggerBookmarkRequestBodyTarget) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "http":
		fallthrough
	case "cli":
		*e = TriggerBookmarkRequestBodyTarget(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TriggerBookmarkRequestBodyTarget: %v", v)
	}
}

type TriggerBookmarkRequestBody struct {
	// Bookmark target
	Target *TriggerBookmarkRequestBodyTarget `json:"target,omitempty"`
}

func (o *TriggerBookmarkRequestBody) GetTarget() *TriggerBookmarkRequestBodyTarget {
	if o == nil {
		return nil
	}
	return o.Target
}

type TriggerBookmarkRequest struct {
	RequestBody TriggerBookmarkRequestBody `request:"mediaType=application/json"`
	ID          string                     `pathParam:"style=simple,explode=false,name=id"`
}

func (o *TriggerBookmarkRequest) GetRequestBody() TriggerBookmarkRequestBody {
	if o == nil {
		return TriggerBookmarkRequestBody{}
	}
	return o.RequestBody
}

func (o *TriggerBookmarkRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type TriggerBookmarkResponse struct {
	// Bad Request
	APIErrorResponse *shared.APIErrorResponse
	// HTTP response content type for this operation
	ContentType string
	// Array of created events
	EventArray []shared.Event
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *TriggerBookmarkResponse) GetAPIErrorResponse() *shared.APIErrorResponse {
	if o == nil {
		return nil
	}
	return o.APIErrorResponse
}

func (o *TriggerBookmarkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TriggerBookmarkResponse) GetEventArray() []shared.Event {
	if o == nil {
		return nil
	}
	return o.EventArray
}

func (o *TriggerBookmarkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TriggerBookmarkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
