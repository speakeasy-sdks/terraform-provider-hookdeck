// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DestinationRateLimitPeriod - Period to rate limit attempts
type DestinationRateLimitPeriod string

const (
	DestinationRateLimitPeriodLessThanNilGreaterThan DestinationRateLimitPeriod = "<nil>"
	DestinationRateLimitPeriodSecond                 DestinationRateLimitPeriod = "second"
	DestinationRateLimitPeriodMinute                 DestinationRateLimitPeriod = "minute"
	DestinationRateLimitPeriodHour                   DestinationRateLimitPeriod = "hour"
)

func (e DestinationRateLimitPeriod) ToPointer() *DestinationRateLimitPeriod {
	return &e
}

func (e *DestinationRateLimitPeriod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "<nil>":
		fallthrough
	case "second":
		fallthrough
	case "minute":
		fallthrough
	case "hour":
		*e = DestinationRateLimitPeriod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationRateLimitPeriod: %v", v)
	}
}
