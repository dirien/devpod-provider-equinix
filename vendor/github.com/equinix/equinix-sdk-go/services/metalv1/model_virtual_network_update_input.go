/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the VirtualNetworkUpdateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualNetworkUpdateInput{}

// VirtualNetworkUpdateInput struct for VirtualNetworkUpdateInput
type VirtualNetworkUpdateInput struct {
	Description          *string  `json:"description,omitempty"`
	Tags                 []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualNetworkUpdateInput VirtualNetworkUpdateInput

// NewVirtualNetworkUpdateInput instantiates a new VirtualNetworkUpdateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualNetworkUpdateInput() *VirtualNetworkUpdateInput {
	this := VirtualNetworkUpdateInput{}
	return &this
}

// NewVirtualNetworkUpdateInputWithDefaults instantiates a new VirtualNetworkUpdateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualNetworkUpdateInputWithDefaults() *VirtualNetworkUpdateInput {
	this := VirtualNetworkUpdateInput{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VirtualNetworkUpdateInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNetworkUpdateInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VirtualNetworkUpdateInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VirtualNetworkUpdateInput) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VirtualNetworkUpdateInput) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNetworkUpdateInput) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VirtualNetworkUpdateInput) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *VirtualNetworkUpdateInput) SetTags(v []string) {
	o.Tags = v
}

func (o VirtualNetworkUpdateInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualNetworkUpdateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualNetworkUpdateInput) UnmarshalJSON(data []byte) (err error) {
	varVirtualNetworkUpdateInput := _VirtualNetworkUpdateInput{}

	err = json.Unmarshal(data, &varVirtualNetworkUpdateInput)

	if err != nil {
		return err
	}

	*o = VirtualNetworkUpdateInput(varVirtualNetworkUpdateInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualNetworkUpdateInput struct {
	value *VirtualNetworkUpdateInput
	isSet bool
}

func (v NullableVirtualNetworkUpdateInput) Get() *VirtualNetworkUpdateInput {
	return v.value
}

func (v *NullableVirtualNetworkUpdateInput) Set(val *VirtualNetworkUpdateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualNetworkUpdateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualNetworkUpdateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualNetworkUpdateInput(val *VirtualNetworkUpdateInput) *NullableVirtualNetworkUpdateInput {
	return &NullableVirtualNetworkUpdateInput{value: val, isSet: true}
}

func (v NullableVirtualNetworkUpdateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualNetworkUpdateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
