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

// IPAssignmentType the model 'IPAssignmentType'
type IPAssignmentType string

// List of IPAssignment_type
const (
	IPASSIGNMENTTYPE_IP_ASSIGNMENT IPAssignmentType = "IPAssignment"
)

// All allowed values of IPAssignmentType enum
var AllowedIPAssignmentTypeEnumValues = []IPAssignmentType{
	"IPAssignment",
}

func (v *IPAssignmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IPAssignmentType(value)
	for _, existing := range AllowedIPAssignmentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IPAssignmentType", value)
}

// NewIPAssignmentTypeFromValue returns a pointer to a valid IPAssignmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIPAssignmentTypeFromValue(v string) (*IPAssignmentType, error) {
	ev := IPAssignmentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IPAssignmentType: valid values are %v", v, AllowedIPAssignmentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IPAssignmentType) IsValid() bool {
	for _, existing := range AllowedIPAssignmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IPAssignment_type value
func (v IPAssignmentType) Ptr() *IPAssignmentType {
	return &v
}

type NullableIPAssignmentType struct {
	value *IPAssignmentType
	isSet bool
}

func (v NullableIPAssignmentType) Get() *IPAssignmentType {
	return v.value
}

func (v *NullableIPAssignmentType) Set(val *IPAssignmentType) {
	v.value = val
	v.isSet = true
}

func (v NullableIPAssignmentType) IsSet() bool {
	return v.isSet
}

func (v *NullableIPAssignmentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPAssignmentType(val *IPAssignmentType) *NullableIPAssignmentType {
	return &NullableIPAssignmentType{value: val, isSet: true}
}

func (v NullableIPAssignmentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPAssignmentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
