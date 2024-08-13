/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the DeviceList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceList{}

// DeviceList struct for DeviceList
type DeviceList struct {
	Devices              []Device `json:"devices,omitempty"`
	Meta                 *Meta    `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceList DeviceList

// NewDeviceList instantiates a new DeviceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceList() *DeviceList {
	this := DeviceList{}
	return &this
}

// NewDeviceListWithDefaults instantiates a new DeviceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceListWithDefaults() *DeviceList {
	this := DeviceList{}
	return &this
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *DeviceList) GetDevices() []Device {
	if o == nil || IsNil(o.Devices) {
		var ret []Device
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceList) GetDevicesOk() ([]Device, bool) {
	if o == nil || IsNil(o.Devices) {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *DeviceList) HasDevices() bool {
	if o != nil && !IsNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []Device and assigns it to the Devices field.
func (o *DeviceList) SetDevices(v []Device) {
	o.Devices = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DeviceList) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceList) GetMetaOk() (*Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DeviceList) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *DeviceList) SetMeta(v Meta) {
	o.Meta = &v
}

func (o DeviceList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceList) UnmarshalJSON(data []byte) (err error) {
	varDeviceList := _DeviceList{}

	err = json.Unmarshal(data, &varDeviceList)

	if err != nil {
		return err
	}

	*o = DeviceList(varDeviceList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "devices")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceList struct {
	value *DeviceList
	isSet bool
}

func (v NullableDeviceList) Get() *DeviceList {
	return v.value
}

func (v *NullableDeviceList) Set(val *DeviceList) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceList) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceList(val *DeviceList) *NullableDeviceList {
	return &NullableDeviceList{value: val, isSet: true}
}

func (v NullableDeviceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
