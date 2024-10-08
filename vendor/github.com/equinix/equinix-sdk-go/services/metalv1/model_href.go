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

// checks if the Href type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Href{}

// Href struct for Href
type Href struct {
	Href                 string `json:"href"`
	AdditionalProperties map[string]interface{}
}

type _Href Href

// NewHref instantiates a new Href object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHref(href string) *Href {
	this := Href{}
	this.Href = href
	return &this
}

// NewHrefWithDefaults instantiates a new Href object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHrefWithDefaults() *Href {
	this := Href{}
	return &this
}

// GetHref returns the Href field value
func (o *Href) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *Href) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *Href) SetHref(v string) {
	o.Href = v
}

func (o Href) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Href) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Href) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"href",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHref := _Href{}

	err = json.Unmarshal(data, &varHref)

	if err != nil {
		return err
	}

	*o = Href(varHref)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHref struct {
	value *Href
	isSet bool
}

func (v NullableHref) Get() *Href {
	return v.value
}

func (v *NullableHref) Set(val *Href) {
	v.value = val
	v.isSet = true
}

func (v NullableHref) IsSet() bool {
	return v.isSet
}

func (v *NullableHref) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHref(val *Href) *NullableHref {
	return &NullableHref{value: val, isSet: true}
}

func (v NullableHref) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHref) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
