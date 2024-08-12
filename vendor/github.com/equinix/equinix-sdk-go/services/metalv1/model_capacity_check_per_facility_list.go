/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the CapacityCheckPerFacilityList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapacityCheckPerFacilityList{}

// CapacityCheckPerFacilityList struct for CapacityCheckPerFacilityList
type CapacityCheckPerFacilityList struct {
	Servers              []CapacityCheckPerFacilityInfo `json:"servers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapacityCheckPerFacilityList CapacityCheckPerFacilityList

// NewCapacityCheckPerFacilityList instantiates a new CapacityCheckPerFacilityList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapacityCheckPerFacilityList() *CapacityCheckPerFacilityList {
	this := CapacityCheckPerFacilityList{}
	return &this
}

// NewCapacityCheckPerFacilityListWithDefaults instantiates a new CapacityCheckPerFacilityList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapacityCheckPerFacilityListWithDefaults() *CapacityCheckPerFacilityList {
	this := CapacityCheckPerFacilityList{}
	return &this
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *CapacityCheckPerFacilityList) GetServers() []CapacityCheckPerFacilityInfo {
	if o == nil || IsNil(o.Servers) {
		var ret []CapacityCheckPerFacilityInfo
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityCheckPerFacilityList) GetServersOk() ([]CapacityCheckPerFacilityInfo, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *CapacityCheckPerFacilityList) HasServers() bool {
	if o != nil && !IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []CapacityCheckPerFacilityInfo and assigns it to the Servers field.
func (o *CapacityCheckPerFacilityList) SetServers(v []CapacityCheckPerFacilityInfo) {
	o.Servers = v
}

func (o CapacityCheckPerFacilityList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapacityCheckPerFacilityList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapacityCheckPerFacilityList) UnmarshalJSON(data []byte) (err error) {
	varCapacityCheckPerFacilityList := _CapacityCheckPerFacilityList{}

	err = json.Unmarshal(data, &varCapacityCheckPerFacilityList)

	if err != nil {
		return err
	}

	*o = CapacityCheckPerFacilityList(varCapacityCheckPerFacilityList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "servers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapacityCheckPerFacilityList struct {
	value *CapacityCheckPerFacilityList
	isSet bool
}

func (v NullableCapacityCheckPerFacilityList) Get() *CapacityCheckPerFacilityList {
	return v.value
}

func (v *NullableCapacityCheckPerFacilityList) Set(val *CapacityCheckPerFacilityList) {
	v.value = val
	v.isSet = true
}

func (v NullableCapacityCheckPerFacilityList) IsSet() bool {
	return v.isSet
}

func (v *NullableCapacityCheckPerFacilityList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapacityCheckPerFacilityList(val *CapacityCheckPerFacilityList) *NullableCapacityCheckPerFacilityList {
	return &NullableCapacityCheckPerFacilityList{value: val, isSet: true}
}

func (v NullableCapacityCheckPerFacilityList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapacityCheckPerFacilityList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
