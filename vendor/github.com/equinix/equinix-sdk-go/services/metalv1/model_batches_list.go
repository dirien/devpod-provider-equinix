/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the BatchesList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchesList{}

// BatchesList struct for BatchesList
type BatchesList struct {
	Batches              []Batch `json:"batches,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BatchesList BatchesList

// NewBatchesList instantiates a new BatchesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchesList() *BatchesList {
	this := BatchesList{}
	return &this
}

// NewBatchesListWithDefaults instantiates a new BatchesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchesListWithDefaults() *BatchesList {
	this := BatchesList{}
	return &this
}

// GetBatches returns the Batches field value if set, zero value otherwise.
func (o *BatchesList) GetBatches() []Batch {
	if o == nil || IsNil(o.Batches) {
		var ret []Batch
		return ret
	}
	return o.Batches
}

// GetBatchesOk returns a tuple with the Batches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchesList) GetBatchesOk() ([]Batch, bool) {
	if o == nil || IsNil(o.Batches) {
		return nil, false
	}
	return o.Batches, true
}

// HasBatches returns a boolean if a field has been set.
func (o *BatchesList) HasBatches() bool {
	if o != nil && !IsNil(o.Batches) {
		return true
	}

	return false
}

// SetBatches gets a reference to the given []Batch and assigns it to the Batches field.
func (o *BatchesList) SetBatches(v []Batch) {
	o.Batches = v
}

func (o BatchesList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchesList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Batches) {
		toSerialize["batches"] = o.Batches
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BatchesList) UnmarshalJSON(data []byte) (err error) {
	varBatchesList := _BatchesList{}

	err = json.Unmarshal(data, &varBatchesList)

	if err != nil {
		return err
	}

	*o = BatchesList(varBatchesList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "batches")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBatchesList struct {
	value *BatchesList
	isSet bool
}

func (v NullableBatchesList) Get() *BatchesList {
	return v.value
}

func (v *NullableBatchesList) Set(val *BatchesList) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchesList) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchesList(val *BatchesList) *NullableBatchesList {
	return &NullableBatchesList{value: val, isSet: true}
}

func (v NullableBatchesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
