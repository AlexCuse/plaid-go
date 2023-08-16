/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.410.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// SandboxTransferTestClockCreateResponse Defines the response schema for `/sandbox/transfer/test_clock/create`
type SandboxTransferTestClockCreateResponse struct {
	TestClock TransferTestClock `json:"test_clock"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _SandboxTransferTestClockCreateResponse SandboxTransferTestClockCreateResponse

// NewSandboxTransferTestClockCreateResponse instantiates a new SandboxTransferTestClockCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxTransferTestClockCreateResponse(testClock TransferTestClock, requestId string) *SandboxTransferTestClockCreateResponse {
	this := SandboxTransferTestClockCreateResponse{}
	this.TestClock = testClock
	this.RequestId = requestId
	return &this
}

// NewSandboxTransferTestClockCreateResponseWithDefaults instantiates a new SandboxTransferTestClockCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxTransferTestClockCreateResponseWithDefaults() *SandboxTransferTestClockCreateResponse {
	this := SandboxTransferTestClockCreateResponse{}
	return &this
}

// GetTestClock returns the TestClock field value
func (o *SandboxTransferTestClockCreateResponse) GetTestClock() TransferTestClock {
	if o == nil {
		var ret TransferTestClock
		return ret
	}

	return o.TestClock
}

// GetTestClockOk returns a tuple with the TestClock field value
// and a boolean to check if the value has been set.
func (o *SandboxTransferTestClockCreateResponse) GetTestClockOk() (*TransferTestClock, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TestClock, true
}

// SetTestClock sets field value
func (o *SandboxTransferTestClockCreateResponse) SetTestClock(v TransferTestClock) {
	o.TestClock = v
}

// GetRequestId returns the RequestId field value
func (o *SandboxTransferTestClockCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *SandboxTransferTestClockCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *SandboxTransferTestClockCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o SandboxTransferTestClockCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["test_clock"] = o.TestClock
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SandboxTransferTestClockCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSandboxTransferTestClockCreateResponse := _SandboxTransferTestClockCreateResponse{}

	if err = json.Unmarshal(bytes, &varSandboxTransferTestClockCreateResponse); err == nil {
		*o = SandboxTransferTestClockCreateResponse(varSandboxTransferTestClockCreateResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "test_clock")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSandboxTransferTestClockCreateResponse struct {
	value *SandboxTransferTestClockCreateResponse
	isSet bool
}

func (v NullableSandboxTransferTestClockCreateResponse) Get() *SandboxTransferTestClockCreateResponse {
	return v.value
}

func (v *NullableSandboxTransferTestClockCreateResponse) Set(val *SandboxTransferTestClockCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxTransferTestClockCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxTransferTestClockCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxTransferTestClockCreateResponse(val *SandboxTransferTestClockCreateResponse) *NullableSandboxTransferTestClockCreateResponse {
	return &NullableSandboxTransferTestClockCreateResponse{value: val, isSet: true}
}

func (v NullableSandboxTransferTestClockCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxTransferTestClockCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


