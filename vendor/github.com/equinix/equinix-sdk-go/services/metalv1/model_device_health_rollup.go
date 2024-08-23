/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
	"time"
)

// checks if the DeviceHealthRollup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceHealthRollup{}

// DeviceHealthRollup Represents a Device Health Status
type DeviceHealthRollup struct {
	HealthRollup *DeviceHealthRollupHealthRollup `json:"health_rollup,omitempty"`
	// Last update of health status.
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceHealthRollup DeviceHealthRollup

// NewDeviceHealthRollup instantiates a new DeviceHealthRollup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceHealthRollup() *DeviceHealthRollup {
	this := DeviceHealthRollup{}
	return &this
}

// NewDeviceHealthRollupWithDefaults instantiates a new DeviceHealthRollup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceHealthRollupWithDefaults() *DeviceHealthRollup {
	this := DeviceHealthRollup{}
	return &this
}

// GetHealthRollup returns the HealthRollup field value if set, zero value otherwise.
func (o *DeviceHealthRollup) GetHealthRollup() DeviceHealthRollupHealthRollup {
	if o == nil || IsNil(o.HealthRollup) {
		var ret DeviceHealthRollupHealthRollup
		return ret
	}
	return *o.HealthRollup
}

// GetHealthRollupOk returns a tuple with the HealthRollup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthRollup) GetHealthRollupOk() (*DeviceHealthRollupHealthRollup, bool) {
	if o == nil || IsNil(o.HealthRollup) {
		return nil, false
	}
	return o.HealthRollup, true
}

// HasHealthRollup returns a boolean if a field has been set.
func (o *DeviceHealthRollup) HasHealthRollup() bool {
	if o != nil && !IsNil(o.HealthRollup) {
		return true
	}

	return false
}

// SetHealthRollup gets a reference to the given DeviceHealthRollupHealthRollup and assigns it to the HealthRollup field.
func (o *DeviceHealthRollup) SetHealthRollup(v DeviceHealthRollupHealthRollup) {
	o.HealthRollup = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DeviceHealthRollup) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceHealthRollup) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DeviceHealthRollup) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DeviceHealthRollup) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o DeviceHealthRollup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceHealthRollup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HealthRollup) {
		toSerialize["health_rollup"] = o.HealthRollup
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceHealthRollup) UnmarshalJSON(data []byte) (err error) {
	varDeviceHealthRollup := _DeviceHealthRollup{}

	err = json.Unmarshal(data, &varDeviceHealthRollup)

	if err != nil {
		return err
	}

	*o = DeviceHealthRollup(varDeviceHealthRollup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "health_rollup")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceHealthRollup struct {
	value *DeviceHealthRollup
	isSet bool
}

func (v NullableDeviceHealthRollup) Get() *DeviceHealthRollup {
	return v.value
}

func (v *NullableDeviceHealthRollup) Set(val *DeviceHealthRollup) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceHealthRollup) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceHealthRollup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceHealthRollup(val *DeviceHealthRollup) *NullableDeviceHealthRollup {
	return &NullableDeviceHealthRollup{value: val, isSet: true}
}

func (v NullableDeviceHealthRollup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceHealthRollup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
