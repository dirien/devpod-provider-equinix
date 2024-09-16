/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the IPReservationFacility type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPReservationFacility{}

// IPReservationFacility struct for IPReservationFacility
type IPReservationFacility struct {
	Address  *Address                `json:"address,omitempty"`
	Code     *string                 `json:"code,omitempty"`
	Features []FacilityFeaturesInner `json:"features,omitempty"`
	Id       *string                 `json:"id,omitempty"`
	// IP ranges registered in facility. Can be used for GeoIP location
	IpRanges             []string     `json:"ip_ranges,omitempty"`
	Metro                *DeviceMetro `json:"metro,omitempty"`
	Name                 *string      `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IPReservationFacility IPReservationFacility

// NewIPReservationFacility instantiates a new IPReservationFacility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPReservationFacility() *IPReservationFacility {
	this := IPReservationFacility{}
	return &this
}

// NewIPReservationFacilityWithDefaults instantiates a new IPReservationFacility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPReservationFacilityWithDefaults() *IPReservationFacility {
	this := IPReservationFacility{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *IPReservationFacility) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservationFacility) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *IPReservationFacility) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *IPReservationFacility) SetAddress(v Address) {
	o.Address = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *IPReservationFacility) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservationFacility) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *IPReservationFacility) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *IPReservationFacility) SetCode(v string) {
	o.Code = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *IPReservationFacility) GetFeatures() []FacilityFeaturesInner {
	if o == nil || IsNil(o.Features) {
		var ret []FacilityFeaturesInner
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservationFacility) GetFeaturesOk() ([]FacilityFeaturesInner, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *IPReservationFacility) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []FacilityFeaturesInner and assigns it to the Features field.
func (o *IPReservationFacility) SetFeatures(v []FacilityFeaturesInner) {
	o.Features = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IPReservationFacility) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservationFacility) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IPReservationFacility) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IPReservationFacility) SetId(v string) {
	o.Id = &v
}

// GetIpRanges returns the IpRanges field value if set, zero value otherwise.
func (o *IPReservationFacility) GetIpRanges() []string {
	if o == nil || IsNil(o.IpRanges) {
		var ret []string
		return ret
	}
	return o.IpRanges
}

// GetIpRangesOk returns a tuple with the IpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservationFacility) GetIpRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.IpRanges) {
		return nil, false
	}
	return o.IpRanges, true
}

// HasIpRanges returns a boolean if a field has been set.
func (o *IPReservationFacility) HasIpRanges() bool {
	if o != nil && !IsNil(o.IpRanges) {
		return true
	}

	return false
}

// SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.
func (o *IPReservationFacility) SetIpRanges(v []string) {
	o.IpRanges = v
}

// GetMetro returns the Metro field value if set, zero value otherwise.
func (o *IPReservationFacility) GetMetro() DeviceMetro {
	if o == nil || IsNil(o.Metro) {
		var ret DeviceMetro
		return ret
	}
	return *o.Metro
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservationFacility) GetMetroOk() (*DeviceMetro, bool) {
	if o == nil || IsNil(o.Metro) {
		return nil, false
	}
	return o.Metro, true
}

// HasMetro returns a boolean if a field has been set.
func (o *IPReservationFacility) HasMetro() bool {
	if o != nil && !IsNil(o.Metro) {
		return true
	}

	return false
}

// SetMetro gets a reference to the given DeviceMetro and assigns it to the Metro field.
func (o *IPReservationFacility) SetMetro(v DeviceMetro) {
	o.Metro = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IPReservationFacility) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservationFacility) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IPReservationFacility) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IPReservationFacility) SetName(v string) {
	o.Name = &v
}

func (o IPReservationFacility) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPReservationFacility) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IpRanges) {
		toSerialize["ip_ranges"] = o.IpRanges
	}
	if !IsNil(o.Metro) {
		toSerialize["metro"] = o.Metro
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IPReservationFacility) UnmarshalJSON(data []byte) (err error) {
	varIPReservationFacility := _IPReservationFacility{}

	err = json.Unmarshal(data, &varIPReservationFacility)

	if err != nil {
		return err
	}

	*o = IPReservationFacility(varIPReservationFacility)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "code")
		delete(additionalProperties, "features")
		delete(additionalProperties, "id")
		delete(additionalProperties, "ip_ranges")
		delete(additionalProperties, "metro")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIPReservationFacility struct {
	value *IPReservationFacility
	isSet bool
}

func (v NullableIPReservationFacility) Get() *IPReservationFacility {
	return v.value
}

func (v *NullableIPReservationFacility) Set(val *IPReservationFacility) {
	v.value = val
	v.isSet = true
}

func (v NullableIPReservationFacility) IsSet() bool {
	return v.isSet
}

func (v *NullableIPReservationFacility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPReservationFacility(val *IPReservationFacility) *NullableIPReservationFacility {
	return &NullableIPReservationFacility{value: val, isSet: true}
}

func (v NullableIPReservationFacility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPReservationFacility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
