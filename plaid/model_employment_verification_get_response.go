/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.77.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// EmploymentVerificationGetResponse EmploymentVerificationGetResponse defines the response schema for `/employment/verification/get`.
type EmploymentVerificationGetResponse struct {
	// A list of employment verification summaries.
	Employments []EmploymentVerification `json:"employments"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _EmploymentVerificationGetResponse EmploymentVerificationGetResponse

// NewEmploymentVerificationGetResponse instantiates a new EmploymentVerificationGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmploymentVerificationGetResponse(employments []EmploymentVerification, requestId string) *EmploymentVerificationGetResponse {
	this := EmploymentVerificationGetResponse{}
	this.Employments = employments
	this.RequestId = requestId
	return &this
}

// NewEmploymentVerificationGetResponseWithDefaults instantiates a new EmploymentVerificationGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmploymentVerificationGetResponseWithDefaults() *EmploymentVerificationGetResponse {
	this := EmploymentVerificationGetResponse{}
	return &this
}

// GetEmployments returns the Employments field value
func (o *EmploymentVerificationGetResponse) GetEmployments() []EmploymentVerification {
	if o == nil {
		var ret []EmploymentVerification
		return ret
	}

	return o.Employments
}

// GetEmploymentsOk returns a tuple with the Employments field value
// and a boolean to check if the value has been set.
func (o *EmploymentVerificationGetResponse) GetEmploymentsOk() (*[]EmploymentVerification, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Employments, true
}

// SetEmployments sets field value
func (o *EmploymentVerificationGetResponse) SetEmployments(v []EmploymentVerification) {
	o.Employments = v
}

// GetRequestId returns the RequestId field value
func (o *EmploymentVerificationGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *EmploymentVerificationGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *EmploymentVerificationGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o EmploymentVerificationGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["employments"] = o.Employments
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmploymentVerificationGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varEmploymentVerificationGetResponse := _EmploymentVerificationGetResponse{}

	if err = json.Unmarshal(bytes, &varEmploymentVerificationGetResponse); err == nil {
		*o = EmploymentVerificationGetResponse(varEmploymentVerificationGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "employments")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmploymentVerificationGetResponse struct {
	value *EmploymentVerificationGetResponse
	isSet bool
}

func (v NullableEmploymentVerificationGetResponse) Get() *EmploymentVerificationGetResponse {
	return v.value
}

func (v *NullableEmploymentVerificationGetResponse) Set(val *EmploymentVerificationGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmploymentVerificationGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmploymentVerificationGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmploymentVerificationGetResponse(val *EmploymentVerificationGetResponse) *NullableEmploymentVerificationGetResponse {
	return &NullableEmploymentVerificationGetResponse{value: val, isSet: true}
}

func (v NullableEmploymentVerificationGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmploymentVerificationGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


