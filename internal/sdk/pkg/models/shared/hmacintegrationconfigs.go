// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type HMACIntegrationConfigsEncoding string

const (
	HMACIntegrationConfigsEncodingBase64 HMACIntegrationConfigsEncoding = "base64"
	HMACIntegrationConfigsEncodingHex    HMACIntegrationConfigsEncoding = "hex"
)

func (e HMACIntegrationConfigsEncoding) ToPointer() *HMACIntegrationConfigsEncoding {
	return &e
}

func (e *HMACIntegrationConfigsEncoding) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "base64":
		fallthrough
	case "hex":
		*e = HMACIntegrationConfigsEncoding(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HMACIntegrationConfigsEncoding: %v", v)
	}
}

// HMACIntegrationConfigs - Decrypted Key/Value object of the associated configuration for that provider
type HMACIntegrationConfigs struct {
	Algorithm        HMACAlgorithms                 `json:"algorithm"`
	Encoding         HMACIntegrationConfigsEncoding `json:"encoding"`
	HeaderKey        string                         `json:"header_key"`
	WebhookSecretKey string                         `json:"webhook_secret_key"`
}
