/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the ProjectUsageList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectUsageList{}

// ProjectUsageList struct for ProjectUsageList
type ProjectUsageList struct {
	Usages               []ProjectUsage `json:"usages,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProjectUsageList ProjectUsageList

// NewProjectUsageList instantiates a new ProjectUsageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectUsageList() *ProjectUsageList {
	this := ProjectUsageList{}
	return &this
}

// NewProjectUsageListWithDefaults instantiates a new ProjectUsageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectUsageListWithDefaults() *ProjectUsageList {
	this := ProjectUsageList{}
	return &this
}

// GetUsages returns the Usages field value if set, zero value otherwise.
func (o *ProjectUsageList) GetUsages() []ProjectUsage {
	if o == nil || IsNil(o.Usages) {
		var ret []ProjectUsage
		return ret
	}
	return o.Usages
}

// GetUsagesOk returns a tuple with the Usages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUsageList) GetUsagesOk() ([]ProjectUsage, bool) {
	if o == nil || IsNil(o.Usages) {
		return nil, false
	}
	return o.Usages, true
}

// HasUsages returns a boolean if a field has been set.
func (o *ProjectUsageList) HasUsages() bool {
	if o != nil && !IsNil(o.Usages) {
		return true
	}

	return false
}

// SetUsages gets a reference to the given []ProjectUsage and assigns it to the Usages field.
func (o *ProjectUsageList) SetUsages(v []ProjectUsage) {
	o.Usages = v
}

func (o ProjectUsageList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectUsageList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Usages) {
		toSerialize["usages"] = o.Usages
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectUsageList) UnmarshalJSON(data []byte) (err error) {
	varProjectUsageList := _ProjectUsageList{}

	err = json.Unmarshal(data, &varProjectUsageList)

	if err != nil {
		return err
	}

	*o = ProjectUsageList(varProjectUsageList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "usages")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectUsageList struct {
	value *ProjectUsageList
	isSet bool
}

func (v NullableProjectUsageList) Get() *ProjectUsageList {
	return v.value
}

func (v *NullableProjectUsageList) Set(val *ProjectUsageList) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectUsageList) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectUsageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectUsageList(val *ProjectUsageList) *NullableProjectUsageList {
	return &NullableProjectUsageList{value: val, isSet: true}
}

func (v NullableProjectUsageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectUsageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
