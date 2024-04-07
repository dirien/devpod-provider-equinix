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
	"fmt"
)

// checks if the MetalGatewayElasticIpCreateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetalGatewayElasticIpCreateInput{}

// MetalGatewayElasticIpCreateInput struct for MetalGatewayElasticIpCreateInput
type MetalGatewayElasticIpCreateInput struct {
	// An IP address (or IP Address range) contained within one of the project's IP Reservations
	Address string `json:"address"`
	// An IP address contained within the Metal Gateways' IP Reservation range.
	NextHop string `json:"next_hop"`
	// Optional User-defined JSON object value.
	Customdata map[string]interface{} `json:"customdata,omitempty"`
	// Optional list of User-defined tags. Can be used by users to provide additional details or context regarding the purpose or usage of this resource.
	Tags                 []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MetalGatewayElasticIpCreateInput MetalGatewayElasticIpCreateInput

// NewMetalGatewayElasticIpCreateInput instantiates a new MetalGatewayElasticIpCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetalGatewayElasticIpCreateInput(address string, nextHop string) *MetalGatewayElasticIpCreateInput {
	this := MetalGatewayElasticIpCreateInput{}
	this.Address = address
	this.NextHop = nextHop
	return &this
}

// NewMetalGatewayElasticIpCreateInputWithDefaults instantiates a new MetalGatewayElasticIpCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetalGatewayElasticIpCreateInputWithDefaults() *MetalGatewayElasticIpCreateInput {
	this := MetalGatewayElasticIpCreateInput{}
	return &this
}

// GetAddress returns the Address field value
func (o *MetalGatewayElasticIpCreateInput) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *MetalGatewayElasticIpCreateInput) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *MetalGatewayElasticIpCreateInput) SetAddress(v string) {
	o.Address = v
}

// GetNextHop returns the NextHop field value
func (o *MetalGatewayElasticIpCreateInput) GetNextHop() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextHop
}

// GetNextHopOk returns a tuple with the NextHop field value
// and a boolean to check if the value has been set.
func (o *MetalGatewayElasticIpCreateInput) GetNextHopOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextHop, true
}

// SetNextHop sets field value
func (o *MetalGatewayElasticIpCreateInput) SetNextHop(v string) {
	o.NextHop = v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *MetalGatewayElasticIpCreateInput) GetCustomdata() map[string]interface{} {
	if o == nil || IsNil(o.Customdata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetalGatewayElasticIpCreateInput) GetCustomdataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Customdata) {
		return map[string]interface{}{}, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *MetalGatewayElasticIpCreateInput) HasCustomdata() bool {
	if o != nil && !IsNil(o.Customdata) {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *MetalGatewayElasticIpCreateInput) SetCustomdata(v map[string]interface{}) {
	o.Customdata = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MetalGatewayElasticIpCreateInput) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetalGatewayElasticIpCreateInput) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MetalGatewayElasticIpCreateInput) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *MetalGatewayElasticIpCreateInput) SetTags(v []string) {
	o.Tags = v
}

func (o MetalGatewayElasticIpCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetalGatewayElasticIpCreateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["next_hop"] = o.NextHop
	if !IsNil(o.Customdata) {
		toSerialize["customdata"] = o.Customdata
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MetalGatewayElasticIpCreateInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"next_hop",
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

	varMetalGatewayElasticIpCreateInput := _MetalGatewayElasticIpCreateInput{}

	err = json.Unmarshal(data, &varMetalGatewayElasticIpCreateInput)

	if err != nil {
		return err
	}

	*o = MetalGatewayElasticIpCreateInput(varMetalGatewayElasticIpCreateInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "next_hop")
		delete(additionalProperties, "customdata")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMetalGatewayElasticIpCreateInput struct {
	value *MetalGatewayElasticIpCreateInput
	isSet bool
}

func (v NullableMetalGatewayElasticIpCreateInput) Get() *MetalGatewayElasticIpCreateInput {
	return v.value
}

func (v *NullableMetalGatewayElasticIpCreateInput) Set(val *MetalGatewayElasticIpCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableMetalGatewayElasticIpCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableMetalGatewayElasticIpCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetalGatewayElasticIpCreateInput(val *MetalGatewayElasticIpCreateInput) *NullableMetalGatewayElasticIpCreateInput {
	return &NullableMetalGatewayElasticIpCreateInput{value: val, isSet: true}
}

func (v NullableMetalGatewayElasticIpCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetalGatewayElasticIpCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
