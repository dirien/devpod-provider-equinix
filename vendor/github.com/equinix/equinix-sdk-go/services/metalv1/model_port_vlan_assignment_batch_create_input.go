/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the PortVlanAssignmentBatchCreateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortVlanAssignmentBatchCreateInput{}

// PortVlanAssignmentBatchCreateInput struct for PortVlanAssignmentBatchCreateInput
type PortVlanAssignmentBatchCreateInput struct {
	VlanAssignments      []PortVlanAssignmentBatchCreateInputVlanAssignmentsInner `json:"vlan_assignments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortVlanAssignmentBatchCreateInput PortVlanAssignmentBatchCreateInput

// NewPortVlanAssignmentBatchCreateInput instantiates a new PortVlanAssignmentBatchCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortVlanAssignmentBatchCreateInput() *PortVlanAssignmentBatchCreateInput {
	this := PortVlanAssignmentBatchCreateInput{}
	return &this
}

// NewPortVlanAssignmentBatchCreateInputWithDefaults instantiates a new PortVlanAssignmentBatchCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortVlanAssignmentBatchCreateInputWithDefaults() *PortVlanAssignmentBatchCreateInput {
	this := PortVlanAssignmentBatchCreateInput{}
	return &this
}

// GetVlanAssignments returns the VlanAssignments field value if set, zero value otherwise.
func (o *PortVlanAssignmentBatchCreateInput) GetVlanAssignments() []PortVlanAssignmentBatchCreateInputVlanAssignmentsInner {
	if o == nil || IsNil(o.VlanAssignments) {
		var ret []PortVlanAssignmentBatchCreateInputVlanAssignmentsInner
		return ret
	}
	return o.VlanAssignments
}

// GetVlanAssignmentsOk returns a tuple with the VlanAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignmentBatchCreateInput) GetVlanAssignmentsOk() ([]PortVlanAssignmentBatchCreateInputVlanAssignmentsInner, bool) {
	if o == nil || IsNil(o.VlanAssignments) {
		return nil, false
	}
	return o.VlanAssignments, true
}

// HasVlanAssignments returns a boolean if a field has been set.
func (o *PortVlanAssignmentBatchCreateInput) HasVlanAssignments() bool {
	if o != nil && !IsNil(o.VlanAssignments) {
		return true
	}

	return false
}

// SetVlanAssignments gets a reference to the given []PortVlanAssignmentBatchCreateInputVlanAssignmentsInner and assigns it to the VlanAssignments field.
func (o *PortVlanAssignmentBatchCreateInput) SetVlanAssignments(v []PortVlanAssignmentBatchCreateInputVlanAssignmentsInner) {
	o.VlanAssignments = v
}

func (o PortVlanAssignmentBatchCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortVlanAssignmentBatchCreateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VlanAssignments) {
		toSerialize["vlan_assignments"] = o.VlanAssignments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortVlanAssignmentBatchCreateInput) UnmarshalJSON(data []byte) (err error) {
	varPortVlanAssignmentBatchCreateInput := _PortVlanAssignmentBatchCreateInput{}

	err = json.Unmarshal(data, &varPortVlanAssignmentBatchCreateInput)

	if err != nil {
		return err
	}

	*o = PortVlanAssignmentBatchCreateInput(varPortVlanAssignmentBatchCreateInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "vlan_assignments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortVlanAssignmentBatchCreateInput struct {
	value *PortVlanAssignmentBatchCreateInput
	isSet bool
}

func (v NullablePortVlanAssignmentBatchCreateInput) Get() *PortVlanAssignmentBatchCreateInput {
	return v.value
}

func (v *NullablePortVlanAssignmentBatchCreateInput) Set(val *PortVlanAssignmentBatchCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullablePortVlanAssignmentBatchCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullablePortVlanAssignmentBatchCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortVlanAssignmentBatchCreateInput(val *PortVlanAssignmentBatchCreateInput) *NullablePortVlanAssignmentBatchCreateInput {
	return &NullablePortVlanAssignmentBatchCreateInput{value: val, isSet: true}
}

func (v NullablePortVlanAssignmentBatchCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortVlanAssignmentBatchCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
