// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"hashicups/internal/sdk/pkg/models/shared"
	"net/http"
)

// UpdateDestinationRequestBodyRateLimitPeriod - Period to rate limit attempts
type UpdateDestinationRequestBodyRateLimitPeriod string

const (
	UpdateDestinationRequestBodyRateLimitPeriodSecond UpdateDestinationRequestBodyRateLimitPeriod = "second"
	UpdateDestinationRequestBodyRateLimitPeriodMinute UpdateDestinationRequestBodyRateLimitPeriod = "minute"
	UpdateDestinationRequestBodyRateLimitPeriodHour   UpdateDestinationRequestBodyRateLimitPeriod = "hour"
)

func (e UpdateDestinationRequestBodyRateLimitPeriod) ToPointer() *UpdateDestinationRequestBodyRateLimitPeriod {
	return &e
}

func (e *UpdateDestinationRequestBodyRateLimitPeriod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "second":
		fallthrough
	case "minute":
		fallthrough
	case "hour":
		*e = UpdateDestinationRequestBodyRateLimitPeriod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateDestinationRequestBodyRateLimitPeriod: %v", v)
	}
}

type UpdateDestinationRequestBody struct {
	// Config for the destination's auth method
	AuthMethod *shared.DestinationAuthMethodConfig `json:"auth_method,omitempty"`
	// Path for the CLI destination
	CliPath *string `json:"cli_path,omitempty"`
	// HTTP method used on requests sent to the destination, overrides the method used on requests sent to the source.
	HTTPMethod *shared.DestinationHTTPMethod `json:"http_method,omitempty"`
	// Name for the destination
	Name                   *string `json:"name,omitempty"`
	PathForwardingDisabled *bool   `json:"path_forwarding_disabled,omitempty"`
	// Limit event attempts to receive per period
	RateLimit interface{} `json:"rate_limit,omitempty"`
	// Period to rate limit attempts
	RateLimitPeriod *UpdateDestinationRequestBodyRateLimitPeriod `json:"rate_limit_period,omitempty"`
	// Endpoint of the destination
	URL *string `json:"url,omitempty"`
}

type UpdateDestinationRequest struct {
	RequestBody UpdateDestinationRequestBody `request:"mediaType=application/json"`
	ID          string                       `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateDestinationResponse struct {
	// Bad Request
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	// A single destination
	Destination *shared.Destination
	StatusCode  int
	RawResponse *http.Response
}
