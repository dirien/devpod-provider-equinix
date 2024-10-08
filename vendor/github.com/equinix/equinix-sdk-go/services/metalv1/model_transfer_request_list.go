/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the TransferRequestList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferRequestList{}

// TransferRequestList struct for TransferRequestList
type TransferRequestList struct {
	Transfers            []TransferRequest `json:"transfers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransferRequestList TransferRequestList

// NewTransferRequestList instantiates a new TransferRequestList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRequestList() *TransferRequestList {
	this := TransferRequestList{}
	return &this
}

// NewTransferRequestListWithDefaults instantiates a new TransferRequestList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRequestListWithDefaults() *TransferRequestList {
	this := TransferRequestList{}
	return &this
}

// GetTransfers returns the Transfers field value if set, zero value otherwise.
func (o *TransferRequestList) GetTransfers() []TransferRequest {
	if o == nil || IsNil(o.Transfers) {
		var ret []TransferRequest
		return ret
	}
	return o.Transfers
}

// GetTransfersOk returns a tuple with the Transfers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRequestList) GetTransfersOk() ([]TransferRequest, bool) {
	if o == nil || IsNil(o.Transfers) {
		return nil, false
	}
	return o.Transfers, true
}

// HasTransfers returns a boolean if a field has been set.
func (o *TransferRequestList) HasTransfers() bool {
	if o != nil && !IsNil(o.Transfers) {
		return true
	}

	return false
}

// SetTransfers gets a reference to the given []TransferRequest and assigns it to the Transfers field.
func (o *TransferRequestList) SetTransfers(v []TransferRequest) {
	o.Transfers = v
}

func (o TransferRequestList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferRequestList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Transfers) {
		toSerialize["transfers"] = o.Transfers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TransferRequestList) UnmarshalJSON(data []byte) (err error) {
	varTransferRequestList := _TransferRequestList{}

	err = json.Unmarshal(data, &varTransferRequestList)

	if err != nil {
		return err
	}

	*o = TransferRequestList(varTransferRequestList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "transfers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferRequestList struct {
	value *TransferRequestList
	isSet bool
}

func (v NullableTransferRequestList) Get() *TransferRequestList {
	return v.value
}

func (v *NullableTransferRequestList) Set(val *TransferRequestList) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRequestList) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRequestList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRequestList(val *TransferRequestList) *NullableTransferRequestList {
	return &NullableTransferRequestList{value: val, isSet: true}
}

func (v NullableTransferRequestList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRequestList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
