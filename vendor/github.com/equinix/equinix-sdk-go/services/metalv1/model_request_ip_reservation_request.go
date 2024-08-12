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

// RequestIPReservationRequest - struct for RequestIPReservationRequest
type RequestIPReservationRequest struct {
	IPReservationRequestInput   *IPReservationRequestInput
	VrfIpReservationCreateInput *VrfIpReservationCreateInput
}

// IPReservationRequestInputAsRequestIPReservationRequest is a convenience function that returns IPReservationRequestInput wrapped in RequestIPReservationRequest
func IPReservationRequestInputAsRequestIPReservationRequest(v *IPReservationRequestInput) RequestIPReservationRequest {
	return RequestIPReservationRequest{
		IPReservationRequestInput: v,
	}
}

// VrfIpReservationCreateInputAsRequestIPReservationRequest is a convenience function that returns VrfIpReservationCreateInput wrapped in RequestIPReservationRequest
func VrfIpReservationCreateInputAsRequestIPReservationRequest(v *VrfIpReservationCreateInput) RequestIPReservationRequest {
	return RequestIPReservationRequest{
		VrfIpReservationCreateInput: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestIPReservationRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IPReservationRequestInput
	err = newStrictDecoder(data).Decode(&dst.IPReservationRequestInput)
	if err == nil {
		jsonIPReservationRequestInput, _ := json.Marshal(dst.IPReservationRequestInput)
		if string(jsonIPReservationRequestInput) == "{}" { // empty struct
			dst.IPReservationRequestInput = nil
		} else {
			match++
		}
	} else {
		dst.IPReservationRequestInput = nil
	}

	// try to unmarshal data into VrfIpReservationCreateInput
	err = newStrictDecoder(data).Decode(&dst.VrfIpReservationCreateInput)
	if err == nil {
		jsonVrfIpReservationCreateInput, _ := json.Marshal(dst.VrfIpReservationCreateInput)
		if string(jsonVrfIpReservationCreateInput) == "{}" { // empty struct
			dst.VrfIpReservationCreateInput = nil
		} else {
			match++
		}
	} else {
		dst.VrfIpReservationCreateInput = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.IPReservationRequestInput = nil
		dst.VrfIpReservationCreateInput = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RequestIPReservationRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RequestIPReservationRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestIPReservationRequest) MarshalJSON() ([]byte, error) {
	if src.IPReservationRequestInput != nil {
		return json.Marshal(&src.IPReservationRequestInput)
	}

	if src.VrfIpReservationCreateInput != nil {
		return json.Marshal(&src.VrfIpReservationCreateInput)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestIPReservationRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IPReservationRequestInput != nil {
		return obj.IPReservationRequestInput
	}

	if obj.VrfIpReservationCreateInput != nil {
		return obj.VrfIpReservationCreateInput
	}

	// all schemas are nil
	return nil
}

type NullableRequestIPReservationRequest struct {
	value *RequestIPReservationRequest
	isSet bool
}

func (v NullableRequestIPReservationRequest) Get() *RequestIPReservationRequest {
	return v.value
}

func (v *NullableRequestIPReservationRequest) Set(val *RequestIPReservationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestIPReservationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestIPReservationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestIPReservationRequest(val *RequestIPReservationRequest) *NullableRequestIPReservationRequest {
	return &NullableRequestIPReservationRequest{value: val, isSet: true}
}

func (v NullableRequestIPReservationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestIPReservationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
