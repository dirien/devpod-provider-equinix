/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the Metadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Metadata{}

// Metadata struct for Metadata
type Metadata struct {
	Class      *string                `json:"class,omitempty"`
	Customdata map[string]interface{} `json:"customdata,omitempty"`
	// The facility code of the instance
	Facility *string `json:"facility,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	Id       *string `json:"id,omitempty"`
	Iqn      *string `json:"iqn,omitempty"`
	// The metro code of the instance
	Metro           *string                `json:"metro,omitempty"`
	Network         *MetadataNetwork       `json:"network,omitempty"`
	OperatingSystem map[string]interface{} `json:"operating_system,omitempty"`
	// The plan slug of the instance
	Plan *string `json:"plan,omitempty"`
	// An array of the private subnets
	PrivateSubnets []string `json:"private_subnets,omitempty"`
	Reserved       *bool    `json:"reserved,omitempty"`
	// The specs of the plan version of the instance
	Specs                map[string]interface{} `json:"specs,omitempty"`
	SshKeys              []string               `json:"ssh_keys,omitempty"`
	SwitchShortId        *string                `json:"switch_short_id,omitempty"`
	State                *DeviceState           `json:"state,omitempty"`
	Tags                 []string               `json:"tags,omitempty"`
	Volumes              []string               `json:"volumes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Metadata Metadata

// NewMetadata instantiates a new Metadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadata() *Metadata {
	this := Metadata{}
	return &this
}

// NewMetadataWithDefaults instantiates a new Metadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataWithDefaults() *Metadata {
	this := Metadata{}
	return &this
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *Metadata) GetClass() string {
	if o == nil || IsNil(o.Class) {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetClassOk() (*string, bool) {
	if o == nil || IsNil(o.Class) {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *Metadata) HasClass() bool {
	if o != nil && !IsNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *Metadata) SetClass(v string) {
	o.Class = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *Metadata) GetCustomdata() map[string]interface{} {
	if o == nil || IsNil(o.Customdata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetCustomdataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Customdata) {
		return map[string]interface{}{}, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *Metadata) HasCustomdata() bool {
	if o != nil && !IsNil(o.Customdata) {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *Metadata) SetCustomdata(v map[string]interface{}) {
	o.Customdata = v
}

// GetFacility returns the Facility field value if set, zero value otherwise.
func (o *Metadata) GetFacility() string {
	if o == nil || IsNil(o.Facility) {
		var ret string
		return ret
	}
	return *o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetFacilityOk() (*string, bool) {
	if o == nil || IsNil(o.Facility) {
		return nil, false
	}
	return o.Facility, true
}

// HasFacility returns a boolean if a field has been set.
func (o *Metadata) HasFacility() bool {
	if o != nil && !IsNil(o.Facility) {
		return true
	}

	return false
}

// SetFacility gets a reference to the given string and assigns it to the Facility field.
func (o *Metadata) SetFacility(v string) {
	o.Facility = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *Metadata) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *Metadata) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *Metadata) SetHostname(v string) {
	o.Hostname = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Metadata) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Metadata) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Metadata) SetId(v string) {
	o.Id = &v
}

// GetIqn returns the Iqn field value if set, zero value otherwise.
func (o *Metadata) GetIqn() string {
	if o == nil || IsNil(o.Iqn) {
		var ret string
		return ret
	}
	return *o.Iqn
}

// GetIqnOk returns a tuple with the Iqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetIqnOk() (*string, bool) {
	if o == nil || IsNil(o.Iqn) {
		return nil, false
	}
	return o.Iqn, true
}

// HasIqn returns a boolean if a field has been set.
func (o *Metadata) HasIqn() bool {
	if o != nil && !IsNil(o.Iqn) {
		return true
	}

	return false
}

// SetIqn gets a reference to the given string and assigns it to the Iqn field.
func (o *Metadata) SetIqn(v string) {
	o.Iqn = &v
}

// GetMetro returns the Metro field value if set, zero value otherwise.
func (o *Metadata) GetMetro() string {
	if o == nil || IsNil(o.Metro) {
		var ret string
		return ret
	}
	return *o.Metro
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetMetroOk() (*string, bool) {
	if o == nil || IsNil(o.Metro) {
		return nil, false
	}
	return o.Metro, true
}

// HasMetro returns a boolean if a field has been set.
func (o *Metadata) HasMetro() bool {
	if o != nil && !IsNil(o.Metro) {
		return true
	}

	return false
}

// SetMetro gets a reference to the given string and assigns it to the Metro field.
func (o *Metadata) SetMetro(v string) {
	o.Metro = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *Metadata) GetNetwork() MetadataNetwork {
	if o == nil || IsNil(o.Network) {
		var ret MetadataNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetNetworkOk() (*MetadataNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *Metadata) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given MetadataNetwork and assigns it to the Network field.
func (o *Metadata) SetNetwork(v MetadataNetwork) {
	o.Network = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *Metadata) GetOperatingSystem() map[string]interface{} {
	if o == nil || IsNil(o.OperatingSystem) {
		var ret map[string]interface{}
		return ret
	}
	return o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetOperatingSystemOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.OperatingSystem) {
		return map[string]interface{}{}, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *Metadata) HasOperatingSystem() bool {
	if o != nil && !IsNil(o.OperatingSystem) {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given map[string]interface{} and assigns it to the OperatingSystem field.
func (o *Metadata) SetOperatingSystem(v map[string]interface{}) {
	o.OperatingSystem = v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *Metadata) GetPlan() string {
	if o == nil || IsNil(o.Plan) {
		var ret string
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetPlanOk() (*string, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *Metadata) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given string and assigns it to the Plan field.
func (o *Metadata) SetPlan(v string) {
	o.Plan = &v
}

// GetPrivateSubnets returns the PrivateSubnets field value if set, zero value otherwise.
func (o *Metadata) GetPrivateSubnets() []string {
	if o == nil || IsNil(o.PrivateSubnets) {
		var ret []string
		return ret
	}
	return o.PrivateSubnets
}

// GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetPrivateSubnetsOk() ([]string, bool) {
	if o == nil || IsNil(o.PrivateSubnets) {
		return nil, false
	}
	return o.PrivateSubnets, true
}

// HasPrivateSubnets returns a boolean if a field has been set.
func (o *Metadata) HasPrivateSubnets() bool {
	if o != nil && !IsNil(o.PrivateSubnets) {
		return true
	}

	return false
}

// SetPrivateSubnets gets a reference to the given []string and assigns it to the PrivateSubnets field.
func (o *Metadata) SetPrivateSubnets(v []string) {
	o.PrivateSubnets = v
}

// GetReserved returns the Reserved field value if set, zero value otherwise.
func (o *Metadata) GetReserved() bool {
	if o == nil || IsNil(o.Reserved) {
		var ret bool
		return ret
	}
	return *o.Reserved
}

// GetReservedOk returns a tuple with the Reserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetReservedOk() (*bool, bool) {
	if o == nil || IsNil(o.Reserved) {
		return nil, false
	}
	return o.Reserved, true
}

// HasReserved returns a boolean if a field has been set.
func (o *Metadata) HasReserved() bool {
	if o != nil && !IsNil(o.Reserved) {
		return true
	}

	return false
}

// SetReserved gets a reference to the given bool and assigns it to the Reserved field.
func (o *Metadata) SetReserved(v bool) {
	o.Reserved = &v
}

// GetSpecs returns the Specs field value if set, zero value otherwise.
func (o *Metadata) GetSpecs() map[string]interface{} {
	if o == nil || IsNil(o.Specs) {
		var ret map[string]interface{}
		return ret
	}
	return o.Specs
}

// GetSpecsOk returns a tuple with the Specs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetSpecsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Specs) {
		return map[string]interface{}{}, false
	}
	return o.Specs, true
}

// HasSpecs returns a boolean if a field has been set.
func (o *Metadata) HasSpecs() bool {
	if o != nil && !IsNil(o.Specs) {
		return true
	}

	return false
}

// SetSpecs gets a reference to the given map[string]interface{} and assigns it to the Specs field.
func (o *Metadata) SetSpecs(v map[string]interface{}) {
	o.Specs = v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *Metadata) GetSshKeys() []string {
	if o == nil || IsNil(o.SshKeys) {
		var ret []string
		return ret
	}
	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetSshKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.SshKeys) {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *Metadata) HasSshKeys() bool {
	if o != nil && !IsNil(o.SshKeys) {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []string and assigns it to the SshKeys field.
func (o *Metadata) SetSshKeys(v []string) {
	o.SshKeys = v
}

// GetSwitchShortId returns the SwitchShortId field value if set, zero value otherwise.
func (o *Metadata) GetSwitchShortId() string {
	if o == nil || IsNil(o.SwitchShortId) {
		var ret string
		return ret
	}
	return *o.SwitchShortId
}

// GetSwitchShortIdOk returns a tuple with the SwitchShortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetSwitchShortIdOk() (*string, bool) {
	if o == nil || IsNil(o.SwitchShortId) {
		return nil, false
	}
	return o.SwitchShortId, true
}

// HasSwitchShortId returns a boolean if a field has been set.
func (o *Metadata) HasSwitchShortId() bool {
	if o != nil && !IsNil(o.SwitchShortId) {
		return true
	}

	return false
}

// SetSwitchShortId gets a reference to the given string and assigns it to the SwitchShortId field.
func (o *Metadata) SetSwitchShortId(v string) {
	o.SwitchShortId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Metadata) GetState() DeviceState {
	if o == nil || IsNil(o.State) {
		var ret DeviceState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetStateOk() (*DeviceState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Metadata) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given DeviceState and assigns it to the State field.
func (o *Metadata) SetState(v DeviceState) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Metadata) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Metadata) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Metadata) SetTags(v []string) {
	o.Tags = v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *Metadata) GetVolumes() []string {
	if o == nil || IsNil(o.Volumes) {
		var ret []string
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetVolumesOk() ([]string, bool) {
	if o == nil || IsNil(o.Volumes) {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *Metadata) HasVolumes() bool {
	if o != nil && !IsNil(o.Volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []string and assigns it to the Volumes field.
func (o *Metadata) SetVolumes(v []string) {
	o.Volumes = v
}

func (o Metadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Metadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Class) {
		toSerialize["class"] = o.Class
	}
	if !IsNil(o.Customdata) {
		toSerialize["customdata"] = o.Customdata
	}
	if !IsNil(o.Facility) {
		toSerialize["facility"] = o.Facility
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Iqn) {
		toSerialize["iqn"] = o.Iqn
	}
	if !IsNil(o.Metro) {
		toSerialize["metro"] = o.Metro
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.OperatingSystem) {
		toSerialize["operating_system"] = o.OperatingSystem
	}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !IsNil(o.PrivateSubnets) {
		toSerialize["private_subnets"] = o.PrivateSubnets
	}
	if !IsNil(o.Reserved) {
		toSerialize["reserved"] = o.Reserved
	}
	if !IsNil(o.Specs) {
		toSerialize["specs"] = o.Specs
	}
	if !IsNil(o.SshKeys) {
		toSerialize["ssh_keys"] = o.SshKeys
	}
	if !IsNil(o.SwitchShortId) {
		toSerialize["switch_short_id"] = o.SwitchShortId
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Volumes) {
		toSerialize["volumes"] = o.Volumes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Metadata) UnmarshalJSON(data []byte) (err error) {
	varMetadata := _Metadata{}

	err = json.Unmarshal(data, &varMetadata)

	if err != nil {
		return err
	}

	*o = Metadata(varMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "class")
		delete(additionalProperties, "customdata")
		delete(additionalProperties, "facility")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "id")
		delete(additionalProperties, "iqn")
		delete(additionalProperties, "metro")
		delete(additionalProperties, "network")
		delete(additionalProperties, "operating_system")
		delete(additionalProperties, "plan")
		delete(additionalProperties, "private_subnets")
		delete(additionalProperties, "reserved")
		delete(additionalProperties, "specs")
		delete(additionalProperties, "ssh_keys")
		delete(additionalProperties, "switch_short_id")
		delete(additionalProperties, "state")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "volumes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMetadata struct {
	value *Metadata
	isSet bool
}

func (v NullableMetadata) Get() *Metadata {
	return v.value
}

func (v *NullableMetadata) Set(val *Metadata) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadata(val *Metadata) *NullableMetadata {
	return &NullableMetadata{value: val, isSet: true}
}

func (v NullableMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
