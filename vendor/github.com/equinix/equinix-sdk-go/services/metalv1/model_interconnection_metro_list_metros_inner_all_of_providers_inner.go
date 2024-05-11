/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. Currently the search parameter is only available on devices, ssh_keys, api_keys and memberships endpoints.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the InterconnectionMetroListMetrosInnerAllOfProvidersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterconnectionMetroListMetrosInnerAllOfProvidersInner{}

// InterconnectionMetroListMetrosInnerAllOfProvidersInner struct for InterconnectionMetroListMetrosInnerAllOfProvidersInner
type InterconnectionMetroListMetrosInnerAllOfProvidersInner struct {
	Type                 *string  `json:"type,omitempty"`
	Name                 *string  `json:"name,omitempty"`
	Locations            []string `json:"locations,omitempty"`
	Bandwidths           []int32  `json:"bandwidths,omitempty"`
	Features             []string `json:"features,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InterconnectionMetroListMetrosInnerAllOfProvidersInner InterconnectionMetroListMetrosInnerAllOfProvidersInner

// NewInterconnectionMetroListMetrosInnerAllOfProvidersInner instantiates a new InterconnectionMetroListMetrosInnerAllOfProvidersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterconnectionMetroListMetrosInnerAllOfProvidersInner() *InterconnectionMetroListMetrosInnerAllOfProvidersInner {
	this := InterconnectionMetroListMetrosInnerAllOfProvidersInner{}
	return &this
}

// NewInterconnectionMetroListMetrosInnerAllOfProvidersInnerWithDefaults instantiates a new InterconnectionMetroListMetrosInnerAllOfProvidersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterconnectionMetroListMetrosInnerAllOfProvidersInnerWithDefaults() *InterconnectionMetroListMetrosInnerAllOfProvidersInner {
	this := InterconnectionMetroListMetrosInnerAllOfProvidersInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InterconnectionMetroListMetrosInnerAllOfProvidersInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionMetroListMetrosInnerAllOfProvidersInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InterconnectionMetroListMetrosInnerAllOfProvidersInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InterconnectionMetroListMetrosInnerAllOfProvidersInner) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InterconnectionMetroListMetrosInnerAllOfProvidersInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionMetroListMetrosInnerAllOfProvidersInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InterconnectionMetroListMetrosInnerAllOfProvidersInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InterconnectionMetroListMetrosInnerAllOfProvidersInner) SetName(v string) {
	o.Name = &v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *InterconnectionMetroListMetrosInnerAllOfProvidersInner) GetLocations() []string {
	if o == nil || IsNil(o.Locations) {
		var ret []string
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionMetroListMetrosInnerAllOfProvidersInner) GetLocationsOk() ([]string, bool) {
	if o == nil || IsNil(o.Locations) {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *InterconnectionMetroListMetrosInnerAllOfProvidersInner) HasLocations() bool {
	if o != nil && !IsNil(o.Locations) {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []string and assigns it to the Locations field.
func (o *InterconnectionMetroListMetrosInnerAllOfProvidersInner) SetLocations(v []string) {
	o.Locations = v
}

// GetBandwidths returns the Bandwidths field value if set, zero value otherwise.
func (o *InterconnectionMetroListMetrosInnerAllOfProvidersInner) GetBandwidths() []int32 {
	if o == nil || IsNil(o.Bandwidths) {
		var ret []int32
		return ret
	}
	return o.Bandwidths
}

// GetBandwidthsOk returns a tuple with the Bandwidths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionMetroListMetrosInnerAllOfProvidersInner) GetBandwidthsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Bandwidths) {
		return nil, false
	}
	return o.Bandwidths, true
}

// HasBandwidths returns a boolean if a field has been set.
func (o *InterconnectionMetroListMetrosInnerAllOfProvidersInner) HasBandwidths() bool {
	if o != nil && !IsNil(o.Bandwidths) {
		return true
	}

	return false
}

// SetBandwidths gets a reference to the given []int32 and assigns it to the Bandwidths field.
func (o *InterconnectionMetroListMetrosInnerAllOfProvidersInner) SetBandwidths(v []int32) {
	o.Bandwidths = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *InterconnectionMetroListMetrosInnerAllOfProvidersInner) GetFeatures() []string {
	if o == nil || IsNil(o.Features) {
		var ret []string
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionMetroListMetrosInnerAllOfProvidersInner) GetFeaturesOk() ([]string, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *InterconnectionMetroListMetrosInnerAllOfProvidersInner) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *InterconnectionMetroListMetrosInnerAllOfProvidersInner) SetFeatures(v []string) {
	o.Features = v
}

func (o InterconnectionMetroListMetrosInnerAllOfProvidersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterconnectionMetroListMetrosInnerAllOfProvidersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Locations) {
		toSerialize["locations"] = o.Locations
	}
	if !IsNil(o.Bandwidths) {
		toSerialize["bandwidths"] = o.Bandwidths
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InterconnectionMetroListMetrosInnerAllOfProvidersInner) UnmarshalJSON(data []byte) (err error) {
	varInterconnectionMetroListMetrosInnerAllOfProvidersInner := _InterconnectionMetroListMetrosInnerAllOfProvidersInner{}

	err = json.Unmarshal(data, &varInterconnectionMetroListMetrosInnerAllOfProvidersInner)

	if err != nil {
		return err
	}

	*o = InterconnectionMetroListMetrosInnerAllOfProvidersInner(varInterconnectionMetroListMetrosInnerAllOfProvidersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "locations")
		delete(additionalProperties, "bandwidths")
		delete(additionalProperties, "features")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInterconnectionMetroListMetrosInnerAllOfProvidersInner struct {
	value *InterconnectionMetroListMetrosInnerAllOfProvidersInner
	isSet bool
}

func (v NullableInterconnectionMetroListMetrosInnerAllOfProvidersInner) Get() *InterconnectionMetroListMetrosInnerAllOfProvidersInner {
	return v.value
}

func (v *NullableInterconnectionMetroListMetrosInnerAllOfProvidersInner) Set(val *InterconnectionMetroListMetrosInnerAllOfProvidersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInterconnectionMetroListMetrosInnerAllOfProvidersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInterconnectionMetroListMetrosInnerAllOfProvidersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterconnectionMetroListMetrosInnerAllOfProvidersInner(val *InterconnectionMetroListMetrosInnerAllOfProvidersInner) *NullableInterconnectionMetroListMetrosInnerAllOfProvidersInner {
	return &NullableInterconnectionMetroListMetrosInnerAllOfProvidersInner{value: val, isSet: true}
}

func (v NullableInterconnectionMetroListMetrosInnerAllOfProvidersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterconnectionMetroListMetrosInnerAllOfProvidersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
