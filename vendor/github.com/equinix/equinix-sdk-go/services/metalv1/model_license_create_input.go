/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the LicenseCreateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseCreateInput{}

// LicenseCreateInput struct for LicenseCreateInput
type LicenseCreateInput struct {
	Description          *string  `json:"description,omitempty"`
	LicenseeProductId    *string  `json:"licensee_product_id,omitempty"`
	Size                 *float32 `json:"size,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LicenseCreateInput LicenseCreateInput

// NewLicenseCreateInput instantiates a new LicenseCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseCreateInput() *LicenseCreateInput {
	this := LicenseCreateInput{}
	return &this
}

// NewLicenseCreateInputWithDefaults instantiates a new LicenseCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseCreateInputWithDefaults() *LicenseCreateInput {
	this := LicenseCreateInput{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LicenseCreateInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseCreateInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LicenseCreateInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LicenseCreateInput) SetDescription(v string) {
	o.Description = &v
}

// GetLicenseeProductId returns the LicenseeProductId field value if set, zero value otherwise.
func (o *LicenseCreateInput) GetLicenseeProductId() string {
	if o == nil || IsNil(o.LicenseeProductId) {
		var ret string
		return ret
	}
	return *o.LicenseeProductId
}

// GetLicenseeProductIdOk returns a tuple with the LicenseeProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseCreateInput) GetLicenseeProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseeProductId) {
		return nil, false
	}
	return o.LicenseeProductId, true
}

// HasLicenseeProductId returns a boolean if a field has been set.
func (o *LicenseCreateInput) HasLicenseeProductId() bool {
	if o != nil && !IsNil(o.LicenseeProductId) {
		return true
	}

	return false
}

// SetLicenseeProductId gets a reference to the given string and assigns it to the LicenseeProductId field.
func (o *LicenseCreateInput) SetLicenseeProductId(v string) {
	o.LicenseeProductId = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *LicenseCreateInput) GetSize() float32 {
	if o == nil || IsNil(o.Size) {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseCreateInput) GetSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *LicenseCreateInput) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *LicenseCreateInput) SetSize(v float32) {
	o.Size = &v
}

func (o LicenseCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseCreateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.LicenseeProductId) {
		toSerialize["licensee_product_id"] = o.LicenseeProductId
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LicenseCreateInput) UnmarshalJSON(data []byte) (err error) {
	varLicenseCreateInput := _LicenseCreateInput{}

	err = json.Unmarshal(data, &varLicenseCreateInput)

	if err != nil {
		return err
	}

	*o = LicenseCreateInput(varLicenseCreateInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "licensee_product_id")
		delete(additionalProperties, "size")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLicenseCreateInput struct {
	value *LicenseCreateInput
	isSet bool
}

func (v NullableLicenseCreateInput) Get() *LicenseCreateInput {
	return v.value
}

func (v *NullableLicenseCreateInput) Set(val *LicenseCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseCreateInput(val *LicenseCreateInput) *NullableLicenseCreateInput {
	return &NullableLicenseCreateInput{value: val, isSet: true}
}

func (v NullableLicenseCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
