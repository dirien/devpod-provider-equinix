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

// checks if the PortVlanAssignment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortVlanAssignment{}

// PortVlanAssignment struct for PortVlanAssignment
type PortVlanAssignment struct {
	CreatedAt            *time.Time               `json:"created_at,omitempty"`
	Id                   *string                  `json:"id,omitempty"`
	Native               *bool                    `json:"native,omitempty"`
	Port                 *Href                    `json:"port,omitempty"`
	State                *PortVlanAssignmentState `json:"state,omitempty"`
	UpdatedAt            *time.Time               `json:"updated_at,omitempty"`
	VirtualNetwork       *Href                    `json:"virtual_network,omitempty"`
	Vlan                 *int32                   `json:"vlan,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortVlanAssignment PortVlanAssignment

// NewPortVlanAssignment instantiates a new PortVlanAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortVlanAssignment() *PortVlanAssignment {
	this := PortVlanAssignment{}
	return &this
}

// NewPortVlanAssignmentWithDefaults instantiates a new PortVlanAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortVlanAssignmentWithDefaults() *PortVlanAssignment {
	this := PortVlanAssignment{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PortVlanAssignment) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PortVlanAssignment) SetId(v string) {
	o.Id = &v
}

// GetNative returns the Native field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetNative() bool {
	if o == nil || IsNil(o.Native) {
		var ret bool
		return ret
	}
	return *o.Native
}

// GetNativeOk returns a tuple with the Native field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetNativeOk() (*bool, bool) {
	if o == nil || IsNil(o.Native) {
		return nil, false
	}
	return o.Native, true
}

// HasNative returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasNative() bool {
	if o != nil && !IsNil(o.Native) {
		return true
	}

	return false
}

// SetNative gets a reference to the given bool and assigns it to the Native field.
func (o *PortVlanAssignment) SetNative(v bool) {
	o.Native = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetPort() Href {
	if o == nil || IsNil(o.Port) {
		var ret Href
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetPortOk() (*Href, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given Href and assigns it to the Port field.
func (o *PortVlanAssignment) SetPort(v Href) {
	o.Port = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetState() PortVlanAssignmentState {
	if o == nil || IsNil(o.State) {
		var ret PortVlanAssignmentState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetStateOk() (*PortVlanAssignmentState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given PortVlanAssignmentState and assigns it to the State field.
func (o *PortVlanAssignment) SetState(v PortVlanAssignmentState) {
	o.State = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PortVlanAssignment) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVirtualNetwork returns the VirtualNetwork field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetVirtualNetwork() Href {
	if o == nil || IsNil(o.VirtualNetwork) {
		var ret Href
		return ret
	}
	return *o.VirtualNetwork
}

// GetVirtualNetworkOk returns a tuple with the VirtualNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetVirtualNetworkOk() (*Href, bool) {
	if o == nil || IsNil(o.VirtualNetwork) {
		return nil, false
	}
	return o.VirtualNetwork, true
}

// HasVirtualNetwork returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasVirtualNetwork() bool {
	if o != nil && !IsNil(o.VirtualNetwork) {
		return true
	}

	return false
}

// SetVirtualNetwork gets a reference to the given Href and assigns it to the VirtualNetwork field.
func (o *PortVlanAssignment) SetVirtualNetwork(v Href) {
	o.VirtualNetwork = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *PortVlanAssignment) GetVlan() int32 {
	if o == nil || IsNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignment) GetVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.Vlan) {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *PortVlanAssignment) HasVlan() bool {
	if o != nil && !IsNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *PortVlanAssignment) SetVlan(v int32) {
	o.Vlan = &v
}

func (o PortVlanAssignment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortVlanAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Native) {
		toSerialize["native"] = o.Native
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.VirtualNetwork) {
		toSerialize["virtual_network"] = o.VirtualNetwork
	}
	if !IsNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortVlanAssignment) UnmarshalJSON(data []byte) (err error) {
	varPortVlanAssignment := _PortVlanAssignment{}

	err = json.Unmarshal(data, &varPortVlanAssignment)

	if err != nil {
		return err
	}

	*o = PortVlanAssignment(varPortVlanAssignment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "native")
		delete(additionalProperties, "port")
		delete(additionalProperties, "state")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "virtual_network")
		delete(additionalProperties, "vlan")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortVlanAssignment struct {
	value *PortVlanAssignment
	isSet bool
}

func (v NullablePortVlanAssignment) Get() *PortVlanAssignment {
	return v.value
}

func (v *NullablePortVlanAssignment) Set(val *PortVlanAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullablePortVlanAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullablePortVlanAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortVlanAssignment(val *PortVlanAssignment) *NullablePortVlanAssignment {
	return &NullablePortVlanAssignment{value: val, isSet: true}
}

func (v NullablePortVlanAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortVlanAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
