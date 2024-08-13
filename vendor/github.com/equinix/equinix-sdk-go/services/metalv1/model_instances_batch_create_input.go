/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the InstancesBatchCreateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstancesBatchCreateInput{}

// InstancesBatchCreateInput struct for InstancesBatchCreateInput
type InstancesBatchCreateInput struct {
	Batches              []InstancesBatchCreateInputBatchesInner `json:"batches,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InstancesBatchCreateInput InstancesBatchCreateInput

// NewInstancesBatchCreateInput instantiates a new InstancesBatchCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstancesBatchCreateInput() *InstancesBatchCreateInput {
	this := InstancesBatchCreateInput{}
	return &this
}

// NewInstancesBatchCreateInputWithDefaults instantiates a new InstancesBatchCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstancesBatchCreateInputWithDefaults() *InstancesBatchCreateInput {
	this := InstancesBatchCreateInput{}
	return &this
}

// GetBatches returns the Batches field value if set, zero value otherwise.
func (o *InstancesBatchCreateInput) GetBatches() []InstancesBatchCreateInputBatchesInner {
	if o == nil || IsNil(o.Batches) {
		var ret []InstancesBatchCreateInputBatchesInner
		return ret
	}
	return o.Batches
}

// GetBatchesOk returns a tuple with the Batches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesBatchCreateInput) GetBatchesOk() ([]InstancesBatchCreateInputBatchesInner, bool) {
	if o == nil || IsNil(o.Batches) {
		return nil, false
	}
	return o.Batches, true
}

// HasBatches returns a boolean if a field has been set.
func (o *InstancesBatchCreateInput) HasBatches() bool {
	if o != nil && !IsNil(o.Batches) {
		return true
	}

	return false
}

// SetBatches gets a reference to the given []InstancesBatchCreateInputBatchesInner and assigns it to the Batches field.
func (o *InstancesBatchCreateInput) SetBatches(v []InstancesBatchCreateInputBatchesInner) {
	o.Batches = v
}

func (o InstancesBatchCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstancesBatchCreateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Batches) {
		toSerialize["batches"] = o.Batches
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InstancesBatchCreateInput) UnmarshalJSON(data []byte) (err error) {
	varInstancesBatchCreateInput := _InstancesBatchCreateInput{}

	err = json.Unmarshal(data, &varInstancesBatchCreateInput)

	if err != nil {
		return err
	}

	*o = InstancesBatchCreateInput(varInstancesBatchCreateInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "batches")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInstancesBatchCreateInput struct {
	value *InstancesBatchCreateInput
	isSet bool
}

func (v NullableInstancesBatchCreateInput) Get() *InstancesBatchCreateInput {
	return v.value
}

func (v *NullableInstancesBatchCreateInput) Set(val *InstancesBatchCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableInstancesBatchCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableInstancesBatchCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstancesBatchCreateInput(val *InstancesBatchCreateInput) *NullableInstancesBatchCreateInput {
	return &NullableInstancesBatchCreateInput{value: val, isSet: true}
}

func (v NullableInstancesBatchCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstancesBatchCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
