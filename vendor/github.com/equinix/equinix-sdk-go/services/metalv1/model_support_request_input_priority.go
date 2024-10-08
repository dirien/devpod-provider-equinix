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

// SupportRequestInputPriority the model 'SupportRequestInputPriority'
type SupportRequestInputPriority string

// List of SupportRequestInput_priority
const (
	SUPPORTREQUESTINPUTPRIORITY_URGENT SupportRequestInputPriority = "urgent"
	SUPPORTREQUESTINPUTPRIORITY_HIGH   SupportRequestInputPriority = "high"
	SUPPORTREQUESTINPUTPRIORITY_MEDIUM SupportRequestInputPriority = "medium"
	SUPPORTREQUESTINPUTPRIORITY_LOW    SupportRequestInputPriority = "low"
)

// All allowed values of SupportRequestInputPriority enum
var AllowedSupportRequestInputPriorityEnumValues = []SupportRequestInputPriority{
	"urgent",
	"high",
	"medium",
	"low",
}

func (v *SupportRequestInputPriority) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SupportRequestInputPriority(value)
	for _, existing := range AllowedSupportRequestInputPriorityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SupportRequestInputPriority", value)
}

// NewSupportRequestInputPriorityFromValue returns a pointer to a valid SupportRequestInputPriority
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSupportRequestInputPriorityFromValue(v string) (*SupportRequestInputPriority, error) {
	ev := SupportRequestInputPriority(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SupportRequestInputPriority: valid values are %v", v, AllowedSupportRequestInputPriorityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SupportRequestInputPriority) IsValid() bool {
	for _, existing := range AllowedSupportRequestInputPriorityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SupportRequestInput_priority value
func (v SupportRequestInputPriority) Ptr() *SupportRequestInputPriority {
	return &v
}

type NullableSupportRequestInputPriority struct {
	value *SupportRequestInputPriority
	isSet bool
}

func (v NullableSupportRequestInputPriority) Get() *SupportRequestInputPriority {
	return v.value
}

func (v *NullableSupportRequestInputPriority) Set(val *SupportRequestInputPriority) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportRequestInputPriority) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportRequestInputPriority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportRequestInputPriority(val *SupportRequestInputPriority) *NullableSupportRequestInputPriority {
	return &NullableSupportRequestInputPriority{value: val, isSet: true}
}

func (v NullableSupportRequestInputPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportRequestInputPriority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
