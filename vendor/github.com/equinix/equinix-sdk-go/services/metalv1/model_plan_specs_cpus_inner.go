/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the PlanSpecsCpusInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanSpecsCpusInner{}

// PlanSpecsCpusInner struct for PlanSpecsCpusInner
type PlanSpecsCpusInner struct {
	Count                *int32  `json:"count,omitempty"`
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PlanSpecsCpusInner PlanSpecsCpusInner

// NewPlanSpecsCpusInner instantiates a new PlanSpecsCpusInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanSpecsCpusInner() *PlanSpecsCpusInner {
	this := PlanSpecsCpusInner{}
	return &this
}

// NewPlanSpecsCpusInnerWithDefaults instantiates a new PlanSpecsCpusInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanSpecsCpusInnerWithDefaults() *PlanSpecsCpusInner {
	this := PlanSpecsCpusInner{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PlanSpecsCpusInner) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanSpecsCpusInner) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PlanSpecsCpusInner) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PlanSpecsCpusInner) SetCount(v int32) {
	o.Count = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PlanSpecsCpusInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanSpecsCpusInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PlanSpecsCpusInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PlanSpecsCpusInner) SetType(v string) {
	o.Type = &v
}

func (o PlanSpecsCpusInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlanSpecsCpusInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PlanSpecsCpusInner) UnmarshalJSON(data []byte) (err error) {
	varPlanSpecsCpusInner := _PlanSpecsCpusInner{}

	err = json.Unmarshal(data, &varPlanSpecsCpusInner)

	if err != nil {
		return err
	}

	*o = PlanSpecsCpusInner(varPlanSpecsCpusInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlanSpecsCpusInner struct {
	value *PlanSpecsCpusInner
	isSet bool
}

func (v NullablePlanSpecsCpusInner) Get() *PlanSpecsCpusInner {
	return v.value
}

func (v *NullablePlanSpecsCpusInner) Set(val *PlanSpecsCpusInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanSpecsCpusInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanSpecsCpusInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanSpecsCpusInner(val *PlanSpecsCpusInner) *NullablePlanSpecsCpusInner {
	return &NullablePlanSpecsCpusInner{value: val, isSet: true}
}

func (v NullablePlanSpecsCpusInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanSpecsCpusInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
