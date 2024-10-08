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

// checks if the Invitation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Invitation{}

// Invitation struct for Invitation
type Invitation struct {
	CreatedAt            *time.Time             `json:"created_at,omitempty"`
	Href                 *string                `json:"href,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	Invitation           *Href                  `json:"invitation,omitempty"`
	InvitedBy            *Href                  `json:"invited_by,omitempty"`
	Invitee              *string                `json:"invitee,omitempty"`
	Nonce                *string                `json:"nonce,omitempty"`
	Organization         *Href                  `json:"organization,omitempty"`
	Projects             []Href                 `json:"projects,omitempty"`
	Roles                []InvitationRolesInner `json:"roles,omitempty"`
	UpdatedAt            *time.Time             `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Invitation Invitation

// NewInvitation instantiates a new Invitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitation() *Invitation {
	this := Invitation{}
	return &this
}

// NewInvitationWithDefaults instantiates a new Invitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationWithDefaults() *Invitation {
	this := Invitation{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Invitation) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Invitation) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Invitation) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Invitation) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Invitation) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Invitation) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Invitation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Invitation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Invitation) SetId(v string) {
	o.Id = &v
}

// GetInvitation returns the Invitation field value if set, zero value otherwise.
func (o *Invitation) GetInvitation() Href {
	if o == nil || IsNil(o.Invitation) {
		var ret Href
		return ret
	}
	return *o.Invitation
}

// GetInvitationOk returns a tuple with the Invitation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetInvitationOk() (*Href, bool) {
	if o == nil || IsNil(o.Invitation) {
		return nil, false
	}
	return o.Invitation, true
}

// HasInvitation returns a boolean if a field has been set.
func (o *Invitation) HasInvitation() bool {
	if o != nil && !IsNil(o.Invitation) {
		return true
	}

	return false
}

// SetInvitation gets a reference to the given Href and assigns it to the Invitation field.
func (o *Invitation) SetInvitation(v Href) {
	o.Invitation = &v
}

// GetInvitedBy returns the InvitedBy field value if set, zero value otherwise.
func (o *Invitation) GetInvitedBy() Href {
	if o == nil || IsNil(o.InvitedBy) {
		var ret Href
		return ret
	}
	return *o.InvitedBy
}

// GetInvitedByOk returns a tuple with the InvitedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetInvitedByOk() (*Href, bool) {
	if o == nil || IsNil(o.InvitedBy) {
		return nil, false
	}
	return o.InvitedBy, true
}

// HasInvitedBy returns a boolean if a field has been set.
func (o *Invitation) HasInvitedBy() bool {
	if o != nil && !IsNil(o.InvitedBy) {
		return true
	}

	return false
}

// SetInvitedBy gets a reference to the given Href and assigns it to the InvitedBy field.
func (o *Invitation) SetInvitedBy(v Href) {
	o.InvitedBy = &v
}

// GetInvitee returns the Invitee field value if set, zero value otherwise.
func (o *Invitation) GetInvitee() string {
	if o == nil || IsNil(o.Invitee) {
		var ret string
		return ret
	}
	return *o.Invitee
}

// GetInviteeOk returns a tuple with the Invitee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetInviteeOk() (*string, bool) {
	if o == nil || IsNil(o.Invitee) {
		return nil, false
	}
	return o.Invitee, true
}

// HasInvitee returns a boolean if a field has been set.
func (o *Invitation) HasInvitee() bool {
	if o != nil && !IsNil(o.Invitee) {
		return true
	}

	return false
}

// SetInvitee gets a reference to the given string and assigns it to the Invitee field.
func (o *Invitation) SetInvitee(v string) {
	o.Invitee = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *Invitation) GetNonce() string {
	if o == nil || IsNil(o.Nonce) {
		var ret string
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetNonceOk() (*string, bool) {
	if o == nil || IsNil(o.Nonce) {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *Invitation) HasNonce() bool {
	if o != nil && !IsNil(o.Nonce) {
		return true
	}

	return false
}

// SetNonce gets a reference to the given string and assigns it to the Nonce field.
func (o *Invitation) SetNonce(v string) {
	o.Nonce = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *Invitation) GetOrganization() Href {
	if o == nil || IsNil(o.Organization) {
		var ret Href
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetOrganizationOk() (*Href, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *Invitation) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Href and assigns it to the Organization field.
func (o *Invitation) SetOrganization(v Href) {
	o.Organization = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *Invitation) GetProjects() []Href {
	if o == nil || IsNil(o.Projects) {
		var ret []Href
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetProjectsOk() ([]Href, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *Invitation) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []Href and assigns it to the Projects field.
func (o *Invitation) SetProjects(v []Href) {
	o.Projects = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *Invitation) GetRoles() []InvitationRolesInner {
	if o == nil || IsNil(o.Roles) {
		var ret []InvitationRolesInner
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetRolesOk() ([]InvitationRolesInner, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *Invitation) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []InvitationRolesInner and assigns it to the Roles field.
func (o *Invitation) SetRoles(v []InvitationRolesInner) {
	o.Roles = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Invitation) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Invitation) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Invitation) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Invitation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Invitation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Invitation) {
		toSerialize["invitation"] = o.Invitation
	}
	if !IsNil(o.InvitedBy) {
		toSerialize["invited_by"] = o.InvitedBy
	}
	if !IsNil(o.Invitee) {
		toSerialize["invitee"] = o.Invitee
	}
	if !IsNil(o.Nonce) {
		toSerialize["nonce"] = o.Nonce
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Invitation) UnmarshalJSON(data []byte) (err error) {
	varInvitation := _Invitation{}

	err = json.Unmarshal(data, &varInvitation)

	if err != nil {
		return err
	}

	*o = Invitation(varInvitation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "href")
		delete(additionalProperties, "id")
		delete(additionalProperties, "invitation")
		delete(additionalProperties, "invited_by")
		delete(additionalProperties, "invitee")
		delete(additionalProperties, "nonce")
		delete(additionalProperties, "organization")
		delete(additionalProperties, "projects")
		delete(additionalProperties, "roles")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvitation struct {
	value *Invitation
	isSet bool
}

func (v NullableInvitation) Get() *Invitation {
	return v.value
}

func (v *NullableInvitation) Set(val *Invitation) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitation) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitation(val *Invitation) *NullableInvitation {
	return &NullableInvitation{value: val, isSet: true}
}

func (v NullableInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
