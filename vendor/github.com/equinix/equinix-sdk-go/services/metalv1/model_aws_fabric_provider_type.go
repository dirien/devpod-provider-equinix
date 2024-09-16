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

// AWSFabricProviderType the model 'AWSFabricProviderType'
type AWSFabricProviderType string

// List of AWSFabricProvider_type
const (
	AWSFABRICPROVIDERTYPE_CSP_AWS AWSFabricProviderType = "CSP_AWS"
)

// All allowed values of AWSFabricProviderType enum
var AllowedAWSFabricProviderTypeEnumValues = []AWSFabricProviderType{
	"CSP_AWS",
}

func (v *AWSFabricProviderType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AWSFabricProviderType(value)
	for _, existing := range AllowedAWSFabricProviderTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AWSFabricProviderType", value)
}

// NewAWSFabricProviderTypeFromValue returns a pointer to a valid AWSFabricProviderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAWSFabricProviderTypeFromValue(v string) (*AWSFabricProviderType, error) {
	ev := AWSFabricProviderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AWSFabricProviderType: valid values are %v", v, AllowedAWSFabricProviderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AWSFabricProviderType) IsValid() bool {
	for _, existing := range AllowedAWSFabricProviderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AWSFabricProvider_type value
func (v AWSFabricProviderType) Ptr() *AWSFabricProviderType {
	return &v
}

type NullableAWSFabricProviderType struct {
	value *AWSFabricProviderType
	isSet bool
}

func (v NullableAWSFabricProviderType) Get() *AWSFabricProviderType {
	return v.value
}

func (v *NullableAWSFabricProviderType) Set(val *AWSFabricProviderType) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSFabricProviderType) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSFabricProviderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSFabricProviderType(val *AWSFabricProviderType) *NullableAWSFabricProviderType {
	return &NullableAWSFabricProviderType{value: val, isSet: true}
}

func (v NullableAWSFabricProviderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSFabricProviderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
