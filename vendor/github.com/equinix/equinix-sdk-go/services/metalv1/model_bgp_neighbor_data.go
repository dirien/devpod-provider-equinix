/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the BgpNeighborData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BgpNeighborData{}

// BgpNeighborData struct for BgpNeighborData
type BgpNeighborData struct {
	// Address Family for IP Address. Accepted values are 4 or 6
	AddressFamily *int32 `json:"address_family,omitempty"`
	// The customer's ASN. In a local BGP deployment, this will be an internal ASN used to route within the data center. For a global BGP deployment, this will be the your own ASN, configured when you set up BGP for your project.
	CustomerAs *int32 `json:"customer_as,omitempty"`
	// The device's IP address. For an IPv4 BGP session, this is typically the private bond0 address for the device.
	CustomerIp *string `json:"customer_ip,omitempty"`
	// True if an MD5 password is configured for the project.
	Md5Enabled *bool `json:"md5_enabled,omitempty"`
	// The MD5 password configured for the project, if set.
	Md5Password *string `json:"md5_password,omitempty"`
	// True when the BGP session should be configured as multihop.
	Multihop *bool `json:"multihop,omitempty"`
	// The Peer ASN to use when configuring BGP on your device.
	PeerAs *int32 `json:"peer_as,omitempty"`
	// A list of one or more IP addresses to use for the Peer IP section of your BGP configuration. For non-multihop sessions, this will typically be a single gateway address for the device. For multihop sessions, it will be a list of IPs.
	PeerIps []string `json:"peer_ips,omitempty"`
	// A list of project subnets
	RoutesIn []BgpRoute `json:"routes_in,omitempty"`
	// A list of outgoing routes. Only populated if the BGP session has default route enabled.
	RoutesOut            []BgpRoute `json:"routes_out,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BgpNeighborData BgpNeighborData

// NewBgpNeighborData instantiates a new BgpNeighborData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpNeighborData() *BgpNeighborData {
	this := BgpNeighborData{}
	return &this
}

// NewBgpNeighborDataWithDefaults instantiates a new BgpNeighborData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpNeighborDataWithDefaults() *BgpNeighborData {
	this := BgpNeighborData{}
	return &this
}

// GetAddressFamily returns the AddressFamily field value if set, zero value otherwise.
func (o *BgpNeighborData) GetAddressFamily() int32 {
	if o == nil || IsNil(o.AddressFamily) {
		var ret int32
		return ret
	}
	return *o.AddressFamily
}

// GetAddressFamilyOk returns a tuple with the AddressFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborData) GetAddressFamilyOk() (*int32, bool) {
	if o == nil || IsNil(o.AddressFamily) {
		return nil, false
	}
	return o.AddressFamily, true
}

// HasAddressFamily returns a boolean if a field has been set.
func (o *BgpNeighborData) HasAddressFamily() bool {
	if o != nil && !IsNil(o.AddressFamily) {
		return true
	}

	return false
}

// SetAddressFamily gets a reference to the given int32 and assigns it to the AddressFamily field.
func (o *BgpNeighborData) SetAddressFamily(v int32) {
	o.AddressFamily = &v
}

// GetCustomerAs returns the CustomerAs field value if set, zero value otherwise.
func (o *BgpNeighborData) GetCustomerAs() int32 {
	if o == nil || IsNil(o.CustomerAs) {
		var ret int32
		return ret
	}
	return *o.CustomerAs
}

// GetCustomerAsOk returns a tuple with the CustomerAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborData) GetCustomerAsOk() (*int32, bool) {
	if o == nil || IsNil(o.CustomerAs) {
		return nil, false
	}
	return o.CustomerAs, true
}

// HasCustomerAs returns a boolean if a field has been set.
func (o *BgpNeighborData) HasCustomerAs() bool {
	if o != nil && !IsNil(o.CustomerAs) {
		return true
	}

	return false
}

// SetCustomerAs gets a reference to the given int32 and assigns it to the CustomerAs field.
func (o *BgpNeighborData) SetCustomerAs(v int32) {
	o.CustomerAs = &v
}

// GetCustomerIp returns the CustomerIp field value if set, zero value otherwise.
func (o *BgpNeighborData) GetCustomerIp() string {
	if o == nil || IsNil(o.CustomerIp) {
		var ret string
		return ret
	}
	return *o.CustomerIp
}

// GetCustomerIpOk returns a tuple with the CustomerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborData) GetCustomerIpOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerIp) {
		return nil, false
	}
	return o.CustomerIp, true
}

// HasCustomerIp returns a boolean if a field has been set.
func (o *BgpNeighborData) HasCustomerIp() bool {
	if o != nil && !IsNil(o.CustomerIp) {
		return true
	}

	return false
}

// SetCustomerIp gets a reference to the given string and assigns it to the CustomerIp field.
func (o *BgpNeighborData) SetCustomerIp(v string) {
	o.CustomerIp = &v
}

// GetMd5Enabled returns the Md5Enabled field value if set, zero value otherwise.
func (o *BgpNeighborData) GetMd5Enabled() bool {
	if o == nil || IsNil(o.Md5Enabled) {
		var ret bool
		return ret
	}
	return *o.Md5Enabled
}

// GetMd5EnabledOk returns a tuple with the Md5Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborData) GetMd5EnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Md5Enabled) {
		return nil, false
	}
	return o.Md5Enabled, true
}

// HasMd5Enabled returns a boolean if a field has been set.
func (o *BgpNeighborData) HasMd5Enabled() bool {
	if o != nil && !IsNil(o.Md5Enabled) {
		return true
	}

	return false
}

// SetMd5Enabled gets a reference to the given bool and assigns it to the Md5Enabled field.
func (o *BgpNeighborData) SetMd5Enabled(v bool) {
	o.Md5Enabled = &v
}

// GetMd5Password returns the Md5Password field value if set, zero value otherwise.
func (o *BgpNeighborData) GetMd5Password() string {
	if o == nil || IsNil(o.Md5Password) {
		var ret string
		return ret
	}
	return *o.Md5Password
}

// GetMd5PasswordOk returns a tuple with the Md5Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborData) GetMd5PasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Md5Password) {
		return nil, false
	}
	return o.Md5Password, true
}

// HasMd5Password returns a boolean if a field has been set.
func (o *BgpNeighborData) HasMd5Password() bool {
	if o != nil && !IsNil(o.Md5Password) {
		return true
	}

	return false
}

// SetMd5Password gets a reference to the given string and assigns it to the Md5Password field.
func (o *BgpNeighborData) SetMd5Password(v string) {
	o.Md5Password = &v
}

// GetMultihop returns the Multihop field value if set, zero value otherwise.
func (o *BgpNeighborData) GetMultihop() bool {
	if o == nil || IsNil(o.Multihop) {
		var ret bool
		return ret
	}
	return *o.Multihop
}

// GetMultihopOk returns a tuple with the Multihop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborData) GetMultihopOk() (*bool, bool) {
	if o == nil || IsNil(o.Multihop) {
		return nil, false
	}
	return o.Multihop, true
}

// HasMultihop returns a boolean if a field has been set.
func (o *BgpNeighborData) HasMultihop() bool {
	if o != nil && !IsNil(o.Multihop) {
		return true
	}

	return false
}

// SetMultihop gets a reference to the given bool and assigns it to the Multihop field.
func (o *BgpNeighborData) SetMultihop(v bool) {
	o.Multihop = &v
}

// GetPeerAs returns the PeerAs field value if set, zero value otherwise.
func (o *BgpNeighborData) GetPeerAs() int32 {
	if o == nil || IsNil(o.PeerAs) {
		var ret int32
		return ret
	}
	return *o.PeerAs
}

// GetPeerAsOk returns a tuple with the PeerAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborData) GetPeerAsOk() (*int32, bool) {
	if o == nil || IsNil(o.PeerAs) {
		return nil, false
	}
	return o.PeerAs, true
}

// HasPeerAs returns a boolean if a field has been set.
func (o *BgpNeighborData) HasPeerAs() bool {
	if o != nil && !IsNil(o.PeerAs) {
		return true
	}

	return false
}

// SetPeerAs gets a reference to the given int32 and assigns it to the PeerAs field.
func (o *BgpNeighborData) SetPeerAs(v int32) {
	o.PeerAs = &v
}

// GetPeerIps returns the PeerIps field value if set, zero value otherwise.
func (o *BgpNeighborData) GetPeerIps() []string {
	if o == nil || IsNil(o.PeerIps) {
		var ret []string
		return ret
	}
	return o.PeerIps
}

// GetPeerIpsOk returns a tuple with the PeerIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborData) GetPeerIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.PeerIps) {
		return nil, false
	}
	return o.PeerIps, true
}

// HasPeerIps returns a boolean if a field has been set.
func (o *BgpNeighborData) HasPeerIps() bool {
	if o != nil && !IsNil(o.PeerIps) {
		return true
	}

	return false
}

// SetPeerIps gets a reference to the given []string and assigns it to the PeerIps field.
func (o *BgpNeighborData) SetPeerIps(v []string) {
	o.PeerIps = v
}

// GetRoutesIn returns the RoutesIn field value if set, zero value otherwise.
func (o *BgpNeighborData) GetRoutesIn() []BgpRoute {
	if o == nil || IsNil(o.RoutesIn) {
		var ret []BgpRoute
		return ret
	}
	return o.RoutesIn
}

// GetRoutesInOk returns a tuple with the RoutesIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborData) GetRoutesInOk() ([]BgpRoute, bool) {
	if o == nil || IsNil(o.RoutesIn) {
		return nil, false
	}
	return o.RoutesIn, true
}

// HasRoutesIn returns a boolean if a field has been set.
func (o *BgpNeighborData) HasRoutesIn() bool {
	if o != nil && !IsNil(o.RoutesIn) {
		return true
	}

	return false
}

// SetRoutesIn gets a reference to the given []BgpRoute and assigns it to the RoutesIn field.
func (o *BgpNeighborData) SetRoutesIn(v []BgpRoute) {
	o.RoutesIn = v
}

// GetRoutesOut returns the RoutesOut field value if set, zero value otherwise.
func (o *BgpNeighborData) GetRoutesOut() []BgpRoute {
	if o == nil || IsNil(o.RoutesOut) {
		var ret []BgpRoute
		return ret
	}
	return o.RoutesOut
}

// GetRoutesOutOk returns a tuple with the RoutesOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpNeighborData) GetRoutesOutOk() ([]BgpRoute, bool) {
	if o == nil || IsNil(o.RoutesOut) {
		return nil, false
	}
	return o.RoutesOut, true
}

// HasRoutesOut returns a boolean if a field has been set.
func (o *BgpNeighborData) HasRoutesOut() bool {
	if o != nil && !IsNil(o.RoutesOut) {
		return true
	}

	return false
}

// SetRoutesOut gets a reference to the given []BgpRoute and assigns it to the RoutesOut field.
func (o *BgpNeighborData) SetRoutesOut(v []BgpRoute) {
	o.RoutesOut = v
}

func (o BgpNeighborData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BgpNeighborData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressFamily) {
		toSerialize["address_family"] = o.AddressFamily
	}
	if !IsNil(o.CustomerAs) {
		toSerialize["customer_as"] = o.CustomerAs
	}
	if !IsNil(o.CustomerIp) {
		toSerialize["customer_ip"] = o.CustomerIp
	}
	if !IsNil(o.Md5Enabled) {
		toSerialize["md5_enabled"] = o.Md5Enabled
	}
	if !IsNil(o.Md5Password) {
		toSerialize["md5_password"] = o.Md5Password
	}
	if !IsNil(o.Multihop) {
		toSerialize["multihop"] = o.Multihop
	}
	if !IsNil(o.PeerAs) {
		toSerialize["peer_as"] = o.PeerAs
	}
	if !IsNil(o.PeerIps) {
		toSerialize["peer_ips"] = o.PeerIps
	}
	if !IsNil(o.RoutesIn) {
		toSerialize["routes_in"] = o.RoutesIn
	}
	if !IsNil(o.RoutesOut) {
		toSerialize["routes_out"] = o.RoutesOut
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BgpNeighborData) UnmarshalJSON(data []byte) (err error) {
	varBgpNeighborData := _BgpNeighborData{}

	err = json.Unmarshal(data, &varBgpNeighborData)

	if err != nil {
		return err
	}

	*o = BgpNeighborData(varBgpNeighborData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address_family")
		delete(additionalProperties, "customer_as")
		delete(additionalProperties, "customer_ip")
		delete(additionalProperties, "md5_enabled")
		delete(additionalProperties, "md5_password")
		delete(additionalProperties, "multihop")
		delete(additionalProperties, "peer_as")
		delete(additionalProperties, "peer_ips")
		delete(additionalProperties, "routes_in")
		delete(additionalProperties, "routes_out")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBgpNeighborData struct {
	value *BgpNeighborData
	isSet bool
}

func (v NullableBgpNeighborData) Get() *BgpNeighborData {
	return v.value
}

func (v *NullableBgpNeighborData) Set(val *BgpNeighborData) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpNeighborData) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpNeighborData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpNeighborData(val *BgpNeighborData) *NullableBgpNeighborData {
	return &NullableBgpNeighborData{value: val, isSet: true}
}

func (v NullableBgpNeighborData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpNeighborData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
