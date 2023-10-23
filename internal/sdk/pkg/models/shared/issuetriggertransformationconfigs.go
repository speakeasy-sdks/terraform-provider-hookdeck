// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"hashicups/internal/sdk/pkg/utils"
)

type IssueTriggerTransformationConfigsTransformationsType string

const (
	IssueTriggerTransformationConfigsTransformationsTypeStr        IssueTriggerTransformationConfigsTransformationsType = "str"
	IssueTriggerTransformationConfigsTransformationsTypeArrayOfstr IssueTriggerTransformationConfigsTransformationsType = "arrayOfstr"
)

type IssueTriggerTransformationConfigsTransformations struct {
	Str        *string
	ArrayOfstr []string

	Type IssueTriggerTransformationConfigsTransformationsType
}

func CreateIssueTriggerTransformationConfigsTransformationsStr(str string) IssueTriggerTransformationConfigsTransformations {
	typ := IssueTriggerTransformationConfigsTransformationsTypeStr

	return IssueTriggerTransformationConfigsTransformations{
		Str:  &str,
		Type: typ,
	}
}

func CreateIssueTriggerTransformationConfigsTransformationsArrayOfstr(arrayOfstr []string) IssueTriggerTransformationConfigsTransformations {
	typ := IssueTriggerTransformationConfigsTransformationsTypeArrayOfstr

	return IssueTriggerTransformationConfigsTransformations{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *IssueTriggerTransformationConfigsTransformations) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = IssueTriggerTransformationConfigsTransformationsTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = IssueTriggerTransformationConfigsTransformationsTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u IssueTriggerTransformationConfigsTransformations) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// IssueTriggerTransformationConfigs - Configurations for a 'Transformation' issue trigger
type IssueTriggerTransformationConfigs struct {
	// The minimum log level to open the issue on
	LogLevel TransformationExecutionLogLevel `json:"log_level"`
	// A pattern to match on the transformation name or array of transformation IDs. Use `*` as wildcard.
	Transformations IssueTriggerTransformationConfigsTransformations `json:"transformations"`
}

func (o *IssueTriggerTransformationConfigs) GetLogLevel() TransformationExecutionLogLevel {
	if o == nil {
		return TransformationExecutionLogLevel("")
	}
	return o.LogLevel
}

func (o *IssueTriggerTransformationConfigs) GetTransformations() IssueTriggerTransformationConfigsTransformations {
	if o == nil {
		return IssueTriggerTransformationConfigsTransformations{}
	}
	return o.Transformations
}
