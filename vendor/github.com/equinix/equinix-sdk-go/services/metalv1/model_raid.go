/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the Raid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Raid{}

// Raid struct for Raid
type Raid struct {
	Devices              []string `json:"devices,omitempty"`
	Level                *string  `json:"level,omitempty"`
	Name                 *string  `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Raid Raid

// NewRaid instantiates a new Raid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRaid() *Raid {
	this := Raid{}
	return &this
}

// NewRaidWithDefaults instantiates a new Raid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRaidWithDefaults() *Raid {
	this := Raid{}
	return &this
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *Raid) GetDevices() []string {
	if o == nil || IsNil(o.Devices) {
		var ret []string
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Raid) GetDevicesOk() ([]string, bool) {
	if o == nil || IsNil(o.Devices) {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *Raid) HasDevices() bool {
	if o != nil && !IsNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []string and assigns it to the Devices field.
func (o *Raid) SetDevices(v []string) {
	o.Devices = v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *Raid) GetLevel() string {
	if o == nil || IsNil(o.Level) {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Raid) GetLevelOk() (*string, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *Raid) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *Raid) SetLevel(v string) {
	o.Level = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Raid) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Raid) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Raid) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Raid) SetName(v string) {
	o.Name = &v
}

func (o Raid) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Raid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Raid) UnmarshalJSON(data []byte) (err error) {
	varRaid := _Raid{}

	err = json.Unmarshal(data, &varRaid)

	if err != nil {
		return err
	}

	*o = Raid(varRaid)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "devices")
		delete(additionalProperties, "level")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRaid struct {
	value *Raid
	isSet bool
}

func (v NullableRaid) Get() *Raid {
	return v.value
}

func (v *NullableRaid) Set(val *Raid) {
	v.value = val
	v.isSet = true
}

func (v NullableRaid) IsSet() bool {
	return v.isSet
}

func (v *NullableRaid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRaid(val *Raid) *NullableRaid {
	return &NullableRaid{value: val, isSet: true}
}

func (v NullableRaid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRaid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
