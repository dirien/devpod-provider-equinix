/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
	"fmt"
)

// FindTrafficDirectionParameter the model 'FindTrafficDirectionParameter'
type FindTrafficDirectionParameter string

// List of findTraffic_direction_parameter
const (
	FINDTRAFFICDIRECTIONPARAMETER_INBOUND  FindTrafficDirectionParameter = "inbound"
	FINDTRAFFICDIRECTIONPARAMETER_OUTBOUND FindTrafficDirectionParameter = "outbound"
)

// All allowed values of FindTrafficDirectionParameter enum
var AllowedFindTrafficDirectionParameterEnumValues = []FindTrafficDirectionParameter{
	"inbound",
	"outbound",
}

func (v *FindTrafficDirectionParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FindTrafficDirectionParameter(value)
	for _, existing := range AllowedFindTrafficDirectionParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FindTrafficDirectionParameter", value)
}

// NewFindTrafficDirectionParameterFromValue returns a pointer to a valid FindTrafficDirectionParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFindTrafficDirectionParameterFromValue(v string) (*FindTrafficDirectionParameter, error) {
	ev := FindTrafficDirectionParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FindTrafficDirectionParameter: valid values are %v", v, AllowedFindTrafficDirectionParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FindTrafficDirectionParameter) IsValid() bool {
	for _, existing := range AllowedFindTrafficDirectionParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to findTraffic_direction_parameter value
func (v FindTrafficDirectionParameter) Ptr() *FindTrafficDirectionParameter {
	return &v
}

type NullableFindTrafficDirectionParameter struct {
	value *FindTrafficDirectionParameter
	isSet bool
}

func (v NullableFindTrafficDirectionParameter) Get() *FindTrafficDirectionParameter {
	return v.value
}

func (v *NullableFindTrafficDirectionParameter) Set(val *FindTrafficDirectionParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableFindTrafficDirectionParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableFindTrafficDirectionParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindTrafficDirectionParameter(val *FindTrafficDirectionParameter) *NullableFindTrafficDirectionParameter {
	return &NullableFindTrafficDirectionParameter{value: val, isSet: true}
}

func (v NullableFindTrafficDirectionParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindTrafficDirectionParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
