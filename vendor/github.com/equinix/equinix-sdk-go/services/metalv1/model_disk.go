/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the Disk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Disk{}

// Disk struct for Disk
type Disk struct {
	Device               *string     `json:"device,omitempty"`
	WipeTable            *bool       `json:"wipeTable,omitempty"`
	Partitions           []Partition `json:"partitions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Disk Disk

// NewDisk instantiates a new Disk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisk() *Disk {
	this := Disk{}
	return &this
}

// NewDiskWithDefaults instantiates a new Disk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskWithDefaults() *Disk {
	this := Disk{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *Disk) GetDevice() string {
	if o == nil || IsNil(o.Device) {
		var ret string
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disk) GetDeviceOk() (*string, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *Disk) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *Disk) SetDevice(v string) {
	o.Device = &v
}

// GetWipeTable returns the WipeTable field value if set, zero value otherwise.
func (o *Disk) GetWipeTable() bool {
	if o == nil || IsNil(o.WipeTable) {
		var ret bool
		return ret
	}
	return *o.WipeTable
}

// GetWipeTableOk returns a tuple with the WipeTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disk) GetWipeTableOk() (*bool, bool) {
	if o == nil || IsNil(o.WipeTable) {
		return nil, false
	}
	return o.WipeTable, true
}

// HasWipeTable returns a boolean if a field has been set.
func (o *Disk) HasWipeTable() bool {
	if o != nil && !IsNil(o.WipeTable) {
		return true
	}

	return false
}

// SetWipeTable gets a reference to the given bool and assigns it to the WipeTable field.
func (o *Disk) SetWipeTable(v bool) {
	o.WipeTable = &v
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *Disk) GetPartitions() []Partition {
	if o == nil || IsNil(o.Partitions) {
		var ret []Partition
		return ret
	}
	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disk) GetPartitionsOk() ([]Partition, bool) {
	if o == nil || IsNil(o.Partitions) {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *Disk) HasPartitions() bool {
	if o != nil && !IsNil(o.Partitions) {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []Partition and assigns it to the Partitions field.
func (o *Disk) SetPartitions(v []Partition) {
	o.Partitions = v
}

func (o Disk) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Disk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.WipeTable) {
		toSerialize["wipeTable"] = o.WipeTable
	}
	if !IsNil(o.Partitions) {
		toSerialize["partitions"] = o.Partitions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Disk) UnmarshalJSON(data []byte) (err error) {
	varDisk := _Disk{}

	err = json.Unmarshal(data, &varDisk)

	if err != nil {
		return err
	}

	*o = Disk(varDisk)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "wipeTable")
		delete(additionalProperties, "partitions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDisk struct {
	value *Disk
	isSet bool
}

func (v NullableDisk) Get() *Disk {
	return v.value
}

func (v *NullableDisk) Set(val *Disk) {
	v.value = val
	v.isSet = true
}

func (v NullableDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisk(val *Disk) *NullableDisk {
	return &NullableDisk{value: val, isSet: true}
}

func (v NullableDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
