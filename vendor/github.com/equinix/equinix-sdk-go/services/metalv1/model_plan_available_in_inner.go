/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the PlanAvailableInInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanAvailableInInner{}

// PlanAvailableInInner struct for PlanAvailableInInner
type PlanAvailableInInner struct {
	// href to the Facility
	Href                 *string                    `json:"href,omitempty"`
	Price                *PlanAvailableInInnerPrice `json:"price,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PlanAvailableInInner PlanAvailableInInner

// NewPlanAvailableInInner instantiates a new PlanAvailableInInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanAvailableInInner() *PlanAvailableInInner {
	this := PlanAvailableInInner{}
	return &this
}

// NewPlanAvailableInInnerWithDefaults instantiates a new PlanAvailableInInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanAvailableInInnerWithDefaults() *PlanAvailableInInner {
	this := PlanAvailableInInner{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *PlanAvailableInInner) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanAvailableInInner) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *PlanAvailableInInner) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *PlanAvailableInInner) SetHref(v string) {
	o.Href = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PlanAvailableInInner) GetPrice() PlanAvailableInInnerPrice {
	if o == nil || IsNil(o.Price) {
		var ret PlanAvailableInInnerPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanAvailableInInner) GetPriceOk() (*PlanAvailableInInnerPrice, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PlanAvailableInInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given PlanAvailableInInnerPrice and assigns it to the Price field.
func (o *PlanAvailableInInner) SetPrice(v PlanAvailableInInnerPrice) {
	o.Price = &v
}

func (o PlanAvailableInInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlanAvailableInInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PlanAvailableInInner) UnmarshalJSON(data []byte) (err error) {
	varPlanAvailableInInner := _PlanAvailableInInner{}

	err = json.Unmarshal(data, &varPlanAvailableInInner)

	if err != nil {
		return err
	}

	*o = PlanAvailableInInner(varPlanAvailableInInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "price")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlanAvailableInInner struct {
	value *PlanAvailableInInner
	isSet bool
}

func (v NullablePlanAvailableInInner) Get() *PlanAvailableInInner {
	return v.value
}

func (v *NullablePlanAvailableInInner) Set(val *PlanAvailableInInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanAvailableInInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanAvailableInInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanAvailableInInner(val *PlanAvailableInInner) *NullablePlanAvailableInInner {
	return &NullablePlanAvailableInInner{value: val, isSet: true}
}

func (v NullablePlanAvailableInInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanAvailableInInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
