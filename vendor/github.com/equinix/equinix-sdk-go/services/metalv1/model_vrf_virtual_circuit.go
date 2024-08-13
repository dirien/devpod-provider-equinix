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

// checks if the VrfVirtualCircuit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VrfVirtualCircuit{}

// VrfVirtualCircuit struct for VrfVirtualCircuit
type VrfVirtualCircuit struct {
	// An IPv4 address from the subnet that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Metal IP. By default, the last usable IP address in the subnet will be used.
	CustomerIp *string `json:"customer_ip,omitempty"`
	// An IPv6 address from the subnet IPv6 that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet IPv6 as the Metal IPv6. By default, the last usable IP address in the subnet IPv6 will be used.
	CustomerIpv6 *string `json:"customer_ipv6,omitempty"`
	Description  *string `json:"description,omitempty"`
	Id           *string `json:"id,omitempty"`
	// The MD5 password for the BGP peering in plaintext (not a checksum).
	Md5 *string `json:"md5,omitempty"`
	// An IPv4 address from the subnet that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Customer IP. By default, the first usable IP address in the subnet will be used.
	MetalIp *string `json:"metal_ip,omitempty"`
	// An IPv6 address from the subnet IPv6 that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IPv6 address in the subnet IPv6 as the Customer IP. By default, the first usable IPv6 address in the subnet IPv6 will be used.
	MetalIpv6 *string              `json:"metal_ipv6,omitempty"`
	Name      *string              `json:"name,omitempty"`
	Port      *InterconnectionPort `json:"port,omitempty"`
	NniVlan   *int32               `json:"nni_vlan,omitempty"`
	// The peer ASN that will be used with the VRF on the Virtual Circuit.
	PeerAsn *int64   `json:"peer_asn,omitempty"`
	Project *Project `json:"project,omitempty"`
	// integer representing bps speed
	Speed  *int64                   `json:"speed,omitempty"`
	Status *VrfVirtualCircuitStatus `json:"status,omitempty"`
	// The /30 or /31 IPv4 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IP and Customer IP must be IPs from this subnet. For /30 subnets, the network and broadcast IPs cannot be used as the Metal or Customer IP.
	Subnet *string `json:"subnet,omitempty"`
	// The /126 or /127 IPv6 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IPv6 and Customer IPv6 must be IPs from this subnet. For /126 subnets, the network and broadcast IPs cannot be used as the Metal IPv6 or Customer IPv6. The subnet specified must be contained within an already-defined IP Range for the VRF.
	SubnetIpv6           *string               `json:"subnet_ipv6,omitempty"`
	Tags                 []string              `json:"tags,omitempty"`
	Type                 *VrfIpReservationType `json:"type,omitempty"`
	Vrf                  Vrf                   `json:"vrf"`
	CreatedAt            *time.Time            `json:"created_at,omitempty"`
	UpdatedAt            *time.Time            `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VrfVirtualCircuit VrfVirtualCircuit

// NewVrfVirtualCircuit instantiates a new VrfVirtualCircuit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVrfVirtualCircuit(vrf Vrf) *VrfVirtualCircuit {
	this := VrfVirtualCircuit{}
	this.Vrf = vrf
	return &this
}

// NewVrfVirtualCircuitWithDefaults instantiates a new VrfVirtualCircuit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVrfVirtualCircuitWithDefaults() *VrfVirtualCircuit {
	this := VrfVirtualCircuit{}
	return &this
}

// GetCustomerIp returns the CustomerIp field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetCustomerIp() string {
	if o == nil || IsNil(o.CustomerIp) {
		var ret string
		return ret
	}
	return *o.CustomerIp
}

// GetCustomerIpOk returns a tuple with the CustomerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetCustomerIpOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerIp) {
		return nil, false
	}
	return o.CustomerIp, true
}

// HasCustomerIp returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasCustomerIp() bool {
	if o != nil && !IsNil(o.CustomerIp) {
		return true
	}

	return false
}

// SetCustomerIp gets a reference to the given string and assigns it to the CustomerIp field.
func (o *VrfVirtualCircuit) SetCustomerIp(v string) {
	o.CustomerIp = &v
}

// GetCustomerIpv6 returns the CustomerIpv6 field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetCustomerIpv6() string {
	if o == nil || IsNil(o.CustomerIpv6) {
		var ret string
		return ret
	}
	return *o.CustomerIpv6
}

// GetCustomerIpv6Ok returns a tuple with the CustomerIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetCustomerIpv6Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomerIpv6) {
		return nil, false
	}
	return o.CustomerIpv6, true
}

// HasCustomerIpv6 returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasCustomerIpv6() bool {
	if o != nil && !IsNil(o.CustomerIpv6) {
		return true
	}

	return false
}

// SetCustomerIpv6 gets a reference to the given string and assigns it to the CustomerIpv6 field.
func (o *VrfVirtualCircuit) SetCustomerIpv6(v string) {
	o.CustomerIpv6 = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VrfVirtualCircuit) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VrfVirtualCircuit) SetId(v string) {
	o.Id = &v
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetMd5() string {
	if o == nil || IsNil(o.Md5) {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetMd5Ok() (*string, bool) {
	if o == nil || IsNil(o.Md5) {
		return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasMd5() bool {
	if o != nil && !IsNil(o.Md5) {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *VrfVirtualCircuit) SetMd5(v string) {
	o.Md5 = &v
}

// GetMetalIp returns the MetalIp field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetMetalIp() string {
	if o == nil || IsNil(o.MetalIp) {
		var ret string
		return ret
	}
	return *o.MetalIp
}

// GetMetalIpOk returns a tuple with the MetalIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetMetalIpOk() (*string, bool) {
	if o == nil || IsNil(o.MetalIp) {
		return nil, false
	}
	return o.MetalIp, true
}

// HasMetalIp returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasMetalIp() bool {
	if o != nil && !IsNil(o.MetalIp) {
		return true
	}

	return false
}

// SetMetalIp gets a reference to the given string and assigns it to the MetalIp field.
func (o *VrfVirtualCircuit) SetMetalIp(v string) {
	o.MetalIp = &v
}

// GetMetalIpv6 returns the MetalIpv6 field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetMetalIpv6() string {
	if o == nil || IsNil(o.MetalIpv6) {
		var ret string
		return ret
	}
	return *o.MetalIpv6
}

// GetMetalIpv6Ok returns a tuple with the MetalIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetMetalIpv6Ok() (*string, bool) {
	if o == nil || IsNil(o.MetalIpv6) {
		return nil, false
	}
	return o.MetalIpv6, true
}

// HasMetalIpv6 returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasMetalIpv6() bool {
	if o != nil && !IsNil(o.MetalIpv6) {
		return true
	}

	return false
}

// SetMetalIpv6 gets a reference to the given string and assigns it to the MetalIpv6 field.
func (o *VrfVirtualCircuit) SetMetalIpv6(v string) {
	o.MetalIpv6 = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VrfVirtualCircuit) SetName(v string) {
	o.Name = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetPort() InterconnectionPort {
	if o == nil || IsNil(o.Port) {
		var ret InterconnectionPort
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetPortOk() (*InterconnectionPort, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given InterconnectionPort and assigns it to the Port field.
func (o *VrfVirtualCircuit) SetPort(v InterconnectionPort) {
	o.Port = &v
}

// GetNniVlan returns the NniVlan field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetNniVlan() int32 {
	if o == nil || IsNil(o.NniVlan) {
		var ret int32
		return ret
	}
	return *o.NniVlan
}

// GetNniVlanOk returns a tuple with the NniVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetNniVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.NniVlan) {
		return nil, false
	}
	return o.NniVlan, true
}

// HasNniVlan returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasNniVlan() bool {
	if o != nil && !IsNil(o.NniVlan) {
		return true
	}

	return false
}

// SetNniVlan gets a reference to the given int32 and assigns it to the NniVlan field.
func (o *VrfVirtualCircuit) SetNniVlan(v int32) {
	o.NniVlan = &v
}

// GetPeerAsn returns the PeerAsn field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetPeerAsn() int64 {
	if o == nil || IsNil(o.PeerAsn) {
		var ret int64
		return ret
	}
	return *o.PeerAsn
}

// GetPeerAsnOk returns a tuple with the PeerAsn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetPeerAsnOk() (*int64, bool) {
	if o == nil || IsNil(o.PeerAsn) {
		return nil, false
	}
	return o.PeerAsn, true
}

// HasPeerAsn returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasPeerAsn() bool {
	if o != nil && !IsNil(o.PeerAsn) {
		return true
	}

	return false
}

// SetPeerAsn gets a reference to the given int64 and assigns it to the PeerAsn field.
func (o *VrfVirtualCircuit) SetPeerAsn(v int64) {
	o.PeerAsn = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetProject() Project {
	if o == nil || IsNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetProjectOk() (*Project, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *VrfVirtualCircuit) SetProject(v Project) {
	o.Project = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetSpeed() int64 {
	if o == nil || IsNil(o.Speed) {
		var ret int64
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetSpeedOk() (*int64, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int64 and assigns it to the Speed field.
func (o *VrfVirtualCircuit) SetSpeed(v int64) {
	o.Speed = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetStatus() VrfVirtualCircuitStatus {
	if o == nil || IsNil(o.Status) {
		var ret VrfVirtualCircuitStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetStatusOk() (*VrfVirtualCircuitStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given VrfVirtualCircuitStatus and assigns it to the Status field.
func (o *VrfVirtualCircuit) SetStatus(v VrfVirtualCircuitStatus) {
	o.Status = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetSubnet() string {
	if o == nil || IsNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.Subnet) {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasSubnet() bool {
	if o != nil && !IsNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *VrfVirtualCircuit) SetSubnet(v string) {
	o.Subnet = &v
}

// GetSubnetIpv6 returns the SubnetIpv6 field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetSubnetIpv6() string {
	if o == nil || IsNil(o.SubnetIpv6) {
		var ret string
		return ret
	}
	return *o.SubnetIpv6
}

// GetSubnetIpv6Ok returns a tuple with the SubnetIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetSubnetIpv6Ok() (*string, bool) {
	if o == nil || IsNil(o.SubnetIpv6) {
		return nil, false
	}
	return o.SubnetIpv6, true
}

// HasSubnetIpv6 returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasSubnetIpv6() bool {
	if o != nil && !IsNil(o.SubnetIpv6) {
		return true
	}

	return false
}

// SetSubnetIpv6 gets a reference to the given string and assigns it to the SubnetIpv6 field.
func (o *VrfVirtualCircuit) SetSubnetIpv6(v string) {
	o.SubnetIpv6 = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *VrfVirtualCircuit) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetType() VrfIpReservationType {
	if o == nil || IsNil(o.Type) {
		var ret VrfIpReservationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetTypeOk() (*VrfIpReservationType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given VrfIpReservationType and assigns it to the Type field.
func (o *VrfVirtualCircuit) SetType(v VrfIpReservationType) {
	o.Type = &v
}

// GetVrf returns the Vrf field value
func (o *VrfVirtualCircuit) GetVrf() Vrf {
	if o == nil {
		var ret Vrf
		return ret
	}

	return o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetVrfOk() (*Vrf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vrf, true
}

// SetVrf sets field value
func (o *VrfVirtualCircuit) SetVrf(v Vrf) {
	o.Vrf = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VrfVirtualCircuit) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *VrfVirtualCircuit) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuit) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *VrfVirtualCircuit) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *VrfVirtualCircuit) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o VrfVirtualCircuit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VrfVirtualCircuit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerIp) {
		toSerialize["customer_ip"] = o.CustomerIp
	}
	if !IsNil(o.CustomerIpv6) {
		toSerialize["customer_ipv6"] = o.CustomerIpv6
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Md5) {
		toSerialize["md5"] = o.Md5
	}
	if !IsNil(o.MetalIp) {
		toSerialize["metal_ip"] = o.MetalIp
	}
	if !IsNil(o.MetalIpv6) {
		toSerialize["metal_ipv6"] = o.MetalIpv6
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.NniVlan) {
		toSerialize["nni_vlan"] = o.NniVlan
	}
	if !IsNil(o.PeerAsn) {
		toSerialize["peer_asn"] = o.PeerAsn
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !IsNil(o.SubnetIpv6) {
		toSerialize["subnet_ipv6"] = o.SubnetIpv6
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["vrf"] = o.Vrf
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VrfVirtualCircuit) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vrf",
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

	varVrfVirtualCircuit := _VrfVirtualCircuit{}

	err = json.Unmarshal(data, &varVrfVirtualCircuit)

	if err != nil {
		return err
	}

	*o = VrfVirtualCircuit(varVrfVirtualCircuit)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customer_ip")
		delete(additionalProperties, "customer_ipv6")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "md5")
		delete(additionalProperties, "metal_ip")
		delete(additionalProperties, "metal_ipv6")
		delete(additionalProperties, "name")
		delete(additionalProperties, "port")
		delete(additionalProperties, "nni_vlan")
		delete(additionalProperties, "peer_asn")
		delete(additionalProperties, "project")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "status")
		delete(additionalProperties, "subnet")
		delete(additionalProperties, "subnet_ipv6")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "type")
		delete(additionalProperties, "vrf")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVrfVirtualCircuit struct {
	value *VrfVirtualCircuit
	isSet bool
}

func (v NullableVrfVirtualCircuit) Get() *VrfVirtualCircuit {
	return v.value
}

func (v *NullableVrfVirtualCircuit) Set(val *VrfVirtualCircuit) {
	v.value = val
	v.isSet = true
}

func (v NullableVrfVirtualCircuit) IsSet() bool {
	return v.isSet
}

func (v *NullableVrfVirtualCircuit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrfVirtualCircuit(val *VrfVirtualCircuit) *NullableVrfVirtualCircuit {
	return &NullableVrfVirtualCircuit{value: val, isSet: true}
}

func (v NullableVrfVirtualCircuit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrfVirtualCircuit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
