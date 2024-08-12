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

// InterconnectionMode The mode of the interconnection (only relevant to Dedicated Ports). Shared connections won't have this field. Can be either 'standard' or 'tunnel'.   The default mode of an interconnection on a Dedicated Port is 'standard'. The mode can only be changed when there are no associated virtual circuits on the interconnection.   In tunnel mode, an 802.1q tunnel is added to a port to send/receive double tagged packets from server instances.
type InterconnectionMode string

// List of Interconnection_mode
const (
	INTERCONNECTIONMODE_STANDARD InterconnectionMode = "standard"
	INTERCONNECTIONMODE_TUNNEL   InterconnectionMode = "tunnel"
)

// All allowed values of InterconnectionMode enum
var AllowedInterconnectionModeEnumValues = []InterconnectionMode{
	"standard",
	"tunnel",
}

func (v *InterconnectionMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InterconnectionMode(value)
	for _, existing := range AllowedInterconnectionModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InterconnectionMode", value)
}

// NewInterconnectionModeFromValue returns a pointer to a valid InterconnectionMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterconnectionModeFromValue(v string) (*InterconnectionMode, error) {
	ev := InterconnectionMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InterconnectionMode: valid values are %v", v, AllowedInterconnectionModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InterconnectionMode) IsValid() bool {
	for _, existing := range AllowedInterconnectionModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Interconnection_mode value
func (v InterconnectionMode) Ptr() *InterconnectionMode {
	return &v
}

type NullableInterconnectionMode struct {
	value *InterconnectionMode
	isSet bool
}

func (v NullableInterconnectionMode) Get() *InterconnectionMode {
	return v.value
}

func (v *NullableInterconnectionMode) Set(val *InterconnectionMode) {
	v.value = val
	v.isSet = true
}

func (v NullableInterconnectionMode) IsSet() bool {
	return v.isSet
}

func (v *NullableInterconnectionMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterconnectionMode(val *InterconnectionMode) *NullableInterconnectionMode {
	return &NullableInterconnectionMode{value: val, isSet: true}
}

func (v NullableInterconnectionMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterconnectionMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
