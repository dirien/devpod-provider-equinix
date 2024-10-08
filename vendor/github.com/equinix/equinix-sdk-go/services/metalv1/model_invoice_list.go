/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the InvoiceList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceList{}

// InvoiceList struct for InvoiceList
type InvoiceList struct {
	Invoices             []Invoice `json:"invoices,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InvoiceList InvoiceList

// NewInvoiceList instantiates a new InvoiceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceList() *InvoiceList {
	this := InvoiceList{}
	return &this
}

// NewInvoiceListWithDefaults instantiates a new InvoiceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceListWithDefaults() *InvoiceList {
	this := InvoiceList{}
	return &this
}

// GetInvoices returns the Invoices field value if set, zero value otherwise.
func (o *InvoiceList) GetInvoices() []Invoice {
	if o == nil || IsNil(o.Invoices) {
		var ret []Invoice
		return ret
	}
	return o.Invoices
}

// GetInvoicesOk returns a tuple with the Invoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceList) GetInvoicesOk() ([]Invoice, bool) {
	if o == nil || IsNil(o.Invoices) {
		return nil, false
	}
	return o.Invoices, true
}

// HasInvoices returns a boolean if a field has been set.
func (o *InvoiceList) HasInvoices() bool {
	if o != nil && !IsNil(o.Invoices) {
		return true
	}

	return false
}

// SetInvoices gets a reference to the given []Invoice and assigns it to the Invoices field.
func (o *InvoiceList) SetInvoices(v []Invoice) {
	o.Invoices = v
}

func (o InvoiceList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Invoices) {
		toSerialize["invoices"] = o.Invoices
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InvoiceList) UnmarshalJSON(data []byte) (err error) {
	varInvoiceList := _InvoiceList{}

	err = json.Unmarshal(data, &varInvoiceList)

	if err != nil {
		return err
	}

	*o = InvoiceList(varInvoiceList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "invoices")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvoiceList struct {
	value *InvoiceList
	isSet bool
}

func (v NullableInvoiceList) Get() *InvoiceList {
	return v.value
}

func (v *NullableInvoiceList) Set(val *InvoiceList) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceList) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceList(val *InvoiceList) *NullableInvoiceList {
	return &NullableInvoiceList{value: val, isSet: true}
}

func (v NullableInvoiceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
