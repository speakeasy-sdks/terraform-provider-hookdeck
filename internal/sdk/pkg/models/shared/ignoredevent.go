// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"hashicups/internal/sdk/pkg/utils"
	"time"
)

type IgnoredEventMetaTransformationFailedMeta struct {
	TransformationID string `json:"transformation_id"`
}

func (o *IgnoredEventMetaTransformationFailedMeta) GetTransformationID() string {
	if o == nil {
		return ""
	}
	return o.TransformationID
}

type IgnoredEventMetaFilteredMeta string

const (
	IgnoredEventMetaFilteredMetaBody    IgnoredEventMetaFilteredMeta = "body"
	IgnoredEventMetaFilteredMetaHeaders IgnoredEventMetaFilteredMeta = "headers"
	IgnoredEventMetaFilteredMetaPath    IgnoredEventMetaFilteredMeta = "path"
	IgnoredEventMetaFilteredMetaQuery   IgnoredEventMetaFilteredMeta = "query"
)

func (e IgnoredEventMetaFilteredMeta) ToPointer() *IgnoredEventMetaFilteredMeta {
	return &e
}

func (e *IgnoredEventMetaFilteredMeta) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "body":
		fallthrough
	case "headers":
		fallthrough
	case "path":
		fallthrough
	case "query":
		*e = IgnoredEventMetaFilteredMeta(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IgnoredEventMetaFilteredMeta: %v", v)
	}
}

type IgnoredEventMetaType string

const (
	IgnoredEventMetaTypeIgnoredEventMetaFilteredMeta             IgnoredEventMetaType = "IgnoredEvent_meta_FilteredMeta"
	IgnoredEventMetaTypeIgnoredEventMetaTransformationFailedMeta IgnoredEventMetaType = "IgnoredEvent_meta_TransformationFailedMeta"
)

type IgnoredEventMeta struct {
	IgnoredEventMetaFilteredMeta             *IgnoredEventMetaFilteredMeta
	IgnoredEventMetaTransformationFailedMeta *IgnoredEventMetaTransformationFailedMeta

	Type IgnoredEventMetaType
}

func CreateIgnoredEventMetaIgnoredEventMetaFilteredMeta(ignoredEventMetaFilteredMeta IgnoredEventMetaFilteredMeta) IgnoredEventMeta {
	typ := IgnoredEventMetaTypeIgnoredEventMetaFilteredMeta

	return IgnoredEventMeta{
		IgnoredEventMetaFilteredMeta: &ignoredEventMetaFilteredMeta,
		Type:                         typ,
	}
}

func CreateIgnoredEventMetaIgnoredEventMetaTransformationFailedMeta(ignoredEventMetaTransformationFailedMeta IgnoredEventMetaTransformationFailedMeta) IgnoredEventMeta {
	typ := IgnoredEventMetaTypeIgnoredEventMetaTransformationFailedMeta

	return IgnoredEventMeta{
		IgnoredEventMetaTransformationFailedMeta: &ignoredEventMetaTransformationFailedMeta,
		Type:                                     typ,
	}
}

func (u *IgnoredEventMeta) UnmarshalJSON(data []byte) error {

	ignoredEventMetaTransformationFailedMeta := new(IgnoredEventMetaTransformationFailedMeta)
	if err := utils.UnmarshalJSON(data, &ignoredEventMetaTransformationFailedMeta, "", true, true); err == nil {
		u.IgnoredEventMetaTransformationFailedMeta = ignoredEventMetaTransformationFailedMeta
		u.Type = IgnoredEventMetaTypeIgnoredEventMetaTransformationFailedMeta
		return nil
	}

	ignoredEventMetaFilteredMeta := new(IgnoredEventMetaFilteredMeta)
	if err := utils.UnmarshalJSON(data, &ignoredEventMetaFilteredMeta, "", true, true); err == nil {
		u.IgnoredEventMetaFilteredMeta = ignoredEventMetaFilteredMeta
		u.Type = IgnoredEventMetaTypeIgnoredEventMetaFilteredMeta
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u IgnoredEventMeta) MarshalJSON() ([]byte, error) {
	if u.IgnoredEventMetaFilteredMeta != nil {
		return utils.MarshalJSON(u.IgnoredEventMetaFilteredMeta, "", true)
	}

	if u.IgnoredEventMetaTransformationFailedMeta != nil {
		return utils.MarshalJSON(u.IgnoredEventMetaTransformationFailedMeta, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type IgnoredEvent struct {
	Cause     IgnoredEventCause `json:"cause"`
	CreatedAt time.Time         `json:"created_at"`
	ID        string            `json:"id"`
	Meta      *IgnoredEventMeta `json:"meta,omitempty"`
	RequestID string            `json:"request_id"`
	TeamID    string            `json:"team_id"`
	UpdatedAt time.Time         `json:"updated_at"`
	WebhookID string            `json:"webhook_id"`
}

func (i IgnoredEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IgnoredEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *IgnoredEvent) GetCause() IgnoredEventCause {
	if o == nil {
		return IgnoredEventCause("")
	}
	return o.Cause
}

func (o *IgnoredEvent) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *IgnoredEvent) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *IgnoredEvent) GetMeta() *IgnoredEventMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *IgnoredEvent) GetRequestID() string {
	if o == nil {
		return ""
	}
	return o.RequestID
}

func (o *IgnoredEvent) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *IgnoredEvent) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *IgnoredEvent) GetWebhookID() string {
	if o == nil {
		return ""
	}
	return o.WebhookID
}
