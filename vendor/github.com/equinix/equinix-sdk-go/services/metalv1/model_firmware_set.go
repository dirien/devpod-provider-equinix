/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the FirmwareSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirmwareSet{}

// FirmwareSet Represents a Firmware Set
type FirmwareSet struct {
	// Firmware Set UUID
	Uuid string `json:"uuid"`
	// Firmware Set Name
	Name string `json:"name"`
	// Datetime when the block was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Datetime when the block was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Represents a list of attributes
	Attributes []Attribute `json:"attributes,omitempty"`
	// List of components versions
	ComponentFirmware    []Component `json:"component_firmware,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareSet FirmwareSet

// NewFirmwareSet instantiates a new FirmwareSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareSet(uuid string, name string) *FirmwareSet {
	this := FirmwareSet{}
	this.Uuid = uuid
	this.Name = name
	return &this
}

// NewFirmwareSetWithDefaults instantiates a new FirmwareSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareSetWithDefaults() *FirmwareSet {
	this := FirmwareSet{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *FirmwareSet) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *FirmwareSet) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *FirmwareSet) SetUuid(v string) {
	o.Uuid = v
}

// GetName returns the Name field value
func (o *FirmwareSet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FirmwareSet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FirmwareSet) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FirmwareSet) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareSet) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FirmwareSet) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FirmwareSet) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FirmwareSet) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareSet) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FirmwareSet) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *FirmwareSet) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *FirmwareSet) GetAttributes() []Attribute {
	if o == nil || IsNil(o.Attributes) {
		var ret []Attribute
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareSet) GetAttributesOk() ([]Attribute, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *FirmwareSet) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []Attribute and assigns it to the Attributes field.
func (o *FirmwareSet) SetAttributes(v []Attribute) {
	o.Attributes = v
}

// GetComponentFirmware returns the ComponentFirmware field value if set, zero value otherwise.
func (o *FirmwareSet) GetComponentFirmware() []Component {
	if o == nil || IsNil(o.ComponentFirmware) {
		var ret []Component
		return ret
	}
	return o.ComponentFirmware
}

// GetComponentFirmwareOk returns a tuple with the ComponentFirmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareSet) GetComponentFirmwareOk() ([]Component, bool) {
	if o == nil || IsNil(o.ComponentFirmware) {
		return nil, false
	}
	return o.ComponentFirmware, true
}

// HasComponentFirmware returns a boolean if a field has been set.
func (o *FirmwareSet) HasComponentFirmware() bool {
	if o != nil && !IsNil(o.ComponentFirmware) {
		return true
	}

	return false
}

// SetComponentFirmware gets a reference to the given []Component and assigns it to the ComponentFirmware field.
func (o *FirmwareSet) SetComponentFirmware(v []Component) {
	o.ComponentFirmware = v
}

func (o FirmwareSet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirmwareSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uuid"] = o.Uuid
	toSerialize["name"] = o.Name
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.ComponentFirmware) {
		toSerialize["component_firmware"] = o.ComponentFirmware
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FirmwareSet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uuid",
		"name",
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

	varFirmwareSet := _FirmwareSet{}

	err = json.Unmarshal(data, &varFirmwareSet)

	if err != nil {
		return err
	}

	*o = FirmwareSet(varFirmwareSet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "component_firmware")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareSet struct {
	value *FirmwareSet
	isSet bool
}

func (v NullableFirmwareSet) Get() *FirmwareSet {
	return v.value
}

func (v *NullableFirmwareSet) Set(val *FirmwareSet) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareSet) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareSet(val *FirmwareSet) *NullableFirmwareSet {
	return &NullableFirmwareSet{value: val, isSet: true}
}

func (v NullableFirmwareSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
