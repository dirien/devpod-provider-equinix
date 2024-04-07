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

// checks if the InvitationInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvitationInput{}

// InvitationInput struct for InvitationInput
type InvitationInput struct {
	Invitee              string                 `json:"invitee"`
	Message              *string                `json:"message,omitempty"`
	OrganizationId       *string                `json:"organization_id,omitempty"`
	ProjectsIds          []string               `json:"projects_ids,omitempty"`
	Roles                []InvitationRolesInner `json:"roles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InvitationInput InvitationInput

// NewInvitationInput instantiates a new InvitationInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitationInput(invitee string) *InvitationInput {
	this := InvitationInput{}
	this.Invitee = invitee
	return &this
}

// NewInvitationInputWithDefaults instantiates a new InvitationInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationInputWithDefaults() *InvitationInput {
	this := InvitationInput{}
	return &this
}

// GetInvitee returns the Invitee field value
func (o *InvitationInput) GetInvitee() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Invitee
}

// GetInviteeOk returns a tuple with the Invitee field value
// and a boolean to check if the value has been set.
func (o *InvitationInput) GetInviteeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Invitee, true
}

// SetInvitee sets field value
func (o *InvitationInput) SetInvitee(v string) {
	o.Invitee = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InvitationInput) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationInput) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InvitationInput) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InvitationInput) SetMessage(v string) {
	o.Message = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *InvitationInput) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationInput) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *InvitationInput) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *InvitationInput) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetProjectsIds returns the ProjectsIds field value if set, zero value otherwise.
func (o *InvitationInput) GetProjectsIds() []string {
	if o == nil || IsNil(o.ProjectsIds) {
		var ret []string
		return ret
	}
	return o.ProjectsIds
}

// GetProjectsIdsOk returns a tuple with the ProjectsIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationInput) GetProjectsIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectsIds) {
		return nil, false
	}
	return o.ProjectsIds, true
}

// HasProjectsIds returns a boolean if a field has been set.
func (o *InvitationInput) HasProjectsIds() bool {
	if o != nil && !IsNil(o.ProjectsIds) {
		return true
	}

	return false
}

// SetProjectsIds gets a reference to the given []string and assigns it to the ProjectsIds field.
func (o *InvitationInput) SetProjectsIds(v []string) {
	o.ProjectsIds = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *InvitationInput) GetRoles() []InvitationRolesInner {
	if o == nil || IsNil(o.Roles) {
		var ret []InvitationRolesInner
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationInput) GetRolesOk() ([]InvitationRolesInner, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *InvitationInput) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []InvitationRolesInner and assigns it to the Roles field.
func (o *InvitationInput) SetRoles(v []InvitationRolesInner) {
	o.Roles = v
}

func (o InvitationInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvitationInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invitee"] = o.Invitee
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !IsNil(o.ProjectsIds) {
		toSerialize["projects_ids"] = o.ProjectsIds
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InvitationInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"invitee",
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

	varInvitationInput := _InvitationInput{}

	err = json.Unmarshal(data, &varInvitationInput)

	if err != nil {
		return err
	}

	*o = InvitationInput(varInvitationInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "invitee")
		delete(additionalProperties, "message")
		delete(additionalProperties, "organization_id")
		delete(additionalProperties, "projects_ids")
		delete(additionalProperties, "roles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvitationInput struct {
	value *InvitationInput
	isSet bool
}

func (v NullableInvitationInput) Get() *InvitationInput {
	return v.value
}

func (v *NullableInvitationInput) Set(val *InvitationInput) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitationInput) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitationInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitationInput(val *InvitationInput) *NullableInvitationInput {
	return &NullableInvitationInput{value: val, isSet: true}
}

func (v NullableInvitationInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitationInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
