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

// checks if the VrfMetalGateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VrfMetalGateway{}

// VrfMetalGateway struct for VrfMetalGateway
type VrfMetalGateway struct {
	CreatedAt            *time.Time         `json:"created_at,omitempty"`
	CreatedBy            *Href              `json:"created_by,omitempty"`
	Href                 *string            `json:"href,omitempty"`
	Id                   *string            `json:"id,omitempty"`
	IpReservation        *VrfIpReservation  `json:"ip_reservation,omitempty"`
	Project              *Project           `json:"project,omitempty"`
	State                *MetalGatewayState `json:"state,omitempty"`
	UpdatedAt            *time.Time         `json:"updated_at,omitempty"`
	VirtualNetwork       *VirtualNetwork    `json:"virtual_network,omitempty"`
	Vrf                  *Vrf               `json:"vrf,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VrfMetalGateway VrfMetalGateway

// NewVrfMetalGateway instantiates a new VrfMetalGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVrfMetalGateway() *VrfMetalGateway {
	this := VrfMetalGateway{}
	return &this
}

// NewVrfMetalGatewayWithDefaults instantiates a new VrfMetalGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVrfMetalGatewayWithDefaults() *VrfMetalGateway {
	this := VrfMetalGateway{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VrfMetalGateway) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfMetalGateway) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VrfMetalGateway) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VrfMetalGateway) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *VrfMetalGateway) GetCreatedBy() Href {
	if o == nil || IsNil(o.CreatedBy) {
		var ret Href
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfMetalGateway) GetCreatedByOk() (*Href, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *VrfMetalGateway) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given Href and assigns it to the CreatedBy field.
func (o *VrfMetalGateway) SetCreatedBy(v Href) {
	o.CreatedBy = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *VrfMetalGateway) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfMetalGateway) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *VrfMetalGateway) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *VrfMetalGateway) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VrfMetalGateway) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfMetalGateway) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VrfMetalGateway) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VrfMetalGateway) SetId(v string) {
	o.Id = &v
}

// GetIpReservation returns the IpReservation field value if set, zero value otherwise.
func (o *VrfMetalGateway) GetIpReservation() VrfIpReservation {
	if o == nil || IsNil(o.IpReservation) {
		var ret VrfIpReservation
		return ret
	}
	return *o.IpReservation
}

// GetIpReservationOk returns a tuple with the IpReservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfMetalGateway) GetIpReservationOk() (*VrfIpReservation, bool) {
	if o == nil || IsNil(o.IpReservation) {
		return nil, false
	}
	return o.IpReservation, true
}

// HasIpReservation returns a boolean if a field has been set.
func (o *VrfMetalGateway) HasIpReservation() bool {
	if o != nil && !IsNil(o.IpReservation) {
		return true
	}

	return false
}

// SetIpReservation gets a reference to the given VrfIpReservation and assigns it to the IpReservation field.
func (o *VrfMetalGateway) SetIpReservation(v VrfIpReservation) {
	o.IpReservation = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *VrfMetalGateway) GetProject() Project {
	if o == nil || IsNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfMetalGateway) GetProjectOk() (*Project, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *VrfMetalGateway) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *VrfMetalGateway) SetProject(v Project) {
	o.Project = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *VrfMetalGateway) GetState() MetalGatewayState {
	if o == nil || IsNil(o.State) {
		var ret MetalGatewayState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfMetalGateway) GetStateOk() (*MetalGatewayState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *VrfMetalGateway) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given MetalGatewayState and assigns it to the State field.
func (o *VrfMetalGateway) SetState(v MetalGatewayState) {
	o.State = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *VrfMetalGateway) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfMetalGateway) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *VrfMetalGateway) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *VrfMetalGateway) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVirtualNetwork returns the VirtualNetwork field value if set, zero value otherwise.
func (o *VrfMetalGateway) GetVirtualNetwork() VirtualNetwork {
	if o == nil || IsNil(o.VirtualNetwork) {
		var ret VirtualNetwork
		return ret
	}
	return *o.VirtualNetwork
}

// GetVirtualNetworkOk returns a tuple with the VirtualNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfMetalGateway) GetVirtualNetworkOk() (*VirtualNetwork, bool) {
	if o == nil || IsNil(o.VirtualNetwork) {
		return nil, false
	}
	return o.VirtualNetwork, true
}

// HasVirtualNetwork returns a boolean if a field has been set.
func (o *VrfMetalGateway) HasVirtualNetwork() bool {
	if o != nil && !IsNil(o.VirtualNetwork) {
		return true
	}

	return false
}

// SetVirtualNetwork gets a reference to the given VirtualNetwork and assigns it to the VirtualNetwork field.
func (o *VrfMetalGateway) SetVirtualNetwork(v VirtualNetwork) {
	o.VirtualNetwork = &v
}

// GetVrf returns the Vrf field value if set, zero value otherwise.
func (o *VrfMetalGateway) GetVrf() Vrf {
	if o == nil || IsNil(o.Vrf) {
		var ret Vrf
		return ret
	}
	return *o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfMetalGateway) GetVrfOk() (*Vrf, bool) {
	if o == nil || IsNil(o.Vrf) {
		return nil, false
	}
	return o.Vrf, true
}

// HasVrf returns a boolean if a field has been set.
func (o *VrfMetalGateway) HasVrf() bool {
	if o != nil && !IsNil(o.Vrf) {
		return true
	}

	return false
}

// SetVrf gets a reference to the given Vrf and assigns it to the Vrf field.
func (o *VrfMetalGateway) SetVrf(v Vrf) {
	o.Vrf = &v
}

func (o VrfMetalGateway) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VrfMetalGateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IpReservation) {
		toSerialize["ip_reservation"] = o.IpReservation
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
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
	if !IsNil(o.Vrf) {
		toSerialize["vrf"] = o.Vrf
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VrfMetalGateway) UnmarshalJSON(data []byte) (err error) {
	varVrfMetalGateway := _VrfMetalGateway{}

	err = json.Unmarshal(data, &varVrfMetalGateway)

	if err != nil {
		return err
	}

	*o = VrfMetalGateway(varVrfMetalGateway)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "href")
		delete(additionalProperties, "id")
		delete(additionalProperties, "ip_reservation")
		delete(additionalProperties, "project")
		delete(additionalProperties, "state")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "virtual_network")
		delete(additionalProperties, "vrf")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVrfMetalGateway struct {
	value *VrfMetalGateway
	isSet bool
}

func (v NullableVrfMetalGateway) Get() *VrfMetalGateway {
	return v.value
}

func (v *NullableVrfMetalGateway) Set(val *VrfMetalGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableVrfMetalGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableVrfMetalGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrfMetalGateway(val *VrfMetalGateway) *NullableVrfMetalGateway {
	return &NullableVrfMetalGateway{value: val, isSet: true}
}

func (v NullableVrfMetalGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrfMetalGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
