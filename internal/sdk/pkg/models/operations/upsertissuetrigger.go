// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"hashicups/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpsertIssueTriggerRequestBodyConfigsType string

const (
	UpsertIssueTriggerRequestBodyConfigsTypeIssueTriggerDeliveryConfigs       UpsertIssueTriggerRequestBodyConfigsType = "IssueTriggerDeliveryConfigs"
	UpsertIssueTriggerRequestBodyConfigsTypeIssueTriggerTransformationConfigs UpsertIssueTriggerRequestBodyConfigsType = "IssueTriggerTransformationConfigs"
	UpsertIssueTriggerRequestBodyConfigsTypeIssueTriggerBackpressureConfigs   UpsertIssueTriggerRequestBodyConfigsType = "IssueTriggerBackpressureConfigs"
)

type UpsertIssueTriggerRequestBodyConfigs struct {
	IssueTriggerDeliveryConfigs       *shared.IssueTriggerDeliveryConfigs
	IssueTriggerTransformationConfigs *shared.IssueTriggerTransformationConfigs
	IssueTriggerBackpressureConfigs   *shared.IssueTriggerBackpressureConfigs

	Type UpsertIssueTriggerRequestBodyConfigsType
}

func CreateUpsertIssueTriggerRequestBodyConfigsIssueTriggerDeliveryConfigs(issueTriggerDeliveryConfigs shared.IssueTriggerDeliveryConfigs) UpsertIssueTriggerRequestBodyConfigs {
	typ := UpsertIssueTriggerRequestBodyConfigsTypeIssueTriggerDeliveryConfigs

	return UpsertIssueTriggerRequestBodyConfigs{
		IssueTriggerDeliveryConfigs: &issueTriggerDeliveryConfigs,
		Type:                        typ,
	}
}

func CreateUpsertIssueTriggerRequestBodyConfigsIssueTriggerTransformationConfigs(issueTriggerTransformationConfigs shared.IssueTriggerTransformationConfigs) UpsertIssueTriggerRequestBodyConfigs {
	typ := UpsertIssueTriggerRequestBodyConfigsTypeIssueTriggerTransformationConfigs

	return UpsertIssueTriggerRequestBodyConfigs{
		IssueTriggerTransformationConfigs: &issueTriggerTransformationConfigs,
		Type:                              typ,
	}
}

func CreateUpsertIssueTriggerRequestBodyConfigsIssueTriggerBackpressureConfigs(issueTriggerBackpressureConfigs shared.IssueTriggerBackpressureConfigs) UpsertIssueTriggerRequestBodyConfigs {
	typ := UpsertIssueTriggerRequestBodyConfigsTypeIssueTriggerBackpressureConfigs

	return UpsertIssueTriggerRequestBodyConfigs{
		IssueTriggerBackpressureConfigs: &issueTriggerBackpressureConfigs,
		Type:                            typ,
	}
}

func (u *UpsertIssueTriggerRequestBodyConfigs) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	issueTriggerDeliveryConfigs := new(shared.IssueTriggerDeliveryConfigs)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&issueTriggerDeliveryConfigs); err == nil {
		u.IssueTriggerDeliveryConfigs = issueTriggerDeliveryConfigs
		u.Type = UpsertIssueTriggerRequestBodyConfigsTypeIssueTriggerDeliveryConfigs
		return nil
	}

	issueTriggerTransformationConfigs := new(shared.IssueTriggerTransformationConfigs)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&issueTriggerTransformationConfigs); err == nil {
		u.IssueTriggerTransformationConfigs = issueTriggerTransformationConfigs
		u.Type = UpsertIssueTriggerRequestBodyConfigsTypeIssueTriggerTransformationConfigs
		return nil
	}

	issueTriggerBackpressureConfigs := new(shared.IssueTriggerBackpressureConfigs)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&issueTriggerBackpressureConfigs); err == nil {
		u.IssueTriggerBackpressureConfigs = issueTriggerBackpressureConfigs
		u.Type = UpsertIssueTriggerRequestBodyConfigsTypeIssueTriggerBackpressureConfigs
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UpsertIssueTriggerRequestBodyConfigs) MarshalJSON() ([]byte, error) {
	if u.IssueTriggerDeliveryConfigs != nil {
		return json.Marshal(u.IssueTriggerDeliveryConfigs)
	}

	if u.IssueTriggerTransformationConfigs != nil {
		return json.Marshal(u.IssueTriggerTransformationConfigs)
	}

	if u.IssueTriggerBackpressureConfigs != nil {
		return json.Marshal(u.IssueTriggerBackpressureConfigs)
	}

	return nil, nil
}

type UpsertIssueTriggerRequestBody struct {
	// Notification channels object for the specific channel type
	Channels shared.IssueTriggerChannels `json:"channels"`
	// Configuration object for the specific issue type selected
	Configs *UpsertIssueTriggerRequestBodyConfigs `json:"configs,omitempty"`
	// Required unique name to use as reference when using the API
	Name string `json:"name"`
	// Issue type
	Type shared.IssueType `json:"type"`
}

type UpsertIssueTriggerResponse struct {
	// Bad Request
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	// A single issue trigger
	IssueTrigger *shared.IssueTrigger
	StatusCode   int
	RawResponse  *http.Response
}
