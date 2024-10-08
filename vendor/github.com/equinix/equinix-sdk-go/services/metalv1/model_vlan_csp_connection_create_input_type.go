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

// VlanCSPConnectionCreateInputType the model 'VlanCSPConnectionCreateInputType'
type VlanCSPConnectionCreateInputType string

// List of VlanCSPConnectionCreateInput_type
const (
	VLANCSPCONNECTIONCREATEINPUTTYPE_SHARED_PORT_VLAN_TO_CSP VlanCSPConnectionCreateInputType = "shared_port_vlan_to_csp"
)

// All allowed values of VlanCSPConnectionCreateInputType enum
var AllowedVlanCSPConnectionCreateInputTypeEnumValues = []VlanCSPConnectionCreateInputType{
	"shared_port_vlan_to_csp",
}

func (v *VlanCSPConnectionCreateInputType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VlanCSPConnectionCreateInputType(value)
	for _, existing := range AllowedVlanCSPConnectionCreateInputTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VlanCSPConnectionCreateInputType", value)
}

// NewVlanCSPConnectionCreateInputTypeFromValue returns a pointer to a valid VlanCSPConnectionCreateInputType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVlanCSPConnectionCreateInputTypeFromValue(v string) (*VlanCSPConnectionCreateInputType, error) {
	ev := VlanCSPConnectionCreateInputType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VlanCSPConnectionCreateInputType: valid values are %v", v, AllowedVlanCSPConnectionCreateInputTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VlanCSPConnectionCreateInputType) IsValid() bool {
	for _, existing := range AllowedVlanCSPConnectionCreateInputTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VlanCSPConnectionCreateInput_type value
func (v VlanCSPConnectionCreateInputType) Ptr() *VlanCSPConnectionCreateInputType {
	return &v
}

type NullableVlanCSPConnectionCreateInputType struct {
	value *VlanCSPConnectionCreateInputType
	isSet bool
}

func (v NullableVlanCSPConnectionCreateInputType) Get() *VlanCSPConnectionCreateInputType {
	return v.value
}

func (v *NullableVlanCSPConnectionCreateInputType) Set(val *VlanCSPConnectionCreateInputType) {
	v.value = val
	v.isSet = true
}

func (v NullableVlanCSPConnectionCreateInputType) IsSet() bool {
	return v.isSet
}

func (v *NullableVlanCSPConnectionCreateInputType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVlanCSPConnectionCreateInputType(val *VlanCSPConnectionCreateInputType) *NullableVlanCSPConnectionCreateInputType {
	return &NullableVlanCSPConnectionCreateInputType{value: val, isSet: true}
}

func (v NullableVlanCSPConnectionCreateInputType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVlanCSPConnectionCreateInputType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
