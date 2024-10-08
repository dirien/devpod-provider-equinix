/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
	"fmt"
)

// checks if the ProjectCreateFromRootInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectCreateFromRootInput{}

// ProjectCreateFromRootInput struct for ProjectCreateFromRootInput
type ProjectCreateFromRootInput struct {
	Customdata map[string]interface{} `json:"customdata,omitempty"`
	// The name of the project. Cannot contain characters encoded in greater than 3 bytes such as emojis.
	Name                 string                          `json:"name"`
	OrganizationId       *string                         `json:"organization_id,omitempty"`
	PaymentMethodId      *string                         `json:"payment_method_id,omitempty"`
	Type                 *ProjectCreateFromRootInputType `json:"type,omitempty"`
	Tags                 []string                        `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProjectCreateFromRootInput ProjectCreateFromRootInput

// NewProjectCreateFromRootInput instantiates a new ProjectCreateFromRootInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectCreateFromRootInput(name string) *ProjectCreateFromRootInput {
	this := ProjectCreateFromRootInput{}
	this.Name = name
	return &this
}

// NewProjectCreateFromRootInputWithDefaults instantiates a new ProjectCreateFromRootInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectCreateFromRootInputWithDefaults() *ProjectCreateFromRootInput {
	this := ProjectCreateFromRootInput{}
	return &this
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *ProjectCreateFromRootInput) GetCustomdata() map[string]interface{} {
	if o == nil || IsNil(o.Customdata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreateFromRootInput) GetCustomdataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Customdata) {
		return map[string]interface{}{}, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *ProjectCreateFromRootInput) HasCustomdata() bool {
	if o != nil && !IsNil(o.Customdata) {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *ProjectCreateFromRootInput) SetCustomdata(v map[string]interface{}) {
	o.Customdata = v
}

// GetName returns the Name field value
func (o *ProjectCreateFromRootInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectCreateFromRootInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectCreateFromRootInput) SetName(v string) {
	o.Name = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ProjectCreateFromRootInput) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreateFromRootInput) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ProjectCreateFromRootInput) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *ProjectCreateFromRootInput) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise.
func (o *ProjectCreateFromRootInput) GetPaymentMethodId() string {
	if o == nil || IsNil(o.PaymentMethodId) {
		var ret string
		return ret
	}
	return *o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreateFromRootInput) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethodId) {
		return nil, false
	}
	return o.PaymentMethodId, true
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *ProjectCreateFromRootInput) HasPaymentMethodId() bool {
	if o != nil && !IsNil(o.PaymentMethodId) {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given string and assigns it to the PaymentMethodId field.
func (o *ProjectCreateFromRootInput) SetPaymentMethodId(v string) {
	o.PaymentMethodId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProjectCreateFromRootInput) GetType() ProjectCreateFromRootInputType {
	if o == nil || IsNil(o.Type) {
		var ret ProjectCreateFromRootInputType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreateFromRootInput) GetTypeOk() (*ProjectCreateFromRootInputType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProjectCreateFromRootInput) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ProjectCreateFromRootInputType and assigns it to the Type field.
func (o *ProjectCreateFromRootInput) SetType(v ProjectCreateFromRootInputType) {
	o.Type = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ProjectCreateFromRootInput) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreateFromRootInput) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ProjectCreateFromRootInput) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ProjectCreateFromRootInput) SetTags(v []string) {
	o.Tags = v
}

func (o ProjectCreateFromRootInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectCreateFromRootInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Customdata) {
		toSerialize["customdata"] = o.Customdata
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !IsNil(o.PaymentMethodId) {
		toSerialize["payment_method_id"] = o.PaymentMethodId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectCreateFromRootInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varProjectCreateFromRootInput := _ProjectCreateFromRootInput{}

	err = json.Unmarshal(data, &varProjectCreateFromRootInput)

	if err != nil {
		return err
	}

	*o = ProjectCreateFromRootInput(varProjectCreateFromRootInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customdata")
		delete(additionalProperties, "name")
		delete(additionalProperties, "organization_id")
		delete(additionalProperties, "payment_method_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectCreateFromRootInput struct {
	value *ProjectCreateFromRootInput
	isSet bool
}

func (v NullableProjectCreateFromRootInput) Get() *ProjectCreateFromRootInput {
	return v.value
}

func (v *NullableProjectCreateFromRootInput) Set(val *ProjectCreateFromRootInput) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectCreateFromRootInput) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectCreateFromRootInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectCreateFromRootInput(val *ProjectCreateFromRootInput) *NullableProjectCreateFromRootInput {
	return &NullableProjectCreateFromRootInput{value: val, isSet: true}
}

func (v NullableProjectCreateFromRootInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectCreateFromRootInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
