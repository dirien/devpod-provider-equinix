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

// checks if the VerifyEmail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyEmail{}

// VerifyEmail struct for VerifyEmail
type VerifyEmail struct {
	// User verification token
	UserToken            string `json:"user_token"`
	AdditionalProperties map[string]interface{}
}

type _VerifyEmail VerifyEmail

// NewVerifyEmail instantiates a new VerifyEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyEmail(userToken string) *VerifyEmail {
	this := VerifyEmail{}
	this.UserToken = userToken
	return &this
}

// NewVerifyEmailWithDefaults instantiates a new VerifyEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyEmailWithDefaults() *VerifyEmail {
	this := VerifyEmail{}
	return &this
}

// GetUserToken returns the UserToken field value
func (o *VerifyEmail) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *VerifyEmail) GetUserTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value
func (o *VerifyEmail) SetUserToken(v string) {
	o.UserToken = v
}

func (o VerifyEmail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyEmail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_token"] = o.UserToken

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VerifyEmail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_token",
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

	varVerifyEmail := _VerifyEmail{}

	err = json.Unmarshal(data, &varVerifyEmail)

	if err != nil {
		return err
	}

	*o = VerifyEmail(varVerifyEmail)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "user_token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVerifyEmail struct {
	value *VerifyEmail
	isSet bool
}

func (v NullableVerifyEmail) Get() *VerifyEmail {
	return v.value
}

func (v *NullableVerifyEmail) Set(val *VerifyEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyEmail(val *VerifyEmail) *NullableVerifyEmail {
	return &NullableVerifyEmail{value: val, isSet: true}
}

func (v NullableVerifyEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
