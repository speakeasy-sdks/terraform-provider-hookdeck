// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// APIKeyType - Type of auth method
type APIKeyType string

const (
	APIKeyTypeAPIKey APIKeyType = "API_KEY"
)

func (e APIKeyType) ToPointer() *APIKeyType {
	return &e
}

func (e *APIKeyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "API_KEY":
		*e = APIKeyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for APIKeyType: %v", v)
	}
}

// APIKey - API Key
type APIKey struct {
	// API key config for the destination's auth method
	Config *DestinationAuthMethodAPIKeyConfig `json:"config,omitempty"`
	// Type of auth method
	Type APIKeyType `json:"type"`
}

func (o *APIKey) GetConfig() *DestinationAuthMethodAPIKeyConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *APIKey) GetType() APIKeyType {
	if o == nil {
		return APIKeyType("")
	}
	return o.Type
}
