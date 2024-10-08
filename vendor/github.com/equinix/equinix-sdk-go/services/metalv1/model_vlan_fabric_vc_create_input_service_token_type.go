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

// VlanFabricVcCreateInputServiceTokenType Either 'a_side' or 'z_side'. Setting this field to 'a_side' will create an interconnection with Fabric VCs (Metal Billed). Setting this field to 'z_side' will create an interconnection with Fabric VCs (Fabric Billed). This is required when the 'type' is 'shared', but this is not applicable when the 'type' is 'dedicated'. This parameter is included in the specification as a developer preview and is generally unavailable. Please contact our Support team for more details.
type VlanFabricVcCreateInputServiceTokenType string

// List of VlanFabricVcCreateInput_service_token_type
const (
	VLANFABRICVCCREATEINPUTSERVICETOKENTYPE_A_SIDE VlanFabricVcCreateInputServiceTokenType = "a_side"
	VLANFABRICVCCREATEINPUTSERVICETOKENTYPE_Z_SIDE VlanFabricVcCreateInputServiceTokenType = "z_side"
)

// All allowed values of VlanFabricVcCreateInputServiceTokenType enum
var AllowedVlanFabricVcCreateInputServiceTokenTypeEnumValues = []VlanFabricVcCreateInputServiceTokenType{
	"a_side",
	"z_side",
}

func (v *VlanFabricVcCreateInputServiceTokenType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VlanFabricVcCreateInputServiceTokenType(value)
	for _, existing := range AllowedVlanFabricVcCreateInputServiceTokenTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VlanFabricVcCreateInputServiceTokenType", value)
}

// NewVlanFabricVcCreateInputServiceTokenTypeFromValue returns a pointer to a valid VlanFabricVcCreateInputServiceTokenType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVlanFabricVcCreateInputServiceTokenTypeFromValue(v string) (*VlanFabricVcCreateInputServiceTokenType, error) {
	ev := VlanFabricVcCreateInputServiceTokenType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VlanFabricVcCreateInputServiceTokenType: valid values are %v", v, AllowedVlanFabricVcCreateInputServiceTokenTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VlanFabricVcCreateInputServiceTokenType) IsValid() bool {
	for _, existing := range AllowedVlanFabricVcCreateInputServiceTokenTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VlanFabricVcCreateInput_service_token_type value
func (v VlanFabricVcCreateInputServiceTokenType) Ptr() *VlanFabricVcCreateInputServiceTokenType {
	return &v
}

type NullableVlanFabricVcCreateInputServiceTokenType struct {
	value *VlanFabricVcCreateInputServiceTokenType
	isSet bool
}

func (v NullableVlanFabricVcCreateInputServiceTokenType) Get() *VlanFabricVcCreateInputServiceTokenType {
	return v.value
}

func (v *NullableVlanFabricVcCreateInputServiceTokenType) Set(val *VlanFabricVcCreateInputServiceTokenType) {
	v.value = val
	v.isSet = true
}

func (v NullableVlanFabricVcCreateInputServiceTokenType) IsSet() bool {
	return v.isSet
}

func (v *NullableVlanFabricVcCreateInputServiceTokenType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVlanFabricVcCreateInputServiceTokenType(val *VlanFabricVcCreateInputServiceTokenType) *NullableVlanFabricVcCreateInputServiceTokenType {
	return &NullableVlanFabricVcCreateInputServiceTokenType{value: val, isSet: true}
}

func (v NullableVlanFabricVcCreateInputServiceTokenType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVlanFabricVcCreateInputServiceTokenType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
